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
// source: google/datastore/v1/aggregation_result.proto

package datastorepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The result of a single bucket from a Datastore aggregation query.
//
// The keys of `aggregate_properties` are the same for all results in an
// aggregation query, unlike entity queries which can have different fields
// present for each result.
type AggregationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result of the aggregation functions, ex: `COUNT(*) AS total_entities`.
	//
	// The key is the
	// [alias][google.datastore.v1.AggregationQuery.Aggregation.alias] assigned to
	// the aggregation function on input and the size of this map equals the
	// number of aggregation functions in the query.
	AggregateProperties map[string]*Value `protobuf:"bytes,2,rep,name=aggregate_properties,json=aggregateProperties,proto3" json:"aggregate_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AggregationResult) Reset() {
	*x = AggregationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1_aggregation_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationResult) ProtoMessage() {}

func (x *AggregationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1_aggregation_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationResult.ProtoReflect.Descriptor instead.
func (*AggregationResult) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1_aggregation_result_proto_rawDescGZIP(), []int{0}
}

func (x *AggregationResult) GetAggregateProperties() map[string]*Value {
	if x != nil {
		return x.AggregateProperties
	}
	return nil
}

// A batch of aggregation results produced by an aggregation query.
type AggregationResultBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The aggregation results for this batch.
	AggregationResults []*AggregationResult `protobuf:"bytes,1,rep,name=aggregation_results,json=aggregationResults,proto3" json:"aggregation_results,omitempty"`
	// The state of the query after the current batch.
	// Only COUNT(*) aggregations are supported in the initial launch. Therefore,
	// expected result type is limited to `NO_MORE_RESULTS`.
	MoreResults QueryResultBatch_MoreResultsType `protobuf:"varint,2,opt,name=more_results,json=moreResults,proto3,enum=google.datastore.v1.QueryResultBatch_MoreResultsType" json:"more_results,omitempty"`
	// Read timestamp this batch was returned from.
	//
	// In a single transaction, subsequent query result batches for the same query
	// can have a greater timestamp. Each batch's read timestamp
	// is valid for all preceding batches.
	ReadTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=read_time,json=readTime,proto3" json:"read_time,omitempty"`
}

func (x *AggregationResultBatch) Reset() {
	*x = AggregationResultBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_datastore_v1_aggregation_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationResultBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationResultBatch) ProtoMessage() {}

func (x *AggregationResultBatch) ProtoReflect() protoreflect.Message {
	mi := &file_google_datastore_v1_aggregation_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationResultBatch.ProtoReflect.Descriptor instead.
func (*AggregationResultBatch) Descriptor() ([]byte, []int) {
	return file_google_datastore_v1_aggregation_result_proto_rawDescGZIP(), []int{1}
}

func (x *AggregationResultBatch) GetAggregationResults() []*AggregationResult {
	if x != nil {
		return x.AggregationResults
	}
	return nil
}

func (x *AggregationResultBatch) GetMoreResults() QueryResultBatch_MoreResultsType {
	if x != nil {
		return x.MoreResults
	}
	return QueryResultBatch_MORE_RESULTS_TYPE_UNSPECIFIED
}

func (x *AggregationResultBatch) GetReadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReadTime
	}
	return nil
}

var File_google_datastore_v1_aggregation_result_proto protoreflect.FileDescriptor

var file_google_datastore_v1_aggregation_result_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x72, 0x0a,
	0x14, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x1a, 0x62, 0x0a, 0x18, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84, 0x02, 0x0a, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x57, 0x0a, 0x13, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x12, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x58, 0x0a, 0x0c, 0x6d, 0x6f, 0x72,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x4d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0xc7, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x62, 0xaa,
	0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_datastore_v1_aggregation_result_proto_rawDescOnce sync.Once
	file_google_datastore_v1_aggregation_result_proto_rawDescData = file_google_datastore_v1_aggregation_result_proto_rawDesc
)

func file_google_datastore_v1_aggregation_result_proto_rawDescGZIP() []byte {
	file_google_datastore_v1_aggregation_result_proto_rawDescOnce.Do(func() {
		file_google_datastore_v1_aggregation_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_datastore_v1_aggregation_result_proto_rawDescData)
	})
	return file_google_datastore_v1_aggregation_result_proto_rawDescData
}

var file_google_datastore_v1_aggregation_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_datastore_v1_aggregation_result_proto_goTypes = []any{
	(*AggregationResult)(nil),             // 0: google.datastore.v1.AggregationResult
	(*AggregationResultBatch)(nil),        // 1: google.datastore.v1.AggregationResultBatch
	nil,                                   // 2: google.datastore.v1.AggregationResult.AggregatePropertiesEntry
	(QueryResultBatch_MoreResultsType)(0), // 3: google.datastore.v1.QueryResultBatch.MoreResultsType
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
	(*Value)(nil),                         // 5: google.datastore.v1.Value
}
var file_google_datastore_v1_aggregation_result_proto_depIdxs = []int32{
	2, // 0: google.datastore.v1.AggregationResult.aggregate_properties:type_name -> google.datastore.v1.AggregationResult.AggregatePropertiesEntry
	0, // 1: google.datastore.v1.AggregationResultBatch.aggregation_results:type_name -> google.datastore.v1.AggregationResult
	3, // 2: google.datastore.v1.AggregationResultBatch.more_results:type_name -> google.datastore.v1.QueryResultBatch.MoreResultsType
	4, // 3: google.datastore.v1.AggregationResultBatch.read_time:type_name -> google.protobuf.Timestamp
	5, // 4: google.datastore.v1.AggregationResult.AggregatePropertiesEntry.value:type_name -> google.datastore.v1.Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_datastore_v1_aggregation_result_proto_init() }
func file_google_datastore_v1_aggregation_result_proto_init() {
	if File_google_datastore_v1_aggregation_result_proto != nil {
		return
	}
	file_google_datastore_v1_entity_proto_init()
	file_google_datastore_v1_query_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_datastore_v1_aggregation_result_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AggregationResult); i {
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
		file_google_datastore_v1_aggregation_result_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AggregationResultBatch); i {
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
			RawDescriptor: file_google_datastore_v1_aggregation_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_datastore_v1_aggregation_result_proto_goTypes,
		DependencyIndexes: file_google_datastore_v1_aggregation_result_proto_depIdxs,
		MessageInfos:      file_google_datastore_v1_aggregation_result_proto_msgTypes,
	}.Build()
	File_google_datastore_v1_aggregation_result_proto = out.File
	file_google_datastore_v1_aggregation_result_proto_rawDesc = nil
	file_google_datastore_v1_aggregation_result_proto_goTypes = nil
	file_google_datastore_v1_aggregation_result_proto_depIdxs = nil
}
