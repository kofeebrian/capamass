// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/amass/enum/enum_service.proto

package enum

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_amass_enum_enum_service_proto protoreflect.FileDescriptor

var file_protos_amass_enum_enum_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6d, 0x61, 0x73, 0x73, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x1a, 0x24, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6d, 0x61, 0x73, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x3d, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x11, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6f, 0x66, 0x65, 0x65, 0x62, 0x72, 0x69, 0x61, 0x6e, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x6d,
	0x61, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6d, 0x61, 0x73, 0x73,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_amass_enum_enum_service_proto_goTypes = []interface{}{
	(*EnumRequest)(nil),  // 0: enum.EnumRequest
	(*EnumResponse)(nil), // 1: enum.EnumResponse
}
var file_protos_amass_enum_enum_service_proto_depIdxs = []int32{
	0, // 0: enum.EnumService.Run:input_type -> enum.EnumRequest
	1, // 1: enum.EnumService.Run:output_type -> enum.EnumResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_amass_enum_enum_service_proto_init() }
func file_protos_amass_enum_enum_service_proto_init() {
	if File_protos_amass_enum_enum_service_proto != nil {
		return
	}
	file_protos_amass_enum_enum_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_amass_enum_enum_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_amass_enum_enum_service_proto_goTypes,
		DependencyIndexes: file_protos_amass_enum_enum_service_proto_depIdxs,
	}.Build()
	File_protos_amass_enum_enum_service_proto = out.File
	file_protos_amass_enum_enum_service_proto_rawDesc = nil
	file_protos_amass_enum_enum_service_proto_goTypes = nil
	file_protos_amass_enum_enum_service_proto_depIdxs = nil
}
