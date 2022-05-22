// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/librarium.proto

package api

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

// The request message containing Autor or a Book name.
type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AskMessage string `protobuf:"bytes,1,opt,name=askMessage,proto3" json:"askMessage,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_librarium_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_librarium_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_proto_librarium_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetAskMessage() string {
	if x != nil {
		return x.AskMessage
	}
	return ""
}

// The response message containing the greetings
type DataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplyMessage []string `protobuf:"bytes,1,rep,name=replyMessage,proto3" json:"replyMessage,omitempty"`
}

func (x *DataReply) Reset() {
	*x = DataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_librarium_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataReply) ProtoMessage() {}

func (x *DataReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_librarium_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataReply.ProtoReflect.Descriptor instead.
func (*DataReply) Descriptor() ([]byte, []int) {
	return file_proto_librarium_proto_rawDescGZIP(), []int{1}
}

func (x *DataReply) GetReplyMessage() []string {
	if x != nil {
		return x.ReplyMessage
	}
	return nil
}

var File_proto_librarium_proto protoreflect.FileDescriptor

var file_proto_librarium_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69,
	0x75, 0x6d, 0x22, 0x2d, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x2f, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x83, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d,
	0x12, 0x3a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_librarium_proto_rawDescOnce sync.Once
	file_proto_librarium_proto_rawDescData = file_proto_librarium_proto_rawDesc
)

func file_proto_librarium_proto_rawDescGZIP() []byte {
	file_proto_librarium_proto_rawDescOnce.Do(func() {
		file_proto_librarium_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_librarium_proto_rawDescData)
	})
	return file_proto_librarium_proto_rawDescData
}

var file_proto_librarium_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_librarium_proto_goTypes = []interface{}{
	(*DataRequest)(nil), // 0: librarium.DataRequest
	(*DataReply)(nil),   // 1: librarium.DataReply
}
var file_proto_librarium_proto_depIdxs = []int32{
	0, // 0: librarium.Librarium.GetAutor:input_type -> librarium.DataRequest
	0, // 1: librarium.Librarium.GetBooks:input_type -> librarium.DataRequest
	1, // 2: librarium.Librarium.GetAutor:output_type -> librarium.DataReply
	1, // 3: librarium.Librarium.GetBooks:output_type -> librarium.DataReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_librarium_proto_init() }
func file_proto_librarium_proto_init() {
	if File_proto_librarium_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_librarium_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_proto_librarium_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataReply); i {
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
			RawDescriptor: file_proto_librarium_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_librarium_proto_goTypes,
		DependencyIndexes: file_proto_librarium_proto_depIdxs,
		MessageInfos:      file_proto_librarium_proto_msgTypes,
	}.Build()
	File_proto_librarium_proto = out.File
	file_proto_librarium_proto_rawDesc = nil
	file_proto_librarium_proto_goTypes = nil
	file_proto_librarium_proto_depIdxs = nil
}
