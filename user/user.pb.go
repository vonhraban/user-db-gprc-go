// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
	return fileDescriptor_116e343673f7ffaf, []int{1, 1, 0}
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
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
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity.Unmarshal(m, b)
}
func (m *UserEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity.Marshal(b, m, deterministic)
}
func (m *UserEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity.Merge(m, src)
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
	return fileDescriptor_116e343673f7ffaf, []int{1, 0}
}

func (m *UserEntity_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity_Email.Unmarshal(m, b)
}
func (m *UserEntity_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity_Email.Marshal(b, m, deterministic)
}
func (m *UserEntity_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity_Email.Merge(m, src)
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
	return fileDescriptor_116e343673f7ffaf, []int{1, 1}
}

func (m *UserEntity_Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserEntity_Phone.Unmarshal(m, b)
}
func (m *UserEntity_Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserEntity_Phone.Marshal(b, m, deterministic)
}
func (m *UserEntity_Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserEntity_Phone.Merge(m, src)
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
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
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
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
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

type GetUserByIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIdRequest) Reset()         { *m = GetUserByIdRequest{} }
func (m *GetUserByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdRequest) ProtoMessage()    {}
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *GetUserByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdRequest.Unmarshal(m, b)
}
func (m *GetUserByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdRequest.Merge(m, src)
}
func (m *GetUserByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdRequest.Size(m)
}
func (m *GetUserByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdRequest proto.InternalMessageInfo

func (m *GetUserByIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	User                 *UserEntity `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                *Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *UserEntity {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("user.UserEntity_Phone_PhoneType", UserEntity_Phone_PhoneType_name, UserEntity_Phone_PhoneType_value)
	proto.RegisterType((*Error)(nil), "user.Error")
	proto.RegisterType((*UserEntity)(nil), "user.UserEntity")
	proto.RegisterType((*UserEntity_Email)(nil), "user.UserEntity.Email")
	proto.RegisterType((*UserEntity_Phone)(nil), "user.UserEntity.Phone")
	proto.RegisterType((*CreateUserRequest)(nil), "user.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "user.CreateUserResponse")
	proto.RegisterType((*GetUserByIdRequest)(nil), "user.GetUserByIdRequest")
	proto.RegisterType((*GetUserResponse)(nil), "user.GetUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0x64, 0x93, 0xec, 0xf6, 0x66, 0x5d, 0xeb, 0x05, 0xd7, 0x10, 0x7c, 0xe8, 0x86, 0x15,
	0xf6, 0x29, 0x0f, 0xd1, 0x17, 0x11, 0xd4, 0x56, 0x83, 0x14, 0x6a, 0x2d, 0x83, 0xbe, 0xf8, 0x22,
	0x69, 0x73, 0xd1, 0x80, 0xf9, 0x70, 0x26, 0x79, 0x08, 0xf8, 0x0f, 0xfc, 0x8b, 0xfe, 0x18, 0x99,
	0x99, 0xa4, 0xa9, 0x8d, 0xc2, 0xbe, 0x84, 0xb9, 0x73, 0xcf, 0xb9, 0xe7, 0x9c, 0x99, 0x09, 0x40,
	0x23, 0x88, 0x87, 0x15, 0x2f, 0xeb, 0x12, 0x2d, 0xb9, 0x0e, 0xae, 0xc1, 0x8e, 0x39, 0x2f, 0x39,
	0x7a, 0x70, 0x96, 0x93, 0x10, 0xc9, 0x57, 0xf2, 0x8c, 0x99, 0x71, 0x3b, 0x61, 0x7d, 0x19, 0xfc,
	0x36, 0x01, 0x3e, 0x09, 0xe2, 0x71, 0x51, 0x67, 0x75, 0x8b, 0x97, 0x60, 0x66, 0xa9, 0xc2, 0xd8,
	0xcc, 0xcc, 0x52, 0x44, 0xb0, 0x8a, 0x24, 0x27, 0xcf, 0x54, 0x2c, 0xb5, 0xc6, 0x10, 0x1c, 0xca,
	0x93, 0xec, 0xbb, 0xf0, 0x4e, 0x67, 0xa7, 0xb7, 0x6e, 0x74, 0x15, 0x2a, 0xe1, 0x61, 0x4a, 0x18,
	0xcb, 0x36, 0xeb, 0x50, 0xf8, 0x02, 0xee, 0x55, 0xdf, 0xca, 0x82, 0xbe, 0x14, 0x4d, 0xbe, 0x25,
	0x2e, 0x3c, 0xeb, 0x3f, 0xb4, 0x8d, 0x44, 0xb1, 0x0b, 0x05, 0x5e, 0x6b, 0xac, 0xff, 0x0a, 0x6c,
	0x35, 0x4d, 0x46, 0x48, 0xd2, 0x94, 0x93, 0x10, 0x7d, 0x84, 0xae, 0xc4, 0xc7, 0x30, 0xc9, 0xc4,
	0x86, 0x67, 0x79, 0xc2, 0x5b, 0x65, 0xf4, 0x9c, 0x0d, 0x1b, 0xfe, 0x4f, 0xb0, 0xd5, 0x5c, 0xbc,
	0x02, 0x47, 0x1b, 0xe8, 0xf8, 0x5d, 0x85, 0xcf, 0xc0, 0xaa, 0xdb, 0x4a, 0x47, 0xbc, 0x8c, 0x66,
	0xff, 0x76, 0xa5, 0xbf, 0x1f, 0xdb, 0x8a, 0x98, 0x42, 0x07, 0x4f, 0x60, 0xb2, 0xdf, 0x42, 0x00,
	0xe7, 0xfd, 0x87, 0xc5, 0x72, 0x15, 0x4f, 0x4f, 0xf0, 0x02, 0xce, 0x57, 0xf3, 0xf5, 0xdb, 0xd5,
	0x72, 0x1d, 0x4f, 0x8d, 0xe0, 0x39, 0x3c, 0x78, 0xc3, 0x29, 0xa9, 0x49, 0x0e, 0x64, 0xf4, 0xa3,
	0x21, 0x51, 0xe3, 0x0d, 0xa8, 0xeb, 0x51, 0x3e, 0xdc, 0x68, 0x7a, 0xac, 0xc8, 0xf4, 0xe5, 0xbd,
	0x04, 0x3c, 0xa4, 0x8a, 0xaa, 0x2c, 0x04, 0x8d, 0x2e, 0xc8, 0x83, 0x33, 0xd1, 0xec, 0x76, 0xf2,
	0x58, 0x74, 0xf4, 0xbe, 0x0c, 0x6e, 0x00, 0xdf, 0x51, 0x2d, 0xc9, 0x8b, 0x76, 0x99, 0xf6, 0xda,
	0x47, 0xfc, 0xe0, 0x33, 0xdc, 0xef, 0x50, 0x7b, 0x89, 0x3b, 0xd9, 0xc3, 0x6b, 0xb0, 0x49, 0xbe,
	0x2d, 0x25, 0xeb, 0x46, 0xae, 0x86, 0xa9, 0xe7, 0xc6, 0x74, 0x27, 0xfa, 0x65, 0x80, 0x25, 0x79,
	0x38, 0x07, 0x18, 0xa2, 0xe0, 0x23, 0x0d, 0x1d, 0x9d, 0x8b, 0xef, 0x8d, 0x1b, 0xda, 0x52, 0x70,
	0x82, 0xaf, 0xc1, 0x3d, 0x48, 0x83, 0x1d, 0x74, 0x1c, 0xd0, 0x7f, 0xf8, 0x57, 0x67, 0x98, 0xb0,
	0x75, 0xd4, 0x9f, 0xf1, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x8d, 0x87, 0xa5, 0x27,
	0x03, 0x00, 0x00,
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
	// Retrieve a user by id
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
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

func (c *userClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/user.User/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// Create user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Retrieve a user by id
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserResponse, error)
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

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*GetUserByIdRequest))
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
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
