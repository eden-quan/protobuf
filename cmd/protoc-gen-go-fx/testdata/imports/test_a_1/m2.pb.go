// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// source: cmd/protoc-gen-go-fx/testdata/imports/test_a_1/m2.proto

package test_a_1

import (
	protoreflect "github.com/eden-quan/protobuf/reflect/protoreflect"
	protoimpl "github.com/eden-quan/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type M2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *M2) Reset() {
	*x = M2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *M2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2) ProtoMessage() {}

func (x *M2) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use M2.ProtoReflect.Descriptor instead.
func (*M2) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescGZIP(), []int{0}
}

var File_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x61, 0x22, 0x04,
	0x0a, 0x02, 0x4d, 0x32, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescData = file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_goTypes = []interface{}{
	(*M2)(nil), // 0: test.a.M2
}
var file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_init() }
func file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_init() {
	if File_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto = out.File
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_imports_test_a_1_m2_proto_depIdxs = nil
}
