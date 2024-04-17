// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package grpctest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type TestRequest struct {
	Input                int32    `protobuf:"varint,1,opt,name=Input,proto3" json:"Input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (dst *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(dst, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetInput() int32 {
	if m != nil {
		return m.Input
	}
	return 0
}

type TestResponse struct {
	Output               int32    `protobuf:"varint,2,opt,name=Output,proto3" json:"Output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{1}
}
func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (dst *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(dst, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetOutput() int32 {
	if m != nil {
		return m.Output
	}
	return 0
}

type PrintKVRequest struct {
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*PrintKVRequest_ValueString
	//	*PrintKVRequest_ValueInt
	Value                isPrintKVRequest_Value `protobuf_oneof:"Value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PrintKVRequest) Reset()         { *m = PrintKVRequest{} }
func (m *PrintKVRequest) String() string { return proto.CompactTextString(m) }
func (*PrintKVRequest) ProtoMessage()    {}
func (*PrintKVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{2}
}
func (m *PrintKVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintKVRequest.Unmarshal(m, b)
}
func (m *PrintKVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintKVRequest.Marshal(b, m, deterministic)
}
func (dst *PrintKVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintKVRequest.Merge(dst, src)
}
func (m *PrintKVRequest) XXX_Size() int {
	return xxx_messageInfo_PrintKVRequest.Size(m)
}
func (m *PrintKVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintKVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrintKVRequest proto.InternalMessageInfo

func (m *PrintKVRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type isPrintKVRequest_Value interface {
	isPrintKVRequest_Value()
}

type PrintKVRequest_ValueString struct {
	ValueString string `protobuf:"bytes,2,opt,name=ValueString,proto3,oneof"`
}

type PrintKVRequest_ValueInt struct {
	ValueInt int32 `protobuf:"varint,3,opt,name=ValueInt,proto3,oneof"`
}

func (*PrintKVRequest_ValueString) isPrintKVRequest_Value() {}

func (*PrintKVRequest_ValueInt) isPrintKVRequest_Value() {}

func (m *PrintKVRequest) GetValue() isPrintKVRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PrintKVRequest) GetValueString() string {
	if x, ok := m.GetValue().(*PrintKVRequest_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *PrintKVRequest) GetValueInt() int32 {
	if x, ok := m.GetValue().(*PrintKVRequest_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PrintKVRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PrintKVRequest_OneofMarshaler, _PrintKVRequest_OneofUnmarshaler, _PrintKVRequest_OneofSizer, []interface{}{
		(*PrintKVRequest_ValueString)(nil),
		(*PrintKVRequest_ValueInt)(nil),
	}
}

func _PrintKVRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PrintKVRequest)
	// Value
	switch x := m.Value.(type) {
	case *PrintKVRequest_ValueString:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ValueString)
	case *PrintKVRequest_ValueInt:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ValueInt))
	case nil:
	default:
		return fmt.Errorf("PrintKVRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _PrintKVRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PrintKVRequest)
	switch tag {
	case 2: // Value.ValueString
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &PrintKVRequest_ValueString{x}
		return true, err
	case 3: // Value.ValueInt
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &PrintKVRequest_ValueInt{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _PrintKVRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PrintKVRequest)
	// Value
	switch x := m.Value.(type) {
	case *PrintKVRequest_ValueString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ValueString)))
		n += len(x.ValueString)
	case *PrintKVRequest_ValueInt:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ValueInt))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PrintKVResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintKVResponse) Reset()         { *m = PrintKVResponse{} }
func (m *PrintKVResponse) String() string { return proto.CompactTextString(m) }
func (*PrintKVResponse) ProtoMessage()    {}
func (*PrintKVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{3}
}
func (m *PrintKVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintKVResponse.Unmarshal(m, b)
}
func (m *PrintKVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintKVResponse.Marshal(b, m, deterministic)
}
func (dst *PrintKVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintKVResponse.Merge(dst, src)
}
func (m *PrintKVResponse) XXX_Size() int {
	return xxx_messageInfo_PrintKVResponse.Size(m)
}
func (m *PrintKVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintKVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintKVResponse proto.InternalMessageInfo

type BidirectionalRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalRequest) Reset()         { *m = BidirectionalRequest{} }
func (m *BidirectionalRequest) String() string { return proto.CompactTextString(m) }
func (*BidirectionalRequest) ProtoMessage()    {}
func (*BidirectionalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{4}
}
func (m *BidirectionalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalRequest.Unmarshal(m, b)
}
func (m *BidirectionalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalRequest.Marshal(b, m, deterministic)
}
func (dst *BidirectionalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalRequest.Merge(dst, src)
}
func (m *BidirectionalRequest) XXX_Size() int {
	return xxx_messageInfo_BidirectionalRequest.Size(m)
}
func (m *BidirectionalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalRequest proto.InternalMessageInfo

func (m *BidirectionalRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type BidirectionalResponse struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidirectionalResponse) Reset()         { *m = BidirectionalResponse{} }
func (m *BidirectionalResponse) String() string { return proto.CompactTextString(m) }
func (*BidirectionalResponse) ProtoMessage()    {}
func (*BidirectionalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{5}
}
func (m *BidirectionalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidirectionalResponse.Unmarshal(m, b)
}
func (m *BidirectionalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidirectionalResponse.Marshal(b, m, deterministic)
}
func (dst *BidirectionalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidirectionalResponse.Merge(dst, src)
}
func (m *BidirectionalResponse) XXX_Size() int {
	return xxx_messageInfo_BidirectionalResponse.Size(m)
}
func (m *BidirectionalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BidirectionalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BidirectionalResponse proto.InternalMessageInfo

func (m *BidirectionalResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PrintStdioRequest struct {
	Stdout               []byte   `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr               []byte   `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintStdioRequest) Reset()         { *m = PrintStdioRequest{} }
func (m *PrintStdioRequest) String() string { return proto.CompactTextString(m) }
func (*PrintStdioRequest) ProtoMessage()    {}
func (*PrintStdioRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{6}
}
func (m *PrintStdioRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintStdioRequest.Unmarshal(m, b)
}
func (m *PrintStdioRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintStdioRequest.Marshal(b, m, deterministic)
}
func (dst *PrintStdioRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintStdioRequest.Merge(dst, src)
}
func (m *PrintStdioRequest) XXX_Size() int {
	return xxx_messageInfo_PrintStdioRequest.Size(m)
}
func (m *PrintStdioRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintStdioRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrintStdioRequest proto.InternalMessageInfo

func (m *PrintStdioRequest) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *PrintStdioRequest) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{7}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PongResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ed149f2304c9fa82, []int{8}
}
func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (dst *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(dst, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "grpctest.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "grpctest.TestResponse")
	proto.RegisterType((*PrintKVRequest)(nil), "grpctest.PrintKVRequest")
	proto.RegisterType((*PrintKVResponse)(nil), "grpctest.PrintKVResponse")
	proto.RegisterType((*BidirectionalRequest)(nil), "grpctest.BidirectionalRequest")
	proto.RegisterType((*BidirectionalResponse)(nil), "grpctest.BidirectionalResponse")
	proto.RegisterType((*PrintStdioRequest)(nil), "grpctest.PrintStdioRequest")
	proto.RegisterType((*PingRequest)(nil), "grpctest.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "grpctest.PongResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Double(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	PrintKV(ctx context.Context, in *PrintKVRequest, opts ...grpc.CallOption) (*PrintKVResponse, error)
	Bidirectional(ctx context.Context, in *BidirectionalRequest, opts ...grpc.CallOption) (*BidirectionalResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Test_StreamClient, error)
	PrintStdio(ctx context.Context, in *PrintStdioRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) Double(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/Double", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) PrintKV(ctx context.Context, in *PrintKVRequest, opts ...grpc.CallOption) (*PrintKVResponse, error) {
	out := new(PrintKVResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/PrintKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Bidirectional(ctx context.Context, in *BidirectionalRequest, opts ...grpc.CallOption) (*BidirectionalResponse, error) {
	out := new(BidirectionalResponse)
	err := c.cc.Invoke(ctx, "/grpctest.Test/Bidirectional", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Test_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/grpctest.Test/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamClient{stream}
	return x, nil
}

type Test_StreamClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testStreamClient struct {
	grpc.ClientStream
}

func (x *testStreamClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testStreamClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) PrintStdio(ctx context.Context, in *PrintStdioRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpctest.Test/PrintStdio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Double(context.Context, *TestRequest) (*TestResponse, error)
	PrintKV(context.Context, *PrintKVRequest) (*PrintKVResponse, error)
	Bidirectional(context.Context, *BidirectionalRequest) (*BidirectionalResponse, error)
	Stream(Test_StreamServer) error
	PrintStdio(context.Context, *PrintStdioRequest) (*empty.Empty, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Double_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Double(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/Double",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Double(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_PrintKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).PrintKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/PrintKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).PrintKV(ctx, req.(*PrintKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Bidirectional_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidirectionalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Bidirectional(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/Bidirectional",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Bidirectional(ctx, req.(*BidirectionalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Stream(&testStreamServer{stream})
}

type Test_StreamServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type testStreamServer struct {
	grpc.ServerStream
}

func (x *testStreamServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testStreamServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Test_PrintStdio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintStdioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).PrintStdio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.Test/PrintStdio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).PrintStdio(ctx, req.(*PrintStdioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Double",
			Handler:    _Test_Double_Handler,
		},
		{
			MethodName: "PrintKV",
			Handler:    _Test_PrintKV_Handler,
		},
		{
			MethodName: "Bidirectional",
			Handler:    _Test_Bidirectional_Handler,
		},
		{
			MethodName: "PrintStdio",
			Handler:    _Test_PrintStdio_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Test_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/grpctest.PingPong/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	Ping(context.Context, *PingRequest) (*PongResponse, error)
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.PingPong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPong_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_ed149f2304c9fa82) }

var fileDescriptor_test_ed149f2304c9fa82 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x6f, 0xda, 0x30,
	0x18, 0x4d, 0x42, 0x49, 0xe9, 0x07, 0x74, 0xad, 0xd5, 0x22, 0x96, 0x4e, 0x5b, 0xe5, 0x49, 0x5d,
	0x4f, 0xe9, 0xd4, 0x1d, 0xa6, 0x1d, 0x26, 0x4d, 0xb0, 0x49, 0x20, 0x0e, 0x43, 0x66, 0xe2, 0x0e,
	0xc4, 0x8b, 0x2c, 0x25, 0x71, 0x66, 0x3b, 0x07, 0x7e, 0xeb, 0xfe, 0xcc, 0x64, 0xc7, 0x09, 0x21,
	0x62, 0x87, 0xde, 0xfc, 0xbd, 0xbc, 0xbc, 0xef, 0xf9, 0x3d, 0x03, 0x28, 0x2a, 0x55, 0x98, 0x0b,
	0xae, 0x38, 0xea, 0xc5, 0x22, 0xdf, 0xe9, 0x39, 0xb8, 0x8b, 0x39, 0x8f, 0x13, 0xfa, 0x64, 0xf0,
	0x6d, 0xf1, 0xfb, 0x89, 0xa6, 0xb9, 0xda, 0x97, 0x34, 0xfc, 0x1e, 0xfa, 0xbf, 0xa8, 0x54, 0x84,
	0xfe, 0x29, 0xa8, 0x54, 0xe8, 0x06, 0xba, 0xf3, 0x2c, 0x2f, 0xd4, 0xd8, 0xbd, 0x77, 0x1f, 0xbb,
	0xa4, 0x1c, 0xf0, 0x03, 0x0c, 0x4a, 0x92, 0xcc, 0x79, 0x26, 0x29, 0x1a, 0x81, 0xff, 0xb3, 0x50,
	0x9a, 0xe6, 0x19, 0x9a, 0x9d, 0x70, 0x0a, 0x97, 0x4b, 0xc1, 0x32, 0xb5, 0x58, 0x57, 0x7a, 0x57,
	0xd0, 0x59, 0xd0, 0xbd, 0x51, 0xbb, 0x20, 0xfa, 0x88, 0x30, 0xf4, 0xd7, 0x9b, 0xa4, 0xa0, 0x2b,
	0x25, 0x58, 0x16, 0x1b, 0x81, 0x8b, 0x99, 0x43, 0x9a, 0x20, 0x7a, 0x03, 0x3d, 0x33, 0xce, 0x33,
	0x35, 0xee, 0xe8, 0x0d, 0x33, 0x87, 0xd4, 0xc8, 0xe4, 0x1c, 0xba, 0xe6, 0x8c, 0xaf, 0xe1, 0x55,
	0xbd, 0xae, 0x74, 0x86, 0x1f, 0xe0, 0x66, 0xc2, 0x22, 0x26, 0xe8, 0x4e, 0x31, 0x9e, 0x6d, 0x92,
	0xca, 0xc7, 0x25, 0x78, 0x2c, 0x32, 0x36, 0x86, 0xc4, 0x63, 0x11, 0xfe, 0x00, 0xb7, 0x2d, 0x9e,
	0xbd, 0x5a, 0x9b, 0x38, 0x85, 0x6b, 0xb3, 0x63, 0xa5, 0x22, 0xc6, 0x2b, 0xb5, 0x11, 0xf8, 0x52,
	0x45, 0xdc, 0xc6, 0x34, 0x20, 0x76, 0xb2, 0x38, 0x15, 0xc2, 0x5c, 0xab, 0xc4, 0xa9, 0x10, 0x78,
	0x08, 0xfd, 0x25, 0xcb, 0x62, 0xfb, 0x3b, 0xbe, 0x87, 0xc1, 0x92, 0xeb, 0xd1, 0xee, 0xbc, 0x82,
	0x4e, 0x2a, 0xe3, 0x2a, 0xa4, 0x54, 0xc6, 0xcf, 0x7f, 0x3d, 0x38, 0xd3, 0x89, 0xa3, 0x2f, 0xe0,
	0x7f, 0xe7, 0xc5, 0x36, 0xa1, 0xe8, 0x36, 0xac, 0x0a, 0x0d, 0x1b, 0x85, 0x05, 0xa3, 0x36, 0x6c,
	0x83, 0x70, 0xd0, 0x37, 0x38, 0xb7, 0xe9, 0xa0, 0xf1, 0x81, 0x74, 0xdc, 0x4f, 0xf0, 0xfa, 0xc4,
	0x97, 0x5a, 0x81, 0xc0, 0xf0, 0x28, 0x24, 0xf4, 0xf6, 0xc0, 0x3e, 0x95, 0x72, 0xf0, 0xee, 0xbf,
	0xdf, 0x6b, 0xcd, 0xaf, 0xe0, 0xaf, 0x94, 0xa0, 0x9b, 0xf4, 0xc5, 0x17, 0x7a, 0x74, 0x3f, 0xba,
	0x68, 0x0a, 0x70, 0xa8, 0x03, 0xdd, 0xb5, 0xdc, 0x37, 0x4b, 0xd2, 0x42, 0xe6, 0xdd, 0x87, 0xd5,
	0xbb, 0x0f, 0x7f, 0xe8, 0x77, 0x8f, 0x9d, 0xe7, 0x29, 0xf4, 0x74, 0x1d, 0xba, 0x03, 0xf4, 0x19,
	0xce, 0xf4, 0xb9, 0xe9, 0xa6, 0x51, 0x55, 0xd3, 0x4d, 0xb3, 0x32, 0xec, 0x6c, 0x7d, 0x23, 0xfb,
	0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x4b, 0x8c, 0x40, 0x74, 0x03, 0x00, 0x00,
}