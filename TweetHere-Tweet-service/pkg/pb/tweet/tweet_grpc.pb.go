// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/tweet/tweet.proto

package tweet

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TweetService_AddTweet_FullMethodName        = "/tweet.TweetService/AddTweet"
	TweetService_AddTweet2_FullMethodName       = "/tweet.TweetService/AddTweet2"
	TweetService_GetOurTweet_FullMethodName     = "/tweet.TweetService/GetOurTweet"
	TweetService_GetOthersTweet_FullMethodName  = "/tweet.TweetService/GetOthersTweet"
	TweetService_EditTweet_FullMethodName       = "/tweet.TweetService/EditTweet"
	TweetService_DeletePost_FullMethodName      = "/tweet.TweetService/DeletePost"
	TweetService_LikePost_FullMethodName        = "/tweet.TweetService/LikePost"
	TweetService_UnLikePost_FullMethodName      = "/tweet.TweetService/UnLikePost"
	TweetService_SavePost_FullMethodName        = "/tweet.TweetService/SavePost"
	TweetService_UnSavePost_FullMethodName      = "/tweet.TweetService/UnSavePost"
	TweetService_RplyCommentPost_FullMethodName = "/tweet.TweetService/RplyCommentPost"
	TweetService_CommentPost_FullMethodName     = "/tweet.TweetService/CommentPost"
	TweetService_GetComments_FullMethodName     = "/tweet.TweetService/GetComments"
)

// TweetServiceClient is the client API for TweetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TweetServiceClient interface {
	AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*AddTweetResponse, error)
	AddTweet2(ctx context.Context, in *AddTweet2Request, opts ...grpc.CallOption) (*AddTweet2Response, error)
	GetOurTweet(ctx context.Context, in *GetOurTweetRequest, opts ...grpc.CallOption) (*GetOurTweetResponse, error)
	GetOthersTweet(ctx context.Context, in *GetOthersTweetRequest, opts ...grpc.CallOption) (*GetOthersTweetResponse, error)
	EditTweet(ctx context.Context, in *EditTweetRequest, opts ...grpc.CallOption) (*EditTweetResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error)
	UnLikePost(ctx context.Context, in *UnLikePostRequest, opts ...grpc.CallOption) (*UnLikePostResponse, error)
	SavePost(ctx context.Context, in *SavePostRequest, opts ...grpc.CallOption) (*SavePostResponse, error)
	UnSavePost(ctx context.Context, in *UnSavePostRequest, opts ...grpc.CallOption) (*UnSavePostRespone, error)
	RplyCommentPost(ctx context.Context, in *RplyCommentPostRequest, opts ...grpc.CallOption) (*RplyCommentPostResponse, error)
	CommentPost(ctx context.Context, in *CommentPostRequest, opts ...grpc.CallOption) (*CommentPostResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
}

type tweetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetServiceClient(cc grpc.ClientConnInterface) TweetServiceClient {
	return &tweetServiceClient{cc}
}

func (c *tweetServiceClient) AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*AddTweetResponse, error) {
	out := new(AddTweetResponse)
	err := c.cc.Invoke(ctx, TweetService_AddTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) AddTweet2(ctx context.Context, in *AddTweet2Request, opts ...grpc.CallOption) (*AddTweet2Response, error) {
	out := new(AddTweet2Response)
	err := c.cc.Invoke(ctx, TweetService_AddTweet2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) GetOurTweet(ctx context.Context, in *GetOurTweetRequest, opts ...grpc.CallOption) (*GetOurTweetResponse, error) {
	out := new(GetOurTweetResponse)
	err := c.cc.Invoke(ctx, TweetService_GetOurTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) GetOthersTweet(ctx context.Context, in *GetOthersTweetRequest, opts ...grpc.CallOption) (*GetOthersTweetResponse, error) {
	out := new(GetOthersTweetResponse)
	err := c.cc.Invoke(ctx, TweetService_GetOthersTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) EditTweet(ctx context.Context, in *EditTweetRequest, opts ...grpc.CallOption) (*EditTweetResponse, error) {
	out := new(EditTweetResponse)
	err := c.cc.Invoke(ctx, TweetService_EditTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, TweetService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) LikePost(ctx context.Context, in *LikePostRequest, opts ...grpc.CallOption) (*LikePostResponse, error) {
	out := new(LikePostResponse)
	err := c.cc.Invoke(ctx, TweetService_LikePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) UnLikePost(ctx context.Context, in *UnLikePostRequest, opts ...grpc.CallOption) (*UnLikePostResponse, error) {
	out := new(UnLikePostResponse)
	err := c.cc.Invoke(ctx, TweetService_UnLikePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) SavePost(ctx context.Context, in *SavePostRequest, opts ...grpc.CallOption) (*SavePostResponse, error) {
	out := new(SavePostResponse)
	err := c.cc.Invoke(ctx, TweetService_SavePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) UnSavePost(ctx context.Context, in *UnSavePostRequest, opts ...grpc.CallOption) (*UnSavePostRespone, error) {
	out := new(UnSavePostRespone)
	err := c.cc.Invoke(ctx, TweetService_UnSavePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) RplyCommentPost(ctx context.Context, in *RplyCommentPostRequest, opts ...grpc.CallOption) (*RplyCommentPostResponse, error) {
	out := new(RplyCommentPostResponse)
	err := c.cc.Invoke(ctx, TweetService_RplyCommentPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) CommentPost(ctx context.Context, in *CommentPostRequest, opts ...grpc.CallOption) (*CommentPostResponse, error) {
	out := new(CommentPostResponse)
	err := c.cc.Invoke(ctx, TweetService_CommentPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, TweetService_GetComments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetServiceServer is the server API for TweetService service.
// All implementations must embed UnimplementedTweetServiceServer
// for forward compatibility
type TweetServiceServer interface {
	AddTweet(context.Context, *AddTweetRequest) (*AddTweetResponse, error)
	AddTweet2(context.Context, *AddTweet2Request) (*AddTweet2Response, error)
	GetOurTweet(context.Context, *GetOurTweetRequest) (*GetOurTweetResponse, error)
	GetOthersTweet(context.Context, *GetOthersTweetRequest) (*GetOthersTweetResponse, error)
	EditTweet(context.Context, *EditTweetRequest) (*EditTweetResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error)
	UnLikePost(context.Context, *UnLikePostRequest) (*UnLikePostResponse, error)
	SavePost(context.Context, *SavePostRequest) (*SavePostResponse, error)
	UnSavePost(context.Context, *UnSavePostRequest) (*UnSavePostRespone, error)
	RplyCommentPost(context.Context, *RplyCommentPostRequest) (*RplyCommentPostResponse, error)
	CommentPost(context.Context, *CommentPostRequest) (*CommentPostResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	mustEmbedUnimplementedTweetServiceServer()
}

// UnimplementedTweetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTweetServiceServer struct {
}

func (UnimplementedTweetServiceServer) AddTweet(context.Context, *AddTweetRequest) (*AddTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTweet not implemented")
}
func (UnimplementedTweetServiceServer) AddTweet2(context.Context, *AddTweet2Request) (*AddTweet2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTweet2 not implemented")
}
func (UnimplementedTweetServiceServer) GetOurTweet(context.Context, *GetOurTweetRequest) (*GetOurTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOurTweet not implemented")
}
func (UnimplementedTweetServiceServer) GetOthersTweet(context.Context, *GetOthersTweetRequest) (*GetOthersTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOthersTweet not implemented")
}
func (UnimplementedTweetServiceServer) EditTweet(context.Context, *EditTweetRequest) (*EditTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTweet not implemented")
}
func (UnimplementedTweetServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedTweetServiceServer) LikePost(context.Context, *LikePostRequest) (*LikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikePost not implemented")
}
func (UnimplementedTweetServiceServer) UnLikePost(context.Context, *UnLikePostRequest) (*UnLikePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnLikePost not implemented")
}
func (UnimplementedTweetServiceServer) SavePost(context.Context, *SavePostRequest) (*SavePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePost not implemented")
}
func (UnimplementedTweetServiceServer) UnSavePost(context.Context, *UnSavePostRequest) (*UnSavePostRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSavePost not implemented")
}
func (UnimplementedTweetServiceServer) RplyCommentPost(context.Context, *RplyCommentPostRequest) (*RplyCommentPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RplyCommentPost not implemented")
}
func (UnimplementedTweetServiceServer) CommentPost(context.Context, *CommentPostRequest) (*CommentPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentPost not implemented")
}
func (UnimplementedTweetServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedTweetServiceServer) mustEmbedUnimplementedTweetServiceServer() {}

// UnsafeTweetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetServiceServer will
// result in compilation errors.
type UnsafeTweetServiceServer interface {
	mustEmbedUnimplementedTweetServiceServer()
}

func RegisterTweetServiceServer(s grpc.ServiceRegistrar, srv TweetServiceServer) {
	s.RegisterService(&TweetService_ServiceDesc, srv)
}

func _TweetService_AddTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).AddTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_AddTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).AddTweet(ctx, req.(*AddTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_AddTweet2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTweet2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).AddTweet2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_AddTweet2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).AddTweet2(ctx, req.(*AddTweet2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_GetOurTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOurTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).GetOurTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_GetOurTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).GetOurTweet(ctx, req.(*GetOurTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_GetOthersTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOthersTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).GetOthersTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_GetOthersTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).GetOthersTweet(ctx, req.(*GetOthersTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_EditTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).EditTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_EditTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).EditTweet(ctx, req.(*EditTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_LikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).LikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_LikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).LikePost(ctx, req.(*LikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_UnLikePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnLikePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).UnLikePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_UnLikePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).UnLikePost(ctx, req.(*UnLikePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_SavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).SavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_SavePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).SavePost(ctx, req.(*SavePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_UnSavePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSavePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).UnSavePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_UnSavePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).UnSavePost(ctx, req.(*UnSavePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_RplyCommentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RplyCommentPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).RplyCommentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_RplyCommentPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).RplyCommentPost(ctx, req.(*RplyCommentPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_CommentPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).CommentPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_CommentPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).CommentPost(ctx, req.(*CommentPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TweetService_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TweetService_ServiceDesc is the grpc.ServiceDesc for TweetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TweetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tweet.TweetService",
	HandlerType: (*TweetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTweet",
			Handler:    _TweetService_AddTweet_Handler,
		},
		{
			MethodName: "AddTweet2",
			Handler:    _TweetService_AddTweet2_Handler,
		},
		{
			MethodName: "GetOurTweet",
			Handler:    _TweetService_GetOurTweet_Handler,
		},
		{
			MethodName: "GetOthersTweet",
			Handler:    _TweetService_GetOthersTweet_Handler,
		},
		{
			MethodName: "EditTweet",
			Handler:    _TweetService_EditTweet_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _TweetService_DeletePost_Handler,
		},
		{
			MethodName: "LikePost",
			Handler:    _TweetService_LikePost_Handler,
		},
		{
			MethodName: "UnLikePost",
			Handler:    _TweetService_UnLikePost_Handler,
		},
		{
			MethodName: "SavePost",
			Handler:    _TweetService_SavePost_Handler,
		},
		{
			MethodName: "UnSavePost",
			Handler:    _TweetService_UnSavePost_Handler,
		},
		{
			MethodName: "RplyCommentPost",
			Handler:    _TweetService_RplyCommentPost_Handler,
		},
		{
			MethodName: "CommentPost",
			Handler:    _TweetService_CommentPost_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _TweetService_GetComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/tweet/tweet.proto",
}
