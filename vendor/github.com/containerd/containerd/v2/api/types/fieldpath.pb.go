// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/types/fieldpath.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_github_com_containerd_containerd_api_types_fieldpath_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         63300,
		Name:          "containerd.types.fieldpath_all",
		Tag:           "varint,63300,opt,name=fieldpath_all",
		Filename:      "github.com/containerd/containerd/api/types/fieldpath.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         64400,
		Name:          "containerd.types.fieldpath",
		Tag:           "varint,64400,opt,name=fieldpath",
		Filename:      "github.com/containerd/containerd/api/types/fieldpath.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional bool fieldpath_all = 63300;
	E_FieldpathAll = &file_github_com_containerd_containerd_api_types_fieldpath_proto_extTypes[0]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool fieldpath = 64400;
	E_Fieldpath = &file_github_com_containerd_containerd_api_types_fieldpath_proto_extTypes[1]
)

var File_github_com_containerd_containerd_api_types_fieldpath_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_types_fieldpath_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x46, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x61, 0x6c,
	0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc4, 0xee, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x61,
	0x74, 0x68, 0x41, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x3a, 0x42, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x90, 0xf7, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_github_com_containerd_containerd_api_types_fieldpath_proto_goTypes = []interface{}{
	(*descriptorpb.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_github_com_containerd_containerd_api_types_fieldpath_proto_depIdxs = []int32{
	0, // 0: containerd.types.fieldpath_all:extendee -> google.protobuf.FileOptions
	1, // 1: containerd.types.fieldpath:extendee -> google.protobuf.MessageOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_types_fieldpath_proto_init() }
func file_github_com_containerd_containerd_api_types_fieldpath_proto_init() {
	if File_github_com_containerd_containerd_api_types_fieldpath_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_containerd_containerd_api_types_fieldpath_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_types_fieldpath_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_types_fieldpath_proto_depIdxs,
		ExtensionInfos:    file_github_com_containerd_containerd_api_types_fieldpath_proto_extTypes,
	}.Build()
	File_github_com_containerd_containerd_api_types_fieldpath_proto = out.File
	file_github_com_containerd_containerd_api_types_fieldpath_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_types_fieldpath_proto_goTypes = nil
	file_github_com_containerd_containerd_api_types_fieldpath_proto_depIdxs = nil
}
