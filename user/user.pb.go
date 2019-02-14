// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type UserEntity_Phone_PhoneType int32

const (
	UserEntity_Phone_MOBILE   UserEntity_Phone_PhoneType = 0
	UserEntity_Phone_LANDLINE UserEntity_Phone_PhoneType = 1
)

var UserEntity_Phone_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "LANDLINE",
}
var UserEntity_Phone_PhoneType_value = map[string]int32{
	"MOBILE":   0,
	"LANDLINE": 1,
}

func (x UserEntity_Phone_PhoneType) String() string {
	return proto.EnumName(UserEntity_Phone_PhoneType_name, int32(x))
}
func (UserEntity_Phone_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{0, 1, 0}
}

type UserEntity struct {
	Id                   int32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Emails               []*UserEntity_Email `protobuf:"bytes,3,rep,name=emails,proto3" json:"emails,omitempty"`
	PhoneNumbers         []*UserEntity_Phone `protobuf:"bytes,4,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserEntity) Reset()         { *m = UserEntity{} }
func (m *UserEntity) String() string { return proto.CompactTextString(m) }
func (*UserEntity) ProtoMessage()    {}
func (*UserEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{0}
}
func (m *UserEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity.Unmarshal(m, b)
}
func (m *UserEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity.Marshal(b, m, deterministic)
}
func (dst *UserEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity.Merge(dst, src)
}
func (m *UserEntity) XXX_Size() int {
	return xxx_messageInfo_UserEntity.Size(m)
}
func (m *UserEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEntity.DiscardUnknown(m)
}

var xxx_messageInfo_UserEntity proto.InternalMessageInfo

func (m *UserEntity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserEntity) GetEmails() []*UserEntity_Email {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *UserEntity) GetPhoneNumbers() []*UserEntity_Phone {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

type UserEntity_Email struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	IsPrimary            bool     `protobuf:"varint,2,opt,name=isPrimary,proto3" json:"isPrimary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserEntity_Email) Reset()         { *m = UserEntity_Email{} }
func (m *UserEntity_Email) String() string { return proto.CompactTextString(m) }
func (*UserEntity_Email) ProtoMessage()    {}
func (*UserEntity_Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{0, 0}
}
func (m *UserEntity_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity_Email.Unmarshal(m, b)
}
func (m *UserEntity_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity_Email.Marshal(b, m, deterministic)
}
func (dst *UserEntity_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity_Email.Merge(dst, src)
}
func (m *UserEntity_Email) XXX_Size() int {
	return xxx_messageInfo_UserEntity_Email.Size(m)
}
func (m *UserEntity_Email) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEntity_Email.DiscardUnknown(m)
}

var xxx_messageInfo_UserEntity_Email proto.InternalMessageInfo

func (m *UserEntity_Email) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserEntity_Email) GetIsPrimary() bool {
	if m != nil {
		return m.IsPrimary
	}
	return false
}

type UserEntity_Phone struct {
	Number               string                     `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 UserEntity_Phone_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=user.UserEntity_Phone_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UserEntity_Phone) Reset()         { *m = UserEntity_Phone{} }
func (m *UserEntity_Phone) String() string { return proto.CompactTextString(m) }
func (*UserEntity_Phone) ProtoMessage()    {}
func (*UserEntity_Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{0, 1}
}
func (m *UserEntity_Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity_Phone.Unmarshal(m, b)
}
func (m *UserEntity_Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity_Phone.Marshal(b, m, deterministic)
}
func (dst *UserEntity_Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity_Phone.Merge(dst, src)
}
func (m *UserEntity_Phone) XXX_Size() int {
	return xxx_messageInfo_UserEntity_Phone.Size(m)
}
func (m *UserEntity_Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_UserEntity_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_UserEntity_Phone proto.InternalMessageInfo

func (m *UserEntity_Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *UserEntity_Phone) GetType() UserEntity_Phone_PhoneType {
	if m != nil {
		return m.Type
	}
	return UserEntity_Phone_MOBILE
}

type CreateUserRequest struct {
	User                 *UserEntity `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{1}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *UserEntity {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_c9d9430a1717c0a3, []int{2}
}
func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (dst *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(dst, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*UserEntity)(nil), "user.UserEntity")
	proto.RegisterType((*UserEntity_Email)(nil), "user.UserEntity.Email")
	proto.RegisterType((*UserEntity_Phone)(nil), "user.UserEntity.Phone")
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
	proto.RegisterEnum("user.UserEntity_Phone_PhoneType", UserEntity_Phone_PhoneType_name, UserEntity_Phone_PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// Create user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// Create user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_c9d9430a1717c0a3) }

var fileDescriptor_user_c9d9430a1717c0a3 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x6d, 0xd2, 0x34, 0x6d, 0xa6, 0xfd, 0x95, 0xfe, 0xe6, 0x50, 0x97, 0xe2, 0x21, 0x04, 0x85,
	0x9e, 0x72, 0x88, 0x5e, 0x44, 0x50, 0xaa, 0xe6, 0x50, 0xa8, 0xb5, 0x2c, 0x7a, 0x96, 0xb4, 0x19,
	0x30, 0x60, 0xfe, 0xb8, 0x9b, 0x1c, 0x02, 0x7e, 0x54, 0x3f, 0x8c, 0x64, 0xb7, 0xb5, 0x62, 0xf5,
	0x12, 0x76, 0x66, 0xde, 0x7b, 0xf3, 0x66, 0x26, 0x00, 0x95, 0x24, 0xe1, 0x17, 0x22, 0x2f, 0x73,
	0xb4, 0x9a, 0xb7, 0xf7, 0x61, 0x02, 0x3c, 0x49, 0x12, 0x61, 0x56, 0x26, 0x65, 0x8d, 0x43, 0x30,
	0x93, 0x98, 0x19, 0xae, 0x31, 0xed, 0x70, 0x33, 0x89, 0x11, 0xc1, 0xca, 0xa2, 0x94, 0x98, 0xe9,
	0x1a, 0x53, 0x87, 0xab, 0x37, 0xfa, 0x60, 0x53, 0x1a, 0x25, 0xaf, 0x92, 0xb5, 0xdd, 0xf6, 0xb4,
	0x1f, 0x8c, 0x7d, 0xa5, 0xba, 0x57, 0xf1, 0xc3, 0xa6, 0xcc, 0xb7, 0x28, 0xbc, 0x84, 0x7f, 0xc5,
	0x4b, 0x9e, 0xd1, 0x73, 0x56, 0xa5, 0x6b, 0x12, 0x92, 0x59, 0x7f, 0xd0, 0x56, 0x0d, 0x8a, 0x0f,
	0x14, 0x78, 0xa9, 0xb1, 0x93, 0x6b, 0xe8, 0x28, 0x35, 0x64, 0xd0, 0x8d, 0xe2, 0x58, 0x90, 0x94,
	0xca, 0x9e, 0xc3, 0x77, 0x21, 0x1e, 0x83, 0x93, 0xc8, 0x95, 0x48, 0xd2, 0x48, 0xd4, 0xca, 0x68,
	0x8f, 0xef, 0x13, 0x93, 0x77, 0xe8, 0x28, 0x5d, 0x1c, 0x83, 0xad, 0x0d, 0x6c, 0xf9, 0xdb, 0x08,
	0xcf, 0xc1, 0x2a, 0xeb, 0x42, 0x8f, 0x38, 0x0c, 0xdc, 0xdf, 0x5d, 0xe9, 0xef, 0x63, 0x5d, 0x10,
	0x57, 0x68, 0xef, 0x14, 0x9c, 0xaf, 0x14, 0x02, 0xd8, 0xf7, 0x0f, 0x37, 0xf3, 0x45, 0x38, 0x6a,
	0xe1, 0x00, 0x7a, 0x8b, 0xd9, 0xf2, 0x6e, 0x31, 0x5f, 0x86, 0x23, 0xc3, 0xbb, 0x80, 0xff, 0xb7,
	0x82, 0xa2, 0x92, 0x1a, 0x41, 0x4e, 0x6f, 0x15, 0xc9, 0x12, 0x4f, 0x40, 0xed, 0x5e, 0xf9, 0xe8,
	0x07, 0xa3, 0x9f, 0x1d, 0xb9, 0xbe, 0xcc, 0x15, 0xe0, 0x77, 0xaa, 0x2c, 0xf2, 0x4c, 0xd2, 0xc1,
	0x81, 0x18, 0x74, 0x65, 0xb5, 0xd9, 0x34, 0x6b, 0xd1, 0xa3, 0xef, 0xc2, 0x60, 0x0e, 0x56, 0xc3,
	0xc4, 0x19, 0xc0, 0x5e, 0x07, 0x8f, 0x74, 0xb7, 0x03, 0x53, 0x13, 0x76, 0x58, 0xd0, 0x2d, 0xbd,
	0xd6, 0xda, 0x56, 0x7f, 0xcc, 0xd9, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x9b, 0x2e, 0x7b,
	0x3f, 0x02, 0x00, 0x00,
}