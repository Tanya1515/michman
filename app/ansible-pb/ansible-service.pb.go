// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ansible-service.proto

package ansible

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ClusterDataRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterDataRequest) Reset()         { *m = ClusterDataRequest{} }
func (m *ClusterDataRequest) String() string { return proto.CompactTextString(m) }
func (*ClusterDataRequest) ProtoMessage()    {}
func (*ClusterDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a650c782ccb80b8, []int{0}
}

func (m *ClusterDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDataRequest.Unmarshal(m, b)
}
func (m *ClusterDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDataRequest.Marshal(b, m, deterministic)
}
func (m *ClusterDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDataRequest.Merge(m, src)
}
func (m *ClusterDataRequest) XXX_Size() int {
	return xxx_messageInfo_ClusterDataRequest.Size(m)
}
func (m *ClusterDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDataRequest proto.InternalMessageInfo

func (m *ClusterDataRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterDataRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ClusterReply struct {
	State                string   `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
	Template             string   `protobuf:"bytes,2,opt,name=Template,proto3" json:"Template,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterReply) Reset()         { *m = ClusterReply{} }
func (m *ClusterReply) String() string { return proto.CompactTextString(m) }
func (*ClusterReply) ProtoMessage()    {}
func (*ClusterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a650c782ccb80b8, []int{1}
}

func (m *ClusterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterReply.Unmarshal(m, b)
}
func (m *ClusterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterReply.Marshal(b, m, deterministic)
}
func (m *ClusterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterReply.Merge(m, src)
}
func (m *ClusterReply) XXX_Size() int {
	return xxx_messageInfo_ClusterReply.Size(m)
}
func (m *ClusterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterReply.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterReply proto.InternalMessageInfo

func (m *ClusterReply) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ClusterReply) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

type OkMessage struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OkMessage) Reset()         { *m = OkMessage{} }
func (m *OkMessage) String() string { return proto.CompactTextString(m) }
func (*OkMessage) ProtoMessage()    {}
func (*OkMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a650c782ccb80b8, []int{2}
}

func (m *OkMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkMessage.Unmarshal(m, b)
}
func (m *OkMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkMessage.Marshal(b, m, deterministic)
}
func (m *OkMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkMessage.Merge(m, src)
}
func (m *OkMessage) XXX_Size() int {
	return xxx_messageInfo_OkMessage.Size(m)
}
func (m *OkMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_OkMessage.DiscardUnknown(m)
}

var xxx_messageInfo_OkMessage proto.InternalMessageInfo

func (m *OkMessage) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*ClusterDataRequest)(nil), "ansible.ClusterDataRequest")
	proto.RegisterType((*ClusterReply)(nil), "ansible.ClusterReply")
	proto.RegisterType((*OkMessage)(nil), "ansible.okMessage")
}

func init() { proto.RegisterFile("ansible-service.proto", fileDescriptor_7a650c782ccb80b8) }

var fileDescriptor_7a650c782ccb80b8 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x14, 0x84, 0x13, 0xd1, 0xc4, 0x3c, 0xb4, 0x79, 0x18, 0x09, 0xb1, 0x91, 0xb5, 0xb1, 0xf1, 0x8a,
	0xd8, 0x5a, 0x08, 0x51, 0x52, 0x19, 0xe4, 0x92, 0xd6, 0xe2, 0x25, 0x8e, 0xb2, 0x78, 0xc9, 0x9e,
	0xbb, 0xef, 0x0e, 0xee, 0xdf, 0xcb, 0xdd, 0xad, 0x57, 0x28, 0xa4, 0xdb, 0x99, 0xe1, 0x63, 0x76,
	0x1e, 0x8d, 0x65, 0x1f, 0xec, 0x26, 0xc3, 0x5d, 0x80, 0x2f, 0xed, 0x16, 0x49, 0xee, 0x9d, 0x3a,
	0x1e, 0x46, 0xdb, 0x3c, 0x10, 0xcf, 0xb3, 0x22, 0x28, 0xfc, 0x93, 0xa8, 0xa4, 0xf8, 0x2e, 0x10,
	0x94, 0x99, 0x8e, 0x97, 0xb2, 0xc3, 0xa4, 0x7f, 0xdd, 0xbf, 0x1d, 0xa5, 0xcd, 0xbb, 0xf6, 0xd6,
	0x55, 0x8e, 0xc9, 0x51, 0xeb, 0xd5, 0x6f, 0xf3, 0x48, 0x67, 0x91, 0x4e, 0x91, 0x67, 0x15, 0x5f,
	0xd0, 0xc9, 0x4a, 0x45, 0x7f, 0xc1, 0x56, 0xf0, 0x94, 0x4e, 0xd7, 0xd8, 0xe5, 0x59, 0x1d, 0xb4,
	0x74, 0xa7, 0xcd, 0x0d, 0x8d, 0xdc, 0xd7, 0x0b, 0x42, 0x90, 0x4f, 0xf0, 0x25, 0x0d, 0x6a, 0xa2,
	0x08, 0x91, 0x8f, 0x6a, 0xf6, 0x4a, 0xc3, 0x85, 0x07, 0x14, 0x9e, 0x9f, 0xe9, 0x7c, 0xee, 0x21,
	0x8a, 0xd8, 0xcb, 0x57, 0x49, 0x9c, 0x92, 0xfc, 0xdf, 0x31, 0x1d, 0xff, 0x0d, 0x9b, 0x6f, 0x9a,
	0xde, 0xec, 0x8d, 0x78, 0x05, 0x5f, 0xc2, 0x2f, 0x9d, 0xda, 0x0f, 0xbb, 0x15, 0xb5, 0x6e, 0xcf,
	0x8b, 0xee, 0x18, 0x29, 0xe4, 0xbd, 0x6a, 0xdb, 0x0f, 0x37, 0x70, 0x17, 0x76, 0x33, 0x4c, 0x6f,
	0x33, 0x68, 0xae, 0x7c, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x32, 0x75, 0x7c, 0x10, 0x7e, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends request to start cluster creation
	CreateCluster(ctx context.Context, in *ClusterDataRequest, opts ...grpc.CallOption) (*ClusterReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) CreateCluster(ctx context.Context, in *ClusterDataRequest, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := c.cc.Invoke(ctx, "/ansible.Greeter/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends request to start cluster creation
	CreateCluster(context.Context, *ClusterDataRequest) (*ClusterReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) CreateCluster(ctx context.Context, req *ClusterDataRequest) (*ClusterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ansible.Greeter/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateCluster(ctx, req.(*ClusterDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ansible.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCluster",
			Handler:    _Greeter_CreateCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ansible-service.proto",
}

// ServerNotificationClient is the client API for ServerNotification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerNotificationClient interface {
	// Sends request after cluster creation ends
	ClusterReadyStatus(ctx context.Context, in *ClusterDataRequest, opts ...grpc.CallOption) (*OkMessage, error)
}

type serverNotificationClient struct {
	cc *grpc.ClientConn
}

func NewServerNotificationClient(cc *grpc.ClientConn) ServerNotificationClient {
	return &serverNotificationClient{cc}
}

func (c *serverNotificationClient) ClusterReadyStatus(ctx context.Context, in *ClusterDataRequest, opts ...grpc.CallOption) (*OkMessage, error) {
	out := new(OkMessage)
	err := c.cc.Invoke(ctx, "/ansible.ServerNotification/ClusterReadyStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerNotificationServer is the server API for ServerNotification service.
type ServerNotificationServer interface {
	// Sends request after cluster creation ends
	ClusterReadyStatus(context.Context, *ClusterDataRequest) (*OkMessage, error)
}

// UnimplementedServerNotificationServer can be embedded to have forward compatible implementations.
type UnimplementedServerNotificationServer struct {
}

func (*UnimplementedServerNotificationServer) ClusterReadyStatus(ctx context.Context, req *ClusterDataRequest) (*OkMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterReadyStatus not implemented")
}

func RegisterServerNotificationServer(s *grpc.Server, srv ServerNotificationServer) {
	s.RegisterService(&_ServerNotification_serviceDesc, srv)
}

func _ServerNotification_ClusterReadyStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNotificationServer).ClusterReadyStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ansible.ServerNotification/ClusterReadyStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNotificationServer).ClusterReadyStatus(ctx, req.(*ClusterDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerNotification_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ansible.ServerNotification",
	HandlerType: (*ServerNotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClusterReadyStatus",
			Handler:    _ServerNotification_ClusterReadyStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ansible-service.proto",
}