// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hello/svc.proto

package hello

import (
	v1 "github.com/sarafanfm/mtserver/example/api/hello/v1"
	v2 "github.com/sarafanfm/mtserver/example/api/hello/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_hello_svc_proto protoreflect.FileDescriptor

var file_hello_svc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x52, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x4c,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x7d, 0x32, 0x50, 0x0a, 0x02,
	0x56, 0x32, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x76, 0x61, 0x6c, 0x7d, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x72,
	0x61, 0x66, 0x61, 0x6e, 0x66, 0x6d, 0x2f, 0x6d, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hello_svc_proto_goTypes = []interface{}{
	(*v1.Request)(nil),  // 0: hello.v1.Request
	(*v2.Request)(nil),  // 1: hello.v2.Request
	(*v1.Response)(nil), // 2: hello.v1.Response
	(*v2.Response)(nil), // 3: hello.v2.Response
}
var file_hello_svc_proto_depIdxs = []int32{
	0, // 0: hello.V1.SayHello:input_type -> hello.v1.Request
	1, // 1: hello.V2.SayHello:input_type -> hello.v2.Request
	2, // 2: hello.V1.SayHello:output_type -> hello.v1.Response
	3, // 3: hello.V2.SayHello:output_type -> hello.v2.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_svc_proto_init() }
func file_hello_svc_proto_init() {
	if File_hello_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_hello_svc_proto_goTypes,
		DependencyIndexes: file_hello_svc_proto_depIdxs,
	}.Build()
	File_hello_svc_proto = out.File
	file_hello_svc_proto_rawDesc = nil
	file_hello_svc_proto_goTypes = nil
	file_hello_svc_proto_depIdxs = nil
}
