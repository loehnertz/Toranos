// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-management.proto

package user_management

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterCustomerRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=LastName,proto3" json:"LastName,omitempty"`
	LicenseId            string   `protobuf:"bytes,5,opt,name=LicenseId,proto3" json:"LicenseId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCustomerRequest) Reset()         { *m = RegisterCustomerRequest{} }
func (m *RegisterCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterCustomerRequest) ProtoMessage()    {}
func (*RegisterCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{0}
}
func (m *RegisterCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCustomerRequest.Unmarshal(m, b)
}
func (m *RegisterCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCustomerRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCustomerRequest.Merge(dst, src)
}
func (m *RegisterCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterCustomerRequest.Size(m)
}
func (m *RegisterCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCustomerRequest proto.InternalMessageInfo

func (m *RegisterCustomerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterCustomerRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterCustomerRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterCustomerRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterCustomerRequest) GetLicenseId() string {
	if m != nil {
		return m.LicenseId
	}
	return ""
}

type RegisterCustomerResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterCustomerResponse) Reset()         { *m = RegisterCustomerResponse{} }
func (m *RegisterCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterCustomerResponse) ProtoMessage()    {}
func (*RegisterCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{1}
}
func (m *RegisterCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterCustomerResponse.Unmarshal(m, b)
}
func (m *RegisterCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterCustomerResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterCustomerResponse.Merge(dst, src)
}
func (m *RegisterCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterCustomerResponse.Size(m)
}
func (m *RegisterCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterCustomerResponse proto.InternalMessageInfo

func (m *RegisterCustomerResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *RegisterCustomerResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type IssueUserTokenRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueUserTokenRequest) Reset()         { *m = IssueUserTokenRequest{} }
func (m *IssueUserTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IssueUserTokenRequest) ProtoMessage()    {}
func (*IssueUserTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{2}
}
func (m *IssueUserTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueUserTokenRequest.Unmarshal(m, b)
}
func (m *IssueUserTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueUserTokenRequest.Marshal(b, m, deterministic)
}
func (dst *IssueUserTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueUserTokenRequest.Merge(dst, src)
}
func (m *IssueUserTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IssueUserTokenRequest.Size(m)
}
func (m *IssueUserTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueUserTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueUserTokenRequest proto.InternalMessageInfo

func (m *IssueUserTokenRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *IssueUserTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type IssueUserTokenResponse struct {
	Successful           bool     `protobuf:"varint,1,opt,name=Successful,proto3" json:"Successful,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueUserTokenResponse) Reset()         { *m = IssueUserTokenResponse{} }
func (m *IssueUserTokenResponse) String() string { return proto.CompactTextString(m) }
func (*IssueUserTokenResponse) ProtoMessage()    {}
func (*IssueUserTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{3}
}
func (m *IssueUserTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueUserTokenResponse.Unmarshal(m, b)
}
func (m *IssueUserTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueUserTokenResponse.Marshal(b, m, deterministic)
}
func (dst *IssueUserTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueUserTokenResponse.Merge(dst, src)
}
func (m *IssueUserTokenResponse) XXX_Size() int {
	return xxx_messageInfo_IssueUserTokenResponse.Size(m)
}
func (m *IssueUserTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueUserTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueUserTokenResponse proto.InternalMessageInfo

func (m *IssueUserTokenResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func (m *IssueUserTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateUserRequest) Reset()         { *m = AuthenticateUserRequest{} }
func (m *AuthenticateUserRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUserRequest) ProtoMessage()    {}
func (*AuthenticateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{4}
}
func (m *AuthenticateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateUserRequest.Unmarshal(m, b)
}
func (m *AuthenticateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateUserRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateUserRequest.Merge(dst, src)
}
func (m *AuthenticateUserRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateUserRequest.Size(m)
}
func (m *AuthenticateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateUserRequest proto.InternalMessageInfo

func (m *AuthenticateUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthenticateUserResponse struct {
	Authenticated        bool     `protobuf:"varint,1,opt,name=Authenticated,proto3" json:"Authenticated,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateUserResponse) Reset()         { *m = AuthenticateUserResponse{} }
func (m *AuthenticateUserResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateUserResponse) ProtoMessage()    {}
func (*AuthenticateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_management_5286a990f96ef7d6, []int{5}
}
func (m *AuthenticateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateUserResponse.Unmarshal(m, b)
}
func (m *AuthenticateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateUserResponse.Marshal(b, m, deterministic)
}
func (dst *AuthenticateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateUserResponse.Merge(dst, src)
}
func (m *AuthenticateUserResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateUserResponse.Size(m)
}
func (m *AuthenticateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateUserResponse proto.InternalMessageInfo

func (m *AuthenticateUserResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

func (m *AuthenticateUserResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthenticateUserResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterCustomerRequest)(nil), "RegisterCustomerRequest")
	proto.RegisterType((*RegisterCustomerResponse)(nil), "RegisterCustomerResponse")
	proto.RegisterType((*IssueUserTokenRequest)(nil), "IssueUserTokenRequest")
	proto.RegisterType((*IssueUserTokenResponse)(nil), "IssueUserTokenResponse")
	proto.RegisterType((*AuthenticateUserRequest)(nil), "AuthenticateUserRequest")
	proto.RegisterType((*AuthenticateUserResponse)(nil), "AuthenticateUserResponse")
}

func init() {
	proto.RegisterFile("user-management.proto", fileDescriptor_user_management_5286a990f96ef7d6)
}

var fileDescriptor_user_management_5286a990f96ef7d6 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xdd, 0x4a, 0xe3, 0x40,
	0x14, 0x6e, 0xba, 0x6d, 0x69, 0x0f, 0x6c, 0x59, 0x86, 0x6d, 0x3b, 0x0d, 0xcb, 0x22, 0xc1, 0x0b,
	0x6f, 0x1c, 0x41, 0x9f, 0x40, 0x8a, 0x42, 0xb0, 0x96, 0x12, 0xf5, 0x01, 0xc6, 0xe4, 0x58, 0x83,
	0x4d, 0xa6, 0xce, 0x99, 0xc1, 0x07, 0xf2, 0xf1, 0x7c, 0x09, 0xc9, 0x4f, 0x6b, 0xfa, 0x13, 0x6f,
	0xbc, 0xcb, 0xf9, 0x0e, 0xdf, 0xf0, 0xfd, 0x9c, 0xc0, 0xc0, 0x12, 0xea, 0xd3, 0x44, 0xa6, 0x72,
	0x81, 0x09, 0xa6, 0x46, 0xac, 0xb4, 0x32, 0xca, 0x7b, 0x77, 0x60, 0x14, 0xe0, 0x22, 0x26, 0x83,
	0x7a, 0x62, 0xc9, 0xa8, 0x04, 0x75, 0x80, 0xaf, 0x16, 0xc9, 0xb0, 0xbf, 0xd0, 0xbe, 0x4a, 0x64,
	0xbc, 0xe4, 0xce, 0x91, 0x73, 0xd2, 0x0b, 0x8a, 0x81, 0xb9, 0xd0, 0x9d, 0x4b, 0xa2, 0x37, 0xa5,
	0x23, 0xde, 0xcc, 0x17, 0x9b, 0x99, 0xfd, 0x83, 0xde, 0x75, 0xac, 0xc9, 0xcc, 0x64, 0x82, 0xfc,
	0x57, 0xbe, 0xfc, 0x02, 0x32, 0xe6, 0x54, 0x96, 0xcb, 0x56, 0xc1, 0x5c, 0xcf, 0x19, 0x73, 0x1a,
	0x87, 0x98, 0x12, 0xfa, 0x11, 0x6f, 0x17, 0xcc, 0x0d, 0xe0, 0xcd, 0x81, 0xef, 0x8b, 0xa4, 0x95,
	0x4a, 0x09, 0xd9, 0x7f, 0x80, 0x3b, 0x1b, 0x86, 0x48, 0xf4, 0x64, 0x0b, 0xa9, 0xdd, 0xa0, 0x82,
	0x64, 0x2e, 0xee, 0xd5, 0x0b, 0xa6, 0xa5, 0xd8, 0x62, 0xf0, 0x6e, 0x60, 0xe0, 0x13, 0x59, 0x7c,
	0x20, 0xd4, 0x39, 0xb2, 0x36, 0x3d, 0x84, 0x4e, 0x86, 0xf9, 0x51, 0xe9, 0xba, 0x9c, 0xbe, 0xb3,
	0xed, 0xcd, 0x60, 0xb8, 0xfb, 0xd8, 0x8f, 0xc4, 0x9d, 0xc1, 0xe8, 0xd2, 0x9a, 0x67, 0x4c, 0x4d,
	0x1c, 0x4a, 0x93, 0x3f, 0x5b, 0xe9, 0xa4, 0x20, 0x38, 0x55, 0xc2, 0x12, 0xf8, 0x3e, 0xa1, 0x94,
	0x70, 0x0c, 0xbf, 0xab, 0xbb, 0xa8, 0x54, 0xb1, 0x0d, 0x56, 0x6c, 0x37, 0xb7, 0x6c, 0x33, 0x68,
	0x05, 0x6a, 0xb9, 0x2e, 0x33, 0xff, 0x3e, 0xff, 0x70, 0xa0, 0x9f, 0xad, 0x6f, 0x37, 0xc7, 0xc4,
	0x7c, 0xf8, 0xb3, 0x5b, 0x10, 0xe3, 0xa2, 0xe6, 0xb0, 0xdc, 0xb1, 0xa8, 0x6b, 0xd3, 0x6b, 0xb0,
	0x09, 0xf4, 0xb7, 0xc3, 0x64, 0x43, 0x71, 0xb0, 0x2a, 0x77, 0x24, 0x0e, 0xa7, 0xee, 0x35, 0x32,
	0x3d, 0xbb, 0x81, 0x30, 0x2e, 0x6a, 0x42, 0x75, 0xc7, 0xa2, 0x2e, 0x3d, 0xaf, 0xf1, 0xd8, 0xc9,
	0x7f, 0x94, 0x8b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x64, 0xd6, 0x49, 0x41, 0x03, 0x00,
	0x00,
}
