package proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Networkid   uint32 `protobuf:"varint,1,opt,name=networkid,proto3" json:"networkid,omitempty"`
	Fn          string `protobuf:"bytes,2,opt,name=fn,proto3" json:"fn,omitempty"`
	Params      string `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	Abi         string `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty"`
	Address     string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	FromAddress string `protobuf:"bytes,6,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"` // python will not allow .from so it has to be fromAddress
	Value       uint64 `protobuf:"varint,7,opt,name=value,proto3" json:"value,omitempty"`
	GasSupply   uint64 `protobuf:"varint,8,opt,name=gasSupply,proto3" json:"gasSupply,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dApp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dApp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_dApp_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetNetworkid() uint32 {
	if x != nil {
		return x.Networkid
	}
	return 0
}

func (x *CallRequest) GetFn() string {
	if x != nil {
		return x.Fn
	}
	return ""
}

func (x *CallRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *CallRequest) GetAbi() string {
	if x != nil {
		return x.Abi
	}
	return ""
}

func (x *CallRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CallRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *CallRequest) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CallRequest) GetGasSupply() uint64 {
	if x != nil {
		return x.GasSupply
	}
	return 0
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dApp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dApp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_dApp_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_dApp_proto protoreflect.FileDescriptor

var file_dApp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x41, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x65, 0x74, 0x68, 0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x66, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x62, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x62, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x67, 0x61, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x22, 0x26,
	0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x52, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x65, 0x74, 0x68, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x74, 0x68, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dApp_proto_rawDescOnce sync.Once
	file_dApp_proto_rawDescData = file_dApp_proto_rawDesc
)

func file_dApp_proto_rawDescGZIP() []byte {
	file_dApp_proto_rawDescOnce.Do(func() {
		file_dApp_proto_rawDescData = protoimpl.X.CompressGZIP(file_dApp_proto_rawDescData)
	})
	return file_dApp_proto_rawDescData
}

var file_dApp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dApp_proto_goTypes = []interface{}{
	(*CallRequest)(nil),  // 0: protoeth.CallRequest
	(*CallResponse)(nil), // 1: protoeth.CallResponse
}
var file_dApp_proto_depIdxs = []int32{
	0, // 0: protoeth.ProtoEthService.ContractCall:input_type -> protoeth.CallRequest
	1, // 1: protoeth.ProtoEthService.ContractCall:output_type -> protoeth.CallResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dApp_proto_init() }
func file_dApp_proto_init() {
	if File_dApp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dApp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_dApp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_dApp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dApp_proto_goTypes,
		DependencyIndexes: file_dApp_proto_depIdxs,
		MessageInfos:      file_dApp_proto_msgTypes,
	}.Build()
	File_dApp_proto = out.File
	file_dApp_proto_rawDesc = nil
	file_dApp_proto_goTypes = nil
	file_dApp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtoEthServiceClient is the client API for ProtoEthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtoEthServiceClient interface {
	ContractCall(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type protoEthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoEthServiceClient(cc grpc.ClientConnInterface) ProtoEthServiceClient {
	return &protoEthServiceClient{cc}
}

func (c *protoEthServiceClient) ContractCall(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/protoeth.ProtoEthService/ContractCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoEthServiceServer is the server API for ProtoEthService service.
type ProtoEthServiceServer interface {
	ContractCall(context.Context, *CallRequest) (*CallResponse, error)
}

// UnimplementedProtoEthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProtoEthServiceServer struct {
}

func (*UnimplementedProtoEthServiceServer) ContractCall(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractCall not implemented")
}

func RegisterProtoEthServiceServer(s *grpc.Server, srv ProtoEthServiceServer) {
	s.RegisterService(&_ProtoEthService_serviceDesc, srv)
}

func _ProtoEthService_ContractCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoEthServiceServer).ContractCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoeth.ProtoEthService/ContractCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoEthServiceServer).ContractCall(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoEthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoeth.ProtoEthService",
	HandlerType: (*ProtoEthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContractCall",
			Handler:    _ProtoEthService_ContractCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dApp.proto",
}
