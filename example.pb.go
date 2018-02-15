// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Parent
	Dependent
	CreateDependencyRequest
	CreateDependencyResponse
*/
package example

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

type Parent struct {
	Name       string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Dependents []*Dependent `protobuf:"bytes,2,rep,name=dependents" json:"dependents,omitempty"`
}

func (m *Parent) Reset()                    { *m = Parent{} }
func (m *Parent) String() string            { return proto.CompactTextString(m) }
func (*Parent) ProtoMessage()               {}
func (*Parent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Parent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parent) GetDependents() []*Dependent {
	if m != nil {
		return m.Dependents
	}
	return nil
}

type Dependent struct {
	// Types that are valid to be assigned to Child:
	//	*Dependent_Son
	//	*Dependent_Daughter
	Child isDependent_Child `protobuf_oneof:"child"`
	Name  string            `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Dependent) Reset()                    { *m = Dependent{} }
func (m *Dependent) String() string            { return proto.CompactTextString(m) }
func (*Dependent) ProtoMessage()               {}
func (*Dependent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isDependent_Child interface {
	isDependent_Child()
}

type Dependent_Son struct {
	Son *Dependent_Male `protobuf:"bytes,1,opt,name=son,oneof"`
}
type Dependent_Daughter struct {
	Daughter *Dependent_Female `protobuf:"bytes,2,opt,name=daughter,oneof"`
}

func (*Dependent_Son) isDependent_Child()      {}
func (*Dependent_Daughter) isDependent_Child() {}

func (m *Dependent) GetChild() isDependent_Child {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *Dependent) GetSon() *Dependent_Male {
	if x, ok := m.GetChild().(*Dependent_Son); ok {
		return x.Son
	}
	return nil
}

func (m *Dependent) GetDaughter() *Dependent_Female {
	if x, ok := m.GetChild().(*Dependent_Daughter); ok {
		return x.Daughter
	}
	return nil
}

func (m *Dependent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Dependent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Dependent_OneofMarshaler, _Dependent_OneofUnmarshaler, _Dependent_OneofSizer, []interface{}{
		(*Dependent_Son)(nil),
		(*Dependent_Daughter)(nil),
	}
}

func _Dependent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Dependent)
	// child
	switch x := m.Child.(type) {
	case *Dependent_Son:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Son); err != nil {
			return err
		}
	case *Dependent_Daughter:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Daughter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Dependent.Child has unexpected type %T", x)
	}
	return nil
}

func _Dependent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Dependent)
	switch tag {
	case 1: // child.son
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Dependent_Male)
		err := b.DecodeMessage(msg)
		m.Child = &Dependent_Son{msg}
		return true, err
	case 2: // child.daughter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Dependent_Female)
		err := b.DecodeMessage(msg)
		m.Child = &Dependent_Daughter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Dependent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Dependent)
	// child
	switch x := m.Child.(type) {
	case *Dependent_Son:
		s := proto.Size(x.Son)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dependent_Daughter:
		s := proto.Size(x.Daughter)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Dependent_Male struct {
}

func (m *Dependent_Male) Reset()                    { *m = Dependent_Male{} }
func (m *Dependent_Male) String() string            { return proto.CompactTextString(m) }
func (*Dependent_Male) ProtoMessage()               {}
func (*Dependent_Male) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Dependent_Female struct {
	Attributes []string `protobuf:"bytes,1,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *Dependent_Female) Reset()                    { *m = Dependent_Female{} }
func (m *Dependent_Female) String() string            { return proto.CompactTextString(m) }
func (*Dependent_Female) ProtoMessage()               {}
func (*Dependent_Female) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Dependent_Female) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type CreateDependencyRequest struct {
	ParentName string       `protobuf:"bytes,1,opt,name=parent_name,json=parentName" json:"parent_name,omitempty"`
	Dependents []*Dependent `protobuf:"bytes,2,rep,name=dependents" json:"dependents,omitempty"`
}

func (m *CreateDependencyRequest) Reset()                    { *m = CreateDependencyRequest{} }
func (m *CreateDependencyRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDependencyRequest) ProtoMessage()               {}
func (*CreateDependencyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateDependencyRequest) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *CreateDependencyRequest) GetDependents() []*Dependent {
	if m != nil {
		return m.Dependents
	}
	return nil
}

type CreateDependencyResponse struct {
}

func (m *CreateDependencyResponse) Reset()                    { *m = CreateDependencyResponse{} }
func (m *CreateDependencyResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDependencyResponse) ProtoMessage()               {}
func (*CreateDependencyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Parent)(nil), "example.Parent")
	proto.RegisterType((*Dependent)(nil), "example.Dependent")
	proto.RegisterType((*Dependent_Male)(nil), "example.Dependent.Male")
	proto.RegisterType((*Dependent_Female)(nil), "example.Dependent.Female")
	proto.RegisterType((*CreateDependencyRequest)(nil), "example.CreateDependencyRequest")
	proto.RegisterType((*CreateDependencyResponse)(nil), "example.CreateDependencyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Family service

type FamilyClient interface {
	CreateDependency(ctx context.Context, in *CreateDependencyRequest, opts ...grpc.CallOption) (*CreateDependencyResponse, error)
}

type familyClient struct {
	cc *grpc.ClientConn
}

func NewFamilyClient(cc *grpc.ClientConn) FamilyClient {
	return &familyClient{cc}
}

func (c *familyClient) CreateDependency(ctx context.Context, in *CreateDependencyRequest, opts ...grpc.CallOption) (*CreateDependencyResponse, error) {
	out := new(CreateDependencyResponse)
	err := grpc.Invoke(ctx, "/example.Family/CreateDependency", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Family service

type FamilyServer interface {
	CreateDependency(context.Context, *CreateDependencyRequest) (*CreateDependencyResponse, error)
}

func RegisterFamilyServer(s *grpc.Server, srv FamilyServer) {
	s.RegisterService(&_Family_serviceDesc, srv)
}

func _Family_CreateDependency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDependencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilyServer).CreateDependency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Family/CreateDependency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilyServer).CreateDependency(ctx, req.(*CreateDependencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Family_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Family",
	HandlerType: (*FamilyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDependency",
			Handler:    _Family_CreateDependency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3f, 0x4f, 0x83, 0x40,
	0x14, 0x2f, 0xa5, 0x52, 0x79, 0xc4, 0xc4, 0xbc, 0xa5, 0xc8, 0xa0, 0xc8, 0x44, 0x62, 0xc2, 0x80,
	0x83, 0xbb, 0x9a, 0xc6, 0x45, 0xd3, 0xb0, 0x39, 0x99, 0x2b, 0xbc, 0x58, 0x12, 0x38, 0xce, 0xbb,
	0x23, 0xb1, 0xdf, 0xce, 0x8f, 0x66, 0x7a, 0x08, 0x12, 0xab, 0x0e, 0xdd, 0xee, 0xde, 0xfb, 0xfd,
	0xcb, 0x2f, 0x0f, 0x4e, 0xe8, 0x9d, 0xd5, 0xa2, 0xa2, 0x44, 0xc8, 0x46, 0x37, 0x38, 0xff, 0xfa,
	0x46, 0x2b, 0x70, 0x56, 0x4c, 0x12, 0xd7, 0x88, 0x30, 0xe3, 0xac, 0x26, 0xdf, 0x0a, 0xad, 0xd8,
	0xcd, 0xcc, 0x1b, 0x53, 0x80, 0x82, 0x04, 0xf1, 0x82, 0xb8, 0x56, 0xfe, 0x34, 0xb4, 0x63, 0x2f,
	0xc5, 0xa4, 0x97, 0xba, 0xef, 0x57, 0xd9, 0x08, 0x15, 0x7d, 0x58, 0xe0, 0x0e, 0x1b, 0xbc, 0x02,
	0x5b, 0x35, 0xdc, 0x88, 0x7a, 0xe9, 0x62, 0x9f, 0x9a, 0x3c, 0xb2, 0x8a, 0x1e, 0x26, 0xd9, 0x0e,
	0x85, 0x37, 0x70, 0x5c, 0xb0, 0xf6, 0x75, 0xa3, 0x49, 0xfa, 0x53, 0xc3, 0x38, 0xfb, 0x85, 0xb1,
	0xa4, 0xba, 0xe3, 0x0c, 0xe0, 0x21, 0xbb, 0xfd, 0x9d, 0x3d, 0x70, 0x60, 0xb6, 0xd3, 0x0e, 0x62,
	0x70, 0x3a, 0x06, 0x9e, 0x03, 0x30, 0xad, 0x65, 0xb9, 0x6e, 0x35, 0x29, 0xdf, 0x0a, 0xed, 0xd8,
	0xcd, 0x46, 0x93, 0xdb, 0x39, 0x1c, 0xe5, 0x9b, 0xb2, 0x2a, 0x22, 0x0e, 0x8b, 0x3b, 0x49, 0x4c,
	0x53, 0x6f, 0x9a, 0x6f, 0x33, 0x7a, 0x6b, 0x49, 0x69, 0xbc, 0x00, 0x4f, 0x98, 0xbe, 0x5e, 0x46,
	0x65, 0x41, 0x37, 0x7a, 0x3a, 0xb4, 0xb2, 0x00, 0xfc, 0x7d, 0x3f, 0x25, 0x1a, 0xae, 0x28, 0xcd,
	0xc1, 0x59, 0xb2, 0xba, 0xac, 0xb6, 0xf8, 0x0c, 0xa7, 0x3f, 0x51, 0x18, 0x0e, 0xca, 0x7f, 0x04,
	0x0e, 0x2e, 0xff, 0x41, 0x74, 0x16, 0xd1, 0x64, 0xed, 0x98, 0xab, 0xb8, 0xfe, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x45, 0xd0, 0x58, 0xae, 0x26, 0x02, 0x00, 0x00,
}
