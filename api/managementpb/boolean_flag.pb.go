// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: managementpb/boolean_flag.proto

package managementpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BooleanFlag represent a command to set some boolean property to true,
// to false, or avoid changing that property.
type BooleanFlag int32

const (
	// Do not change boolean property. Default value.
	BooleanFlag_DO_NOT_CHANGE BooleanFlag = 0
	// True.
	BooleanFlag_TRUE BooleanFlag = 1
	// False.
	BooleanFlag_FALSE BooleanFlag = 2
)

// Enum value maps for BooleanFlag.
var (
	BooleanFlag_name = map[int32]string{
		0: "DO_NOT_CHANGE",
		1: "TRUE",
		2: "FALSE",
	}
	BooleanFlag_value = map[string]int32{
		"DO_NOT_CHANGE": 0,
		"TRUE":          1,
		"FALSE":         2,
	}
)

func (x BooleanFlag) Enum() *BooleanFlag {
	p := new(BooleanFlag)
	*p = x
	return p
}

func (x BooleanFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BooleanFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_boolean_flag_proto_enumTypes[0].Descriptor()
}

func (BooleanFlag) Type() protoreflect.EnumType {
	return &file_managementpb_boolean_flag_proto_enumTypes[0]
}

func (x BooleanFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BooleanFlag.Descriptor instead.
func (BooleanFlag) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_boolean_flag_proto_rawDescGZIP(), []int{0}
}

var File_managementpb_boolean_flag_proto protoreflect.FileDescriptor

var file_managementpb_boolean_flag_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2a,
	0x35, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x11,
	0x0a, 0x0d, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x52, 0x55, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x46,
	0x41, 0x4c, 0x53, 0x45, 0x10, 0x02, 0x42, 0x9d, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x42, 0x10, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63,
	0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02,
	0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xca, 0x02, 0x0c,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xe2, 0x02, 0x18, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_boolean_flag_proto_rawDescOnce sync.Once
	file_managementpb_boolean_flag_proto_rawDescData = file_managementpb_boolean_flag_proto_rawDesc
)

func file_managementpb_boolean_flag_proto_rawDescGZIP() []byte {
	file_managementpb_boolean_flag_proto_rawDescOnce.Do(func() {
		file_managementpb_boolean_flag_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_boolean_flag_proto_rawDescData)
	})
	return file_managementpb_boolean_flag_proto_rawDescData
}

var (
	file_managementpb_boolean_flag_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_managementpb_boolean_flag_proto_goTypes   = []interface{}{
		(BooleanFlag)(0), // 0: managementpb.BooleanFlag
	}
)

var file_managementpb_boolean_flag_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_boolean_flag_proto_init() }
func file_managementpb_boolean_flag_proto_init() {
	if File_managementpb_boolean_flag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_boolean_flag_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_boolean_flag_proto_goTypes,
		DependencyIndexes: file_managementpb_boolean_flag_proto_depIdxs,
		EnumInfos:         file_managementpb_boolean_flag_proto_enumTypes,
	}.Build()
	File_managementpb_boolean_flag_proto = out.File
	file_managementpb_boolean_flag_proto_rawDesc = nil
	file_managementpb_boolean_flag_proto_goTypes = nil
	file_managementpb_boolean_flag_proto_depIdxs = nil
}
