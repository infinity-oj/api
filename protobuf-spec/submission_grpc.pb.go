// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf_spec

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubmissionsClient is the client API for Submissions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionsClient interface {
	CreateSubmission(ctx context.Context, in *CreateSubmissionRequest, opts ...grpc.CallOption) (*CreateSubmissionResponse, error)
}

type submissionsClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionsClient(cc grpc.ClientConnInterface) SubmissionsClient {
	return &submissionsClient{cc}
}

func (c *submissionsClient) CreateSubmission(ctx context.Context, in *CreateSubmissionRequest, opts ...grpc.CallOption) (*CreateSubmissionResponse, error) {
	out := new(CreateSubmissionResponse)
	err := c.cc.Invoke(ctx, "/Submissions/CreateSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionsServer is the server API for Submissions service.
// All implementations must embed UnimplementedSubmissionsServer
// for forward compatibility
type SubmissionsServer interface {
	CreateSubmission(context.Context, *CreateSubmissionRequest) (*CreateSubmissionResponse, error)
	mustEmbedUnimplementedSubmissionsServer()
}

// UnimplementedSubmissionsServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionsServer struct {
}

func (*UnimplementedSubmissionsServer) CreateSubmission(context.Context, *CreateSubmissionRequest) (*CreateSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubmission not implemented")
}
func (*UnimplementedSubmissionsServer) mustEmbedUnimplementedSubmissionsServer() {}

func RegisterSubmissionsServer(s *grpc.Server, srv SubmissionsServer) {
	s.RegisterService(&_Submissions_serviceDesc, srv)
}

func _Submissions_CreateSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionsServer).CreateSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Submissions/CreateSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionsServer).CreateSubmission(ctx, req.(*CreateSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Submissions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Submissions",
	HandlerType: (*SubmissionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSubmission",
			Handler:    _Submissions_CreateSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "submission.proto",
}
