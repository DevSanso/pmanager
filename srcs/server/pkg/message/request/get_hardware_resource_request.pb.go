// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: message/request/get_hardware_resource_request.proto

package request

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "pkg/message/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHardwareResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client    *common.ClientInfo        `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Resources []common.HardwareResource `protobuf:"varint,2,rep,packed,name=resources,proto3,enum=common.HardwareResource" json:"resources,omitempty"`
}

func (x *GetHardwareResourceRequest) Reset() {
	*x = GetHardwareResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_request_get_hardware_resource_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHardwareResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHardwareResourceRequest) ProtoMessage() {}

func (x *GetHardwareResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_request_get_hardware_resource_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHardwareResourceRequest.ProtoReflect.Descriptor instead.
func (*GetHardwareResourceRequest) Descriptor() ([]byte, []int) {
	return file_message_request_get_hardware_resource_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetHardwareResourceRequest) GetClient() *common.ClientInfo {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *GetHardwareResourceRequest) GetResources() []common.HardwareResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_message_request_get_hardware_resource_request_proto protoreflect.FileDescriptor

var file_message_request_get_hardware_resource_request_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_request_get_hardware_resource_request_proto_rawDescOnce sync.Once
	file_message_request_get_hardware_resource_request_proto_rawDescData = file_message_request_get_hardware_resource_request_proto_rawDesc
)

func file_message_request_get_hardware_resource_request_proto_rawDescGZIP() []byte {
	file_message_request_get_hardware_resource_request_proto_rawDescOnce.Do(func() {
		file_message_request_get_hardware_resource_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_request_get_hardware_resource_request_proto_rawDescData)
	})
	return file_message_request_get_hardware_resource_request_proto_rawDescData
}

var file_message_request_get_hardware_resource_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_request_get_hardware_resource_request_proto_goTypes = []interface{}{
	(*GetHardwareResourceRequest)(nil), // 0: request.GetHardwareResourceRequest
	(*common.ClientInfo)(nil),          // 1: common.ClientInfo
	(common.HardwareResource)(0),       // 2: common.HardwareResource
}
var file_message_request_get_hardware_resource_request_proto_depIdxs = []int32{
	1, // 0: request.GetHardwareResourceRequest.client:type_name -> common.ClientInfo
	2, // 1: request.GetHardwareResourceRequest.resources:type_name -> common.HardwareResource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_request_get_hardware_resource_request_proto_init() }
func file_message_request_get_hardware_resource_request_proto_init() {
	if File_message_request_get_hardware_resource_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_request_get_hardware_resource_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHardwareResourceRequest); i {
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
			RawDescriptor: file_message_request_get_hardware_resource_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_request_get_hardware_resource_request_proto_goTypes,
		DependencyIndexes: file_message_request_get_hardware_resource_request_proto_depIdxs,
		MessageInfos:      file_message_request_get_hardware_resource_request_proto_msgTypes,
	}.Build()
	File_message_request_get_hardware_resource_request_proto = out.File
	file_message_request_get_hardware_resource_request_proto_rawDesc = nil
	file_message_request_get_hardware_resource_request_proto_goTypes = nil
	file_message_request_get_hardware_resource_request_proto_depIdxs = nil
}
