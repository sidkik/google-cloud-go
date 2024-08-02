// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/chat/v1/action_status.proto

package chatpb

import (
	reflect "reflect"
	sync "sync"

	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents the status for a request to either invoke or submit a
// [dialog](https://developers.google.com/workspace/chat/dialogs).
type ActionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code.
	StatusCode code.Code `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=google.rpc.Code" json:"status_code,omitempty"`
	// The message to send users about the status of their request.
	// If unset, a generic message based on the `status_code` is sent.
	UserFacingMessage string `protobuf:"bytes,2,opt,name=user_facing_message,json=userFacingMessage,proto3" json:"user_facing_message,omitempty"`
}

func (x *ActionStatus) Reset() {
	*x = ActionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_chat_v1_action_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionStatus) ProtoMessage() {}

func (x *ActionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_action_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionStatus.ProtoReflect.Descriptor instead.
func (*ActionStatus) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_action_status_proto_rawDescGZIP(), []int{0}
}

func (x *ActionStatus) GetStatusCode() code.Code {
	if x != nil {
		return x.StatusCode
	}
	return code.Code(0)
}

func (x *ActionStatus) GetUserFacingMessage() string {
	if x != nil {
		return x.UserFacingMessage
	}
	return ""
}

var File_google_chat_v1_action_status_proto protoreflect.FileDescriptor

var file_google_chat_v1_action_status_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0c, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e,
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x73, 0x65,
	0x72, 0x46, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0xaa,
	0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x0b, 0x44, 0x59, 0x4e, 0x41, 0x50,
	0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70,
	0x73, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_chat_v1_action_status_proto_rawDescOnce sync.Once
	file_google_chat_v1_action_status_proto_rawDescData = file_google_chat_v1_action_status_proto_rawDesc
)

func file_google_chat_v1_action_status_proto_rawDescGZIP() []byte {
	file_google_chat_v1_action_status_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_action_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_action_status_proto_rawDescData)
	})
	return file_google_chat_v1_action_status_proto_rawDescData
}

var file_google_chat_v1_action_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_chat_v1_action_status_proto_goTypes = []any{
	(*ActionStatus)(nil), // 0: google.chat.v1.ActionStatus
	(code.Code)(0),       // 1: google.rpc.Code
}
var file_google_chat_v1_action_status_proto_depIdxs = []int32{
	1, // 0: google.chat.v1.ActionStatus.status_code:type_name -> google.rpc.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_chat_v1_action_status_proto_init() }
func file_google_chat_v1_action_status_proto_init() {
	if File_google_chat_v1_action_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_chat_v1_action_status_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActionStatus); i {
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
			RawDescriptor: file_google_chat_v1_action_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_action_status_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_action_status_proto_depIdxs,
		MessageInfos:      file_google_chat_v1_action_status_proto_msgTypes,
	}.Build()
	File_google_chat_v1_action_status_proto = out.File
	file_google_chat_v1_action_status_proto_rawDesc = nil
	file_google_chat_v1_action_status_proto_goTypes = nil
	file_google_chat_v1_action_status_proto_depIdxs = nil
}
