// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: staff_tarif_service.proto

package staff_service

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

var File_staff_tarif_service_proto protoreflect.FileDescriptor

var file_staff_tarif_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x03,
	0x0a, 0x12, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_staff_tarif_service_proto_goTypes = []interface{}{
	(*TariffCreateReq)(nil),   // 0: staff_service.TariffCreateReq
	(*TariffGetListReq)(nil),  // 1: staff_service.TariffGetListReq
	(*TariffIdReq)(nil),       // 2: staff_service.TariffIdReq
	(*TariffUpdateReq)(nil),   // 3: staff_service.TariffUpdateReq
	(*TariffCreateResp)(nil),  // 4: staff_service.TariffCreateResp
	(*TariffGetListResp)(nil), // 5: staff_service.TariffGetListResp
	(*StaffTariff)(nil),       // 6: staff_service.StaffTariff
	(*TariffUpdateResp)(nil),  // 7: staff_service.TariffUpdateResp
	(*TariffDeleteResp)(nil),  // 8: staff_service.TariffDeleteResp
}
var file_staff_tarif_service_proto_depIdxs = []int32{
	0, // 0: staff_service.StaffTariffService.Create:input_type -> staff_service.TariffCreateReq
	1, // 1: staff_service.StaffTariffService.GetList:input_type -> staff_service.TariffGetListReq
	2, // 2: staff_service.StaffTariffService.GetById:input_type -> staff_service.TariffIdReq
	3, // 3: staff_service.StaffTariffService.Update:input_type -> staff_service.TariffUpdateReq
	2, // 4: staff_service.StaffTariffService.Delete:input_type -> staff_service.TariffIdReq
	4, // 5: staff_service.StaffTariffService.Create:output_type -> staff_service.TariffCreateResp
	5, // 6: staff_service.StaffTariffService.GetList:output_type -> staff_service.TariffGetListResp
	6, // 7: staff_service.StaffTariffService.GetById:output_type -> staff_service.StaffTariff
	7, // 8: staff_service.StaffTariffService.Update:output_type -> staff_service.TariffUpdateResp
	8, // 9: staff_service.StaffTariffService.Delete:output_type -> staff_service.TariffDeleteResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_staff_tarif_service_proto_init() }
func file_staff_tarif_service_proto_init() {
	if File_staff_tarif_service_proto != nil {
		return
	}
	file_staff_tarif_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_staff_tarif_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_tarif_service_proto_goTypes,
		DependencyIndexes: file_staff_tarif_service_proto_depIdxs,
	}.Build()
	File_staff_tarif_service_proto = out.File
	file_staff_tarif_service_proto_rawDesc = nil
	file_staff_tarif_service_proto_goTypes = nil
	file_staff_tarif_service_proto_depIdxs = nil
}
