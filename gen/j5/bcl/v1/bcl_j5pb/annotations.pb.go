// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/bcl/v1/annotations.proto

package bcl_j5pb

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

type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children    map[string]*SourceLocation `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StartLine   int32                      `protobuf:"varint,2,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	StartColumn int32                      `protobuf:"varint,3,opt,name=start_column,json=startColumn,proto3" json:"start_column,omitempty"`
	EndLine     int32                      `protobuf:"varint,4,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
	EndColumn   int32                      `protobuf:"varint,5,opt,name=end_column,json=endColumn,proto3" json:"end_column,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_bcl_v1_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_j5_bcl_v1_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_j5_bcl_v1_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *SourceLocation) GetChildren() map[string]*SourceLocation {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *SourceLocation) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *SourceLocation) GetStartColumn() int32 {
	if x != nil {
		return x.StartColumn
	}
	return 0
}

func (x *SourceLocation) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *SourceLocation) GetEndColumn() int32 {
	if x != nil {
		return x.EndColumn
	}
	return 0
}

var File_j5_bcl_v1_annotations_proto protoreflect.FileDescriptor

var file_j5_bcl_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6a, 0x35, 0x2f, 0x62, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a,
	0x35, 0x2e, 0x62, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xa9, 0x02, 0x0a, 0x0e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x6a, 0x35, 0x2e, 0x62, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x1a, 0x56, 0x0a, 0x0d,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x62, 0x63, 0x6c, 0x2e, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x35, 0x2f, 0x62, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x63, 0x6c, 0x5f, 0x6a, 0x35, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_j5_bcl_v1_annotations_proto_rawDescOnce sync.Once
	file_j5_bcl_v1_annotations_proto_rawDescData = file_j5_bcl_v1_annotations_proto_rawDesc
)

func file_j5_bcl_v1_annotations_proto_rawDescGZIP() []byte {
	file_j5_bcl_v1_annotations_proto_rawDescOnce.Do(func() {
		file_j5_bcl_v1_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_bcl_v1_annotations_proto_rawDescData)
	})
	return file_j5_bcl_v1_annotations_proto_rawDescData
}

var file_j5_bcl_v1_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_j5_bcl_v1_annotations_proto_goTypes = []any{
	(*SourceLocation)(nil), // 0: j5.bcl.v1.SourceLocation
	nil,                    // 1: j5.bcl.v1.SourceLocation.ChildrenEntry
}
var file_j5_bcl_v1_annotations_proto_depIdxs = []int32{
	1, // 0: j5.bcl.v1.SourceLocation.children:type_name -> j5.bcl.v1.SourceLocation.ChildrenEntry
	0, // 1: j5.bcl.v1.SourceLocation.ChildrenEntry.value:type_name -> j5.bcl.v1.SourceLocation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_j5_bcl_v1_annotations_proto_init() }
func file_j5_bcl_v1_annotations_proto_init() {
	if File_j5_bcl_v1_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_bcl_v1_annotations_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SourceLocation); i {
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
			RawDescriptor: file_j5_bcl_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_j5_bcl_v1_annotations_proto_goTypes,
		DependencyIndexes: file_j5_bcl_v1_annotations_proto_depIdxs,
		MessageInfos:      file_j5_bcl_v1_annotations_proto_msgTypes,
	}.Build()
	File_j5_bcl_v1_annotations_proto = out.File
	file_j5_bcl_v1_annotations_proto_rawDesc = nil
	file_j5_bcl_v1_annotations_proto_goTypes = nil
	file_j5_bcl_v1_annotations_proto_depIdxs = nil
}
