// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: inboxapi/app_versions.proto

package inboxapi

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

type GetVersionsDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
}

func (x *GetVersionsDetailsRequest) Reset() {
	*x = GetVersionsDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_app_versions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsDetailsRequest) ProtoMessage() {}

func (x *GetVersionsDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_app_versions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetVersionsDetailsRequest) Descriptor() ([]byte, []int) {
	return file_inboxapi_app_versions_proto_rawDescGZIP(), []int{0}
}

func (x *GetVersionsDetailsRequest) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

// AppVersionDetails describe specific application version
type AppVersionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// semver
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// platform: ios, android, web
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	// markdown details
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AppVersionDetails) Reset() {
	*x = AppVersionDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_app_versions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppVersionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppVersionDetails) ProtoMessage() {}

func (x *AppVersionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_app_versions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppVersionDetails.ProtoReflect.Descriptor instead.
func (*AppVersionDetails) Descriptor() ([]byte, []int) {
	return file_inboxapi_app_versions_proto_rawDescGZIP(), []int{1}
}

func (x *AppVersionDetails) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AppVersionDetails) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *AppVersionDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetVersionsDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*AppVersionDetails `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *GetVersionsDetailsResponse) Reset() {
	*x = GetVersionsDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inboxapi_app_versions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsDetailsResponse) ProtoMessage() {}

func (x *GetVersionsDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inboxapi_app_versions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetVersionsDetailsResponse) Descriptor() ([]byte, []int) {
	return file_inboxapi_app_versions_proto_rawDescGZIP(), []int{2}
}

func (x *GetVersionsDetailsResponse) GetDetails() []*AppVersionDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_inboxapi_app_versions_proto protoreflect.FileDescriptor

var file_inboxapi_app_versions_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x22, 0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x22, 0x6b, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x32, 0x6e, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x5f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x69,
	0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inboxapi_app_versions_proto_rawDescOnce sync.Once
	file_inboxapi_app_versions_proto_rawDescData = file_inboxapi_app_versions_proto_rawDesc
)

func file_inboxapi_app_versions_proto_rawDescGZIP() []byte {
	file_inboxapi_app_versions_proto_rawDescOnce.Do(func() {
		file_inboxapi_app_versions_proto_rawDescData = protoimpl.X.CompressGZIP(file_inboxapi_app_versions_proto_rawDescData)
	})
	return file_inboxapi_app_versions_proto_rawDescData
}

var file_inboxapi_app_versions_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inboxapi_app_versions_proto_goTypes = []any{
	(*GetVersionsDetailsRequest)(nil),  // 0: inboxapi.GetVersionsDetailsRequest
	(*AppVersionDetails)(nil),          // 1: inboxapi.AppVersionDetails
	(*GetVersionsDetailsResponse)(nil), // 2: inboxapi.GetVersionsDetailsResponse
}
var file_inboxapi_app_versions_proto_depIdxs = []int32{
	1, // 0: inboxapi.GetVersionsDetailsResponse.details:type_name -> inboxapi.AppVersionDetails
	0, // 1: inboxapi.AppVersions.GetVersionsDetails:input_type -> inboxapi.GetVersionsDetailsRequest
	2, // 2: inboxapi.AppVersions.GetVersionsDetails:output_type -> inboxapi.GetVersionsDetailsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inboxapi_app_versions_proto_init() }
func file_inboxapi_app_versions_proto_init() {
	if File_inboxapi_app_versions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inboxapi_app_versions_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetVersionsDetailsRequest); i {
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
		file_inboxapi_app_versions_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AppVersionDetails); i {
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
		file_inboxapi_app_versions_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetVersionsDetailsResponse); i {
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
			RawDescriptor: file_inboxapi_app_versions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inboxapi_app_versions_proto_goTypes,
		DependencyIndexes: file_inboxapi_app_versions_proto_depIdxs,
		MessageInfos:      file_inboxapi_app_versions_proto_msgTypes,
	}.Build()
	File_inboxapi_app_versions_proto = out.File
	file_inboxapi_app_versions_proto_rawDesc = nil
	file_inboxapi_app_versions_proto_goTypes = nil
	file_inboxapi_app_versions_proto_depIdxs = nil
}
