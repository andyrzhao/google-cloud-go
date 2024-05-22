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
// source: google/monitoring/dashboard/v1/widget.proto

package dashboardpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Widget contains a single dashboard component and configuration of how to
// present the component in the dashboard.
type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The title of the widget.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Content defines the component used to populate the widget.
	//
	// Types that are assignable to Content:
	//
	//	*Widget_XyChart
	//	*Widget_Scorecard
	//	*Widget_Text
	//	*Widget_Blank
	//	*Widget_AlertChart
	//	*Widget_TimeSeriesTable
	//	*Widget_CollapsibleGroup
	//	*Widget_LogsPanel
	//	*Widget_IncidentList
	//	*Widget_PieChart
	//	*Widget_ErrorReportingPanel
	//	*Widget_SectionHeader
	//	*Widget_SingleViewGroup
	Content isWidget_Content `protobuf_oneof:"content"`
	// Optional. The widget id. Ids may be made up of alphanumerics, dashes and
	// underscores. Widget ids are optional.
	Id string `protobuf:"bytes,17,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_widget_proto_rawDescGZIP(), []int{0}
}

func (x *Widget) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (m *Widget) GetContent() isWidget_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Widget) GetXyChart() *XyChart {
	if x, ok := x.GetContent().(*Widget_XyChart); ok {
		return x.XyChart
	}
	return nil
}

func (x *Widget) GetScorecard() *Scorecard {
	if x, ok := x.GetContent().(*Widget_Scorecard); ok {
		return x.Scorecard
	}
	return nil
}

func (x *Widget) GetText() *Text {
	if x, ok := x.GetContent().(*Widget_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Widget) GetBlank() *emptypb.Empty {
	if x, ok := x.GetContent().(*Widget_Blank); ok {
		return x.Blank
	}
	return nil
}

func (x *Widget) GetAlertChart() *AlertChart {
	if x, ok := x.GetContent().(*Widget_AlertChart); ok {
		return x.AlertChart
	}
	return nil
}

func (x *Widget) GetTimeSeriesTable() *TimeSeriesTable {
	if x, ok := x.GetContent().(*Widget_TimeSeriesTable); ok {
		return x.TimeSeriesTable
	}
	return nil
}

func (x *Widget) GetCollapsibleGroup() *CollapsibleGroup {
	if x, ok := x.GetContent().(*Widget_CollapsibleGroup); ok {
		return x.CollapsibleGroup
	}
	return nil
}

func (x *Widget) GetLogsPanel() *LogsPanel {
	if x, ok := x.GetContent().(*Widget_LogsPanel); ok {
		return x.LogsPanel
	}
	return nil
}

func (x *Widget) GetIncidentList() *IncidentList {
	if x, ok := x.GetContent().(*Widget_IncidentList); ok {
		return x.IncidentList
	}
	return nil
}

func (x *Widget) GetPieChart() *PieChart {
	if x, ok := x.GetContent().(*Widget_PieChart); ok {
		return x.PieChart
	}
	return nil
}

func (x *Widget) GetErrorReportingPanel() *ErrorReportingPanel {
	if x, ok := x.GetContent().(*Widget_ErrorReportingPanel); ok {
		return x.ErrorReportingPanel
	}
	return nil
}

func (x *Widget) GetSectionHeader() *SectionHeader {
	if x, ok := x.GetContent().(*Widget_SectionHeader); ok {
		return x.SectionHeader
	}
	return nil
}

func (x *Widget) GetSingleViewGroup() *SingleViewGroup {
	if x, ok := x.GetContent().(*Widget_SingleViewGroup); ok {
		return x.SingleViewGroup
	}
	return nil
}

func (x *Widget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type isWidget_Content interface {
	isWidget_Content()
}

type Widget_XyChart struct {
	// A chart of time series data.
	XyChart *XyChart `protobuf:"bytes,2,opt,name=xy_chart,json=xyChart,proto3,oneof"`
}

type Widget_Scorecard struct {
	// A scorecard summarizing time series data.
	Scorecard *Scorecard `protobuf:"bytes,3,opt,name=scorecard,proto3,oneof"`
}

type Widget_Text struct {
	// A raw string or markdown displaying textual content.
	Text *Text `protobuf:"bytes,4,opt,name=text,proto3,oneof"`
}

type Widget_Blank struct {
	// A blank space.
	Blank *emptypb.Empty `protobuf:"bytes,5,opt,name=blank,proto3,oneof"`
}

type Widget_AlertChart struct {
	// A chart of alert policy data.
	AlertChart *AlertChart `protobuf:"bytes,7,opt,name=alert_chart,json=alertChart,proto3,oneof"`
}

type Widget_TimeSeriesTable struct {
	// A widget that displays time series data in a tabular format.
	TimeSeriesTable *TimeSeriesTable `protobuf:"bytes,8,opt,name=time_series_table,json=timeSeriesTable,proto3,oneof"`
}

type Widget_CollapsibleGroup struct {
	// A widget that groups the other widgets. All widgets that are within
	// the area spanned by the grouping widget are considered member widgets.
	CollapsibleGroup *CollapsibleGroup `protobuf:"bytes,9,opt,name=collapsible_group,json=collapsibleGroup,proto3,oneof"`
}

type Widget_LogsPanel struct {
	// A widget that shows a stream of logs.
	LogsPanel *LogsPanel `protobuf:"bytes,10,opt,name=logs_panel,json=logsPanel,proto3,oneof"`
}

type Widget_IncidentList struct {
	// A widget that shows list of incidents.
	IncidentList *IncidentList `protobuf:"bytes,12,opt,name=incident_list,json=incidentList,proto3,oneof"`
}

type Widget_PieChart struct {
	// A widget that displays timeseries data as a pie chart.
	PieChart *PieChart `protobuf:"bytes,14,opt,name=pie_chart,json=pieChart,proto3,oneof"`
}

type Widget_ErrorReportingPanel struct {
	// A widget that displays a list of error groups.
	ErrorReportingPanel *ErrorReportingPanel `protobuf:"bytes,19,opt,name=error_reporting_panel,json=errorReportingPanel,proto3,oneof"`
}

type Widget_SectionHeader struct {
	// A widget that defines a section header for easier navigation of the
	// dashboard.
	SectionHeader *SectionHeader `protobuf:"bytes,21,opt,name=section_header,json=sectionHeader,proto3,oneof"`
}

type Widget_SingleViewGroup struct {
	// A widget that groups the other widgets by using a dropdown menu.
	SingleViewGroup *SingleViewGroup `protobuf:"bytes,22,opt,name=single_view_group,json=singleViewGroup,proto3,oneof"`
}

func (*Widget_XyChart) isWidget_Content() {}

func (*Widget_Scorecard) isWidget_Content() {}

func (*Widget_Text) isWidget_Content() {}

func (*Widget_Blank) isWidget_Content() {}

func (*Widget_AlertChart) isWidget_Content() {}

func (*Widget_TimeSeriesTable) isWidget_Content() {}

func (*Widget_CollapsibleGroup) isWidget_Content() {}

func (*Widget_LogsPanel) isWidget_Content() {}

func (*Widget_IncidentList) isWidget_Content() {}

func (*Widget_PieChart) isWidget_Content() {}

func (*Widget_ErrorReportingPanel) isWidget_Content() {}

func (*Widget_SectionHeader) isWidget_Content() {}

func (*Widget_SingleViewGroup) isWidget_Content() {}

var File_google_monitoring_dashboard_v1_widget_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_widget_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x65, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x79, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x08, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x44, 0x0a, 0x08,
	0x78, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x58, 0x79, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x07, 0x78, 0x79, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x12, 0x3a, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78,
	0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x62, 0x6c, 0x61,
	0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x12, 0x4d, 0x0a, 0x0b, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x5d, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x5f, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x70, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x70, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4a, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73,
	0x5f, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x73, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x73, 0x50,
	0x61, 0x6e, 0x65, 0x6c, 0x12, 0x53, 0x0a, 0x0d, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x09, 0x70, 0x69, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x08, 0x70, 0x69, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x12, 0x69, 0x0a, 0x15, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x12, 0x56, 0x0a,
	0x0e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x11, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0xf4, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x70, 0x62, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31,
	0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_monitoring_dashboard_v1_widget_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_widget_proto_rawDescData = file_google_monitoring_dashboard_v1_widget_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_widget_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_widget_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_widget_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_widget_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_monitoring_dashboard_v1_widget_proto_goTypes = []interface{}{
	(*Widget)(nil),              // 0: google.monitoring.dashboard.v1.Widget
	(*XyChart)(nil),             // 1: google.monitoring.dashboard.v1.XyChart
	(*Scorecard)(nil),           // 2: google.monitoring.dashboard.v1.Scorecard
	(*Text)(nil),                // 3: google.monitoring.dashboard.v1.Text
	(*emptypb.Empty)(nil),       // 4: google.protobuf.Empty
	(*AlertChart)(nil),          // 5: google.monitoring.dashboard.v1.AlertChart
	(*TimeSeriesTable)(nil),     // 6: google.monitoring.dashboard.v1.TimeSeriesTable
	(*CollapsibleGroup)(nil),    // 7: google.monitoring.dashboard.v1.CollapsibleGroup
	(*LogsPanel)(nil),           // 8: google.monitoring.dashboard.v1.LogsPanel
	(*IncidentList)(nil),        // 9: google.monitoring.dashboard.v1.IncidentList
	(*PieChart)(nil),            // 10: google.monitoring.dashboard.v1.PieChart
	(*ErrorReportingPanel)(nil), // 11: google.monitoring.dashboard.v1.ErrorReportingPanel
	(*SectionHeader)(nil),       // 12: google.monitoring.dashboard.v1.SectionHeader
	(*SingleViewGroup)(nil),     // 13: google.monitoring.dashboard.v1.SingleViewGroup
}
var file_google_monitoring_dashboard_v1_widget_proto_depIdxs = []int32{
	1,  // 0: google.monitoring.dashboard.v1.Widget.xy_chart:type_name -> google.monitoring.dashboard.v1.XyChart
	2,  // 1: google.monitoring.dashboard.v1.Widget.scorecard:type_name -> google.monitoring.dashboard.v1.Scorecard
	3,  // 2: google.monitoring.dashboard.v1.Widget.text:type_name -> google.monitoring.dashboard.v1.Text
	4,  // 3: google.monitoring.dashboard.v1.Widget.blank:type_name -> google.protobuf.Empty
	5,  // 4: google.monitoring.dashboard.v1.Widget.alert_chart:type_name -> google.monitoring.dashboard.v1.AlertChart
	6,  // 5: google.monitoring.dashboard.v1.Widget.time_series_table:type_name -> google.monitoring.dashboard.v1.TimeSeriesTable
	7,  // 6: google.monitoring.dashboard.v1.Widget.collapsible_group:type_name -> google.monitoring.dashboard.v1.CollapsibleGroup
	8,  // 7: google.monitoring.dashboard.v1.Widget.logs_panel:type_name -> google.monitoring.dashboard.v1.LogsPanel
	9,  // 8: google.monitoring.dashboard.v1.Widget.incident_list:type_name -> google.monitoring.dashboard.v1.IncidentList
	10, // 9: google.monitoring.dashboard.v1.Widget.pie_chart:type_name -> google.monitoring.dashboard.v1.PieChart
	11, // 10: google.monitoring.dashboard.v1.Widget.error_reporting_panel:type_name -> google.monitoring.dashboard.v1.ErrorReportingPanel
	12, // 11: google.monitoring.dashboard.v1.Widget.section_header:type_name -> google.monitoring.dashboard.v1.SectionHeader
	13, // 12: google.monitoring.dashboard.v1.Widget.single_view_group:type_name -> google.monitoring.dashboard.v1.SingleViewGroup
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_widget_proto_init() }
func file_google_monitoring_dashboard_v1_widget_proto_init() {
	if File_google_monitoring_dashboard_v1_widget_proto != nil {
		return
	}
	file_google_monitoring_dashboard_v1_alertchart_proto_init()
	file_google_monitoring_dashboard_v1_collapsible_group_proto_init()
	file_google_monitoring_dashboard_v1_error_reporting_panel_proto_init()
	file_google_monitoring_dashboard_v1_incident_list_proto_init()
	file_google_monitoring_dashboard_v1_logs_panel_proto_init()
	file_google_monitoring_dashboard_v1_piechart_proto_init()
	file_google_monitoring_dashboard_v1_scorecard_proto_init()
	file_google_monitoring_dashboard_v1_section_header_proto_init()
	file_google_monitoring_dashboard_v1_single_view_group_proto_init()
	file_google_monitoring_dashboard_v1_table_proto_init()
	file_google_monitoring_dashboard_v1_text_proto_init()
	file_google_monitoring_dashboard_v1_xychart_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget); i {
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
	file_google_monitoring_dashboard_v1_widget_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Widget_XyChart)(nil),
		(*Widget_Scorecard)(nil),
		(*Widget_Text)(nil),
		(*Widget_Blank)(nil),
		(*Widget_AlertChart)(nil),
		(*Widget_TimeSeriesTable)(nil),
		(*Widget_CollapsibleGroup)(nil),
		(*Widget_LogsPanel)(nil),
		(*Widget_IncidentList)(nil),
		(*Widget_PieChart)(nil),
		(*Widget_ErrorReportingPanel)(nil),
		(*Widget_SectionHeader)(nil),
		(*Widget_SingleViewGroup)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_monitoring_dashboard_v1_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_widget_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_widget_proto_depIdxs,
		MessageInfos:      file_google_monitoring_dashboard_v1_widget_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_widget_proto = out.File
	file_google_monitoring_dashboard_v1_widget_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_widget_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_widget_proto_depIdxs = nil
}
