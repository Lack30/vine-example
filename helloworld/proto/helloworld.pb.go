// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/lack-io/vine-example/helloworld/proto/helloworld.proto

package testdata

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	api "github.com/lack-io/vine-example/goproto/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HelloWorldRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (m *HelloWorldRequest) Reset()         { *m = HelloWorldRequest{} }
func (m *HelloWorldRequest) String() string { return proto.CompactTextString(m) }
func (*HelloWorldRequest) ProtoMessage()    {}
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_406dc886660036aa, []int{0}
}
func (m *HelloWorldRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloWorldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloWorldRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloWorldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldRequest.Merge(m, src)
}
func (m *HelloWorldRequest) XXX_Size() int {
	return m.Size()
}
func (m *HelloWorldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldRequest proto.InternalMessageInfo

func (m *HelloWorldRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloWorldRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type HelloWorldResponse struct {
	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (m *HelloWorldResponse) Reset()         { *m = HelloWorldResponse{} }
func (m *HelloWorldResponse) String() string { return proto.CompactTextString(m) }
func (*HelloWorldResponse) ProtoMessage()    {}
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_406dc886660036aa, []int{1}
}
func (m *HelloWorldResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloWorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloWorldResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloWorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloWorldResponse.Merge(m, src)
}
func (m *HelloWorldResponse) XXX_Size() int {
	return m.Size()
}
func (m *HelloWorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloWorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloWorldResponse proto.InternalMessageInfo

func (m *HelloWorldResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

type MulPathRequest struct {
}

func (m *MulPathRequest) Reset()         { *m = MulPathRequest{} }
func (m *MulPathRequest) String() string { return proto.CompactTextString(m) }
func (*MulPathRequest) ProtoMessage()    {}
func (*MulPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_406dc886660036aa, []int{2}
}
func (m *MulPathRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MulPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MulPathRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MulPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulPathRequest.Merge(m, src)
}
func (m *MulPathRequest) XXX_Size() int {
	return m.Size()
}
func (m *MulPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MulPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MulPathRequest proto.InternalMessageInfo

type MulPathResponse struct {
	Data []*api.App `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *MulPathResponse) Reset()         { *m = MulPathResponse{} }
func (m *MulPathResponse) String() string { return proto.CompactTextString(m) }
func (*MulPathResponse) ProtoMessage()    {}
func (*MulPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_406dc886660036aa, []int{3}
}
func (m *MulPathResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MulPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MulPathResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MulPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MulPathResponse.Merge(m, src)
}
func (m *MulPathResponse) XXX_Size() int {
	return m.Size()
}
func (m *MulPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MulPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MulPathResponse proto.InternalMessageInfo

func (m *MulPathResponse) GetData() []*api.App {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloWorldRequest)(nil), "testdata.HelloWorldRequest")
	proto.RegisterType((*HelloWorldResponse)(nil), "testdata.HelloWorldResponse")
	proto.RegisterType((*MulPathRequest)(nil), "testdata.MulPathRequest")
	proto.RegisterType((*MulPathResponse)(nil), "testdata.MulPathResponse")
}

func init() {
	proto.RegisterFile("github.com/lack-io/vine-example/helloworld/proto/helloworld.proto", fileDescriptor_406dc886660036aa)
}

var fileDescriptor_406dc886660036aa = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xc7, 0xbb, 0x02, 0x7e, 0x8c, 0x89, 0xe2, 0xc6, 0x43, 0x45, 0xd3, 0x90, 0x9e, 0x88, 0x09,
	0x5b, 0x83, 0x27, 0xe3, 0x45, 0xc4, 0x03, 0x17, 0x13, 0xd3, 0x0b, 0x47, 0xb2, 0xc0, 0xa4, 0x34,
	0x2e, 0xdd, 0xb5, 0x5d, 0xfc, 0x78, 0x0b, 0x13, 0x5f, 0xca, 0x23, 0x47, 0x8f, 0x06, 0x5e, 0xc4,
	0x6c, 0x59, 0x3e, 0x8c, 0x24, 0xde, 0xa6, 0x33, 0xfd, 0xff, 0x76, 0x7e, 0x03, 0xcd, 0x28, 0xd6,
	0xc3, 0x71, 0x8f, 0xf5, 0xe5, 0x28, 0x10, 0xbc, 0xff, 0x58, 0x8f, 0x65, 0xf0, 0x1c, 0x27, 0x58,
	0xc7, 0x57, 0x3e, 0x52, 0x02, 0x83, 0x21, 0x0a, 0x21, 0x5f, 0x64, 0x2a, 0x06, 0x81, 0x4a, 0xa5,
	0x96, 0x6b, 0x0d, 0x96, 0x37, 0xe8, 0xae, 0xc6, 0x4c, 0x0f, 0xb8, 0xe6, 0x95, 0xeb, 0xff, 0x60,
	0x91, 0x9c, 0x33, 0xb8, 0x8a, 0x83, 0x08, 0x13, 0x4c, 0xb9, 0x46, 0x8b, 0xf1, 0xaf, 0xe0, 0xa8,
	0x6d, 0xd0, 0x1d, 0x83, 0x0e, 0xf1, 0x69, 0x8c, 0x99, 0xa6, 0x14, 0x8a, 0x09, 0x1f, 0xa1, 0x4b,
	0xaa, 0xa4, 0xb6, 0x17, 0xe6, 0x35, 0x2d, 0x43, 0x81, 0x47, 0xe8, 0x6e, 0x55, 0x49, 0xad, 0x14,
	0x9a, 0xd2, 0x3f, 0x07, 0xba, 0x1e, 0xcd, 0x94, 0x4c, 0x32, 0xa4, 0xc7, 0x50, 0x4a, 0x51, 0x89,
	0x37, 0x1b, 0x9e, 0x7f, 0xf8, 0x65, 0x38, 0xb8, 0x1f, 0x8b, 0x07, 0xae, 0x87, 0xf6, 0x0d, 0xbf,
	0x03, 0x87, 0xcb, 0x8e, 0x8d, 0xde, 0x41, 0xd1, 0x08, 0xb9, 0xa4, 0x5a, 0xa8, 0xed, 0x37, 0x2e,
	0xd8, 0xca, 0x8b, 0x19, 0xaf, 0x6e, 0x2c, 0x99, 0xf1, 0xea, 0x5a, 0x2f, 0x66, 0xbd, 0x18, 0x57,
	0x31, 0x6b, 0x2a, 0x15, 0xe6, 0xe9, 0xc6, 0x07, 0x01, 0x68, 0x2f, 0xaf, 0x45, 0x5b, 0x50, 0x6c,
	0x71, 0x21, 0xe8, 0x29, 0x5b, 0x1c, 0x8c, 0xfd, 0x11, 0xae, 0x9c, 0x6d, 0x1e, 0xce, 0xf7, 0xf2,
	0x1d, 0x7a, 0x03, 0x3b, 0x76, 0x59, 0xea, 0xae, 0x7e, 0xfd, 0x6d, 0x54, 0x39, 0xd9, 0x30, 0x59,
	0x10, 0x6e, 0xdd, 0xcf, 0xa9, 0x47, 0x26, 0x53, 0x8f, 0x7c, 0x4f, 0x3d, 0xf2, 0x3e, 0xf3, 0x9c,
	0xc9, 0xcc, 0x73, 0xbe, 0x66, 0x9e, 0xd3, 0xdb, 0xce, 0x25, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x70, 0xad, 0x49, 0xb9, 0x14, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloworldClient interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
	MulPath(ctx context.Context, in *MulPathRequest, opts ...grpc.CallOption) (*MulPathResponse, error)
}

type helloworldClient struct {
	cc *grpc.ClientConn
}

func NewHelloworldClient(cc *grpc.ClientConn) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) Call(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/testdata.Helloworld/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) MulPath(ctx context.Context, in *MulPathRequest, opts ...grpc.CallOption) (*MulPathResponse, error) {
	out := new(MulPathResponse)
	err := c.cc.Invoke(ctx, "/testdata.Helloworld/MulPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServer is the server API for Helloworld service.
type HelloworldServer interface {
	// +gen:get=/api/v1/call
	// +gen:body=*
	Call(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	MulPath(context.Context, *MulPathRequest) (*MulPathResponse, error)
}

// UnimplementedHelloworldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (*UnimplementedHelloworldServer) Call(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedHelloworldServer) MulPath(ctx context.Context, req *MulPathRequest) (*MulPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MulPath not implemented")
}

func RegisterHelloworldServer(s *grpc.Server, srv HelloworldServer) {
	s.RegisterService(&_Helloworld_serviceDesc, srv)
}

func _Helloworld_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Helloworld/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).Call(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_MulPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MulPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).MulPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Helloworld/MulPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).MulPath(ctx, req.(*MulPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Helloworld_Call_Handler,
		},
		{
			MethodName: "MulPath",
			Handler:    _Helloworld_MulPath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/lack-io/vine-example/helloworld/proto/helloworld.proto",
}

func (m *HelloWorldRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloWorldRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloWorldRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Age != 0 {
		i = encodeVarintHelloworld(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintHelloworld(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloWorldResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloWorldResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloWorldResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reply) > 0 {
		i -= len(m.Reply)
		copy(dAtA[i:], m.Reply)
		i = encodeVarintHelloworld(dAtA, i, uint64(len(m.Reply)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MulPathRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MulPathRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MulPathRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MulPathResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MulPathResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MulPathResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHelloworld(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintHelloworld(dAtA []byte, offset int, v uint64) int {
	offset -= sovHelloworld(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HelloWorldRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovHelloworld(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovHelloworld(uint64(m.Age))
	}
	return n
}

func (m *HelloWorldResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reply)
	if l > 0 {
		n += 1 + l + sovHelloworld(uint64(l))
	}
	return n
}

func (m *MulPathRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MulPathResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovHelloworld(uint64(l))
		}
	}
	return n
}

func sovHelloworld(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHelloworld(x uint64) (n int) {
	return sovHelloworld(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HelloWorldRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHelloworld
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloWorldRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloWorldRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHelloworld
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHelloworld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHelloworld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloWorldResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHelloworld
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloWorldResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloWorldResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHelloworld
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHelloworld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHelloworld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MulPathRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHelloworld
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MulPathRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MulPathRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHelloworld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MulPathResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHelloworld
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MulPathResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MulPathResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHelloworld
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHelloworld
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &api.App{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHelloworld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHelloworld
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHelloworld(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHelloworld
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHelloworld
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHelloworld
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHelloworld
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHelloworld
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHelloworld        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHelloworld          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHelloworld = fmt.Errorf("proto: unexpected end of group")
)
