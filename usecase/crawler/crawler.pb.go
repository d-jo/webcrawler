// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: crawler.proto

package webcrawler

import (
	entity "github.com/d-jo/webcrawler/entity"
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

var File_crawler_proto protoreflect.FileDescriptor

var file_crawler_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc3, 0x01, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x70, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x17, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x14, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x2d, 0x6a, 0x6f, 0x2f,
	0x77, 0x65, 0x62, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x63, 0x61,
	0x73, 0x65, 0x2f, 0x77, 0x65, 0x62, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_crawler_proto_goTypes = []interface{}{
	(*entity.StartCommand)(nil),    // 0: entity.StartCommand
	(*entity.StopCommand)(nil),     // 1: entity.StopCommand
	(*entity.ListCommand)(nil),     // 2: entity.ListCommand
	(*entity.GenericResponse)(nil), // 3: entity.GenericResponse
	(*entity.ListResponse)(nil),    // 4: entity.ListResponse
}
var file_crawler_proto_depIdxs = []int32{
	0, // 0: crawler.WebCrawler.StartCrawling:input_type -> entity.StartCommand
	1, // 1: crawler.WebCrawler.StopCrawling:input_type -> entity.StopCommand
	2, // 2: crawler.WebCrawler.List:input_type -> entity.ListCommand
	3, // 3: crawler.WebCrawler.StartCrawling:output_type -> entity.GenericResponse
	3, // 4: crawler.WebCrawler.StopCrawling:output_type -> entity.GenericResponse
	4, // 5: crawler.WebCrawler.List:output_type -> entity.ListResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crawler_proto_init() }
func file_crawler_proto_init() {
	if File_crawler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crawler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crawler_proto_goTypes,
		DependencyIndexes: file_crawler_proto_depIdxs,
	}.Build()
	File_crawler_proto = out.File
	file_crawler_proto_rawDesc = nil
	file_crawler_proto_goTypes = nil
	file_crawler_proto_depIdxs = nil
}
