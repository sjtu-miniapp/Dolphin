// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type OnLoginRequest struct {
	Code                 *string  `protobuf:"bytes,1,req,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnLoginRequest) Reset()         { *m = OnLoginRequest{} }
func (m *OnLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OnLoginRequest) ProtoMessage()    {}
func (*OnLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *OnLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnLoginRequest.Unmarshal(m, b)
}
func (m *OnLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnLoginRequest.Marshal(b, m, deterministic)
}
func (m *OnLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnLoginRequest.Merge(m, src)
}
func (m *OnLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OnLoginRequest.Size(m)
}
func (m *OnLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnLoginRequest proto.InternalMessageInfo

func (m *OnLoginRequest) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

type OnLoginResponse struct {
	Openid               *string  `protobuf:"bytes,1,req,name=openid" json:"openid,omitempty"`
	Sid                  *string  `protobuf:"bytes,2,req,name=sid" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnLoginResponse) Reset()         { *m = OnLoginResponse{} }
func (m *OnLoginResponse) String() string { return proto.CompactTextString(m) }
func (*OnLoginResponse) ProtoMessage()    {}
func (*OnLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *OnLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnLoginResponse.Unmarshal(m, b)
}
func (m *OnLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnLoginResponse.Marshal(b, m, deterministic)
}
func (m *OnLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnLoginResponse.Merge(m, src)
}
func (m *OnLoginResponse) XXX_Size() int {
	return xxx_messageInfo_OnLoginResponse.Size(m)
}
func (m *OnLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnLoginResponse proto.InternalMessageInfo

func (m *OnLoginResponse) GetOpenid() string {
	if m != nil && m.Openid != nil {
		return *m.Openid
	}
	return ""
}

func (m *OnLoginResponse) GetSid() string {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return ""
}

type PutUserRequest struct {
	Openid               *string  `protobuf:"bytes,1,req,name=openid" json:"openid,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Gender               *int32   `protobuf:"varint,3,opt,name=gender" json:"gender,omitempty"`
	Avatar               *string  `protobuf:"bytes,4,opt,name=avatar" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserRequest) Reset()         { *m = PutUserRequest{} }
func (m *PutUserRequest) String() string { return proto.CompactTextString(m) }
func (*PutUserRequest) ProtoMessage()    {}
func (*PutUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *PutUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserRequest.Unmarshal(m, b)
}
func (m *PutUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserRequest.Marshal(b, m, deterministic)
}
func (m *PutUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserRequest.Merge(m, src)
}
func (m *PutUserRequest) XXX_Size() int {
	return xxx_messageInfo_PutUserRequest.Size(m)
}
func (m *PutUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserRequest proto.InternalMessageInfo

func (m *PutUserRequest) GetOpenid() string {
	if m != nil && m.Openid != nil {
		return *m.Openid
	}
	return ""
}

func (m *PutUserRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PutUserRequest) GetGender() int32 {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return 0
}

func (m *PutUserRequest) GetAvatar() string {
	if m != nil && m.Avatar != nil {
		return *m.Avatar
	}
	return ""
}

type PutUserResponse struct {
	Err                  *int32   `protobuf:"varint,1,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutUserResponse) Reset()         { *m = PutUserResponse{} }
func (m *PutUserResponse) String() string { return proto.CompactTextString(m) }
func (*PutUserResponse) ProtoMessage()    {}
func (*PutUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *PutUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutUserResponse.Unmarshal(m, b)
}
func (m *PutUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutUserResponse.Marshal(b, m, deterministic)
}
func (m *PutUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutUserResponse.Merge(m, src)
}
func (m *PutUserResponse) XXX_Size() int {
	return xxx_messageInfo_PutUserResponse.Size(m)
}
func (m *PutUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutUserResponse proto.InternalMessageInfo

func (m *PutUserResponse) GetErr() int32 {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return 0
}

type GetUserRequest struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type GetUserResponse struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Gender               *int32   `protobuf:"varint,2,opt,name=gender" json:"gender,omitempty"`
	Avatar               *string  `protobuf:"bytes,3,opt,name=avatar" json:"avatar,omitempty"`
	SelfGroupId          *int32   `protobuf:"varint,4,req,name=self_group_id,json=selfGroupId" json:"self_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
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

func (m *GetUserResponse) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GetUserResponse) GetGender() int32 {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return 0
}

func (m *GetUserResponse) GetAvatar() string {
	if m != nil && m.Avatar != nil {
		return *m.Avatar
	}
	return ""
}

func (m *GetUserResponse) GetSelfGroupId() int32 {
	if m != nil && m.SelfGroupId != nil {
		return *m.SelfGroupId
	}
	return 0
}

type CheckAuthRequest struct {
	Openid               *string  `protobuf:"bytes,1,req,name=openid" json:"openid,omitempty"`
	Sid                  *string  `protobuf:"bytes,2,req,name=sid" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthRequest) Reset()         { *m = CheckAuthRequest{} }
func (m *CheckAuthRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAuthRequest) ProtoMessage()    {}
func (*CheckAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *CheckAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthRequest.Unmarshal(m, b)
}
func (m *CheckAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthRequest.Marshal(b, m, deterministic)
}
func (m *CheckAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthRequest.Merge(m, src)
}
func (m *CheckAuthRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAuthRequest.Size(m)
}
func (m *CheckAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthRequest proto.InternalMessageInfo

func (m *CheckAuthRequest) GetOpenid() string {
	if m != nil && m.Openid != nil {
		return *m.Openid
	}
	return ""
}

func (m *CheckAuthRequest) GetSid() string {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return ""
}

type CheckAuthResponse struct {
	Ok                   *bool    `protobuf:"varint,1,req,name=ok" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthResponse) Reset()         { *m = CheckAuthResponse{} }
func (m *CheckAuthResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAuthResponse) ProtoMessage()    {}
func (*CheckAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *CheckAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthResponse.Unmarshal(m, b)
}
func (m *CheckAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthResponse.Marshal(b, m, deterministic)
}
func (m *CheckAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthResponse.Merge(m, src)
}
func (m *CheckAuthResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAuthResponse.Size(m)
}
func (m *CheckAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthResponse proto.InternalMessageInfo

func (m *CheckAuthResponse) GetOk() bool {
	if m != nil && m.Ok != nil {
		return *m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*OnLoginRequest)(nil), "OnLoginRequest")
	proto.RegisterType((*OnLoginResponse)(nil), "OnLoginResponse")
	proto.RegisterType((*PutUserRequest)(nil), "PutUserRequest")
	proto.RegisterType((*PutUserResponse)(nil), "PutUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "GetUserResponse")
	proto.RegisterType((*CheckAuthRequest)(nil), "CheckAuthRequest")
	proto.RegisterType((*CheckAuthResponse)(nil), "CheckAuthResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0xcb, 0x42, 0xab, 0x1d, 0x23, 0xd0, 0x3d, 0x18, 0xc2, 0x89, 0x6c, 0x3d, 0xf4, 0xb4,
	0x07, 0xe3, 0xad, 0x5e, 0xd4, 0x43, 0x63, 0x62, 0xa2, 0x21, 0xf1, 0xe2, 0xa5, 0xc1, 0xee, 0xd8,
	0x92, 0x56, 0x16, 0x17, 0x30, 0xf1, 0x3f, 0xfa, 0xa3, 0xcc, 0xd2, 0x2d, 0x29, 0xd4, 0xc6, 0xdb,
	0xcc, 0xe4, 0xbd, 0x9d, 0x8f, 0x37, 0x00, 0x24, 0x55, 0xb9, 0xe2, 0xb9, 0x92, 0xa5, 0x64, 0x97,
	0xe0, 0x3e, 0x65, 0x8f, 0x72, 0x99, 0x66, 0x31, 0x7e, 0x56, 0x58, 0x94, 0x94, 0x82, 0xb3, 0x90,
	0x02, 0x03, 0x2b, 0x22, 0x93, 0x61, 0x5c, 0xd7, 0x6c, 0x0a, 0x5e, 0xa3, 0x2a, 0x72, 0x99, 0x15,
	0x48, 0x2f, 0x60, 0x20, 0x73, 0xcc, 0x52, 0x61, 0x84, 0xa6, 0xa3, 0x3e, 0xd8, 0x45, 0x2a, 0x02,
	0x52, 0x0f, 0x75, 0xc9, 0x36, 0xe0, 0x3e, 0x57, 0xe5, 0x4b, 0x81, 0x6a, 0xb7, 0xe2, 0x98, 0x97,
	0x82, 0x93, 0x25, 0x1f, 0x68, 0xcc, 0x75, 0xad, 0xb5, 0x4b, 0xcc, 0x04, 0xaa, 0xc0, 0x8e, 0xac,
	0x49, 0x3f, 0x36, 0x9d, 0x9e, 0x27, 0x5f, 0x49, 0x99, 0xa8, 0xc0, 0x89, 0x2c, 0xfd, 0xc6, 0xb6,
	0x63, 0x63, 0xf0, 0x9a, 0x6d, 0x06, 0xd5, 0x07, 0x1b, 0x95, 0x0a, 0xac, 0xda, 0xaf, 0x4b, 0x16,
	0x81, 0x3b, 0xc3, 0x16, 0x92, 0x0b, 0xa4, 0xc1, 0x21, 0xa9, 0x60, 0xdf, 0xe0, 0x35, 0x0a, 0xf3,
	0xcc, 0x8e, 0xce, 0xfa, 0x93, 0x8e, 0x1c, 0xa1, 0xb3, 0xf7, 0xe9, 0x28, 0x83, 0xf3, 0x02, 0x37,
	0xef, 0xf3, 0xa5, 0x92, 0x55, 0x3e, 0x4f, 0x45, 0xe0, 0x44, 0x64, 0xd2, 0x8f, 0xcf, 0xf4, 0x70,
	0xa6, 0x67, 0x0f, 0x82, 0xdd, 0x80, 0x7f, 0xbf, 0xc2, 0xc5, 0xfa, 0xb6, 0x2a, 0x57, 0xff, 0x25,
	0x76, 0x98, 0xf6, 0x18, 0x46, 0x7b, 0x6e, 0x83, 0xee, 0x02, 0x91, 0xeb, 0xda, 0x7a, 0x1a, 0x13,
	0xb9, 0xbe, 0xfa, 0xb1, 0xc0, 0xd1, 0x02, 0xca, 0xe1, 0xc4, 0x1c, 0x96, 0x7a, 0xbc, 0xfd, 0x23,
	0x84, 0x3e, 0xef, 0xdc, 0x9c, 0xf5, 0xb4, 0xde, 0xa4, 0x4b, 0x3d, 0xde, 0xbe, 0x6a, 0xe8, 0xf3,
	0x4e, 0xf0, 0x5b, 0xbd, 0x89, 0x91, 0x7a, 0xbc, 0x1d, 0x79, 0xe8, 0xf3, 0x4e, 0xc2, 0xac, 0x47,
	0xaf, 0x61, 0xd8, 0xd0, 0xd3, 0x11, 0xef, 0xe6, 0x10, 0x52, 0x7e, 0xf0, 0x71, 0xac, 0x77, 0x37,
	0x78, 0x75, 0xf8, 0x34, 0x7f, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xf7, 0x73, 0x8f, 0xd9,
	0x02, 0x00, 0x00,
}
