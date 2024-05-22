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
// source: google/cloud/video/stitcher/v1/live_configs.proto

package stitcherpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Determines the ad tracking policy.
type AdTracking int32

const (
	// The ad tracking policy is not specified.
	AdTracking_AD_TRACKING_UNSPECIFIED AdTracking = 0
	// Client-side ad tracking is specified. The client player is expected to
	// trigger playback and activity events itself.
	AdTracking_CLIENT AdTracking = 1
	// The Video Stitcher API will trigger playback events on behalf of
	// the client player.
	AdTracking_SERVER AdTracking = 2
)

// Enum value maps for AdTracking.
var (
	AdTracking_name = map[int32]string{
		0: "AD_TRACKING_UNSPECIFIED",
		1: "CLIENT",
		2: "SERVER",
	}
	AdTracking_value = map[string]int32{
		"AD_TRACKING_UNSPECIFIED": 0,
		"CLIENT":                  1,
		"SERVER":                  2,
	}
)

func (x AdTracking) Enum() *AdTracking {
	p := new(AdTracking)
	*p = x
	return p
}

func (x AdTracking) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdTracking) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[0].Descriptor()
}

func (AdTracking) Type() protoreflect.EnumType {
	return &file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[0]
}

func (x AdTracking) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdTracking.Descriptor instead.
func (AdTracking) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{0}
}

// State of the live config.
type LiveConfig_State int32

const (
	// State is not specified.
	LiveConfig_STATE_UNSPECIFIED LiveConfig_State = 0
	// Live config is being created.
	LiveConfig_CREATING LiveConfig_State = 1
	// Live config is ready for use.
	LiveConfig_READY LiveConfig_State = 2
	// Live config is queued up for deletion.
	LiveConfig_DELETING LiveConfig_State = 3
)

// Enum value maps for LiveConfig_State.
var (
	LiveConfig_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
	}
	LiveConfig_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
	}
)

func (x LiveConfig_State) Enum() *LiveConfig_State {
	p := new(LiveConfig_State)
	*p = x
	return p
}

func (x LiveConfig_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LiveConfig_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[1].Descriptor()
}

func (LiveConfig_State) Type() protoreflect.EnumType {
	return &file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[1]
}

func (x LiveConfig_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LiveConfig_State.Descriptor instead.
func (LiveConfig_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{0, 0}
}

// Defines the ad stitching behavior in case the ad duration does not align
// exactly with the ad break boundaries. If not specified, the default is
// `CUT_CURRENT`.
type LiveConfig_StitchingPolicy int32

const (
	// Stitching policy is not specified.
	LiveConfig_STITCHING_POLICY_UNSPECIFIED LiveConfig_StitchingPolicy = 0
	// Cuts an ad short and returns to content in the middle of the ad.
	LiveConfig_CUT_CURRENT LiveConfig_StitchingPolicy = 1
	// Finishes stitching the current ad before returning to content.
	LiveConfig_COMPLETE_AD LiveConfig_StitchingPolicy = 2
)

// Enum value maps for LiveConfig_StitchingPolicy.
var (
	LiveConfig_StitchingPolicy_name = map[int32]string{
		0: "STITCHING_POLICY_UNSPECIFIED",
		1: "CUT_CURRENT",
		2: "COMPLETE_AD",
	}
	LiveConfig_StitchingPolicy_value = map[string]int32{
		"STITCHING_POLICY_UNSPECIFIED": 0,
		"CUT_CURRENT":                  1,
		"COMPLETE_AD":                  2,
	}
)

func (x LiveConfig_StitchingPolicy) Enum() *LiveConfig_StitchingPolicy {
	p := new(LiveConfig_StitchingPolicy)
	*p = x
	return p
}

func (x LiveConfig_StitchingPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LiveConfig_StitchingPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[2].Descriptor()
}

func (LiveConfig_StitchingPolicy) Type() protoreflect.EnumType {
	return &file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes[2]
}

func (x LiveConfig_StitchingPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LiveConfig_StitchingPolicy.Descriptor instead.
func (LiveConfig_StitchingPolicy) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{0, 1}
}

// Metadata for used to register live configs.
type LiveConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the live config, in the form of
	// `projects/{project}/locations/{location}/liveConfigs/{id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Source URI for the live stream manifest.
	SourceUri string `protobuf:"bytes,2,opt,name=source_uri,json=sourceUri,proto3" json:"source_uri,omitempty"`
	// The default ad tag associated with this live stream config.
	AdTagUri string `protobuf:"bytes,3,opt,name=ad_tag_uri,json=adTagUri,proto3" json:"ad_tag_uri,omitempty"`
	// Additional metadata used to register a live stream with Google Ad Manager
	// (GAM)
	GamLiveConfig *GamLiveConfig `protobuf:"bytes,4,opt,name=gam_live_config,json=gamLiveConfig,proto3" json:"gam_live_config,omitempty"`
	// Output only. State of the live config.
	State LiveConfig_State `protobuf:"varint,5,opt,name=state,proto3,enum=google.cloud.video.stitcher.v1.LiveConfig_State" json:"state,omitempty"`
	// Required. Determines how the ads are tracked. If
	// [gam_live_config][google.cloud.video.stitcher.v1.LiveConfig.gam_live_config]
	// is set, the value must be `CLIENT` because the IMA SDK handles ad tracking.
	AdTracking AdTracking `protobuf:"varint,6,opt,name=ad_tracking,json=adTracking,proto3,enum=google.cloud.video.stitcher.v1.AdTracking" json:"ad_tracking,omitempty"`
	// This must refer to a slate in the same
	// project. If Google Ad Manager (GAM) is used for ads, this string sets the
	// value of `slateCreativeId` in
	// https://developers.google.com/ad-manager/api/reference/v202211/LiveStreamEventService.LiveStreamEvent#slateCreativeId
	DefaultSlate string `protobuf:"bytes,7,opt,name=default_slate,json=defaultSlate,proto3" json:"default_slate,omitempty"`
	// Defines the stitcher behavior in case an ad does not align exactly with
	// the ad break boundaries. If not specified, the default is `CUT_CURRENT`.
	StitchingPolicy LiveConfig_StitchingPolicy `protobuf:"varint,8,opt,name=stitching_policy,json=stitchingPolicy,proto3,enum=google.cloud.video.stitcher.v1.LiveConfig_StitchingPolicy" json:"stitching_policy,omitempty"`
	// The configuration for prefetching ads.
	PrefetchConfig *PrefetchConfig `protobuf:"bytes,10,opt,name=prefetch_config,json=prefetchConfig,proto3" json:"prefetch_config,omitempty"`
}

func (x *LiveConfig) Reset() {
	*x = LiveConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveConfig) ProtoMessage() {}

func (x *LiveConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveConfig.ProtoReflect.Descriptor instead.
func (*LiveConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{0}
}

func (x *LiveConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LiveConfig) GetSourceUri() string {
	if x != nil {
		return x.SourceUri
	}
	return ""
}

func (x *LiveConfig) GetAdTagUri() string {
	if x != nil {
		return x.AdTagUri
	}
	return ""
}

func (x *LiveConfig) GetGamLiveConfig() *GamLiveConfig {
	if x != nil {
		return x.GamLiveConfig
	}
	return nil
}

func (x *LiveConfig) GetState() LiveConfig_State {
	if x != nil {
		return x.State
	}
	return LiveConfig_STATE_UNSPECIFIED
}

func (x *LiveConfig) GetAdTracking() AdTracking {
	if x != nil {
		return x.AdTracking
	}
	return AdTracking_AD_TRACKING_UNSPECIFIED
}

func (x *LiveConfig) GetDefaultSlate() string {
	if x != nil {
		return x.DefaultSlate
	}
	return ""
}

func (x *LiveConfig) GetStitchingPolicy() LiveConfig_StitchingPolicy {
	if x != nil {
		return x.StitchingPolicy
	}
	return LiveConfig_STITCHING_POLICY_UNSPECIFIED
}

func (x *LiveConfig) GetPrefetchConfig() *PrefetchConfig {
	if x != nil {
		return x.PrefetchConfig
	}
	return nil
}

// The configuration for prefetch ads.
type PrefetchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Indicates whether the option to prefetch ad requests is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The duration in seconds of the part of the break to be prefetched.
	// This field is only relevant if prefetch is enabled.
	// You should set this duration to as long as possible to increase the
	// benefits of prefetching, but not longer than the shortest ad break
	// expected. For example, for a live event with 30s and 60s ad breaks, the
	// initial duration should be set to 30s.
	InitialAdRequestDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=initial_ad_request_duration,json=initialAdRequestDuration,proto3" json:"initial_ad_request_duration,omitempty"`
}

func (x *PrefetchConfig) Reset() {
	*x = PrefetchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefetchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefetchConfig) ProtoMessage() {}

func (x *PrefetchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefetchConfig.ProtoReflect.Descriptor instead.
func (*PrefetchConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{1}
}

func (x *PrefetchConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PrefetchConfig) GetInitialAdRequestDuration() *durationpb.Duration {
	if x != nil {
		return x.InitialAdRequestDuration
	}
	return nil
}

// Metadata used to register a live stream with Google Ad Manager (GAM)
type GamLiveConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Ad Manager network code to associate with the live config.
	NetworkCode string `protobuf:"bytes,1,opt,name=network_code,json=networkCode,proto3" json:"network_code,omitempty"`
	// Output only. The asset key identifier generated for the live config.
	AssetKey string `protobuf:"bytes,2,opt,name=asset_key,json=assetKey,proto3" json:"asset_key,omitempty"`
	// Output only. The custom asset key identifier generated for the live config.
	CustomAssetKey string `protobuf:"bytes,3,opt,name=custom_asset_key,json=customAssetKey,proto3" json:"custom_asset_key,omitempty"`
}

func (x *GamLiveConfig) Reset() {
	*x = GamLiveConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GamLiveConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GamLiveConfig) ProtoMessage() {}

func (x *GamLiveConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GamLiveConfig.ProtoReflect.Descriptor instead.
func (*GamLiveConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP(), []int{2}
}

func (x *GamLiveConfig) GetNetworkCode() string {
	if x != nil {
		return x.NetworkCode
	}
	return ""
}

func (x *GamLiveConfig) GetAssetKey() string {
	if x != nil {
		return x.AssetKey
	}
	return ""
}

func (x *GamLiveConfig) GetCustomAssetKey() string {
	if x != nil {
		return x.CustomAssetKey
	}
	return ""
}

var File_google_cloud_video_stitcher_v1_live_configs_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_live_configs_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfa, 0x06, 0x0a, 0x0a, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1c, 0x0a, 0x0a, 0x61,
	0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x64, 0x54, 0x61, 0x67, 0x55, 0x72, 0x69, 0x12, 0x55, 0x0a, 0x0f, 0x67, 0x61, 0x6d,
	0x5f, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x6d, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x4b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a,
	0x0b, 0x61, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x61, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x4c, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x65, 0x0a,
	0x10, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x0f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x57, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x45, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x22, 0x55, 0x0a, 0x0f, 0x53, 0x74, 0x69, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x49, 0x54, 0x43,
	0x48, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x55, 0x54,
	0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x41, 0x44, 0x10, 0x02, 0x3a, 0x6f, 0xea, 0x41, 0x6c,
	0x0a, 0x27, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b,
	0x6c, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x22, 0x89, 0x01, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1d, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x58,
	0x0a, 0x1b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x0d, 0x47, 0x61, 0x6d,
	0x4c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x2a, 0x41, 0x0a, 0x0a, 0x41, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x10, 0x02, 0x42, 0x78, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x4c, 0x69,
	0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescData = file_google_cloud_video_stitcher_v1_live_configs_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_live_configs_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_video_stitcher_v1_live_configs_proto_goTypes = []interface{}{
	(AdTracking)(0),                 // 0: google.cloud.video.stitcher.v1.AdTracking
	(LiveConfig_State)(0),           // 1: google.cloud.video.stitcher.v1.LiveConfig.State
	(LiveConfig_StitchingPolicy)(0), // 2: google.cloud.video.stitcher.v1.LiveConfig.StitchingPolicy
	(*LiveConfig)(nil),              // 3: google.cloud.video.stitcher.v1.LiveConfig
	(*PrefetchConfig)(nil),          // 4: google.cloud.video.stitcher.v1.PrefetchConfig
	(*GamLiveConfig)(nil),           // 5: google.cloud.video.stitcher.v1.GamLiveConfig
	(*durationpb.Duration)(nil),     // 6: google.protobuf.Duration
}
var file_google_cloud_video_stitcher_v1_live_configs_proto_depIdxs = []int32{
	5, // 0: google.cloud.video.stitcher.v1.LiveConfig.gam_live_config:type_name -> google.cloud.video.stitcher.v1.GamLiveConfig
	1, // 1: google.cloud.video.stitcher.v1.LiveConfig.state:type_name -> google.cloud.video.stitcher.v1.LiveConfig.State
	0, // 2: google.cloud.video.stitcher.v1.LiveConfig.ad_tracking:type_name -> google.cloud.video.stitcher.v1.AdTracking
	2, // 3: google.cloud.video.stitcher.v1.LiveConfig.stitching_policy:type_name -> google.cloud.video.stitcher.v1.LiveConfig.StitchingPolicy
	4, // 4: google.cloud.video.stitcher.v1.LiveConfig.prefetch_config:type_name -> google.cloud.video.stitcher.v1.PrefetchConfig
	6, // 5: google.cloud.video.stitcher.v1.PrefetchConfig.initial_ad_request_duration:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_live_configs_proto_init() }
func file_google_cloud_video_stitcher_v1_live_configs_proto_init() {
	if File_google_cloud_video_stitcher_v1_live_configs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveConfig); i {
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
		file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefetchConfig); i {
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
		file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GamLiveConfig); i {
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
			RawDescriptor: file_google_cloud_video_stitcher_v1_live_configs_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_live_configs_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_live_configs_proto_depIdxs,
		EnumInfos:         file_google_cloud_video_stitcher_v1_live_configs_proto_enumTypes,
		MessageInfos:      file_google_cloud_video_stitcher_v1_live_configs_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_live_configs_proto = out.File
	file_google_cloud_video_stitcher_v1_live_configs_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_live_configs_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_live_configs_proto_depIdxs = nil
}
