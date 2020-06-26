// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: submission.proto

package protobuf_spec

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateSubmissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmitterId uint64 `protobuf:"varint,1,opt,name=submitter_id,json=submitterId,proto3" json:"submitter_id,omitempty"`
	ProblemId   string `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserSpace   string `protobuf:"bytes,3,opt,name=user_space,json=userSpace,proto3" json:"user_space,omitempty"`
}

func (x *CreateSubmissionRequest) Reset() {
	*x = CreateSubmissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubmissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubmissionRequest) ProtoMessage() {}

func (x *CreateSubmissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubmissionRequest.ProtoReflect.Descriptor instead.
func (*CreateSubmissionRequest) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubmissionRequest) GetSubmitterId() uint64 {
	if x != nil {
		return x.SubmitterId
	}
	return 0
}

func (x *CreateSubmissionRequest) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *CreateSubmissionRequest) GetUserSpace() string {
	if x != nil {
		return x.UserSpace
	}
	return ""
}

type CreateSubmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *CreateSubmissionResponse) Reset() {
	*x = CreateSubmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubmissionResponse) ProtoMessage() {}

func (x *CreateSubmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubmissionResponse.ProtoReflect.Descriptor instead.
func (*CreateSubmissionResponse) Descriptor() ([]byte, []int) {
	return file_submission_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubmissionResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

var File_submission_proto protoreflect.FileDescriptor

var file_submission_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7a, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x70, 0x61, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x56, 0x0a, 0x0b, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submission_proto_rawDescOnce sync.Once
	file_submission_proto_rawDescData = file_submission_proto_rawDesc
)

func file_submission_proto_rawDescGZIP() []byte {
	file_submission_proto_rawDescOnce.Do(func() {
		file_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_submission_proto_rawDescData)
	})
	return file_submission_proto_rawDescData
}

var file_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_submission_proto_goTypes = []interface{}{
	(*CreateSubmissionRequest)(nil),  // 0: CreateSubmissionRequest
	(*CreateSubmissionResponse)(nil), // 1: CreateSubmissionResponse
	(Status)(0),                      // 2: Status
}
var file_submission_proto_depIdxs = []int32{
	2, // 0: CreateSubmissionResponse.status:type_name -> Status
	0, // 1: Submissions.CreateSubmission:input_type -> CreateSubmissionRequest
	1, // 2: Submissions.CreateSubmission:output_type -> CreateSubmissionResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_submission_proto_init() }
func file_submission_proto_init() {
	if File_submission_proto != nil {
		return
	}
	file_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_submission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubmissionRequest); i {
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
		file_submission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubmissionResponse); i {
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
			RawDescriptor: file_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submission_proto_goTypes,
		DependencyIndexes: file_submission_proto_depIdxs,
		MessageInfos:      file_submission_proto_msgTypes,
	}.Build()
	File_submission_proto = out.File
	file_submission_proto_rawDesc = nil
	file_submission_proto_goTypes = nil
	file_submission_proto_depIdxs = nil
}
