// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: personInfo.proto

package apis

import (
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

type PersonalInformationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PersonalInformation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PersonalInformationList) Reset() {
	*x = PersonalInformationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformationList) ProtoMessage() {}

func (x *PersonalInformationList) ProtoReflect() protoreflect.Message {
	mi := &file_personInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformationList.ProtoReflect.Descriptor instead.
func (*PersonalInformationList) Descriptor() ([]byte, []int) {
	return file_personInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalInformationList) GetItems() []*PersonalInformation {
	if x != nil {
		return x.Items
	}
	return nil
}

type PersonalInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Account  int64  `protobuf:"varint,4,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *PersonalInformation) Reset() {
	*x = PersonalInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformation) ProtoMessage() {}

func (x *PersonalInformation) ProtoReflect() protoreflect.Message {
	mi := &file_personInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformation.ProtoReflect.Descriptor instead.
func (*PersonalInformation) Descriptor() ([]byte, []int) {
	return file_personInfo_proto_rawDescGZIP(), []int{1}
}

func (x *PersonalInformation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonalInformation) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PersonalInformation) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PersonalInformation) GetAccount() int64 {
	if x != nil {
		return x.Account
	}
	return 0
}

var File_personInfo_proto protoreflect.FileDescriptor

var file_personInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x70, 0x69, 0x73, 0x22, 0x4a, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x77, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x51, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_personInfo_proto_rawDescOnce sync.Once
	file_personInfo_proto_rawDescData = file_personInfo_proto_rawDesc
)

func file_personInfo_proto_rawDescGZIP() []byte {
	file_personInfo_proto_rawDescOnce.Do(func() {
		file_personInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_personInfo_proto_rawDescData)
	})
	return file_personInfo_proto_rawDescData
}

var file_personInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_personInfo_proto_goTypes = []interface{}{
	(*PersonalInformationList)(nil), // 0: apis.PersonalInformationList
	(*PersonalInformation)(nil),     // 1: apis.PersonalInformation
}
var file_personInfo_proto_depIdxs = []int32{
	1, // 0: apis.PersonalInformationList.items:type_name -> apis.PersonalInformation
	1, // 1: apis.ChatService.Register:input_type -> apis.PersonalInformation
	1, // 2: apis.ChatService.Register:output_type -> apis.PersonalInformation
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_personInfo_proto_init() }
func file_personInfo_proto_init() {
	if File_personInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_personInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformationList); i {
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
		file_personInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformation); i {
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
			RawDescriptor: file_personInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personInfo_proto_goTypes,
		DependencyIndexes: file_personInfo_proto_depIdxs,
		MessageInfos:      file_personInfo_proto_msgTypes,
	}.Build()
	File_personInfo_proto = out.File
	file_personInfo_proto_rawDesc = nil
	file_personInfo_proto_goTypes = nil
	file_personInfo_proto_depIdxs = nil
}
