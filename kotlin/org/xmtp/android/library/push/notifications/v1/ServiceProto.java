// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: notifications/v1/service.proto
// Protobuf Java Version: 4.26.0

package org.xmtp.android.library.push.notifications.v1;

public final class ServiceProto {
  private ServiceProto() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 26,
      /* patch= */ 0,
      /* suffix= */ "",
      ServiceProto.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_DeliveryMechanism_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_DeliveryMechanism_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_RegisterInstallationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_RegisterInstallationRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_RegisterInstallationResponse_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_RegisterInstallationResponse_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_DeleteInstallationRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_DeleteInstallationRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_Subscription_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_Subscription_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_Subscription_HmacKey_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_Subscription_HmacKey_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_SubscribeWithMetadataRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_SubscribeWithMetadataRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_SubscribeRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_SubscribeRequest_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_notifications_v1_UnsubscribeRequest_descriptor;
  static final 
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_notifications_v1_UnsubscribeRequest_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\036notifications/v1/service.proto\022\020notifi" +
      "cations.v1\032\033google/protobuf/empty.proto\"" +
      "\267\001\n\021DeliveryMechanism\022,\n\021apns_device_tok" +
      "en\030\001 \001(\tH\000R\017apnsDeviceToken\0224\n\025firebase_" +
      "device_token\030\002 \001(\tH\000R\023firebaseDeviceToke" +
      "n\022#\n\014custom_token\030\003 \001(\tH\000R\013customTokenB\031" +
      "\n\027delivery_mechanism_type\"\232\001\n\033RegisterIn" +
      "stallationRequest\022\'\n\017installation_id\030\001 \001" +
      "(\tR\016installationId\022R\n\022delivery_mechanism" +
      "\030\002 \001(\0132#.notifications.v1.DeliveryMechan" +
      "ismR\021deliveryMechanism\"h\n\034RegisterInstal" +
      "lationResponse\022\'\n\017installation_id\030\001 \001(\tR" +
      "\016installationId\022\037\n\013valid_until\030\002 \001(\004R\nva" +
      "lidUntil\"D\n\031DeleteInstallationRequest\022\'\n" +
      "\017installation_id\030\001 \001(\tR\016installationId\"\347" +
      "\001\n\014Subscription\022\024\n\005topic\030\001 \001(\tR\005topic\022C\n" +
      "\thmac_keys\030\002 \003(\0132&.notifications.v1.Subs" +
      "cription.HmacKeyR\010hmacKeys\022\033\n\tis_silent\030" +
      "\003 \001(\010R\010isSilent\032_\n\007HmacKey\022B\n\036thirty_day" +
      "_periods_since_epoch\030\001 \001(\rR\032thirtyDayPer" +
      "iodsSinceEpoch\022\020\n\003key\030\002 \001(\014R\003key\"\215\001\n\034Sub" +
      "scribeWithMetadataRequest\022\'\n\017installatio" +
      "n_id\030\001 \001(\tR\016installationId\022D\n\rsubscripti" +
      "ons\030\002 \003(\0132\036.notifications.v1.Subscriptio" +
      "nR\rsubscriptions\"S\n\020SubscribeRequest\022\'\n\017" +
      "installation_id\030\001 \001(\tR\016installationId\022\026\n" +
      "\006topics\030\002 \003(\tR\006topics\"U\n\022UnsubscribeRequ" +
      "est\022\'\n\017installation_id\030\001 \001(\tR\016installati" +
      "onId\022\026\n\006topics\030\002 \003(\tR\006topics2\330\003\n\rNotific" +
      "ations\022u\n\024RegisterInstallation\022-.notific" +
      "ations.v1.RegisterInstallationRequest\032.." +
      "notifications.v1.RegisterInstallationRes" +
      "ponse\022Y\n\022DeleteInstallation\022+.notificati" +
      "ons.v1.DeleteInstallationRequest\032\026.googl" +
      "e.protobuf.Empty\022G\n\tSubscribe\022\".notifica" +
      "tions.v1.SubscribeRequest\032\026.google.proto" +
      "buf.Empty\022_\n\025SubscribeWithMetadata\022..not" +
      "ifications.v1.SubscribeWithMetadataReque" +
      "st\032\026.google.protobuf.Empty\022K\n\013Unsubscrib" +
      "e\022$.notifications.v1.UnsubscribeRequest\032" +
      "\026.google.protobuf.EmptyB\372\001\n.org.xmtp.and" +
      "roid.library.push.notifications.v1B\014Serv" +
      "iceProtoP\001ZYgithub.com/xmtp/example-noti" +
      "fication-server-go/pkg/proto/notificatio" +
      "ns/v1;notificationsv1\242\002\003NXX\252\002\020Notificati" +
      "ons.V1\312\002\020Notifications\\V1\342\002\034Notification" +
      "s\\V1\\GPBMetadata\352\002\021Notifications::V1b\006pr" +
      "oto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          com.google.protobuf.EmptyProto.getDescriptor(),
        });
    internal_static_notifications_v1_DeliveryMechanism_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_notifications_v1_DeliveryMechanism_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_DeliveryMechanism_descriptor,
        new java.lang.String[] { "ApnsDeviceToken", "FirebaseDeviceToken", "CustomToken", "DeliveryMechanismType", });
    internal_static_notifications_v1_RegisterInstallationRequest_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_notifications_v1_RegisterInstallationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_RegisterInstallationRequest_descriptor,
        new java.lang.String[] { "InstallationId", "DeliveryMechanism", });
    internal_static_notifications_v1_RegisterInstallationResponse_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_notifications_v1_RegisterInstallationResponse_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_RegisterInstallationResponse_descriptor,
        new java.lang.String[] { "InstallationId", "ValidUntil", });
    internal_static_notifications_v1_DeleteInstallationRequest_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_notifications_v1_DeleteInstallationRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_DeleteInstallationRequest_descriptor,
        new java.lang.String[] { "InstallationId", });
    internal_static_notifications_v1_Subscription_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_notifications_v1_Subscription_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_Subscription_descriptor,
        new java.lang.String[] { "Topic", "HmacKeys", "IsSilent", });
    internal_static_notifications_v1_Subscription_HmacKey_descriptor =
      internal_static_notifications_v1_Subscription_descriptor.getNestedTypes().get(0);
    internal_static_notifications_v1_Subscription_HmacKey_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_Subscription_HmacKey_descriptor,
        new java.lang.String[] { "ThirtyDayPeriodsSinceEpoch", "Key", });
    internal_static_notifications_v1_SubscribeWithMetadataRequest_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_notifications_v1_SubscribeWithMetadataRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_SubscribeWithMetadataRequest_descriptor,
        new java.lang.String[] { "InstallationId", "Subscriptions", });
    internal_static_notifications_v1_SubscribeRequest_descriptor =
      getDescriptor().getMessageTypes().get(6);
    internal_static_notifications_v1_SubscribeRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_SubscribeRequest_descriptor,
        new java.lang.String[] { "InstallationId", "Topics", });
    internal_static_notifications_v1_UnsubscribeRequest_descriptor =
      getDescriptor().getMessageTypes().get(7);
    internal_static_notifications_v1_UnsubscribeRequest_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_notifications_v1_UnsubscribeRequest_descriptor,
        new java.lang.String[] { "InstallationId", "Topics", });
    descriptor.resolveAllFeaturesImmutable();
    com.google.protobuf.EmptyProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
