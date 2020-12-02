// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: nodo.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vacio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Vacio) Reset() {
	*x = Vacio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vacio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vacio) ProtoMessage() {}

func (x *Vacio) ProtoReflect() protoreflect.Message {
	mi := &file_nodo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vacio.ProtoReflect.Descriptor instead.
func (*Vacio) Descriptor() ([]byte, []int) {
	return file_nodo_proto_rawDescGZIP(), []int{0}
}

type Estado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *Estado) Reset() {
	*x = Estado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estado) ProtoMessage() {}

func (x *Estado) ProtoReflect() protoreflect.Message {
	mi := &file_nodo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estado.ProtoReflect.Descriptor instead.
func (*Estado) Descriptor() ([]byte, []int) {
	return file_nodo_proto_rawDescGZIP(), []int{1}
}

func (x *Estado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type Chunck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Parts   uint64 `protobuf:"varint,2,opt,name=parts,proto3" json:"parts,omitempty"`
}

func (x *Chunck) Reset() {
	*x = Chunck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunck) ProtoMessage() {}

func (x *Chunck) ProtoReflect() protoreflect.Message {
	mi := &file_nodo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunck.ProtoReflect.Descriptor instead.
func (*Chunck) Descriptor() ([]byte, []int) {
	return file_nodo_proto_rawDescGZIP(), []int{2}
}

func (x *Chunck) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunck) GetParts() uint64 {
	if x != nil {
		return x.Parts
	}
	return 0
}

var File_nodo_proto protoreflect.FileDescriptor

var file_nodo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x56, 0x61, 0x63, 0x69, 0x6f, 0x22, 0x20, 0x0a, 0x06,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x38,
	0x0a, 0x06, 0x43, 0x68, 0x75, 0x6e, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x73, 0x32, 0x6a, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x69, 0x6f, 0x4e, 0x6f, 0x64, 0x6f, 0x12, 0x2c, 0x0a, 0x0d, 0x4f, 0x62, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x61, 0x63, 0x69, 0x6f, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x2c, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72,
	0x43, 0x68, 0x75, 0x6e, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x68, 0x75, 0x6e, 0x63, 0x6b, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x73,
	0x74, 0x61, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodo_proto_rawDescOnce sync.Once
	file_nodo_proto_rawDescData = file_nodo_proto_rawDesc
)

func file_nodo_proto_rawDescGZIP() []byte {
	file_nodo_proto_rawDescOnce.Do(func() {
		file_nodo_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodo_proto_rawDescData)
	})
	return file_nodo_proto_rawDescData
}

var file_nodo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nodo_proto_goTypes = []interface{}{
	(*Vacio)(nil),  // 0: proto.Vacio
	(*Estado)(nil), // 1: proto.Estado
	(*Chunck)(nil), // 2: proto.Chunck
}
var file_nodo_proto_depIdxs = []int32{
	0, // 0: proto.ServicioNodo.ObtenerEstado:input_type -> proto.Vacio
	2, // 1: proto.ServicioNodo.EnviarChunck:input_type -> proto.Chunck
	1, // 2: proto.ServicioNodo.ObtenerEstado:output_type -> proto.Estado
	1, // 3: proto.ServicioNodo.EnviarChunck:output_type -> proto.Estado
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nodo_proto_init() }
func file_nodo_proto_init() {
	if File_nodo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vacio); i {
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
		file_nodo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estado); i {
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
		file_nodo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunck); i {
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
			RawDescriptor: file_nodo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodo_proto_goTypes,
		DependencyIndexes: file_nodo_proto_depIdxs,
		MessageInfos:      file_nodo_proto_msgTypes,
	}.Build()
	File_nodo_proto = out.File
	file_nodo_proto_rawDesc = nil
	file_nodo_proto_goTypes = nil
	file_nodo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServicioNodoClient is the client API for ServicioNodo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicioNodoClient interface {
	ObtenerEstado(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Estado, error)
	EnviarChunck(ctx context.Context, in *Chunck, opts ...grpc.CallOption) (*Estado, error)
}

type servicioNodoClient struct {
	cc grpc.ClientConnInterface
}

func NewServicioNodoClient(cc grpc.ClientConnInterface) ServicioNodoClient {
	return &servicioNodoClient{cc}
}

func (c *servicioNodoClient) ObtenerEstado(ctx context.Context, in *Vacio, opts ...grpc.CallOption) (*Estado, error) {
	out := new(Estado)
	err := c.cc.Invoke(ctx, "/proto.ServicioNodo/ObtenerEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicioNodoClient) EnviarChunck(ctx context.Context, in *Chunck, opts ...grpc.CallOption) (*Estado, error) {
	out := new(Estado)
	err := c.cc.Invoke(ctx, "/proto.ServicioNodo/EnviarChunck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicioNodoServer is the server API for ServicioNodo service.
type ServicioNodoServer interface {
	ObtenerEstado(context.Context, *Vacio) (*Estado, error)
	EnviarChunck(context.Context, *Chunck) (*Estado, error)
}

// UnimplementedServicioNodoServer can be embedded to have forward compatible implementations.
type UnimplementedServicioNodoServer struct {
}

func (*UnimplementedServicioNodoServer) ObtenerEstado(context.Context, *Vacio) (*Estado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtenerEstado not implemented")
}
func (*UnimplementedServicioNodoServer) EnviarChunck(context.Context, *Chunck) (*Estado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarChunck not implemented")
}

func RegisterServicioNodoServer(s *grpc.Server, srv ServicioNodoServer) {
	s.RegisterService(&_ServicioNodo_serviceDesc, srv)
}

func _ServicioNodo_ObtenerEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vacio)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioNodoServer).ObtenerEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServicioNodo/ObtenerEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioNodoServer).ObtenerEstado(ctx, req.(*Vacio))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServicioNodo_EnviarChunck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicioNodoServer).EnviarChunck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServicioNodo/EnviarChunck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicioNodoServer).EnviarChunck(ctx, req.(*Chunck))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServicioNodo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServicioNodo",
	HandlerType: (*ServicioNodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtenerEstado",
			Handler:    _ServicioNodo_ObtenerEstado_Handler,
		},
		{
			MethodName: "EnviarChunck",
			Handler:    _ServicioNodo_EnviarChunck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodo.proto",
}
