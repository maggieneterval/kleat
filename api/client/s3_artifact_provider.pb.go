// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.2
// source: s3_artifact_provider.proto

package client

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the S3 artifact provider.
type S3ArtifactProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the S3 artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured S3 artifact accounts.
	Accounts []*S3ArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *S3ArtifactProvider) Reset() {
	*x = S3ArtifactProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_artifact_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ArtifactProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ArtifactProvider) ProtoMessage() {}

func (x *S3ArtifactProvider) ProtoReflect() protoreflect.Message {
	mi := &file_s3_artifact_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ArtifactProvider.ProtoReflect.Descriptor instead.
func (*S3ArtifactProvider) Descriptor() ([]byte, []int) {
	return file_s3_artifact_provider_proto_rawDescGZIP(), []int{0}
}

func (x *S3ArtifactProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *S3ArtifactProvider) GetAccounts() []*S3ArtifactAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for an S3 artifact account.
type S3ArtifactAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The S3 API endpoint. Only required when using an S3 clone such as Minio.
	ApiEndpoint string `protobuf:"bytes,2,opt,name=apiEndpoint,proto3" json:"apiEndpoint,omitempty"`
	// The S3 API region. Only required when using an S3 clone such as Minio.
	ApiRegion string `protobuf:"bytes,3,opt,name=apiRegion,proto3" json:"apiRegion,omitempty"`
	// The AWS Access Key ID. If not provided, Spinnaker will try to find AWS
	// credentials as described at
	// http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default.
	AwsAccessKeyId string `protobuf:"bytes,4,opt,name=awsAccessKeyId,proto3" json:"awsAccessKeyId,omitempty"`
	// The AWS Secret Key.
	AwsSecretAccessKey string `protobuf:"bytes,5,opt,name=awsSecretAccessKey,proto3" json:"awsSecretAccessKey,omitempty"`
	// The S3 region.
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *S3ArtifactAccount) Reset() {
	*x = S3ArtifactAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_artifact_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ArtifactAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ArtifactAccount) ProtoMessage() {}

func (x *S3ArtifactAccount) ProtoReflect() protoreflect.Message {
	mi := &file_s3_artifact_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ArtifactAccount.ProtoReflect.Descriptor instead.
func (*S3ArtifactAccount) Descriptor() ([]byte, []int) {
	return file_s3_artifact_provider_proto_rawDescGZIP(), []int{1}
}

func (x *S3ArtifactAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *S3ArtifactAccount) GetApiEndpoint() string {
	if x != nil {
		return x.ApiEndpoint
	}
	return ""
}

func (x *S3ArtifactAccount) GetApiRegion() string {
	if x != nil {
		return x.ApiRegion
	}
	return ""
}

func (x *S3ArtifactAccount) GetAwsAccessKeyId() string {
	if x != nil {
		return x.AwsAccessKeyId
	}
	return ""
}

func (x *S3ArtifactAccount) GetAwsSecretAccessKey() string {
	if x != nil {
		return x.AwsSecretAccessKey
	}
	return ""
}

func (x *S3ArtifactAccount) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_s3_artifact_provider_proto protoreflect.FileDescriptor

var file_s3_artifact_provider_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x33, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x12, 0x53, 0x33, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x33,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x53, 0x33,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x77, 0x73,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x61,
	0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s3_artifact_provider_proto_rawDescOnce sync.Once
	file_s3_artifact_provider_proto_rawDescData = file_s3_artifact_provider_proto_rawDesc
)

func file_s3_artifact_provider_proto_rawDescGZIP() []byte {
	file_s3_artifact_provider_proto_rawDescOnce.Do(func() {
		file_s3_artifact_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_s3_artifact_provider_proto_rawDescData)
	})
	return file_s3_artifact_provider_proto_rawDescData
}

var file_s3_artifact_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_s3_artifact_provider_proto_goTypes = []interface{}{
	(*S3ArtifactProvider)(nil), // 0: proto.S3ArtifactProvider
	(*S3ArtifactAccount)(nil),  // 1: proto.S3ArtifactAccount
}
var file_s3_artifact_provider_proto_depIdxs = []int32{
	1, // 0: proto.S3ArtifactProvider.accounts:type_name -> proto.S3ArtifactAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_s3_artifact_provider_proto_init() }
func file_s3_artifact_provider_proto_init() {
	if File_s3_artifact_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s3_artifact_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ArtifactProvider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_s3_artifact_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ArtifactAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_s3_artifact_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_s3_artifact_provider_proto_goTypes,
		DependencyIndexes: file_s3_artifact_provider_proto_depIdxs,
		MessageInfos:      file_s3_artifact_provider_proto_msgTypes,
	}.Build()
	File_s3_artifact_provider_proto = out.File
	file_s3_artifact_provider_proto_rawDesc = nil
	file_s3_artifact_provider_proto_goTypes = nil
	file_s3_artifact_provider_proto_depIdxs = nil
}
