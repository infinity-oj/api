// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: problems.proto

package protobuf_spec

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Locale string `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"` //    Problem problem = 1;
}

func (x *CreateProblemRequest) Reset() {
	*x = CreateProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problems_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProblemRequest) ProtoMessage() {}

func (x *CreateProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_problems_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProblemRequest.ProtoReflect.Descriptor instead.
func (*CreateProblemRequest) Descriptor() ([]byte, []int) {
	return file_problems_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProblemRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateProblemRequest) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

type CreateProblemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *CreateProblemResponse) Reset() {
	*x = CreateProblemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problems_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProblemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProblemResponse) ProtoMessage() {}

func (x *CreateProblemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_problems_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProblemResponse.ProtoReflect.Descriptor instead.
func (*CreateProblemResponse) Descriptor() ([]byte, []int) {
	return file_problems_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProblemResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

type FetchProblemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
}

func (x *FetchProblemRequest) Reset() {
	*x = FetchProblemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problems_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchProblemRequest) ProtoMessage() {}

func (x *FetchProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_problems_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchProblemRequest.ProtoReflect.Descriptor instead.
func (*FetchProblemRequest) Descriptor() ([]byte, []int) {
	return file_problems_proto_rawDescGZIP(), []int{2}
}

func (x *FetchProblemRequest) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

type FetchProblemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Problem *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
}

func (x *FetchProblemResponse) Reset() {
	*x = FetchProblemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problems_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchProblemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchProblemResponse) ProtoMessage() {}

func (x *FetchProblemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_problems_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchProblemResponse.ProtoReflect.Descriptor instead.
func (*FetchProblemResponse) Descriptor() ([]byte, []int) {
	return file_problems_proto_rawDescGZIP(), []int{3}
}

func (x *FetchProblemResponse) GetProblem() *Problem {
	if x != nil {
		return x.Problem
	}
	return nil
}

type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group        int32  `protobuf:"varint,1,opt,name=group,proto3" json:"group,omitempty"`
	Locale       string `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	ProblemId    string `protobuf:"bytes,3,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	PublicSpace  string `protobuf:"bytes,4,opt,name=public_space,json=publicSpace,proto3" json:"public_space,omitempty"`
	PrivateSpace string `protobuf:"bytes,5,opt,name=private_space,json=privateSpace,proto3" json:"private_space,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_problems_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_problems_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_problems_proto_rawDescGZIP(), []int{4}
}

func (x *Problem) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Problem) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Problem) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Problem) GetPublicSpace() string {
	if x != nil {
		return x.PublicSpace
	}
	return ""
}

func (x *Problem) GetPrivateSpace() string {
	if x != nil {
		return x.PrivateSpace
	}
	return ""
}

var File_problems_proto protoreflect.FileDescriptor

var file_problems_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x34,
	0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x22, 0x9e, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x32, 0x87, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x12, 0x3e,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12,
	0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x14,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_problems_proto_rawDescOnce sync.Once
	file_problems_proto_rawDescData = file_problems_proto_rawDesc
)

func file_problems_proto_rawDescGZIP() []byte {
	file_problems_proto_rawDescOnce.Do(func() {
		file_problems_proto_rawDescData = protoimpl.X.CompressGZIP(file_problems_proto_rawDescData)
	})
	return file_problems_proto_rawDescData
}

var file_problems_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_problems_proto_goTypes = []interface{}{
	(*CreateProblemRequest)(nil),  // 0: CreateProblemRequest
	(*CreateProblemResponse)(nil), // 1: CreateProblemResponse
	(*FetchProblemRequest)(nil),   // 2: FetchProblemRequest
	(*FetchProblemResponse)(nil),  // 3: FetchProblemResponse
	(*Problem)(nil),               // 4: Problem
	(Status)(0),                   // 5: Status
}
var file_problems_proto_depIdxs = []int32{
	5, // 0: CreateProblemResponse.status:type_name -> Status
	4, // 1: FetchProblemResponse.problem:type_name -> Problem
	0, // 2: Problems.CreateProblem:input_type -> CreateProblemRequest
	2, // 3: Problems.FetchProblem:input_type -> FetchProblemRequest
	1, // 4: Problems.CreateProblem:output_type -> CreateProblemResponse
	3, // 5: Problems.FetchProblem:output_type -> FetchProblemResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_problems_proto_init() }
func file_problems_proto_init() {
	if File_problems_proto != nil {
		return
	}
	file_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_problems_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProblemRequest); i {
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
		file_problems_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProblemResponse); i {
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
		file_problems_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchProblemRequest); i {
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
		file_problems_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchProblemResponse); i {
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
		file_problems_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
			RawDescriptor: file_problems_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_problems_proto_goTypes,
		DependencyIndexes: file_problems_proto_depIdxs,
		MessageInfos:      file_problems_proto_msgTypes,
	}.Build()
	File_problems_proto = out.File
	file_problems_proto_rawDesc = nil
	file_problems_proto_goTypes = nil
	file_problems_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProblemsClient is the client API for Problems service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProblemsClient interface {
	CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*CreateProblemResponse, error)
	FetchProblem(ctx context.Context, in *FetchProblemRequest, opts ...grpc.CallOption) (*FetchProblemResponse, error)
}

type problemsClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemsClient(cc grpc.ClientConnInterface) ProblemsClient {
	return &problemsClient{cc}
}

func (c *problemsClient) CreateProblem(ctx context.Context, in *CreateProblemRequest, opts ...grpc.CallOption) (*CreateProblemResponse, error) {
	out := new(CreateProblemResponse)
	err := c.cc.Invoke(ctx, "/Problems/CreateProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemsClient) FetchProblem(ctx context.Context, in *FetchProblemRequest, opts ...grpc.CallOption) (*FetchProblemResponse, error) {
	out := new(FetchProblemResponse)
	err := c.cc.Invoke(ctx, "/Problems/FetchProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemsServer is the server API for Problems service.
type ProblemsServer interface {
	CreateProblem(context.Context, *CreateProblemRequest) (*CreateProblemResponse, error)
	FetchProblem(context.Context, *FetchProblemRequest) (*FetchProblemResponse, error)
}

// UnimplementedProblemsServer can be embedded to have forward compatible implementations.
type UnimplementedProblemsServer struct {
}

func (*UnimplementedProblemsServer) CreateProblem(context.Context, *CreateProblemRequest) (*CreateProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProblem not implemented")
}
func (*UnimplementedProblemsServer) FetchProblem(context.Context, *FetchProblemRequest) (*FetchProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchProblem not implemented")
}

func RegisterProblemsServer(s *grpc.Server, srv ProblemsServer) {
	s.RegisterService(&_Problems_serviceDesc, srv)
}

func _Problems_CreateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemsServer).CreateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Problems/CreateProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemsServer).CreateProblem(ctx, req.(*CreateProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problems_FetchProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemsServer).FetchProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Problems/FetchProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemsServer).FetchProblem(ctx, req.(*FetchProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Problems_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Problems",
	HandlerType: (*ProblemsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProblem",
			Handler:    _Problems_CreateProblem_Handler,
		},
		{
			MethodName: "FetchProblem",
			Handler:    _Problems_FetchProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "problems.proto",
}
