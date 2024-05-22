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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.3
// source: google/firestore/admin/v1/index.proto

package adminpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Query Scope defines the scope at which a query is run. This is specified on
// a StructuredQuery's `from` field.
type Index_QueryScope int32

const (
	// The query scope is unspecified. Not a valid option.
	Index_QUERY_SCOPE_UNSPECIFIED Index_QueryScope = 0
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified
	// at query time, and that has the collection id specified by the index.
	Index_COLLECTION Index_QueryScope = 1
	// Indexes with a collection group query scope specified allow queries
	// against all collections that has the collection id specified by the
	// index.
	Index_COLLECTION_GROUP Index_QueryScope = 2
	// Include all the collections's ancestor in the index. Only available for
	// Datastore Mode databases.
	Index_COLLECTION_RECURSIVE Index_QueryScope = 3
)

// Enum value maps for Index_QueryScope.
var (
	Index_QueryScope_name = map[int32]string{
		0: "QUERY_SCOPE_UNSPECIFIED",
		1: "COLLECTION",
		2: "COLLECTION_GROUP",
		3: "COLLECTION_RECURSIVE",
	}
	Index_QueryScope_value = map[string]int32{
		"QUERY_SCOPE_UNSPECIFIED": 0,
		"COLLECTION":              1,
		"COLLECTION_GROUP":        2,
		"COLLECTION_RECURSIVE":    3,
	}
)

func (x Index_QueryScope) Enum() *Index_QueryScope {
	p := new(Index_QueryScope)
	*p = x
	return p
}

func (x Index_QueryScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_QueryScope) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_index_proto_enumTypes[0].Descriptor()
}

func (Index_QueryScope) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_index_proto_enumTypes[0]
}

func (x Index_QueryScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_QueryScope.Descriptor instead.
func (Index_QueryScope) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0}
}

// API Scope defines the APIs (Firestore Native, or Firestore in
// Datastore Mode) that are supported for queries.
type Index_ApiScope int32

const (
	// The index can only be used by the Firestore Native query API.
	// This is the default.
	Index_ANY_API Index_ApiScope = 0
	// The index can only be used by the Firestore in Datastore Mode query API.
	Index_DATASTORE_MODE_API Index_ApiScope = 1
)

// Enum value maps for Index_ApiScope.
var (
	Index_ApiScope_name = map[int32]string{
		0: "ANY_API",
		1: "DATASTORE_MODE_API",
	}
	Index_ApiScope_value = map[string]int32{
		"ANY_API":            0,
		"DATASTORE_MODE_API": 1,
	}
)

func (x Index_ApiScope) Enum() *Index_ApiScope {
	p := new(Index_ApiScope)
	*p = x
	return p
}

func (x Index_ApiScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_ApiScope) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_index_proto_enumTypes[1].Descriptor()
}

func (Index_ApiScope) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_index_proto_enumTypes[1]
}

func (x Index_ApiScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_ApiScope.Descriptor instead.
func (Index_ApiScope) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 1}
}

// The state of an index. During index creation, an index will be in the
// `CREATING` state. If the index is created successfully, it will transition
// to the `READY` state. If the index creation encounters a problem, the index
// will transition to the `NEEDS_REPAIR` state.
type Index_State int32

const (
	// The state is unspecified.
	Index_STATE_UNSPECIFIED Index_State = 0
	// The index is being created.
	// There is an active long-running operation for the index.
	// The index is updated when writing a document.
	// Some index data may exist.
	Index_CREATING Index_State = 1
	// The index is ready to be used.
	// The index is updated when writing a document.
	// The index is fully populated from all stored documents it applies to.
	Index_READY Index_State = 2
	// The index was being created, but something went wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing a document.
	// Some index data may exist.
	// Use the google.longrunning.Operations API to determine why the operation
	// that last attempted to create this index failed, then re-create the
	// index.
	Index_NEEDS_REPAIR Index_State = 3
)

// Enum value maps for Index_State.
var (
	Index_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "NEEDS_REPAIR",
	}
	Index_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"NEEDS_REPAIR":      3,
	}
)

func (x Index_State) Enum() *Index_State {
	p := new(Index_State)
	*p = x
	return p
}

func (x Index_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_index_proto_enumTypes[2].Descriptor()
}

func (Index_State) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_index_proto_enumTypes[2]
}

func (x Index_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_State.Descriptor instead.
func (Index_State) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 2}
}

// The supported orderings.
type Index_IndexField_Order int32

const (
	// The ordering is unspecified. Not a valid option.
	Index_IndexField_ORDER_UNSPECIFIED Index_IndexField_Order = 0
	// The field is ordered by ascending field value.
	Index_IndexField_ASCENDING Index_IndexField_Order = 1
	// The field is ordered by descending field value.
	Index_IndexField_DESCENDING Index_IndexField_Order = 2
)

// Enum value maps for Index_IndexField_Order.
var (
	Index_IndexField_Order_name = map[int32]string{
		0: "ORDER_UNSPECIFIED",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	Index_IndexField_Order_value = map[string]int32{
		"ORDER_UNSPECIFIED": 0,
		"ASCENDING":         1,
		"DESCENDING":        2,
	}
)

func (x Index_IndexField_Order) Enum() *Index_IndexField_Order {
	p := new(Index_IndexField_Order)
	*p = x
	return p
}

func (x Index_IndexField_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_IndexField_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_index_proto_enumTypes[3].Descriptor()
}

func (Index_IndexField_Order) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_index_proto_enumTypes[3]
}

func (x Index_IndexField_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_IndexField_Order.Descriptor instead.
func (Index_IndexField_Order) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0, 0}
}

// The supported array value configurations.
type Index_IndexField_ArrayConfig int32

const (
	// The index does not support additional array queries.
	Index_IndexField_ARRAY_CONFIG_UNSPECIFIED Index_IndexField_ArrayConfig = 0
	// The index supports array containment queries.
	Index_IndexField_CONTAINS Index_IndexField_ArrayConfig = 1
)

// Enum value maps for Index_IndexField_ArrayConfig.
var (
	Index_IndexField_ArrayConfig_name = map[int32]string{
		0: "ARRAY_CONFIG_UNSPECIFIED",
		1: "CONTAINS",
	}
	Index_IndexField_ArrayConfig_value = map[string]int32{
		"ARRAY_CONFIG_UNSPECIFIED": 0,
		"CONTAINS":                 1,
	}
)

func (x Index_IndexField_ArrayConfig) Enum() *Index_IndexField_ArrayConfig {
	p := new(Index_IndexField_ArrayConfig)
	*p = x
	return p
}

func (x Index_IndexField_ArrayConfig) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Index_IndexField_ArrayConfig) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_index_proto_enumTypes[4].Descriptor()
}

func (Index_IndexField_ArrayConfig) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_index_proto_enumTypes[4]
}

func (x Index_IndexField_ArrayConfig) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Index_IndexField_ArrayConfig.Descriptor instead.
func (Index_IndexField_ArrayConfig) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0, 1}
}

// Cloud Firestore indexes enable simple and complex queries against
// documents in a database.
type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. A server defined name for this index.
	// The form of this name for composite indexes will be:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}`
	// For single field indexes, this field will be empty.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indexes with a collection query scope specified allow queries
	// against a collection that is the child of a specific document, specified at
	// query time, and that has the same collection id.
	//
	// Indexes with a collection group query scope specified allow queries against
	// all collections descended from a specific document, specified at query
	// time, and that have the same collection id as this index.
	QueryScope Index_QueryScope `protobuf:"varint,2,opt,name=query_scope,json=queryScope,proto3,enum=google.firestore.admin.v1.Index_QueryScope" json:"query_scope,omitempty"`
	// The API scope supported by this index.
	ApiScope Index_ApiScope `protobuf:"varint,5,opt,name=api_scope,json=apiScope,proto3,enum=google.firestore.admin.v1.Index_ApiScope" json:"api_scope,omitempty"`
	// The fields supported by this index.
	//
	// For composite indexes, this requires a minimum of 2 and a maximum of 100
	// fields. The last field entry is always for the field path `__name__`. If,
	// on creation, `__name__` was not specified as the last field, it will be
	// added automatically with the same direction as that of the last field
	// defined. If the final field in a composite index is not directional, the
	// `__name__` will be ordered ASCENDING (unless explicitly specified).
	//
	// For single field indexes, this will always be exactly one entry with a
	// field path equal to the field path of the associated field.
	Fields []*Index_IndexField `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// Output only. The serving state of the index.
	State Index_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.firestore.admin.v1.Index_State" json:"state,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0}
}

func (x *Index) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Index) GetQueryScope() Index_QueryScope {
	if x != nil {
		return x.QueryScope
	}
	return Index_QUERY_SCOPE_UNSPECIFIED
}

func (x *Index) GetApiScope() Index_ApiScope {
	if x != nil {
		return x.ApiScope
	}
	return Index_ANY_API
}

func (x *Index) GetFields() []*Index_IndexField {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Index) GetState() Index_State {
	if x != nil {
		return x.State
	}
	return Index_STATE_UNSPECIFIED
}

// A field in an index.
// The field_path describes which field is indexed, the value_mode describes
// how the field value is indexed.
type Index_IndexField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Can be __name__.
	// For single field indexes, this must match the name of the field or may
	// be omitted.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// How the field value is indexed.
	//
	// Types that are assignable to ValueMode:
	//
	//	*Index_IndexField_Order_
	//	*Index_IndexField_ArrayConfig_
	//	*Index_IndexField_VectorConfig_
	ValueMode isIndex_IndexField_ValueMode `protobuf_oneof:"value_mode"`
}

func (x *Index_IndexField) Reset() {
	*x = Index_IndexField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index_IndexField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index_IndexField) ProtoMessage() {}

func (x *Index_IndexField) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index_IndexField.ProtoReflect.Descriptor instead.
func (*Index_IndexField) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Index_IndexField) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (m *Index_IndexField) GetValueMode() isIndex_IndexField_ValueMode {
	if m != nil {
		return m.ValueMode
	}
	return nil
}

func (x *Index_IndexField) GetOrder() Index_IndexField_Order {
	if x, ok := x.GetValueMode().(*Index_IndexField_Order_); ok {
		return x.Order
	}
	return Index_IndexField_ORDER_UNSPECIFIED
}

func (x *Index_IndexField) GetArrayConfig() Index_IndexField_ArrayConfig {
	if x, ok := x.GetValueMode().(*Index_IndexField_ArrayConfig_); ok {
		return x.ArrayConfig
	}
	return Index_IndexField_ARRAY_CONFIG_UNSPECIFIED
}

func (x *Index_IndexField) GetVectorConfig() *Index_IndexField_VectorConfig {
	if x, ok := x.GetValueMode().(*Index_IndexField_VectorConfig_); ok {
		return x.VectorConfig
	}
	return nil
}

type isIndex_IndexField_ValueMode interface {
	isIndex_IndexField_ValueMode()
}

type Index_IndexField_Order_ struct {
	// Indicates that this field supports ordering by the specified order or
	// comparing using =, !=, <, <=, >, >=.
	Order Index_IndexField_Order `protobuf:"varint,2,opt,name=order,proto3,enum=google.firestore.admin.v1.Index_IndexField_Order,oneof"`
}

type Index_IndexField_ArrayConfig_ struct {
	// Indicates that this field supports operations on `array_value`s.
	ArrayConfig Index_IndexField_ArrayConfig `protobuf:"varint,3,opt,name=array_config,json=arrayConfig,proto3,enum=google.firestore.admin.v1.Index_IndexField_ArrayConfig,oneof"`
}

type Index_IndexField_VectorConfig_ struct {
	// Indicates that this field supports nearest neighbors and distance
	// operations on vector.
	VectorConfig *Index_IndexField_VectorConfig `protobuf:"bytes,4,opt,name=vector_config,json=vectorConfig,proto3,oneof"`
}

func (*Index_IndexField_Order_) isIndex_IndexField_ValueMode() {}

func (*Index_IndexField_ArrayConfig_) isIndex_IndexField_ValueMode() {}

func (*Index_IndexField_VectorConfig_) isIndex_IndexField_ValueMode() {}

// The index configuration to support vector search operations
type Index_IndexField_VectorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The vector dimension this configuration applies to.
	//
	// The resulting index will only include vectors of this dimension, and
	// can be used for vector search with the same dimension.
	Dimension int32 `protobuf:"varint,1,opt,name=dimension,proto3" json:"dimension,omitempty"`
	// The type of index used.
	//
	// Types that are assignable to Type:
	//
	//	*Index_IndexField_VectorConfig_Flat
	Type isIndex_IndexField_VectorConfig_Type `protobuf_oneof:"type"`
}

func (x *Index_IndexField_VectorConfig) Reset() {
	*x = Index_IndexField_VectorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index_IndexField_VectorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index_IndexField_VectorConfig) ProtoMessage() {}

func (x *Index_IndexField_VectorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index_IndexField_VectorConfig.ProtoReflect.Descriptor instead.
func (*Index_IndexField_VectorConfig) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Index_IndexField_VectorConfig) GetDimension() int32 {
	if x != nil {
		return x.Dimension
	}
	return 0
}

func (m *Index_IndexField_VectorConfig) GetType() isIndex_IndexField_VectorConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Index_IndexField_VectorConfig) GetFlat() *Index_IndexField_VectorConfig_FlatIndex {
	if x, ok := x.GetType().(*Index_IndexField_VectorConfig_Flat); ok {
		return x.Flat
	}
	return nil
}

type isIndex_IndexField_VectorConfig_Type interface {
	isIndex_IndexField_VectorConfig_Type()
}

type Index_IndexField_VectorConfig_Flat struct {
	// Indicates the vector index is a flat index.
	Flat *Index_IndexField_VectorConfig_FlatIndex `protobuf:"bytes,2,opt,name=flat,proto3,oneof"`
}

func (*Index_IndexField_VectorConfig_Flat) isIndex_IndexField_VectorConfig_Type() {}

// An index that stores vectors in a flat data structure, and supports
// exhaustive search.
type Index_IndexField_VectorConfig_FlatIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Index_IndexField_VectorConfig_FlatIndex) Reset() {
	*x = Index_IndexField_VectorConfig_FlatIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index_IndexField_VectorConfig_FlatIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index_IndexField_VectorConfig_FlatIndex) ProtoMessage() {}

func (x *Index_IndexField_VectorConfig_FlatIndex) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index_IndexField_VectorConfig_FlatIndex.ProtoReflect.Descriptor instead.
func (*Index_IndexField_VectorConfig_FlatIndex) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_index_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

var File_google_firestore_admin_v1_index_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1_index_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa,
	0x09, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0a,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x41, 0x70, 0x69, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x61, 0x70, 0x69, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0xe0, 0x04, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x49, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x5c,
	0x0a, 0x0c, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5f, 0x0a, 0x0d,
	0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x0c, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xa0, 0x01,
	0x0a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21,
	0x0a, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x58, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x6c, 0x61, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x74, 0x1a, 0x0b, 0x0a, 0x09, 0x46,
	0x6c, 0x61, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x3d, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x22,
	0x39, 0x0a, 0x0b, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c,
	0x0a, 0x18, 0x41, 0x52, 0x52, 0x41, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x69, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4c,
	0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x55, 0x52, 0x53, 0x49, 0x56,
	0x45, 0x10, 0x03, 0x22, 0x2f, 0x0a, 0x08, 0x41, 0x70, 0x69, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x59, 0x5f, 0x41, 0x50, 0x49, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x44, 0x41, 0x54, 0x41, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x41,
	0x50, 0x49, 0x10, 0x01, 0x22, 0x49, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x10, 0x0a,
	0x0c, 0x4e, 0x45, 0x45, 0x44, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x41, 0x49, 0x52, 0x10, 0x03, 0x3a,
	0x7a, 0xea, 0x41, 0x77, 0x0a, 0x1e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x55, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x7d, 0x42, 0xd9, 0x01, 0x0a, 0x1d,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53, 0xaa, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x46, 0x69, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x46,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1_index_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1_index_proto_rawDescData = file_google_firestore_admin_v1_index_proto_rawDesc
)

func file_google_firestore_admin_v1_index_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1_index_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1_index_proto_rawDescData)
	})
	return file_google_firestore_admin_v1_index_proto_rawDescData
}

var file_google_firestore_admin_v1_index_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_google_firestore_admin_v1_index_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_firestore_admin_v1_index_proto_goTypes = []interface{}{
	(Index_QueryScope)(0),                           // 0: google.firestore.admin.v1.Index.QueryScope
	(Index_ApiScope)(0),                             // 1: google.firestore.admin.v1.Index.ApiScope
	(Index_State)(0),                                // 2: google.firestore.admin.v1.Index.State
	(Index_IndexField_Order)(0),                     // 3: google.firestore.admin.v1.Index.IndexField.Order
	(Index_IndexField_ArrayConfig)(0),               // 4: google.firestore.admin.v1.Index.IndexField.ArrayConfig
	(*Index)(nil),                                   // 5: google.firestore.admin.v1.Index
	(*Index_IndexField)(nil),                        // 6: google.firestore.admin.v1.Index.IndexField
	(*Index_IndexField_VectorConfig)(nil),           // 7: google.firestore.admin.v1.Index.IndexField.VectorConfig
	(*Index_IndexField_VectorConfig_FlatIndex)(nil), // 8: google.firestore.admin.v1.Index.IndexField.VectorConfig.FlatIndex
}
var file_google_firestore_admin_v1_index_proto_depIdxs = []int32{
	0, // 0: google.firestore.admin.v1.Index.query_scope:type_name -> google.firestore.admin.v1.Index.QueryScope
	1, // 1: google.firestore.admin.v1.Index.api_scope:type_name -> google.firestore.admin.v1.Index.ApiScope
	6, // 2: google.firestore.admin.v1.Index.fields:type_name -> google.firestore.admin.v1.Index.IndexField
	2, // 3: google.firestore.admin.v1.Index.state:type_name -> google.firestore.admin.v1.Index.State
	3, // 4: google.firestore.admin.v1.Index.IndexField.order:type_name -> google.firestore.admin.v1.Index.IndexField.Order
	4, // 5: google.firestore.admin.v1.Index.IndexField.array_config:type_name -> google.firestore.admin.v1.Index.IndexField.ArrayConfig
	7, // 6: google.firestore.admin.v1.Index.IndexField.vector_config:type_name -> google.firestore.admin.v1.Index.IndexField.VectorConfig
	8, // 7: google.firestore.admin.v1.Index.IndexField.VectorConfig.flat:type_name -> google.firestore.admin.v1.Index.IndexField.VectorConfig.FlatIndex
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1_index_proto_init() }
func file_google_firestore_admin_v1_index_proto_init() {
	if File_google_firestore_admin_v1_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_admin_v1_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_google_firestore_admin_v1_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index_IndexField); i {
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
		file_google_firestore_admin_v1_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index_IndexField_VectorConfig); i {
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
		file_google_firestore_admin_v1_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index_IndexField_VectorConfig_FlatIndex); i {
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
	file_google_firestore_admin_v1_index_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Index_IndexField_Order_)(nil),
		(*Index_IndexField_ArrayConfig_)(nil),
		(*Index_IndexField_VectorConfig_)(nil),
	}
	file_google_firestore_admin_v1_index_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Index_IndexField_VectorConfig_Flat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_admin_v1_index_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1_index_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1_index_proto_depIdxs,
		EnumInfos:         file_google_firestore_admin_v1_index_proto_enumTypes,
		MessageInfos:      file_google_firestore_admin_v1_index_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1_index_proto = out.File
	file_google_firestore_admin_v1_index_proto_rawDesc = nil
	file_google_firestore_admin_v1_index_proto_goTypes = nil
	file_google_firestore_admin_v1_index_proto_depIdxs = nil
}
