// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: message/response/get_group_id_response.proto

package response

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

type GetGroupIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*GetGroupIdResponse_Success_
	//	*GetGroupIdResponse_Error
	Data isGetGroupIdResponse_Data `protobuf_oneof:"data"`
}

func (x *GetGroupIdResponse) Reset() {
	*x = GetGroupIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_response_get_group_id_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupIdResponse) ProtoMessage() {}

func (x *GetGroupIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_response_get_group_id_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupIdResponse.ProtoReflect.Descriptor instead.
func (*GetGroupIdResponse) Descriptor() ([]byte, []int) {
	return file_message_response_get_group_id_response_proto_rawDescGZIP(), []int{0}
}

func (m *GetGroupIdResponse) GetData() isGetGroupIdResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *GetGroupIdResponse) GetSuccess() *GetGroupIdResponse_Success {
	if x, ok := x.GetData().(*GetGroupIdResponse_Success_); ok {
		return x.Success
	}
	return nil
}

func (x *GetGroupIdResponse) GetError() *common.ErrorMessage {
	if x, ok := x.GetData().(*GetGroupIdResponse_Error); ok {
		return x.Error
	}
	return nil
}

type isGetGroupIdResponse_Data interface {
	isGetGroupIdResponse_Data()
}

type GetGroupIdResponse_Success_ struct {
	Success *GetGroupIdResponse_Success `protobuf:"bytes,1,opt,name=success,proto3,oneof"`
}

type GetGroupIdResponse_Error struct {
	Error *common.ErrorMessage `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetGroupIdResponse_Success_) isGetGroupIdResponse_Data() {}

func (*GetGroupIdResponse_Error) isGetGroupIdResponse_Data() {}

type GetGroupIdResponse_Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetGroupIdResponse_Success) Reset() {
	*x = GetGroupIdResponse_Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_response_get_group_id_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupIdResponse_Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupIdResponse_Success) ProtoMessage() {}

func (x *GetGroupIdResponse_Success) ProtoReflect() protoreflect.Message {
	mi := &file_message_response_get_group_id_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupIdResponse_Success.ProtoReflect.Descriptor instead.
func (*GetGroupIdResponse_Success) Descriptor() ([]byte, []int) {
	return file_message_response_get_group_id_response_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GetGroupIdResponse_Success) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_message_response_get_group_id_response_proto protoreflect.FileDescriptor

var file_message_response_get_group_id_response_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x1a, 0x24, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_message_response_get_group_id_response_proto_rawDescOnce sync.Once
	file_message_response_get_group_id_response_proto_rawDescData = file_message_response_get_group_id_response_proto_rawDesc
)

func file_message_response_get_group_id_response_proto_rawDescGZIP() []byte {
	file_message_response_get_group_id_response_proto_rawDescOnce.Do(func() {
		file_message_response_get_group_id_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_response_get_group_id_response_proto_rawDescData)
	})
	return file_message_response_get_group_id_response_proto_rawDescData
}

var file_message_response_get_group_id_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_response_get_group_id_response_proto_goTypes = []interface{}{
	(*GetGroupIdResponse)(nil),         // 0: response.GetGroupIdResponse
	(*GetGroupIdResponse_Success)(nil), // 1: response.GetGroupIdResponse.Success
	(*common.ErrorMessage)(nil),        // 2: common.ErrorMessage
}
var file_message_response_get_group_id_response_proto_depIdxs = []int32{
	1, // 0: response.GetGroupIdResponse.success:type_name -> response.GetGroupIdResponse.Success
	2, // 1: response.GetGroupIdResponse.error:type_name -> common.ErrorMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_response_get_group_id_response_proto_init() }
func file_message_response_get_group_id_response_proto_init() {
	if File_message_response_get_group_id_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_response_get_group_id_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupIdResponse); i {
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
		file_message_response_get_group_id_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupIdResponse_Success); i {
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
	file_message_response_get_group_id_response_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetGroupIdResponse_Success_)(nil),
		(*GetGroupIdResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_response_get_group_id_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_response_get_group_id_response_proto_goTypes,
		DependencyIndexes: file_message_response_get_group_id_response_proto_depIdxs,
		MessageInfos:      file_message_response_get_group_id_response_proto_msgTypes,
	}.Build()
	File_message_response_get_group_id_response_proto = out.File
	file_message_response_get_group_id_response_proto_rawDesc = nil
	file_message_response_get_group_id_response_proto_goTypes = nil
	file_message_response_get_group_id_response_proto_depIdxs = nil
}
