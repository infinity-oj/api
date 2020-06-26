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

// JudgementsClient is the client API for Judgements service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgementsClient interface {
	SubmitJudgement(ctx context.Context, in *SubmitJudgementRequest, opts ...grpc.CallOption) (*SubmitJudgementResponse, error)
	FetchHashFile(ctx context.Context, in *FetchFileHashRequest, opts ...grpc.CallOption) (*FetchFileHashResponse, error)
	FetchFile(ctx context.Context, in *FetchJudgeFileRequest, opts ...grpc.CallOption) (*FetchJudgeFileResponse, error)
	FetchJudgement(ctx context.Context, in *FetchJudgementRequest, opts ...grpc.CallOption) (*FetchJudgementResponse, error)
	ReturnJudgement(ctx context.Context, in *ReturnJudgementRequest, opts ...grpc.CallOption) (*ReturnJudgementResponse, error)
	ListJudgements(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type judgementsClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgementsClient(cc grpc.ClientConnInterface) JudgementsClient {
	return &judgementsClient{cc}
}

func (c *judgementsClient) SubmitJudgement(ctx context.Context, in *SubmitJudgementRequest, opts ...grpc.CallOption) (*SubmitJudgementResponse, error) {
	out := new(SubmitJudgementResponse)
	err := c.cc.Invoke(ctx, "/Judgements/SubmitJudgement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgementsClient) FetchHashFile(ctx context.Context, in *FetchFileHashRequest, opts ...grpc.CallOption) (*FetchFileHashResponse, error) {
	out := new(FetchFileHashResponse)
	err := c.cc.Invoke(ctx, "/Judgements/FetchHashFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgementsClient) FetchFile(ctx context.Context, in *FetchJudgeFileRequest, opts ...grpc.CallOption) (*FetchJudgeFileResponse, error) {
	out := new(FetchJudgeFileResponse)
	err := c.cc.Invoke(ctx, "/Judgements/FetchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgementsClient) FetchJudgement(ctx context.Context, in *FetchJudgementRequest, opts ...grpc.CallOption) (*FetchJudgementResponse, error) {
	out := new(FetchJudgementResponse)
	err := c.cc.Invoke(ctx, "/Judgements/FetchJudgement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgementsClient) ReturnJudgement(ctx context.Context, in *ReturnJudgementRequest, opts ...grpc.CallOption) (*ReturnJudgementResponse, error) {
	out := new(ReturnJudgementResponse)
	err := c.cc.Invoke(ctx, "/Judgements/ReturnJudgement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgementsClient) ListJudgements(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/Judgements/ListJudgements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgementsServer is the server API for Judgements service.
// All implementations must embed UnimplementedJudgementsServer
// for forward compatibility
type JudgementsServer interface {
	SubmitJudgement(context.Context, *SubmitJudgementRequest) (*SubmitJudgementResponse, error)
	FetchHashFile(context.Context, *FetchFileHashRequest) (*FetchFileHashResponse, error)
	FetchFile(context.Context, *FetchJudgeFileRequest) (*FetchJudgeFileResponse, error)
	FetchJudgement(context.Context, *FetchJudgementRequest) (*FetchJudgementResponse, error)
	ReturnJudgement(context.Context, *ReturnJudgementRequest) (*ReturnJudgementResponse, error)
	ListJudgements(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedJudgementsServer()
}

// UnimplementedJudgementsServer must be embedded to have forward compatible implementations.
type UnimplementedJudgementsServer struct {
}

func (*UnimplementedJudgementsServer) SubmitJudgement(context.Context, *SubmitJudgementRequest) (*SubmitJudgementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJudgement not implemented")
}
func (*UnimplementedJudgementsServer) FetchHashFile(context.Context, *FetchFileHashRequest) (*FetchFileHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHashFile not implemented")
}
func (*UnimplementedJudgementsServer) FetchFile(context.Context, *FetchJudgeFileRequest) (*FetchJudgeFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchFile not implemented")
}
func (*UnimplementedJudgementsServer) FetchJudgement(context.Context, *FetchJudgementRequest) (*FetchJudgementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchJudgement not implemented")
}
func (*UnimplementedJudgementsServer) ReturnJudgement(context.Context, *ReturnJudgementRequest) (*ReturnJudgementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnJudgement not implemented")
}
func (*UnimplementedJudgementsServer) ListJudgements(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJudgements not implemented")
}
func (*UnimplementedJudgementsServer) mustEmbedUnimplementedJudgementsServer() {}

func RegisterJudgementsServer(s *grpc.Server, srv JudgementsServer) {
	s.RegisterService(&_Judgements_serviceDesc, srv)
}

func _Judgements_SubmitJudgement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJudgementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).SubmitJudgement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/SubmitJudgement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).SubmitJudgement(ctx, req.(*SubmitJudgementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judgements_FetchHashFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchFileHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).FetchHashFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/FetchHashFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).FetchHashFile(ctx, req.(*FetchFileHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judgements_FetchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchJudgeFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).FetchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/FetchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).FetchFile(ctx, req.(*FetchJudgeFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judgements_FetchJudgement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchJudgementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).FetchJudgement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/FetchJudgement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).FetchJudgement(ctx, req.(*FetchJudgementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judgements_ReturnJudgement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnJudgementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).ReturnJudgement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/ReturnJudgement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).ReturnJudgement(ctx, req.(*ReturnJudgementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judgements_ListJudgements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgementsServer).ListJudgements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Judgements/ListJudgements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgementsServer).ListJudgements(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Judgements_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Judgements",
	HandlerType: (*JudgementsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJudgement",
			Handler:    _Judgements_SubmitJudgement_Handler,
		},
		{
			MethodName: "FetchHashFile",
			Handler:    _Judgements_FetchHashFile_Handler,
		},
		{
			MethodName: "FetchFile",
			Handler:    _Judgements_FetchFile_Handler,
		},
		{
			MethodName: "FetchJudgement",
			Handler:    _Judgements_FetchJudgement_Handler,
		},
		{
			MethodName: "ReturnJudgement",
			Handler:    _Judgements_ReturnJudgement_Handler,
		},
		{
			MethodName: "ListJudgements",
			Handler:    _Judgements_ListJudgements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judgements.proto",
}