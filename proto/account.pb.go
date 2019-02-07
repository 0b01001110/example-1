// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package proto // import "github.com/dogmatiq/example/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OpenAccountRequest is the request initiated by the client to open an account.
type OpenAccountRequest struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenAccountRequest) Reset()         { *m = OpenAccountRequest{} }
func (m *OpenAccountRequest) String() string { return proto.CompactTextString(m) }
func (*OpenAccountRequest) ProtoMessage()    {}
func (*OpenAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_e0feec780f121f4c, []int{0}
}
func (m *OpenAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenAccountRequest.Unmarshal(m, b)
}
func (m *OpenAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenAccountRequest.Marshal(b, m, deterministic)
}
func (dst *OpenAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenAccountRequest.Merge(dst, src)
}
func (m *OpenAccountRequest) XXX_Size() int {
	return xxx_messageInfo_OpenAccountRequest.Size(m)
}
func (m *OpenAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenAccountRequest proto.InternalMessageInfo

func (m *OpenAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *OpenAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// OpenAccountRequest is the request initiated by the client to open an account.
type OpenAccountResponse struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenAccountResponse) Reset()         { *m = OpenAccountResponse{} }
func (m *OpenAccountResponse) String() string { return proto.CompactTextString(m) }
func (*OpenAccountResponse) ProtoMessage()    {}
func (*OpenAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_e0feec780f121f4c, []int{1}
}
func (m *OpenAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenAccountResponse.Unmarshal(m, b)
}
func (m *OpenAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenAccountResponse.Marshal(b, m, deterministic)
}
func (dst *OpenAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenAccountResponse.Merge(dst, src)
}
func (m *OpenAccountResponse) XXX_Size() int {
	return xxx_messageInfo_OpenAccountResponse.Size(m)
}
func (m *OpenAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenAccountResponse proto.InternalMessageInfo

func (m *OpenAccountResponse) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *OpenAccountResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestStreamingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestStreamingRequest) Reset()         { *m = TestStreamingRequest{} }
func (m *TestStreamingRequest) String() string { return proto.CompactTextString(m) }
func (*TestStreamingRequest) ProtoMessage()    {}
func (*TestStreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_e0feec780f121f4c, []int{2}
}
func (m *TestStreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestStreamingRequest.Unmarshal(m, b)
}
func (m *TestStreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestStreamingRequest.Marshal(b, m, deterministic)
}
func (dst *TestStreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestStreamingRequest.Merge(dst, src)
}
func (m *TestStreamingRequest) XXX_Size() int {
	return xxx_messageInfo_TestStreamingRequest.Size(m)
}
func (m *TestStreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestStreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestStreamingRequest proto.InternalMessageInfo

type TestStreamingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestStreamingResponse) Reset()         { *m = TestStreamingResponse{} }
func (m *TestStreamingResponse) String() string { return proto.CompactTextString(m) }
func (*TestStreamingResponse) ProtoMessage()    {}
func (*TestStreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_e0feec780f121f4c, []int{3}
}
func (m *TestStreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestStreamingResponse.Unmarshal(m, b)
}
func (m *TestStreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestStreamingResponse.Marshal(b, m, deterministic)
}
func (dst *TestStreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestStreamingResponse.Merge(dst, src)
}
func (m *TestStreamingResponse) XXX_Size() int {
	return xxx_messageInfo_TestStreamingResponse.Size(m)
}
func (m *TestStreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestStreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestStreamingResponse proto.InternalMessageInfo

func (m *TestStreamingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenAccountRequest)(nil), "proto.OpenAccountRequest")
	proto.RegisterType((*OpenAccountResponse)(nil), "proto.OpenAccountResponse")
	proto.RegisterType((*TestStreamingRequest)(nil), "proto.TestStreamingRequest")
	proto.RegisterType((*TestStreamingResponse)(nil), "proto.TestStreamingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	// OpenAccount is a service handler to process OpenAccountRequest.
	OpenAccount(ctx context.Context, in *OpenAccountRequest, opts ...grpc.CallOption) (*OpenAccountResponse, error)
	TestStreaming(ctx context.Context, in *TestStreamingRequest, opts ...grpc.CallOption) (Account_TestStreamingClient, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) OpenAccount(ctx context.Context, in *OpenAccountRequest, opts ...grpc.CallOption) (*OpenAccountResponse, error) {
	out := new(OpenAccountResponse)
	err := c.cc.Invoke(ctx, "/proto.Account/OpenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) TestStreaming(ctx context.Context, in *TestStreamingRequest, opts ...grpc.CallOption) (Account_TestStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Account_serviceDesc.Streams[0], "/proto.Account/TestStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountTestStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Account_TestStreamingClient interface {
	Recv() (*TestStreamingResponse, error)
	grpc.ClientStream
}

type accountTestStreamingClient struct {
	grpc.ClientStream
}

func (x *accountTestStreamingClient) Recv() (*TestStreamingResponse, error) {
	m := new(TestStreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	// OpenAccount is a service handler to process OpenAccountRequest.
	OpenAccount(context.Context, *OpenAccountRequest) (*OpenAccountResponse, error)
	TestStreaming(*TestStreamingRequest, Account_TestStreamingServer) error
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_OpenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).OpenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Account/OpenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).OpenAccount(ctx, req.(*OpenAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_TestStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestStreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountServer).TestStreaming(m, &accountTestStreamingServer{stream})
}

type Account_TestStreamingServer interface {
	Send(*TestStreamingResponse) error
	grpc.ServerStream
}

type accountTestStreamingServer struct {
	grpc.ServerStream
}

func (x *accountTestStreamingServer) Send(m *TestStreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenAccount",
			Handler:    _Account_OpenAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStreaming",
			Handler:       _Account_TestStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/account.proto",
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_account_e0feec780f121f4c) }

var fileDescriptor_account_e0feec780f121f4c = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x92, 0x3b, 0x97, 0x90, 0x7f, 0x41, 0x6a, 0x9e, 0x23, 0x44, 0x2e, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x48, 0x96, 0x8b, 0x0b, 0xaa, 0x3a, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x88, 0x13, 0x2a, 0xe2, 0x99, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0x04, 0x96, 0x00, 0xb3, 0x95, 0x3c, 0xb8, 0x84, 0x51, 0x0c, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e,
	0x25, 0xc7, 0x24, 0x31, 0x2e, 0x91, 0x90, 0xd4, 0xe2, 0x92, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0xdc,
	0xcc, 0xbc, 0x74, 0xa8, 0xa3, 0x94, 0x0c, 0xb9, 0x44, 0xd1, 0xc4, 0xa1, 0x76, 0x48, 0x70, 0xb1,
	0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x2d, 0x80, 0x71, 0x8d, 0x16, 0x32, 0x72, 0xb1,
	0x43, 0x5d, 0x24, 0xe4, 0xc6, 0xc5, 0x8d, 0xe4, 0x40, 0x21, 0x49, 0x48, 0x38, 0xe8, 0x61, 0xfa,
	0x5e, 0x4a, 0x0a, 0x9b, 0x14, 0xc4, 0x2e, 0x25, 0x06, 0x21, 0x3f, 0x2e, 0x5e, 0x14, 0x67, 0x08,
	0x49, 0x43, 0x95, 0x63, 0x73, 0xb4, 0x94, 0x0c, 0x76, 0x49, 0x98, 0x69, 0x06, 0x8c, 0x4e, 0xca,
	0x51, 0x8a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x29, 0xf9, 0xe9,
	0xb9, 0x89, 0x25, 0x99, 0x85, 0xfa, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xfa, 0x60, 0xed,
	0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xde, 0xbf, 0xf4, 0xa9, 0xcb, 0x01,
	0x00, 0x00,
}
