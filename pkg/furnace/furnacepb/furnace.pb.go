// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: furnace.proto

package furnacepb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Packages []*Package `protobuf:"bytes,1,rep,name=packages,proto3" json:"packages,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furnace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_furnace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_furnace_proto_rawDescGZIP(), []int{0}
}

func (x *BuildRequest) GetPackages() []*Package {
	if x != nil {
		return x.Packages
	}
	return nil
}

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furnace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_furnace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_furnace_proto_rawDescGZIP(), []int{1}
}

func (x *BuildResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type IsQueuedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package *Package `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *IsQueuedRequest) Reset() {
	*x = IsQueuedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furnace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsQueuedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsQueuedRequest) ProtoMessage() {}

func (x *IsQueuedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_furnace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsQueuedRequest.ProtoReflect.Descriptor instead.
func (*IsQueuedRequest) Descriptor() ([]byte, []int) {
	return file_furnace_proto_rawDescGZIP(), []int{2}
}

func (x *IsQueuedRequest) GetPackage() *Package {
	if x != nil {
		return x.Package
	}
	return nil
}

type IsQueuedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IsQueuedResponse) Reset() {
	*x = IsQueuedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furnace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsQueuedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsQueuedResponse) ProtoMessage() {}

func (x *IsQueuedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_furnace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsQueuedResponse.ProtoReflect.Descriptor instead.
func (*IsQueuedResponse) Descriptor() ([]byte, []int) {
	return file_furnace_proto_rawDescGZIP(), []int{3}
}

func (x *IsQueuedResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_furnace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_furnace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_furnace_proto_rawDescGZIP(), []int{4}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_furnace_proto protoreflect.FileDescriptor

var file_furnace_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x72, 0x79, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x08, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x49, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x22, 0x2a, 0x0a, 0x10, 0x49, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x07,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x82, 0x01, 0x0a, 0x07,
	0x46, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x12, 0x15, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72,
	0x79, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x08, 0x49, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x12, 0x18, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e, 0x49, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2e,
	0x49, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x17, 0x5a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x2f,
	0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_furnace_proto_rawDescOnce sync.Once
	file_furnace_proto_rawDescData = file_furnace_proto_rawDesc
)

func file_furnace_proto_rawDescGZIP() []byte {
	file_furnace_proto_rawDescOnce.Do(func() {
		file_furnace_proto_rawDescData = protoimpl.X.CompressGZIP(file_furnace_proto_rawDescData)
	})
	return file_furnace_proto_rawDescData
}

var file_furnace_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_furnace_proto_goTypes = []interface{}{
	(*BuildRequest)(nil),     // 0: foundry.BuildRequest
	(*BuildResponse)(nil),    // 1: foundry.BuildResponse
	(*IsQueuedRequest)(nil),  // 2: foundry.IsQueuedRequest
	(*IsQueuedResponse)(nil), // 3: foundry.IsQueuedResponse
	(*Package)(nil),          // 4: foundry.Package
}
var file_furnace_proto_depIdxs = []int32{
	4, // 0: foundry.BuildRequest.packages:type_name -> foundry.Package
	4, // 1: foundry.IsQueuedRequest.package:type_name -> foundry.Package
	0, // 2: foundry.Furnace.Build:input_type -> foundry.BuildRequest
	2, // 3: foundry.Furnace.IsQueued:input_type -> foundry.IsQueuedRequest
	1, // 4: foundry.Furnace.Build:output_type -> foundry.BuildResponse
	3, // 5: foundry.Furnace.IsQueued:output_type -> foundry.IsQueuedResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_furnace_proto_init() }
func file_furnace_proto_init() {
	if File_furnace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_furnace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
		file_furnace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
		file_furnace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsQueuedRequest); i {
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
		file_furnace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsQueuedResponse); i {
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
		file_furnace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
			RawDescriptor: file_furnace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_furnace_proto_goTypes,
		DependencyIndexes: file_furnace_proto_depIdxs,
		MessageInfos:      file_furnace_proto_msgTypes,
	}.Build()
	File_furnace_proto = out.File
	file_furnace_proto_rawDesc = nil
	file_furnace_proto_goTypes = nil
	file_furnace_proto_depIdxs = nil
}
