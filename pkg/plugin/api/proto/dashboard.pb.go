// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/dashboard.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type KeyRequest struct {
	Namespace            string               `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ApiVersion           string               `protobuf:"bytes,2,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	Kind                 string               `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	LabelSelector        *wrappers.BytesValue `protobuf:"bytes,5,opt,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{1}
}
func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRequest.Unmarshal(m, b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
}
func (dst *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(dst, src)
}
func (m *KeyRequest) XXX_Size() int {
	return xxx_messageInfo_KeyRequest.Size(m)
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KeyRequest) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *KeyRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *KeyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyRequest) GetLabelSelector() *wrappers.BytesValue {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

type ListResponse struct {
	Objects              [][]byte `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{2}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetObjects() [][]byte {
	if m != nil {
		return m.Objects
	}
	return nil
}

type GetResponse struct {
	Object               []byte   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{3}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

type PortForwardRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName              string   `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
	ContainerName        string   `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	PortNumber           uint32   `protobuf:"varint,4,opt,name=portNumber,proto3" json:"portNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortForwardRequest) Reset()         { *m = PortForwardRequest{} }
func (m *PortForwardRequest) String() string { return proto.CompactTextString(m) }
func (*PortForwardRequest) ProtoMessage()    {}
func (*PortForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{4}
}
func (m *PortForwardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortForwardRequest.Unmarshal(m, b)
}
func (m *PortForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortForwardRequest.Marshal(b, m, deterministic)
}
func (dst *PortForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardRequest.Merge(dst, src)
}
func (m *PortForwardRequest) XXX_Size() int {
	return xxx_messageInfo_PortForwardRequest.Size(m)
}
func (m *PortForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardRequest proto.InternalMessageInfo

func (m *PortForwardRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PortForwardRequest) GetPodName() string {
	if m != nil {
		return m.PodName
	}
	return ""
}

func (m *PortForwardRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *PortForwardRequest) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type PortForwardResponse struct {
	PortForwardID        string   `protobuf:"bytes,1,opt,name=portForwardID,proto3" json:"portForwardID,omitempty"`
	PortNumber           uint32   `protobuf:"varint,2,opt,name=portNumber,proto3" json:"portNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortForwardResponse) Reset()         { *m = PortForwardResponse{} }
func (m *PortForwardResponse) String() string { return proto.CompactTextString(m) }
func (*PortForwardResponse) ProtoMessage()    {}
func (*PortForwardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{5}
}
func (m *PortForwardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortForwardResponse.Unmarshal(m, b)
}
func (m *PortForwardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortForwardResponse.Marshal(b, m, deterministic)
}
func (dst *PortForwardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardResponse.Merge(dst, src)
}
func (m *PortForwardResponse) XXX_Size() int {
	return xxx_messageInfo_PortForwardResponse.Size(m)
}
func (m *PortForwardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardResponse proto.InternalMessageInfo

func (m *PortForwardResponse) GetPortForwardID() string {
	if m != nil {
		return m.PortForwardID
	}
	return ""
}

func (m *PortForwardResponse) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type CancelPortForwardRequest struct {
	PortForwardID        string   `protobuf:"bytes,1,opt,name=portForwardID,proto3" json:"portForwardID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelPortForwardRequest) Reset()         { *m = CancelPortForwardRequest{} }
func (m *CancelPortForwardRequest) String() string { return proto.CompactTextString(m) }
func (*CancelPortForwardRequest) ProtoMessage()    {}
func (*CancelPortForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_456106977de738a3, []int{6}
}
func (m *CancelPortForwardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelPortForwardRequest.Unmarshal(m, b)
}
func (m *CancelPortForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelPortForwardRequest.Marshal(b, m, deterministic)
}
func (dst *CancelPortForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelPortForwardRequest.Merge(dst, src)
}
func (m *CancelPortForwardRequest) XXX_Size() int {
	return xxx_messageInfo_CancelPortForwardRequest.Size(m)
}
func (m *CancelPortForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelPortForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelPortForwardRequest proto.InternalMessageInfo

func (m *CancelPortForwardRequest) GetPortForwardID() string {
	if m != nil {
		return m.PortForwardID
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*KeyRequest)(nil), "proto.KeyRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*GetResponse)(nil), "proto.GetResponse")
	proto.RegisterType((*PortForwardRequest)(nil), "proto.PortForwardRequest")
	proto.RegisterType((*PortForwardResponse)(nil), "proto.PortForwardResponse")
	proto.RegisterType((*CancelPortForwardRequest)(nil), "proto.CancelPortForwardRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DashboardClient is the client API for Dashboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DashboardClient interface {
	List(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResponse, error)
	PortForward(ctx context.Context, in *PortForwardRequest, opts ...grpc.CallOption) (*PortForwardResponse, error)
	CancelPortForward(ctx context.Context, in *CancelPortForwardRequest, opts ...grpc.CallOption) (*Empty, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) List(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) PortForward(ctx context.Context, in *PortForwardRequest, opts ...grpc.CallOption) (*PortForwardResponse, error) {
	out := new(PortForwardResponse)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/PortForward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) CancelPortForward(ctx context.Context, in *CancelPortForwardRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Dashboard/CancelPortForward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DashboardServer is the server API for Dashboard service.
type DashboardServer interface {
	List(context.Context, *KeyRequest) (*ListResponse, error)
	Get(context.Context, *KeyRequest) (*GetResponse, error)
	PortForward(context.Context, *PortForwardRequest) (*PortForwardResponse, error)
	CancelPortForward(context.Context, *CancelPortForwardRequest) (*Empty, error)
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).List(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_PortForward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).PortForward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/PortForward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).PortForward(ctx, req.(*PortForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dashboard_CancelPortForward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPortForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).CancelPortForward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Dashboard/CancelPortForward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).CancelPortForward(ctx, req.(*CancelPortForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Dashboard_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Dashboard_Get_Handler,
		},
		{
			MethodName: "PortForward",
			Handler:    _Dashboard_PortForward_Handler,
		},
		{
			MethodName: "CancelPortForward",
			Handler:    _Dashboard_CancelPortForward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dashboard.proto",
}

func init() { proto.RegisterFile("proto/dashboard.proto", fileDescriptor_dashboard_456106977de738a3) }

var fileDescriptor_dashboard_456106977de738a3 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0xdb, 0x76, 0x4b, 0x6f, 0x93, 0x87, 0x9d, 0x45, 0x19, 0xa3, 0xac, 0x25, 0xac, 0x90,
	0x07, 0x49, 0x61, 0xfd, 0x01, 0x5d, 0xab, 0x8b, 0x28, 0x8b, 0x44, 0xe8, 0x8b, 0x4f, 0x93, 0xe4,
	0x5a, 0xa3, 0x69, 0x66, 0x9c, 0x99, 0x50, 0xfa, 0x1b, 0xfe, 0x8b, 0x5f, 0xe6, 0x0f, 0x48, 0x26,
	0x13, 0x92, 0xd8, 0x2e, 0xf4, 0xa9, 0x73, 0xce, 0xbd, 0x73, 0x7b, 0xce, 0xb9, 0x13, 0x78, 0x24,
	0x24, 0xd7, 0x7c, 0x99, 0x31, 0xf5, 0x3d, 0xe1, 0x4c, 0x66, 0x91, 0xc1, 0x64, 0x62, 0x7e, 0xfc,
	0xab, 0x0d, 0xe7, 0x9b, 0x02, 0x97, 0x06, 0x25, 0xd5, 0xb7, 0xe5, 0x4e, 0x32, 0x21, 0x50, 0xaa,
	0xa6, 0x2d, 0x98, 0xc2, 0xe4, 0xdd, 0x56, 0xe8, 0x7d, 0xf0, 0xc7, 0x01, 0xf8, 0x88, 0xfb, 0x18,
	0x7f, 0x55, 0xa8, 0x34, 0x79, 0x06, 0xb3, 0x92, 0x6d, 0x51, 0x09, 0x96, 0x22, 0x75, 0x16, 0x4e,
	0x38, 0x8b, 0x3b, 0x82, 0x5c, 0x01, 0x30, 0x91, 0xaf, 0x51, 0xaa, 0x9c, 0x97, 0xf4, 0xcc, 0x94,
	0x7b, 0x0c, 0x21, 0x30, 0xfe, 0x99, 0x97, 0x19, 0x1d, 0x99, 0x8a, 0x39, 0xd7, 0x5c, 0x3d, 0x80,
	0x8e, 0x1b, 0xae, 0x3e, 0x93, 0x37, 0xe0, 0x15, 0x2c, 0xc1, 0xe2, 0x0b, 0x16, 0x98, 0x6a, 0x2e,
	0xe9, 0x64, 0xe1, 0x84, 0xf3, 0x9b, 0xa7, 0x51, 0xa3, 0x3a, 0x6a, 0x55, 0x47, 0xb7, 0x7b, 0x8d,
	0x6a, 0xcd, 0x8a, 0x0a, 0xe3, 0xe1, 0x8d, 0x20, 0x04, 0xf7, 0x53, 0xae, 0x74, 0x8c, 0x4a, 0xf0,
	0x52, 0x21, 0xa1, 0x30, 0xe5, 0xc9, 0x0f, 0x4c, 0xb5, 0xa2, 0xce, 0x62, 0x14, 0xba, 0x71, 0x0b,
	0x83, 0x17, 0x30, 0xbf, 0xc3, 0xae, 0xf1, 0x31, 0x9c, 0x37, 0x15, 0x63, 0xcf, 0x8d, 0x2d, 0x0a,
	0x7e, 0x3b, 0x40, 0x3e, 0x73, 0xa9, 0xdf, 0x73, 0xb9, 0x63, 0x32, 0x3b, 0x2d, 0x10, 0x0a, 0x53,
	0xc1, 0xb3, 0xfb, 0xda, 0x5f, 0x93, 0x46, 0x0b, 0xc9, 0x35, 0x78, 0x29, 0x2f, 0x35, 0xcb, 0x4b,
	0x94, 0xa6, 0xde, 0x64, 0x32, 0x24, 0xeb, 0x40, 0x05, 0x97, 0xfa, 0xbe, 0xda, 0x26, 0x28, 0x4d,
	0x44, 0x5e, 0xdc, 0x63, 0x82, 0xaf, 0x70, 0x39, 0xd0, 0x64, 0x3d, 0x5c, 0x83, 0x27, 0x3a, 0xfa,
	0xc3, 0xca, 0x0a, 0x1b, 0x92, 0xff, 0x0d, 0x3f, 0x3b, 0x18, 0xfe, 0x1a, 0xe8, 0x5b, 0x56, 0xa6,
	0x58, 0x1c, 0xb1, 0x7d, 0xd2, 0x3f, 0xdc, 0xfc, 0x75, 0x60, 0xb6, 0x6a, 0x1f, 0x20, 0x89, 0x60,
	0x5c, 0xaf, 0x84, 0x5c, 0x34, 0xfb, 0x8b, 0xba, 0x67, 0xe5, 0x5f, 0x5a, 0x6a, 0xb0, 0xb2, 0x97,
	0x30, 0xba, 0xc3, 0xa3, 0xed, 0xc4, 0x52, 0xfd, 0xbd, 0xad, 0x60, 0xde, 0xd3, 0x49, 0x9e, 0xd8,
	0x96, 0x43, 0xed, 0xbe, 0x7f, 0xac, 0x64, 0xa7, 0xdc, 0xc2, 0xc5, 0x81, 0x67, 0xf2, 0xdc, 0x5e,
	0x78, 0x28, 0x0d, 0xdf, 0xb5, 0x0d, 0xe6, 0x93, 0x49, 0xce, 0x0d, 0x78, 0xf5, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x9d, 0x20, 0xf8, 0x82, 0x03, 0x00, 0x00,
}
