// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package authpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestEmailCodeRequest struct {
	// 电子邮箱
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// 操作
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEmailCodeRequest) Reset()         { *m = RequestEmailCodeRequest{} }
func (m *RequestEmailCodeRequest) String() string { return proto.CompactTextString(m) }
func (*RequestEmailCodeRequest) ProtoMessage()    {}
func (*RequestEmailCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *RequestEmailCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEmailCodeRequest.Unmarshal(m, b)
}
func (m *RequestEmailCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEmailCodeRequest.Marshal(b, m, deterministic)
}
func (m *RequestEmailCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEmailCodeRequest.Merge(m, src)
}
func (m *RequestEmailCodeRequest) XXX_Size() int {
	return xxx_messageInfo_RequestEmailCodeRequest.Size(m)
}
func (m *RequestEmailCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEmailCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEmailCodeRequest proto.InternalMessageInfo

func (m *RequestEmailCodeRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RequestEmailCodeRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type RequestEmailCodeResponse struct {
	// success
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestEmailCodeResponse) Reset()         { *m = RequestEmailCodeResponse{} }
func (m *RequestEmailCodeResponse) String() string { return proto.CompactTextString(m) }
func (*RequestEmailCodeResponse) ProtoMessage()    {}
func (*RequestEmailCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *RequestEmailCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEmailCodeResponse.Unmarshal(m, b)
}
func (m *RequestEmailCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEmailCodeResponse.Marshal(b, m, deterministic)
}
func (m *RequestEmailCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEmailCodeResponse.Merge(m, src)
}
func (m *RequestEmailCodeResponse) XXX_Size() int {
	return xxx_messageInfo_RequestEmailCodeResponse.Size(m)
}
func (m *RequestEmailCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEmailCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEmailCodeResponse proto.InternalMessageInfo

func (m *RequestEmailCodeResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type RequestSmSCodeRequest struct {
	// 手机号码
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 操作
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSmSCodeRequest) Reset()         { *m = RequestSmSCodeRequest{} }
func (m *RequestSmSCodeRequest) String() string { return proto.CompactTextString(m) }
func (*RequestSmSCodeRequest) ProtoMessage()    {}
func (*RequestSmSCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *RequestSmSCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSmSCodeRequest.Unmarshal(m, b)
}
func (m *RequestSmSCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSmSCodeRequest.Marshal(b, m, deterministic)
}
func (m *RequestSmSCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSmSCodeRequest.Merge(m, src)
}
func (m *RequestSmSCodeRequest) XXX_Size() int {
	return xxx_messageInfo_RequestSmSCodeRequest.Size(m)
}
func (m *RequestSmSCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSmSCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSmSCodeRequest proto.InternalMessageInfo

func (m *RequestSmSCodeRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *RequestSmSCodeRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type RequestSmSCodeResponse struct {
	// success
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestSmSCodeResponse) Reset()         { *m = RequestSmSCodeResponse{} }
func (m *RequestSmSCodeResponse) String() string { return proto.CompactTextString(m) }
func (*RequestSmSCodeResponse) ProtoMessage()    {}
func (*RequestSmSCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *RequestSmSCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestSmSCodeResponse.Unmarshal(m, b)
}
func (m *RequestSmSCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestSmSCodeResponse.Marshal(b, m, deterministic)
}
func (m *RequestSmSCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSmSCodeResponse.Merge(m, src)
}
func (m *RequestSmSCodeResponse) XXX_Size() int {
	return xxx_messageInfo_RequestSmSCodeResponse.Size(m)
}
func (m *RequestSmSCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSmSCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSmSCodeResponse proto.InternalMessageInfo

func (m *RequestSmSCodeResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*RequestEmailCodeRequest)(nil), "auth.RequestEmailCodeRequest")
	proto.RegisterType((*RequestEmailCodeResponse)(nil), "auth.RequestEmailCodeResponse")
	proto.RegisterType((*RequestSmSCodeRequest)(nil), "auth.RequestSmSCodeRequest")
	proto.RegisterType((*RequestSmSCodeResponse)(nil), "auth.RequestSmSCodeResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xdc, 0xb9, 0xc4, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x5c, 0x73, 0x13, 0x33, 0x73, 0x9c, 0xf3, 0x53, 0x52, 0xa1, 0x7c,
	0x21, 0x11, 0x2e, 0xd6, 0x54, 0x90, 0x98, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23,
	0x24, 0xc6, 0xc5, 0x96, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0x04, 0x16, 0x86, 0xf2, 0x94,
	0xb4, 0xb8, 0x24, 0x30, 0x0d, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xe2, 0xe3, 0x62, 0xca,
	0xcf, 0x06, 0x1b, 0xc3, 0x11, 0xc4, 0x94, 0x9f, 0xad, 0xe4, 0xce, 0x25, 0x0a, 0x55, 0x1b, 0x9c,
	0x1b, 0x8c, 0x6c, 0xa5, 0x18, 0x17, 0x5b, 0x6e, 0x7e, 0x52, 0x66, 0x4e, 0x2a, 0xd4, 0x4e, 0x28,
	0x0f, 0xa7, 0xa5, 0x1a, 0x5c, 0x62, 0xe8, 0x06, 0x61, 0xb7, 0xd2, 0x68, 0x19, 0x23, 0x17, 0x8b,
	0x63, 0x69, 0x49, 0x86, 0x50, 0x20, 0x97, 0x00, 0xba, 0x3b, 0x85, 0x64, 0xf5, 0xc0, 0xe1, 0x82,
	0x23, 0x20, 0xa4, 0xe4, 0x70, 0x49, 0x43, 0xed, 0xf2, 0xe6, 0xe2, 0x43, 0x75, 0x85, 0x90, 0x34,
	0x8a, 0x0e, 0x54, 0x4f, 0x4a, 0xc9, 0x60, 0x97, 0x84, 0x18, 0xe6, 0xc4, 0x11, 0xc5, 0x06, 0x92,
	0x2e, 0x48, 0x4a, 0x62, 0x03, 0xc7, 0x93, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x52, 0x3d, 0x13,
	0xb5, 0xb5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	// 请求电子邮箱验证码
	RequestEmailCode(ctx context.Context, in *RequestEmailCodeRequest, opts ...grpc.CallOption) (*RequestEmailCodeResponse, error)
	// 请求短信验证码
	RequestSmSCode(ctx context.Context, in *RequestSmSCodeRequest, opts ...grpc.CallOption) (*RequestSmSCodeResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) RequestEmailCode(ctx context.Context, in *RequestEmailCodeRequest, opts ...grpc.CallOption) (*RequestEmailCodeResponse, error) {
	out := new(RequestEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RequestEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RequestSmSCode(ctx context.Context, in *RequestSmSCodeRequest, opts ...grpc.CallOption) (*RequestSmSCodeResponse, error) {
	out := new(RequestSmSCodeResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RequestSmSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	// 请求电子邮箱验证码
	RequestEmailCode(context.Context, *RequestEmailCodeRequest) (*RequestEmailCodeResponse, error)
	// 请求短信验证码
	RequestSmSCode(context.Context, *RequestSmSCodeRequest) (*RequestSmSCodeResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_RequestEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RequestEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RequestEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RequestEmailCode(ctx, req.(*RequestEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RequestSmSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSmSCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RequestSmSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RequestSmSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RequestSmSCode(ctx, req.(*RequestSmSCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestEmailCode",
			Handler:    _Auth_RequestEmailCode_Handler,
		},
		{
			MethodName: "RequestSmSCode",
			Handler:    _Auth_RequestSmSCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
