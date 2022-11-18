// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: service.proto

package protoconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	proto "github.com/xmtp/example-notification-server-go/pkg/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// NotificationsName is the fully-qualified name of the Notifications service.
	NotificationsName = "notifications.v1.Notifications"
)

// NotificationsClient is a client for the notifications.v1.Notifications service.
type NotificationsClient interface {
	RegisterInstallation(context.Context, *connect_go.Request[proto.RegisterInstallationRequest]) (*connect_go.Response[proto.RegisterInstallationResponse], error)
	DeleteInstallation(context.Context, *connect_go.Request[proto.DeleteInstallationRequest]) (*connect_go.Response[emptypb.Empty], error)
	Subscribe(context.Context, *connect_go.Request[proto.SubscribeRequest]) (*connect_go.Response[emptypb.Empty], error)
	Unsubscribe(context.Context, *connect_go.Request[proto.UnsubscribeRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewNotificationsClient constructs a client for the notifications.v1.Notifications service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNotificationsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NotificationsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &notificationsClient{
		registerInstallation: connect_go.NewClient[proto.RegisterInstallationRequest, proto.RegisterInstallationResponse](
			httpClient,
			baseURL+"/notifications.v1.Notifications/RegisterInstallation",
			opts...,
		),
		deleteInstallation: connect_go.NewClient[proto.DeleteInstallationRequest, emptypb.Empty](
			httpClient,
			baseURL+"/notifications.v1.Notifications/DeleteInstallation",
			opts...,
		),
		subscribe: connect_go.NewClient[proto.SubscribeRequest, emptypb.Empty](
			httpClient,
			baseURL+"/notifications.v1.Notifications/Subscribe",
			opts...,
		),
		unsubscribe: connect_go.NewClient[proto.UnsubscribeRequest, emptypb.Empty](
			httpClient,
			baseURL+"/notifications.v1.Notifications/Unsubscribe",
			opts...,
		),
	}
}

// notificationsClient implements NotificationsClient.
type notificationsClient struct {
	registerInstallation *connect_go.Client[proto.RegisterInstallationRequest, proto.RegisterInstallationResponse]
	deleteInstallation   *connect_go.Client[proto.DeleteInstallationRequest, emptypb.Empty]
	subscribe            *connect_go.Client[proto.SubscribeRequest, emptypb.Empty]
	unsubscribe          *connect_go.Client[proto.UnsubscribeRequest, emptypb.Empty]
}

// RegisterInstallation calls notifications.v1.Notifications.RegisterInstallation.
func (c *notificationsClient) RegisterInstallation(ctx context.Context, req *connect_go.Request[proto.RegisterInstallationRequest]) (*connect_go.Response[proto.RegisterInstallationResponse], error) {
	return c.registerInstallation.CallUnary(ctx, req)
}

// DeleteInstallation calls notifications.v1.Notifications.DeleteInstallation.
func (c *notificationsClient) DeleteInstallation(ctx context.Context, req *connect_go.Request[proto.DeleteInstallationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteInstallation.CallUnary(ctx, req)
}

// Subscribe calls notifications.v1.Notifications.Subscribe.
func (c *notificationsClient) Subscribe(ctx context.Context, req *connect_go.Request[proto.SubscribeRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.subscribe.CallUnary(ctx, req)
}

// Unsubscribe calls notifications.v1.Notifications.Unsubscribe.
func (c *notificationsClient) Unsubscribe(ctx context.Context, req *connect_go.Request[proto.UnsubscribeRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.unsubscribe.CallUnary(ctx, req)
}

// NotificationsHandler is an implementation of the notifications.v1.Notifications service.
type NotificationsHandler interface {
	RegisterInstallation(context.Context, *connect_go.Request[proto.RegisterInstallationRequest]) (*connect_go.Response[proto.RegisterInstallationResponse], error)
	DeleteInstallation(context.Context, *connect_go.Request[proto.DeleteInstallationRequest]) (*connect_go.Response[emptypb.Empty], error)
	Subscribe(context.Context, *connect_go.Request[proto.SubscribeRequest]) (*connect_go.Response[emptypb.Empty], error)
	Unsubscribe(context.Context, *connect_go.Request[proto.UnsubscribeRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewNotificationsHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNotificationsHandler(svc NotificationsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/notifications.v1.Notifications/RegisterInstallation", connect_go.NewUnaryHandler(
		"/notifications.v1.Notifications/RegisterInstallation",
		svc.RegisterInstallation,
		opts...,
	))
	mux.Handle("/notifications.v1.Notifications/DeleteInstallation", connect_go.NewUnaryHandler(
		"/notifications.v1.Notifications/DeleteInstallation",
		svc.DeleteInstallation,
		opts...,
	))
	mux.Handle("/notifications.v1.Notifications/Subscribe", connect_go.NewUnaryHandler(
		"/notifications.v1.Notifications/Subscribe",
		svc.Subscribe,
		opts...,
	))
	mux.Handle("/notifications.v1.Notifications/Unsubscribe", connect_go.NewUnaryHandler(
		"/notifications.v1.Notifications/Unsubscribe",
		svc.Unsubscribe,
		opts...,
	))
	return "/notifications.v1.Notifications/", mux
}

// UnimplementedNotificationsHandler returns CodeUnimplemented from all methods.
type UnimplementedNotificationsHandler struct{}

func (UnimplementedNotificationsHandler) RegisterInstallation(context.Context, *connect_go.Request[proto.RegisterInstallationRequest]) (*connect_go.Response[proto.RegisterInstallationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notifications.v1.Notifications.RegisterInstallation is not implemented"))
}

func (UnimplementedNotificationsHandler) DeleteInstallation(context.Context, *connect_go.Request[proto.DeleteInstallationRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notifications.v1.Notifications.DeleteInstallation is not implemented"))
}

func (UnimplementedNotificationsHandler) Subscribe(context.Context, *connect_go.Request[proto.SubscribeRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notifications.v1.Notifications.Subscribe is not implemented"))
}

func (UnimplementedNotificationsHandler) Unsubscribe(context.Context, *connect_go.Request[proto.UnsubscribeRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("notifications.v1.Notifications.Unsubscribe is not implemented"))
}