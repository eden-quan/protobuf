// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Messages in this file are used to test wire encoding order.

// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// source: internal/testprotos/order/order.proto

package order

import (
	protoreflect "gitlab.lainuoniao.cn/eden-quan/protobuf/reflect/protoreflect"
	protoimpl "gitlab.lainuoniao.cn/eden-quan/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Message struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields

	Field_2 *string `protobuf:"bytes,2,opt,name=field_2,json=field2" json:"field_2,omitempty"`
	Field_1 *string `protobuf:"bytes,1,opt,name=field_1,json=field1" json:"field_1,omitempty"`
	// Types that are assignable to Oneof_1:
	//
	//	*Message_Field_10
	Oneof_1  isMessage_Oneof_1 `protobuf_oneof:"oneof_1"`
	Field_20 *string           `protobuf:"bytes,20,opt,name=field_20,json=field20" json:"field_20,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetField_2() string {
	if x != nil && x.Field_2 != nil {
		return *x.Field_2
	}
	return ""
}

func (x *Message) GetField_1() string {
	if x != nil && x.Field_1 != nil {
		return *x.Field_1
	}
	return ""
}

func (m *Message) GetOneof_1() isMessage_Oneof_1 {
	if m != nil {
		return m.Oneof_1
	}
	return nil
}

func (x *Message) GetField_10() string {
	if x, ok := x.GetOneof_1().(*Message_Field_10); ok {
		return x.Field_10
	}
	return ""
}

func (x *Message) GetField_20() string {
	if x != nil && x.Field_20 != nil {
		return *x.Field_20
	}
	return ""
}

type isMessage_Oneof_1 interface {
	isMessage_Oneof_1()
}

type Message_Field_10 struct {
	Field_10 string `protobuf:"bytes,10,opt,name=field_10,json=field10,oneof"`
}

func (*Message_Field_10) isMessage_Oneof_1() {}

var file_internal_testprotos_order_order_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         30,
		Name:          "goproto.proto.order.field_30",
		Tag:           "bytes,30,opt,name=field_30",
		Filename:      "internal/testprotos/order/order.proto",
	},
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         31,
		Name:          "goproto.proto.order.field_31",
		Tag:           "bytes,31,opt,name=field_31",
		Filename:      "internal/testprotos/order/order.proto",
	},
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         32,
		Name:          "goproto.proto.order.field_32",
		Tag:           "bytes,32,opt,name=field_32",
		Filename:      "internal/testprotos/order/order.proto",
	},
}

// Extension fields to Message.
var (
	// optional string field_30 = 30;
	E_Field_30 = &file_internal_testprotos_order_order_proto_extTypes[0]
	// optional string field_31 = 31;
	E_Field_31 = &file_internal_testprotos_order_order_proto_extTypes[1]
	// optional string field_32 = 32;
	E_Field_32 = &file_internal_testprotos_order_order_proto_extTypes[2]
)

var File_internal_testprotos_order_order_proto protoreflect.FileDescriptor

var file_internal_testprotos_order_order_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x30, 0x2a, 0x04, 0x08, 0x1e, 0x10, 0x29, 0x42, 0x09, 0x0a, 0x07, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x31, 0x3a, 0x37, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33, 0x30, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x3a, 0x37, 0x0a, 0x08,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33, 0x31, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x33, 0x31, 0x3a, 0x37, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x33,
	0x32, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x32, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
}

var (
	file_internal_testprotos_order_order_proto_rawDescOnce sync.Once
	file_internal_testprotos_order_order_proto_rawDescData = file_internal_testprotos_order_order_proto_rawDesc
)

func file_internal_testprotos_order_order_proto_rawDescGZIP() []byte {
	file_internal_testprotos_order_order_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_order_order_proto_rawDescData)
	})
	return file_internal_testprotos_order_order_proto_rawDescData
}

var file_internal_testprotos_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_order_order_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: goproto.proto.order.Message
}
var file_internal_testprotos_order_order_proto_depIdxs = []int32{
	0, // 0: goproto.proto.order.field_30:extendee -> goproto.proto.order.Message
	0, // 1: goproto.proto.order.field_31:extendee -> goproto.proto.order.Message
	0, // 2: goproto.proto.order.field_32:extendee -> goproto.proto.order.Message
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_order_order_proto_init() }
func file_internal_testprotos_order_order_proto_init() {
	if File_internal_testprotos_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
	}
	file_internal_testprotos_order_order_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Field_10)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_order_order_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_order_order_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_order_order_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_order_order_proto_extTypes,
	}.Build()
	File_internal_testprotos_order_order_proto = out.File
	file_internal_testprotos_order_order_proto_rawDesc = nil
	file_internal_testprotos_order_order_proto_goTypes = nil
	file_internal_testprotos_order_order_proto_depIdxs = nil
}
