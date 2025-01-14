// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_file.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigFileGroup struct {
	Id                   *wrapperspb.UInt64Value   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Comment              *wrapperspb.StringValue   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue   `protobuf:"bytes,6,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue   `protobuf:"bytes,7,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue   `protobuf:"bytes,8,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	FileCount            *wrapperspb.UInt64Value   `protobuf:"bytes,9,opt,name=fileCount,proto3" json:"fileCount,omitempty"`
	UserIds              []*wrapperspb.StringValue `protobuf:"bytes,10,rep,name=user_ids,proto3" json:"user_ids,omitempty"`
	GroupIds             []*wrapperspb.StringValue `protobuf:"bytes,11,rep,name=group_ids,proto3" json:"group_ids,omitempty"`
	RemoveUserIds        []*wrapperspb.StringValue `protobuf:"bytes,13,rep,name=remove_user_ids,proto3" json:"remove_user_ids,omitempty"`
	RemoveGroupIds       []*wrapperspb.StringValue `protobuf:"bytes,14,rep,name=remove_group_ids,proto3" json:"remove_group_ids,omitempty"`
	Editable             *wrapperspb.BoolValue     `protobuf:"bytes,15,opt,name=editable,proto3" json:"editable,omitempty"`
	Owner                *wrapperspb.StringValue   `protobuf:"bytes,16,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigFileGroup) Reset()         { *m = ConfigFileGroup{} }
func (m *ConfigFileGroup) String() string { return proto.CompactTextString(m) }
func (*ConfigFileGroup) ProtoMessage()    {}
func (*ConfigFileGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{0}
}
func (m *ConfigFileGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileGroup.Unmarshal(m, b)
}
func (m *ConfigFileGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileGroup.Marshal(b, m, deterministic)
}
func (dst *ConfigFileGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileGroup.Merge(dst, src)
}
func (m *ConfigFileGroup) XXX_Size() int {
	return xxx_messageInfo_ConfigFileGroup.Size(m)
}
func (m *ConfigFileGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileGroup proto.InternalMessageInfo

func (m *ConfigFileGroup) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileGroup) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileGroup) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileGroup) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileGroup) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileGroup) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFileGroup) GetFileCount() *wrapperspb.UInt64Value {
	if m != nil {
		return m.FileCount
	}
	return nil
}

func (m *ConfigFileGroup) GetUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.GroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveUserIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveUserIds
	}
	return nil
}

func (m *ConfigFileGroup) GetRemoveGroupIds() []*wrapperspb.StringValue {
	if m != nil {
		return m.RemoveGroupIds
	}
	return nil
}

func (m *ConfigFileGroup) GetEditable() *wrapperspb.BoolValue {
	if m != nil {
		return m.Editable
	}
	return nil
}

func (m *ConfigFileGroup) GetOwner() *wrapperspb.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

type ConfigFile struct {
	Id          *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group       *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Content     *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Format      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=format,proto3" json:"format,omitempty"`
	Comment     *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Status      *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Tags        []*ConfigFileTag        `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime  *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy    *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime  *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy    *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	ReleaseTime *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	ReleaseBy   *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=release_by,json=releaseBy,proto3" json:"release_by,omitempty"`
	// 是否为加密配置文件
	IsEncrypted          *wrapperspb.BoolValue `protobuf:"bytes,16,opt,name=is_encrypted,json=isEncrypted,proto3" json:"is_encrypted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigFile) Reset()         { *m = ConfigFile{} }
func (m *ConfigFile) String() string { return proto.CompactTextString(m) }
func (*ConfigFile) ProtoMessage()    {}
func (*ConfigFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{1}
}
func (m *ConfigFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFile.Unmarshal(m, b)
}
func (m *ConfigFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFile.Marshal(b, m, deterministic)
}
func (dst *ConfigFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFile.Merge(dst, src)
}
func (m *ConfigFile) XXX_Size() int {
	return xxx_messageInfo_ConfigFile.Size(m)
}
func (m *ConfigFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFile.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFile proto.InternalMessageInfo

func (m *ConfigFile) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFile) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFile) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFile) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFile) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFile) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFile) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFile) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFile) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFile) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFile) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFile) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFile) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

func (m *ConfigFile) GetReleaseTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseTime
	}
	return nil
}

func (m *ConfigFile) GetReleaseBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ReleaseBy
	}
	return nil
}

func (m *ConfigFile) GetIsEncrypted() *wrapperspb.BoolValue {
	if m != nil {
		return m.IsEncrypted
	}
	return nil
}

type ConfigFileTag struct {
	Key                  *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTag) Reset()         { *m = ConfigFileTag{} }
func (m *ConfigFileTag) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTag) ProtoMessage()    {}
func (*ConfigFileTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{2}
}
func (m *ConfigFileTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTag.Unmarshal(m, b)
}
func (m *ConfigFileTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTag.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTag.Merge(dst, src)
}
func (m *ConfigFileTag) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTag.Size(m)
}
func (m *ConfigFileTag) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTag.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTag proto.InternalMessageInfo

func (m *ConfigFileTag) GetKey() *wrapperspb.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ConfigFileTag) GetValue() *wrapperspb.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ConfigFileRelease struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=md5,proto3" json:"md5,omitempty"`
	Version              *wrapperspb.UInt64Value `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,12,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileRelease) Reset()         { *m = ConfigFileRelease{} }
func (m *ConfigFileRelease) String() string { return proto.CompactTextString(m) }
func (*ConfigFileRelease) ProtoMessage()    {}
func (*ConfigFileRelease) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{3}
}
func (m *ConfigFileRelease) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileRelease.Unmarshal(m, b)
}
func (m *ConfigFileRelease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileRelease.Marshal(b, m, deterministic)
}
func (dst *ConfigFileRelease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileRelease.Merge(dst, src)
}
func (m *ConfigFileRelease) XXX_Size() int {
	return xxx_messageInfo_ConfigFileRelease.Size(m)
}
func (m *ConfigFileRelease) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileRelease.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileRelease proto.InternalMessageInfo

func (m *ConfigFileRelease) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileRelease) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileRelease) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileRelease) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileRelease) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileRelease) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileRelease) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileRelease) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileRelease) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileRelease) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileRelease) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileReleaseHistory struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group                *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	FileName             *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Md5                  *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=md5,proto3" json:"md5,omitempty"`
	Type                 *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               *wrapperspb.StringValue `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []*ConfigFileTag        `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,15,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,16,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileReleaseHistory) Reset()         { *m = ConfigFileReleaseHistory{} }
func (m *ConfigFileReleaseHistory) String() string { return proto.CompactTextString(m) }
func (*ConfigFileReleaseHistory) ProtoMessage()    {}
func (*ConfigFileReleaseHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{4}
}
func (m *ConfigFileReleaseHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileReleaseHistory.Unmarshal(m, b)
}
func (m *ConfigFileReleaseHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileReleaseHistory.Marshal(b, m, deterministic)
}
func (dst *ConfigFileReleaseHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileReleaseHistory.Merge(dst, src)
}
func (m *ConfigFileReleaseHistory) XXX_Size() int {
	return xxx_messageInfo_ConfigFileReleaseHistory.Size(m)
}
func (m *ConfigFileReleaseHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileReleaseHistory.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileReleaseHistory proto.InternalMessageInfo

func (m *ConfigFileReleaseHistory) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetType() *wrapperspb.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetStatus() *wrapperspb.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetTags() []*ConfigFileTag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileReleaseHistory) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ConfigFileTemplate struct {
	Id                   *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Format               *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Comment              *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	CreateTime           *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	CreateBy             *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	ModifyTime           *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=modify_time,json=modifyTime,proto3" json:"modify_time,omitempty"`
	ModifyBy             *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=modify_by,json=modifyBy,proto3" json:"modify_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigFileTemplate) Reset()         { *m = ConfigFileTemplate{} }
func (m *ConfigFileTemplate) String() string { return proto.CompactTextString(m) }
func (*ConfigFileTemplate) ProtoMessage()    {}
func (*ConfigFileTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{5}
}
func (m *ConfigFileTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileTemplate.Unmarshal(m, b)
}
func (m *ConfigFileTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileTemplate.Marshal(b, m, deterministic)
}
func (dst *ConfigFileTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileTemplate.Merge(dst, src)
}
func (m *ConfigFileTemplate) XXX_Size() int {
	return xxx_messageInfo_ConfigFileTemplate.Size(m)
}
func (m *ConfigFileTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileTemplate proto.InternalMessageInfo

func (m *ConfigFileTemplate) GetId() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ConfigFileTemplate) GetName() *wrapperspb.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigFileTemplate) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ConfigFileTemplate) GetFormat() *wrapperspb.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigFileTemplate) GetComment() *wrapperspb.StringValue {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateTime() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetCreateBy() *wrapperspb.StringValue {
	if m != nil {
		return m.CreateBy
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyTime() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyTime
	}
	return nil
}

func (m *ConfigFileTemplate) GetModifyBy() *wrapperspb.StringValue {
	if m != nil {
		return m.ModifyBy
	}
	return nil
}

type ClientConfigFileInfo struct {
	Namespace *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Group     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	FileName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Version   *wrapperspb.UInt64Value `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Md5       *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	// 是否为加密配置文件
	IsEncrypted *wrapperspb.BoolValue `protobuf:"bytes,7,opt,name=is_encrypted,json=isEncrypted,proto3" json:"is_encrypted,omitempty"`
	// 数据密钥，用于加密配置文件
	DataKey *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=data_key,json=dataKey,proto3" json:"data_key,omitempty"`
	// 公钥，用于加密数据密钥
	PublicKey            *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientConfigFileInfo) Reset()         { *m = ClientConfigFileInfo{} }
func (m *ClientConfigFileInfo) String() string { return proto.CompactTextString(m) }
func (*ClientConfigFileInfo) ProtoMessage()    {}
func (*ClientConfigFileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{6}
}
func (m *ClientConfigFileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientConfigFileInfo.Unmarshal(m, b)
}
func (m *ClientConfigFileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientConfigFileInfo.Marshal(b, m, deterministic)
}
func (dst *ClientConfigFileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConfigFileInfo.Merge(dst, src)
}
func (m *ClientConfigFileInfo) XXX_Size() int {
	return xxx_messageInfo_ClientConfigFileInfo.Size(m)
}
func (m *ClientConfigFileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConfigFileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConfigFileInfo proto.InternalMessageInfo

func (m *ClientConfigFileInfo) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ClientConfigFileInfo) GetGroup() *wrapperspb.StringValue {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ClientConfigFileInfo) GetFileName() *wrapperspb.StringValue {
	if m != nil {
		return m.FileName
	}
	return nil
}

func (m *ClientConfigFileInfo) GetContent() *wrapperspb.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ClientConfigFileInfo) GetVersion() *wrapperspb.UInt64Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ClientConfigFileInfo) GetMd5() *wrapperspb.StringValue {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *ClientConfigFileInfo) GetIsEncrypted() *wrapperspb.BoolValue {
	if m != nil {
		return m.IsEncrypted
	}
	return nil
}

func (m *ClientConfigFileInfo) GetDataKey() *wrapperspb.StringValue {
	if m != nil {
		return m.DataKey
	}
	return nil
}

func (m *ClientConfigFileInfo) GetPublicKey() *wrapperspb.StringValue {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type ClientWatchConfigFileRequest struct {
	ClientIp             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	ServiceName          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	WatchFiles           []*ClientConfigFileInfo `protobuf:"bytes,3,rep,name=watch_files,json=watchFiles,proto3" json:"watch_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClientWatchConfigFileRequest) Reset()         { *m = ClientWatchConfigFileRequest{} }
func (m *ClientWatchConfigFileRequest) String() string { return proto.CompactTextString(m) }
func (*ClientWatchConfigFileRequest) ProtoMessage()    {}
func (*ClientWatchConfigFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{7}
}
func (m *ClientWatchConfigFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Unmarshal(m, b)
}
func (m *ClientWatchConfigFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Marshal(b, m, deterministic)
}
func (dst *ClientWatchConfigFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientWatchConfigFileRequest.Merge(dst, src)
}
func (m *ClientWatchConfigFileRequest) XXX_Size() int {
	return xxx_messageInfo_ClientWatchConfigFileRequest.Size(m)
}
func (m *ClientWatchConfigFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientWatchConfigFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientWatchConfigFileRequest proto.InternalMessageInfo

func (m *ClientWatchConfigFileRequest) GetClientIp() *wrapperspb.StringValue {
	if m != nil {
		return m.ClientIp
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetServiceName() *wrapperspb.StringValue {
	if m != nil {
		return m.ServiceName
	}
	return nil
}

func (m *ClientWatchConfigFileRequest) GetWatchFiles() []*ClientConfigFileInfo {
	if m != nil {
		return m.WatchFiles
	}
	return nil
}

type ConfigFileExportRequest struct {
	Namespace            *wrapperspb.StringValue   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Groups               []*wrapperspb.StringValue `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Names                []*wrapperspb.StringValue `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ConfigFileExportRequest) Reset()         { *m = ConfigFileExportRequest{} }
func (m *ConfigFileExportRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigFileExportRequest) ProtoMessage()    {}
func (*ConfigFileExportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_file_7b3903662e8306db, []int{8}
}
func (m *ConfigFileExportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigFileExportRequest.Unmarshal(m, b)
}
func (m *ConfigFileExportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigFileExportRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigFileExportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigFileExportRequest.Merge(dst, src)
}
func (m *ConfigFileExportRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigFileExportRequest.Size(m)
}
func (m *ConfigFileExportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigFileExportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigFileExportRequest proto.InternalMessageInfo

func (m *ConfigFileExportRequest) GetNamespace() *wrapperspb.StringValue {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *ConfigFileExportRequest) GetGroups() []*wrapperspb.StringValue {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *ConfigFileExportRequest) GetNames() []*wrapperspb.StringValue {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigFileGroup)(nil), "v1.ConfigFileGroup")
	proto.RegisterType((*ConfigFile)(nil), "v1.ConfigFile")
	proto.RegisterType((*ConfigFileTag)(nil), "v1.ConfigFileTag")
	proto.RegisterType((*ConfigFileRelease)(nil), "v1.ConfigFileRelease")
	proto.RegisterType((*ConfigFileReleaseHistory)(nil), "v1.ConfigFileReleaseHistory")
	proto.RegisterType((*ConfigFileTemplate)(nil), "v1.ConfigFileTemplate")
	proto.RegisterType((*ClientConfigFileInfo)(nil), "v1.ClientConfigFileInfo")
	proto.RegisterType((*ClientWatchConfigFileRequest)(nil), "v1.ClientWatchConfigFileRequest")
	proto.RegisterType((*ConfigFileExportRequest)(nil), "v1.ConfigFileExportRequest")
}

func init() { proto.RegisterFile("config_file.proto", fileDescriptor_config_file_7b3903662e8306db) }

var fileDescriptor_config_file_7b3903662e8306db = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0xdd, 0x8a, 0xe4, 0x44,
	0x14, 0xa6, 0xff, 0xbb, 0x4f, 0xcf, 0x6f, 0x10, 0x0c, 0xcb, 0x22, 0x4b, 0x83, 0xe0, 0x85, 0xa4,
	0x77, 0xc6, 0x71, 0x74, 0x14, 0x11, 0x7a, 0xd8, 0x75, 0x87, 0x05, 0x91, 0x71, 0x55, 0xf0, 0xa6,
	0xa9, 0x4e, 0x9f, 0xce, 0x14, 0x26, 0xa9, 0x58, 0x55, 0xe9, 0x31, 0x8f, 0xe0, 0xcd, 0xe2, 0x95,
	0xaf, 0xe3, 0x53, 0xf8, 0x12, 0x3e, 0x84, 0x48, 0xfd, 0xc4, 0xf4, 0xcc, 0xb8, 0x4e, 0x75, 0x82,
	0x20, 0xba, 0x57, 0x4d, 0x27, 0xdf, 0x39, 0xa7, 0x7e, 0xbe, 0xef, 0x3b, 0x95, 0x82, 0xc3, 0x90,
	0xa5, 0x2b, 0x1a, 0xcd, 0x57, 0x34, 0xc6, 0x20, 0xe3, 0x4c, 0x32, 0xaf, 0xbd, 0x3e, 0x7a, 0xf0,
	0x56, 0xc4, 0x58, 0x14, 0xe3, 0x54, 0x3f, 0x59, 0xe4, 0xab, 0xe9, 0x35, 0x27, 0x59, 0x86, 0x5c,
	0x18, 0xcc, 0xe4, 0xc7, 0x01, 0xec, 0x9f, 0xeb, 0xc8, 0xa7, 0x34, 0xc6, 0xcf, 0x38, 0xcb, 0x33,
	0xef, 0x5d, 0x68, 0xd3, 0xa5, 0xdf, 0x7a, 0xd4, 0x7a, 0x67, 0x7c, 0xfc, 0x30, 0x30, 0x09, 0x82,
	0x32, 0x41, 0xf0, 0xd5, 0x45, 0x2a, 0x4f, 0x4f, 0xbe, 0x26, 0x71, 0x8e, 0x97, 0x6d, 0xba, 0xf4,
	0x1e, 0x43, 0x37, 0x25, 0x09, 0xfa, 0xed, 0x57, 0xe0, 0xbf, 0x94, 0x9c, 0xa6, 0x91, 0xc1, 0x6b,
	0xa4, 0xf7, 0x11, 0x8c, 0xd4, 0xaf, 0xc8, 0x48, 0x88, 0x7e, 0xc7, 0x21, 0xac, 0x82, 0x7b, 0xa7,
	0x30, 0x08, 0x59, 0x92, 0x60, 0x2a, 0xfd, 0xae, 0x43, 0x64, 0x09, 0xf6, 0x3e, 0x81, 0x71, 0xc8,
	0x91, 0x48, 0x9c, 0x4b, 0x9a, 0xa0, 0xdf, 0x73, 0x88, 0x05, 0x13, 0xf0, 0x82, 0x26, 0xe8, 0x9d,
	0xc1, 0xc8, 0x86, 0x2f, 0x0a, 0xbf, 0xef, 0x10, 0x3c, 0x34, 0xf0, 0x59, 0xa1, 0x2a, 0x27, 0x6c,
	0x49, 0x57, 0x85, 0xa9, 0x3c, 0x70, 0xa9, 0x6c, 0x02, 0xca, 0xca, 0x36, 0x7c, 0x51, 0xf8, 0x43,
	0x97, 0xca, 0x06, 0x3e, 0x2b, 0xd4, 0x3a, 0x2b, 0x36, 0x9c, 0xb3, 0x3c, 0x95, 0xfe, 0xc8, 0x61,
	0x3b, 0x2b, 0xb8, 0xf7, 0x21, 0x0c, 0x73, 0x81, 0x7c, 0x4e, 0x97, 0xc2, 0x87, 0x47, 0x9d, 0xfb,
	0xab, 0x96, 0x68, 0x55, 0x35, 0x52, 0x34, 0xd2, 0xa1, 0x63, 0x87, 0xd0, 0x0a, 0xee, 0x3d, 0x85,
	0x7d, 0x8e, 0x09, 0x5b, 0xe3, 0xfc, 0xcf, 0xe2, 0xbb, 0x0e, 0x19, 0x6e, 0x07, 0x79, 0xcf, 0xe0,
	0xc0, 0x3e, 0xaa, 0x86, 0xb2, 0xe7, 0x90, 0xe8, 0x4e, 0x94, 0x77, 0x0a, 0x43, 0x5c, 0x52, 0x49,
	0x16, 0x31, 0xfa, 0xfb, 0x7a, 0x09, 0x1f, 0xdc, 0xc9, 0x30, 0x63, 0x2c, 0xb6, 0xab, 0x50, 0x62,
	0xbd, 0x63, 0xe8, 0xb1, 0xeb, 0x14, 0xb9, 0x7f, 0xe0, 0xb0, 0x65, 0x06, 0x3a, 0xf9, 0x69, 0x00,
	0x50, 0x69, 0xf1, 0x5f, 0x2d, 0xc3, 0x63, 0xe8, 0xe9, 0x35, 0x72, 0x12, 0xa1, 0x81, 0x1a, 0xe9,
	0xa6, 0x52, 0x49, 0xb7, 0xe7, 0x26, 0x5d, 0x0d, 0xf6, 0x4e, 0xa0, 0xbf, 0x62, 0x3c, 0x21, 0xd2,
	0x49, 0x78, 0x16, 0xbb, 0x69, 0x14, 0x83, 0x6d, 0x8c, 0xe2, 0x04, 0xfa, 0x42, 0x12, 0x99, 0x0b,
	0x27, 0xb1, 0x59, 0xac, 0xf7, 0x36, 0x74, 0x25, 0x89, 0x84, 0x3f, 0xd2, 0x24, 0x3b, 0x0c, 0xd6,
	0x47, 0x41, 0xb5, 0x93, 0x2f, 0x48, 0x74, 0xa9, 0x5f, 0xdf, 0x76, 0x21, 0x68, 0xe2, 0x42, 0xe3,
	0x26, 0x2e, 0xb4, 0xd3, 0xc4, 0x85, 0x76, 0xb7, 0x72, 0xa1, 0x4f, 0x61, 0x87, 0x63, 0x8c, 0x44,
	0xd8, 0x49, 0xef, 0x39, 0x44, 0x8f, 0x6d, 0x84, 0xae, 0xfd, 0x31, 0x40, 0x99, 0x60, 0x51, 0x58,
	0x11, 0xde, 0x43, 0x54, 0x8b, 0xd7, 0xf3, 0xde, 0xa1, 0x62, 0x8e, 0x69, 0xc8, 0x8b, 0x4c, 0xe2,
	0xd2, 0xca, 0xf1, 0xef, 0x34, 0x3c, 0xa6, 0xe2, 0x49, 0x09, 0x9f, 0x08, 0xd8, 0xbd, 0xb1, 0x8f,
	0x5e, 0x00, 0x9d, 0xef, 0xb0, 0x78, 0xa5, 0x2a, 0x37, 0x47, 0xa1, 0x80, 0x4a, 0x28, 0x6b, 0xf5,
	0xcf, 0x49, 0x97, 0x06, 0x3a, 0xf9, 0xad, 0x07, 0x87, 0x55, 0xd5, 0x4b, 0x33, 0x97, 0xff, 0x9c,
	0x1d, 0x9c, 0x99, 0xee, 0x34, 0xd7, 0xc3, 0x74, 0x31, 0x84, 0xa1, 0x82, 0x7f, 0xae, 0x86, 0xba,
	0xe1, 0x24, 0xfd, 0x6d, 0x9c, 0xa4, 0xae, 0x27, 0x04, 0xd0, 0x49, 0x96, 0xef, 0x3b, 0x19, 0x82,
	0x02, 0xaa, 0x3a, 0x6b, 0xe4, 0x82, 0xb2, 0xd4, 0xa9, 0xed, 0x96, 0xe0, 0xff, 0xa3, 0x3d, 0x4c,
	0x7e, 0xef, 0x83, 0x7f, 0x87, 0xec, 0xcf, 0xa8, 0x90, 0x8c, 0x17, 0xaf, 0x39, 0xdf, 0x9c, 0xf3,
	0x55, 0xf7, 0x1c, 0xd4, 0xeb, 0x9e, 0xc3, 0x1a, 0x4a, 0x19, 0xb9, 0x2a, 0xe5, 0x31, 0x74, 0x65,
	0x91, 0xb9, 0x51, 0x5d, 0x23, 0x37, 0xfa, 0xf3, 0xb8, 0x46, 0x7f, 0xde, 0xd9, 0xaa, 0x3f, 0xef,
	0x36, 0x11, 0xe0, 0x5e, 0x13, 0x01, 0xee, 0x37, 0x11, 0xe0, 0xc1, 0x56, 0x02, 0x7c, 0xd9, 0x05,
	0x6f, 0x63, 0x2d, 0x30, 0xc9, 0x62, 0x22, 0xff, 0xf9, 0x76, 0xb3, 0xc1, 0xe7, 0x4e, 0x3d, 0x3e,
	0x77, 0xeb, 0xf1, 0xb9, 0xd7, 0xe0, 0xb3, 0xb1, 0xdf, 0x84, 0x10, 0x83, 0x26, 0x84, 0x18, 0x36,
	0x21, 0xc4, 0x68, 0x2b, 0x42, 0xfc, 0xdc, 0x85, 0x37, 0xce, 0x63, 0x8a, 0xa9, 0xac, 0x68, 0x71,
	0x91, 0xae, 0xd8, 0x4d, 0xb7, 0x6c, 0xd5, 0x74, 0xcb, 0x76, 0x4d, 0xb7, 0xec, 0xd4, 0x75, 0xcb,
	0xee, 0x96, 0x27, 0x84, 0xb2, 0x73, 0xf7, 0xb6, 0xe9, 0xdc, 0xd6, 0xf7, 0xfa, 0xae, 0xbe, 0x77,
	0xfb, 0x58, 0x3a, 0xd8, 0xea, 0x58, 0xea, 0x7d, 0x00, 0xc3, 0x25, 0x91, 0x64, 0xae, 0x8e, 0xa2,
	0x4e, 0xfe, 0xac, 0xd0, 0xcf, 0xb1, 0x50, 0x67, 0xe9, 0x2c, 0x5f, 0xc4, 0x34, 0xd4, 0xa1, 0x2e,
	0xbc, 0x18, 0x19, 0xfc, 0x73, 0x2c, 0x26, 0xbf, 0xb6, 0xe0, 0xa1, 0x21, 0xc6, 0x37, 0x44, 0x86,
	0x57, 0x9b, 0x5d, 0xfb, 0xfb, 0x1c, 0x85, 0xd4, 0x74, 0xd7, 0xef, 0xe7, 0x34, 0x73, 0x22, 0xc8,
	0xd0, 0xc0, 0x2f, 0x32, 0xf5, 0x95, 0x20, 0x90, 0xaf, 0x69, 0x68, 0xb7, 0xdb, 0x85, 0x26, 0x63,
	0x1b, 0xa1, 0x77, 0xfc, 0x0c, 0xc6, 0xd7, 0x6a, 0x54, 0xfa, 0x02, 0x4c, 0xf8, 0x1d, 0x6d, 0xf4,
	0xbe, 0x36, 0xfa, 0xbf, 0xe0, 0xf2, 0x25, 0x68, 0xb0, 0xfa, 0x2b, 0x26, 0xbf, 0xb4, 0xe0, 0xcd,
	0xea, 0xf5, 0x93, 0x1f, 0x32, 0xc6, 0x65, 0x39, 0xa5, 0x26, 0x9c, 0x3f, 0x81, 0xbe, 0x26, 0xb2,
	0xf0, 0xdb, 0x0e, 0x77, 0x0f, 0x16, 0xab, 0x94, 0xa2, 0x53, 0xd8, 0x29, 0xdc, 0xa3, 0x14, 0x0d,
	0x9d, 0xbd, 0x6c, 0xc1, 0x69, 0xc8, 0x92, 0x40, 0x62, 0x1a, 0x62, 0x2a, 0x83, 0x8c, 0xc5, 0x84,
	0x53, 0x11, 0x88, 0x0c, 0x43, 0xba, 0xa2, 0x21, 0x91, 0x94, 0xa5, 0x01, 0xc9, 0xa8, 0x5a, 0x0f,
	0x73, 0x51, 0x18, 0x24, 0x24, 0x25, 0x11, 0xce, 0x36, 0x6e, 0xff, 0xbe, 0x50, 0x15, 0xbe, 0x3d,
	0x8f, 0xa8, 0xbc, 0xca, 0x17, 0x41, 0xc8, 0x92, 0xa9, 0xcd, 0x93, 0xa0, 0xb8, 0x9a, 0xde, 0xc8,
	0x35, 0x15, 0x2c, 0xe7, 0x21, 0x4e, 0x23, 0x36, 0x25, 0x19, 0x9d, 0xae, 0x8f, 0xa6, 0xf6, 0xfa,
	0xd1, 0x64, 0x5d, 0xf4, 0xf5, 0x68, 0xdf, 0xfb, 0x23, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x3a, 0xd9,
	0xd6, 0x96, 0x14, 0x00, 0x00,
}
