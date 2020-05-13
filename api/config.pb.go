// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package api

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

// Config file structure
type Config struct {
	Source               *SourceRepo `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target               *TargetRepo `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSource() *SourceRepo {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Config) GetTarget() *TargetRepo {
	if m != nil {
		return m.Target
	}
	return nil
}

// Generic repo representation
type Repo struct {
	// Types that are valid to be assigned to Repo:
	//	*Repo_SourceRepo
	//	*Repo_TargetRepo
	Repo                 isRepo_Repo `protobuf_oneof:"repo"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

type isRepo_Repo interface {
	isRepo_Repo()
}

type Repo_SourceRepo struct {
	SourceRepo *SourceRepo `protobuf:"bytes,1,opt,name=source_repo,json=sourceRepo,proto3,oneof"`
}

type Repo_TargetRepo struct {
	TargetRepo *TargetRepo `protobuf:"bytes,2,opt,name=target_repo,json=targetRepo,proto3,oneof"`
}

func (*Repo_SourceRepo) isRepo_Repo() {}

func (*Repo_TargetRepo) isRepo_Repo() {}

func (m *Repo) GetRepo() isRepo_Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *Repo) GetSourceRepo() *SourceRepo {
	if x, ok := m.GetRepo().(*Repo_SourceRepo); ok {
		return x.SourceRepo
	}
	return nil
}

func (m *Repo) GetTargetRepo() *TargetRepo {
	if x, ok := m.GetRepo().(*Repo_TargetRepo); ok {
		return x.TargetRepo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Repo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Repo_SourceRepo)(nil),
		(*Repo_TargetRepo)(nil),
	}
}

// Auth contains credentials to login to a chart repository
type Auth struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Auth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// SourceRepo contains the required information of the source chart repository
type SourceRepo struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Auth                 *Auth    `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceRepo) Reset()         { *m = SourceRepo{} }
func (m *SourceRepo) String() string { return proto.CompactTextString(m) }
func (*SourceRepo) ProtoMessage()    {}
func (*SourceRepo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}

func (m *SourceRepo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceRepo.Unmarshal(m, b)
}
func (m *SourceRepo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceRepo.Marshal(b, m, deterministic)
}
func (m *SourceRepo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceRepo.Merge(m, src)
}
func (m *SourceRepo) XXX_Size() int {
	return xxx_messageInfo_SourceRepo.Size(m)
}
func (m *SourceRepo) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceRepo.DiscardUnknown(m)
}

var xxx_messageInfo_SourceRepo proto.InternalMessageInfo

func (m *SourceRepo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SourceRepo) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *SourceRepo) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

// TargetRepo contains the required information of the target chart repository
type TargetRepo struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	ContainerRegistry    string   `protobuf:"bytes,2,opt,name=container_registry,json=containerRegistry,proto3" json:"container_registry,omitempty"`
	ContainerRepository  string   `protobuf:"bytes,3,opt,name=container_repository,json=containerRepository,proto3" json:"container_repository,omitempty"`
	Auth                 *Auth    `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetRepo) Reset()         { *m = TargetRepo{} }
func (m *TargetRepo) String() string { return proto.CompactTextString(m) }
func (*TargetRepo) ProtoMessage()    {}
func (*TargetRepo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}

func (m *TargetRepo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRepo.Unmarshal(m, b)
}
func (m *TargetRepo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRepo.Marshal(b, m, deterministic)
}
func (m *TargetRepo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRepo.Merge(m, src)
}
func (m *TargetRepo) XXX_Size() int {
	return xxx_messageInfo_TargetRepo.Size(m)
}
func (m *TargetRepo) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRepo.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRepo proto.InternalMessageInfo

func (m *TargetRepo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *TargetRepo) GetContainerRegistry() string {
	if m != nil {
		return m.ContainerRegistry
	}
	return ""
}

func (m *TargetRepo) GetContainerRepository() string {
	if m != nil {
		return m.ContainerRepository
	}
	return ""
}

func (m *TargetRepo) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "api.Config")
	proto.RegisterType((*Repo)(nil), "api.Repo")
	proto.RegisterType((*Auth)(nil), "api.Auth")
	proto.RegisterType((*SourceRepo)(nil), "api.SourceRepo")
	proto.RegisterType((*TargetRepo)(nil), "api.TargetRepo")
}

func init() {
	proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4)
}

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x3f, 0x4b, 0xc4, 0x30,
	0x14, 0xb7, 0x36, 0x14, 0xfb, 0x4e, 0x50, 0xa3, 0x43, 0x11, 0x04, 0xe9, 0xa2, 0x8b, 0x05, 0xcf,
	0x5d, 0x50, 0x17, 0x57, 0xa3, 0x93, 0xcb, 0x11, 0x6b, 0xec, 0x05, 0xb5, 0x09, 0x2f, 0x89, 0x72,
	0x9f, 0xc6, 0xaf, 0x2a, 0x49, 0x7a, 0xb9, 0x0e, 0xbd, 0x2d, 0xbf, 0xbf, 0xef, 0x57, 0x28, 0xec,
	0xb7, 0xaa, 0xff, 0x90, 0x5d, 0xa3, 0x51, 0x59, 0x45, 0x73, 0xae, 0x65, 0xfd, 0x0a, 0xc5, 0x43,
	0x20, 0xe9, 0x05, 0x14, 0x46, 0x39, 0x6c, 0x45, 0x95, 0x9d, 0x67, 0x97, 0xb3, 0xf9, 0x41, 0xc3,
	0xb5, 0x6c, 0x9e, 0x03, 0xc5, 0x84, 0x56, 0x6c, 0x90, 0xbd, 0xd1, 0x72, 0xec, 0x84, 0xad, 0x76,
	0x47, 0xc6, 0x97, 0x40, 0x45, 0x63, 0x94, 0xeb, 0x1f, 0x20, 0x1e, 0xd3, 0x39, 0xcc, 0x62, 0x74,
	0x81, 0x42, 0xab, 0x2d, 0xf5, 0x8f, 0x3b, 0x0c, 0x4c, 0x42, 0x3e, 0x13, 0x5b, 0x62, 0x66, 0xfa,
	0x92, 0xcf, 0xd8, 0x84, 0xee, 0x0b, 0x20, 0xde, 0x5c, 0xdf, 0x02, 0xb9, 0x73, 0x76, 0x49, 0x4f,
	0x61, 0xcf, 0x19, 0x81, 0x3d, 0xff, 0x8e, 0xdf, 0x54, 0xb2, 0x84, 0xbd, 0xa6, 0xb9, 0x31, 0xbf,
	0x0a, 0xdf, 0x43, 0x79, 0xc9, 0x12, 0xae, 0x9f, 0x00, 0x36, 0xbb, 0xe8, 0x21, 0xe4, 0x0e, 0xbf,
	0x86, 0x02, 0xff, 0xa4, 0x14, 0xc8, 0xa7, 0xec, 0xd7, 0xb9, 0xf0, 0xa6, 0x67, 0x40, 0xb8, 0xb3,
	0xcb, 0x2a, 0x0f, 0x43, 0xcb, 0x30, 0xd4, 0x8f, 0x60, 0x81, 0xae, 0xff, 0x32, 0x80, 0xcd, 0xee,
	0x89, 0xce, 0x2b, 0xa0, 0xad, 0xea, 0x2d, 0x97, 0xbd, 0xc0, 0x05, 0x8a, 0x4e, 0x1a, 0x8b, 0xab,
	0xe1, 0xc2, 0x51, 0x52, 0xd8, 0x20, 0xd0, 0x6b, 0x38, 0x19, 0xdb, 0xb5, 0x32, 0xd2, 0x2a, 0x5c,
	0x85, 0xf3, 0x25, 0x3b, 0x1e, 0x05, 0xd6, 0x52, 0x5a, 0x48, 0x26, 0x17, 0xbe, 0x15, 0xe1, 0xa7,
	0xb8, 0xf9, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xda, 0x8f, 0x85, 0xbd, 0x24, 0x02, 0x00, 0x00,
}
