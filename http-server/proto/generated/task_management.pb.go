// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: proto/task_management.proto

package generated

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

var File_proto_task_management_proto protoreflect.FileDescriptor

var file_proto_task_management_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd7, 0x02, 0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x12, 0x2e, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x2f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_task_management_proto_goTypes = []any{
	(*Task)(nil),    // 0: proto.Task
	(*NoParam)(nil), // 1: proto.NoParam
	(*TaskId)(nil),  // 2: proto.TaskId
	(*Message)(nil), // 3: proto.Message
}
var file_proto_task_management_proto_depIdxs = []int32{
	0, // 0: proto.TaskManagementService.CreateTask:input_type -> proto.Task
	0, // 1: proto.TaskManagementService.CreateTasks:input_type -> proto.Task
	0, // 2: proto.TaskManagementService.CreateTaskList:input_type -> proto.Task
	1, // 3: proto.TaskManagementService.GetAllTask:input_type -> proto.NoParam
	2, // 4: proto.TaskManagementService.GetTaskById:input_type -> proto.TaskId
	2, // 5: proto.TaskManagementService.UpdateTaskById:input_type -> proto.TaskId
	2, // 6: proto.TaskManagementService.DeleteTaskById:input_type -> proto.TaskId
	3, // 7: proto.TaskManagementService.CreateTask:output_type -> proto.Message
	3, // 8: proto.TaskManagementService.CreateTasks:output_type -> proto.Message
	0, // 9: proto.TaskManagementService.CreateTaskList:output_type -> proto.Task
	0, // 10: proto.TaskManagementService.GetAllTask:output_type -> proto.Task
	0, // 11: proto.TaskManagementService.GetTaskById:output_type -> proto.Task
	0, // 12: proto.TaskManagementService.UpdateTaskById:output_type -> proto.Task
	3, // 13: proto.TaskManagementService.DeleteTaskById:output_type -> proto.Message
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_task_management_proto_init() }
func file_proto_task_management_proto_init() {
	if File_proto_task_management_proto != nil {
		return
	}
	file_proto_task_proto_init()
	file_proto_no_param_proto_init()
	file_proto_message_proto_init()
	file_proto_task_id_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_task_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_management_proto_goTypes,
		DependencyIndexes: file_proto_task_management_proto_depIdxs,
	}.Build()
	File_proto_task_management_proto = out.File
	file_proto_task_management_proto_rawDesc = nil
	file_proto_task_management_proto_goTypes = nil
	file_proto_task_management_proto_depIdxs = nil
}
