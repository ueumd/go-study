// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: hello.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_hello_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_hello_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 数字代表编号
	Url     string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`   // 数字代表编号
	Sex     Gender                 `protobuf:"varint,3,opt,name=sex,proto3,enum=Gender" json:"sex,omitempty"`
	List    map[string]string      `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AddTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HelloRequest) GetSex() Gender {
	if x != nil {
		return x.Sex
	}
	return Gender_MALE
}

func (x *HelloRequest) GetList() map[string]string {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *HelloRequest) GetAddTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AddTime
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 数字代表编号
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{2}
}

func (x *Pong) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb,
	0x01, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x29, 0x0a, 0x0d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a,
	0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x32,
	0x38, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hello_proto_goTypes = []interface{}{
	(Gender)(0),                   // 0: Gender
	(*HelloRequest)(nil),          // 1: HelloRequest
	(*HelloResponse)(nil),         // 2: HelloResponse
	(*Pong)(nil),                  // 3: Pong
	nil,                           // 4: HelloRequest.ListEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: HelloRequest.sex:type_name -> Gender
	4, // 1: HelloRequest.list:type_name -> HelloRequest.ListEntry
	5, // 2: HelloRequest.addTime:type_name -> google.protobuf.Timestamp
	1, // 3: GreeterTest.SayHello:input_type -> HelloRequest
	2, // 4: GreeterTest.SayHello:output_type -> HelloResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		EnumInfos:         file_hello_proto_enumTypes,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterTestClient is the client API for GreeterTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterTestClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterTestClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterTestClient(cc grpc.ClientConnInterface) GreeterTestClient {
	return &greeterTestClient{cc}
}

func (c *greeterTestClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/GreeterTest/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterTestServer is the server API for GreeterTest service.
type GreeterTestServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedGreeterTestServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterTestServer struct {
}

func (*UnimplementedGreeterTestServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterTestServer(s *grpc.Server, srv GreeterTestServer) {
	s.RegisterService(&_GreeterTest_serviceDesc, srv)
}

func _GreeterTest_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterTestServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreeterTest/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterTestServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GreeterTest",
	HandlerType: (*GreeterTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreeterTest_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
