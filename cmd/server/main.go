package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"github.com/xmtp/example-notification-server-go/pkg/api"
	"github.com/xmtp/example-notification-server-go/pkg/db"
	database "github.com/xmtp/example-notification-server-go/pkg/db"
	"github.com/xmtp/example-notification-server-go/pkg/db/migrations"
	"github.com/xmtp/example-notification-server-go/pkg/delivery"
	"github.com/xmtp/example-notification-server-go/pkg/installations"
	"github.com/xmtp/example-notification-server-go/pkg/logging"
	"github.com/xmtp/example-notification-server-go/pkg/options"
	"github.com/xmtp/example-notification-server-go/pkg/subscriptions"
	"github.com/xmtp/example-notification-server-go/pkg/xmtp"
	"go.uber.org/zap"
)

var opts options.Options
var logger *zap.Logger

func main() {
	var err error
	if _, err = flags.Parse(&opts); err != nil {
		if err, ok := err.(*flags.Error); !ok || err.Type != flags.ErrHelp {
			log.Fatalf("Could not parse options: %s", err)
		}
		return
	}

	logger = logging.CreateLogger(opts.LogEncoding, opts.LogLevel)

	if opts.CreateMigration != "" {
		if err = createMigration(); err != nil {
			logger.Fatal("failed to create migration", zap.Error(err))
		}
		return
	}

	if !opts.Xmtp.ListenerEnabled && !opts.Api.Enabled {
		logger.Fatal("no --api or --xmtp-listener flags applied")
	}

	db := initDb()
	ctx, cancel := context.WithCancel(context.Background())
	installationsService := installations.NewInstallationsService(logger, db)
	subscriptionsService := subscriptions.NewSubscriptionsService(logger, db)
	var listener *xmtp.Listener
	var apiServer *api.ApiServer

	if opts.Xmtp.ListenerEnabled {
		var apns *delivery.ApnsDelivery
		var fcm *delivery.FcmDelivery
		var err error

		if opts.Apns.KeyId != "" {
			apns, err = delivery.NewApnsDelivery(logger, opts.Apns)
			if err != nil {
				logger.Fatal("failed to initialize APNS", zap.Error(err))
			}
		}

		if opts.Fcm.CredentialsJson != "" {
			fcm, err = delivery.NewFcmDelivery(ctx, logger, opts.Fcm)
			if err != nil {
				logger.Fatal("failed to initialize FCM", zap.Error(err))
			}
		}

		deliveryService := delivery.NewDeliveryService(logger, apns, fcm)
		listener, err = xmtp.NewListener(ctx, logger, opts.Xmtp, installationsService, subscriptionsService, deliveryService)
		if err != nil {
			logger.Fatal("failed to initialize listener", zap.Error(err))
		}
		listener.Start()
	}

	if opts.Api.Enabled {
		apiServer = api.NewApiServer(logger, installationsService, subscriptionsService)
		apiServer.Start()
	}

	waitForShutdown()

	if apiServer != nil {
		apiServer.Stop()
	}

	if listener != nil {
		listener.Stop()
	}

	cancel()
}

// Commenting out as these are currently unused
func waitForShutdown() {
	termChannel := make(chan os.Signal, 1)
	signal.Notify(termChannel, syscall.SIGINT, syscall.SIGTERM)
	<-termChannel
}

func initDb() *bun.DB {
	db, err := database.CreateBunDB(opts.DbConnectionString, 10*time.Second)
	if err != nil {
		log.Fatal("db creation error", zap.Error(err))
	}

	err = database.Migrate(context.Background(), db)
	if err != nil {
		log.Fatal("db migration error", zap.Error(err))
	}

	return db
}

func createMigration() error {
	db, err := db.CreateBunDB(opts.DbConnectionString, 30*time.Second)
	if err != nil {
		return err
	}

	migrator := migrate.NewMigrator(db, migrations.Migrations)
	files, err := migrator.CreateSQLMigrations(context.Background(), opts.CreateMigration)
	for _, mf := range files {
		fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
	}

	return err
}
