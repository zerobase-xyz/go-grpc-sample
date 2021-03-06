// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hostname.proto

package grpc_health_v1

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_412c89e20cb3d4bd, []int{3, 0}
}

type PodHostnameResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PodHostnameResponse) Reset()         { *m = PodHostnameResponse{} }
func (m *PodHostnameResponse) String() string { return proto.CompactTextString(m) }
func (*PodHostnameResponse) ProtoMessage()    {}
func (*PodHostnameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_412c89e20cb3d4bd, []int{0}
}

func (m *PodHostnameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodHostnameResponse.Unmarshal(m, b)
}
func (m *PodHostnameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodHostnameResponse.Marshal(b, m, deterministic)
}
func (m *PodHostnameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodHostnameResponse.Merge(m, src)
}
func (m *PodHostnameResponse) XXX_Size() int {
	return xxx_messageInfo_PodHostnameResponse.Size(m)
}
func (m *PodHostnameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PodHostnameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PodHostnameResponse proto.InternalMessageInfo

func (m *PodHostnameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_412c89e20cb3d4bd, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_412c89e20cb3d4bd, []int{2}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=grpc.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_412c89e20cb3d4bd, []int{3}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("grpc.health.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*PodHostnameResponse)(nil), "grpc.health.v1.PodHostnameResponse")
	proto.RegisterType((*Empty)(nil), "grpc.health.v1.Empty")
	proto.RegisterType((*HealthCheckRequest)(nil), "grpc.health.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "grpc.health.v1.HealthCheckResponse")
}

func init() { proto.RegisterFile("hostname.proto", fileDescriptor_412c89e20cb3d4bd) }

var fileDescriptor_412c89e20cb3d4bd = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x6c, 0xc4, 0x26, 0xf8, 0x8a, 0xb1, 0xbc, 0x22, 0x14, 0x4f, 0xb2, 0x5e, 0xf4, 0xb2, 0xd0,
	0x78, 0xf3, 0x2a, 0xa5, 0x2d, 0xc2, 0x36, 0x6c, 0xfc, 0x00, 0x2f, 0x12, 0x93, 0x47, 0x23, 0xda,
	0x6c, 0xcc, 0x6e, 0x0b, 0xfe, 0x17, 0x7f, 0xac, 0x64, 0x9b, 0x80, 0xa9, 0xc1, 0xdb, 0xce, 0xdb,
	0x9d, 0xd9, 0x99, 0x79, 0xe0, 0x67, 0x4a, 0x9b, 0x3c, 0x5e, 0x13, 0x2f, 0x4a, 0x65, 0x14, 0xfa,
	0xab, 0xb2, 0x48, 0x78, 0x46, 0xf1, 0x87, 0xc9, 0xf8, 0x76, 0xc2, 0xae, 0x60, 0x14, 0xaa, 0x74,
	0x5e, 0x3f, 0x92, 0xa4, 0x0b, 0x95, 0x6b, 0x42, 0x84, 0xc3, 0x0a, 0x8f, 0x9d, 0x73, 0xe7, 0xf2,
	0x48, 0xda, 0x33, 0xf3, 0xa0, 0x3f, 0x5d, 0x17, 0xe6, 0x8b, 0x71, 0xc0, 0xb9, 0x15, 0xb8, 0xcd,
	0x28, 0x79, 0x97, 0xf4, 0xb9, 0x21, 0x6d, 0x70, 0x0c, 0x9e, 0xa6, 0x72, 0xfb, 0x96, 0x34, 0xac,
	0x06, 0xb2, 0x6f, 0x07, 0x46, 0x2d, 0x42, 0xfd, 0xc9, 0x02, 0x5c, 0x6d, 0x62, 0xb3, 0xd1, 0x96,
	0xe0, 0x07, 0x13, 0xde, 0x36, 0xc7, 0x3b, 0x48, 0x3c, 0xaa, 0x44, 0xf3, 0x55, 0x64, 0x89, 0xb2,
	0x16, 0x60, 0x37, 0x70, 0xdc, 0xba, 0xc0, 0x01, 0x78, 0x0f, 0xe2, 0x4e, 0x2c, 0x9f, 0xc4, 0xb0,
	0x57, 0x81, 0x68, 0x2a, 0x1f, 0x17, 0x62, 0x36, 0x74, 0xf0, 0x04, 0x06, 0x62, 0x79, 0xff, 0xd2,
	0x0c, 0x0e, 0x82, 0x14, 0xb0, 0xc9, 0x1f, 0xaa, 0x34, 0xda, 0x99, 0x46, 0x01, 0xfe, 0x8c, 0xcc,
	0xaf, 0x6e, 0xf0, 0x74, 0xdf, 0x9e, 0x6d, 0xe3, 0xec, 0x62, 0x7f, 0xdc, 0xd1, 0x27, 0xeb, 0x05,
	0xcf, 0xe0, 0xee, 0xe2, 0x60, 0x08, 0x7d, 0x1b, 0x09, 0xd9, 0xbf, 0x79, 0x6d, 0xab, 0x7f, 0xd5,
	0x3b, 0x3a, 0x79, 0x75, 0xed, 0x6e, 0xaf, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x60, 0xb5,
	0x6d, 0xed, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HostnamePodServiceClient is the client API for HostnamePodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostnamePodServiceClient interface {
	GetPodHostname(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodHostnameResponse, error)
}

type hostnamePodServiceClient struct {
	cc *grpc.ClientConn
}

func NewHostnamePodServiceClient(cc *grpc.ClientConn) HostnamePodServiceClient {
	return &hostnamePodServiceClient{cc}
}

func (c *hostnamePodServiceClient) GetPodHostname(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PodHostnameResponse, error) {
	out := new(PodHostnameResponse)
	err := c.cc.Invoke(ctx, "/grpc.health.v1.HostnamePodService/GetPodHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostnamePodServiceServer is the server API for HostnamePodService service.
type HostnamePodServiceServer interface {
	GetPodHostname(context.Context, *Empty) (*PodHostnameResponse, error)
}

// UnimplementedHostnamePodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostnamePodServiceServer struct {
}

func (*UnimplementedHostnamePodServiceServer) GetPodHostname(ctx context.Context, req *Empty) (*PodHostnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodHostname not implemented")
}

func RegisterHostnamePodServiceServer(s *grpc.Server, srv HostnamePodServiceServer) {
	s.RegisterService(&_HostnamePodService_serviceDesc, srv)
}

func _HostnamePodService_GetPodHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostnamePodServiceServer).GetPodHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.health.v1.HostnamePodService/GetPodHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostnamePodServiceServer).GetPodHostname(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostnamePodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.HostnamePodService",
	HandlerType: (*HostnamePodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPodHostname",
			Handler:    _HostnamePodService_GetPodHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostname.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/grpc.health.v1.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.health.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostname.proto",
}
