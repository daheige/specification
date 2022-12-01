// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/polarismesh/specification/source/go/api/v1/model"
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

type Client_ClientType int32

const (
	Client_UNKNOWN Client_ClientType = 0
	Client_SDK     Client_ClientType = 1
	Client_AGENT   Client_ClientType = 2
)

var Client_ClientType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SDK",
	2: "AGENT",
}
var Client_ClientType_value = map[string]int32{
	"UNKNOWN": 0,
	"SDK":     1,
	"AGENT":   2,
}

func (x Client_ClientType) String() string {
	return proto.EnumName(Client_ClientType_name, int32(x))
}
func (Client_ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_144fd70933cdc64e, []int{0, 0}
}

type Client struct {
	Host                 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Type                 Client_ClientType       `protobuf:"varint,2,opt,name=type,proto3,enum=v1.Client_ClientType" json:"type,omitempty"`
	Version              *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Location             *model.Location         `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Id                   *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	Stat                 []*StatInfo             `protobuf:"bytes,6,rep,name=stat,proto3" json:"stat,omitempty"`
	Ctime                *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Mtime                *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_144fd70933cdc64e, []int{0}
}
func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (dst *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(dst, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetHost() *wrapperspb.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Client) GetType() Client_ClientType {
	if m != nil {
		return m.Type
	}
	return Client_UNKNOWN
}

func (m *Client) GetVersion() *wrapperspb.StringValue {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Client) GetLocation() *model.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Client) GetId() *wrapperspb.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Client) GetStat() []*StatInfo {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *Client) GetCtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Ctime
	}
	return nil
}

func (m *Client) GetMtime() *wrapperspb.StringValue {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type StatInfo struct {
	Target               *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Port                 *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Path                 *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Protocol             *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StatInfo) Reset()         { *m = StatInfo{} }
func (m *StatInfo) String() string { return proto.CompactTextString(m) }
func (*StatInfo) ProtoMessage()    {}
func (*StatInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_144fd70933cdc64e, []int{1}
}
func (m *StatInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatInfo.Unmarshal(m, b)
}
func (m *StatInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatInfo.Marshal(b, m, deterministic)
}
func (dst *StatInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatInfo.Merge(dst, src)
}
func (m *StatInfo) XXX_Size() int {
	return xxx_messageInfo_StatInfo.Size(m)
}
func (m *StatInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StatInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StatInfo proto.InternalMessageInfo

func (m *StatInfo) GetTarget() *wrapperspb.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *StatInfo) GetPort() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *StatInfo) GetPath() *wrapperspb.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *StatInfo) GetProtocol() *wrapperspb.StringValue {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func init() {
	proto.RegisterType((*Client)(nil), "v1.Client")
	proto.RegisterType((*StatInfo)(nil), "v1.StatInfo")
	proto.RegisterEnum("v1.Client_ClientType", Client_ClientType_name, Client_ClientType_value)
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_client_144fd70933cdc64e) }

var fileDescriptor_client_144fd70933cdc64e = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xfa, 0xd7, 0xd3, 0x45, 0xca, 0x80, 0x10, 0x16, 0x91, 0xd2, 0xab, 0x0a, 0x9a,
	0xb1, 0x5d, 0x11, 0x6f, 0xd5, 0x55, 0x59, 0x56, 0x2a, 0xb4, 0xbb, 0x0a, 0xde, 0xc8, 0x74, 0x7a,
	0x9a, 0x0c, 0x24, 0x99, 0x61, 0xe6, 0x24, 0xb2, 0x2f, 0xe5, 0x23, 0xf9, 0x2c, 0x92, 0x49, 0xb2,
	0xde, 0x88, 0xe4, 0x2a, 0xe4, 0x9c, 0xdf, 0xf7, 0x9b, 0x61, 0x3e, 0x38, 0x93, 0x99, 0xc2, 0x82,
	0x62, 0x63, 0x35, 0x69, 0x16, 0x56, 0xeb, 0xf3, 0xa7, 0x89, 0xd6, 0x49, 0x86, 0xdc, 0x4f, 0x0e,
	0xe5, 0x89, 0xff, 0xb4, 0xc2, 0x18, 0xb4, 0xae, 0x61, 0xce, 0x67, 0xb9, 0x3e, 0x62, 0xd6, 0xfc,
	0x2c, 0x7f, 0x0d, 0x60, 0xfc, 0xde, 0x1b, 0xd8, 0x4b, 0x18, 0xa6, 0xda, 0x51, 0x14, 0x2c, 0x82,
	0xd5, 0x6c, 0xf3, 0x24, 0x6e, 0x34, 0x71, 0xa7, 0x89, 0xf7, 0x64, 0x55, 0x91, 0x7c, 0x15, 0x59,
	0x89, 0x3b, 0x4f, 0xb2, 0x67, 0x30, 0xa4, 0x3b, 0x83, 0x51, 0xb8, 0x08, 0x56, 0x8f, 0x36, 0x8f,
	0xe3, 0x6a, 0x1d, 0x37, 0xae, 0xf6, 0x73, 0x73, 0x67, 0x70, 0xe7, 0x11, 0xf6, 0x1a, 0x26, 0x15,
	0x5a, 0xa7, 0x74, 0x11, 0x0d, 0x7a, 0xf8, 0x3b, 0x98, 0xad, 0x60, 0x9a, 0x69, 0x29, 0xa8, 0x0e,
	0x0e, 0x7d, 0xf0, 0xac, 0x3e, 0xe6, 0x73, 0x3b, 0xdb, 0xdd, 0x6f, 0xd9, 0x73, 0x08, 0xd5, 0x31,
	0x1a, 0xf5, 0x90, 0x87, 0xea, 0xc8, 0x16, 0x30, 0x74, 0x24, 0x28, 0x1a, 0x2f, 0x06, 0x9d, 0x73,
	0x4f, 0x82, 0xae, 0x8a, 0x93, 0xde, 0xf9, 0x0d, 0xdb, 0xc0, 0x48, 0x92, 0xca, 0x31, 0x9a, 0xf4,
	0x50, 0x36, 0x68, 0x9d, 0xc9, 0x7d, 0x66, 0xda, 0x27, 0xe3, 0xd1, 0xe5, 0x0b, 0x80, 0xbf, 0xaf,
	0xc5, 0x66, 0x30, 0xb9, 0xdd, 0x5e, 0x6f, 0xbf, 0x7c, 0xdb, 0xce, 0x1f, 0xb0, 0x09, 0x0c, 0xf6,
	0x97, 0xd7, 0xf3, 0x80, 0x3d, 0x84, 0xd1, 0xdb, 0x4f, 0x1f, 0xb6, 0x37, 0xf3, 0x70, 0xf9, 0x3b,
	0x80, 0x69, 0x77, 0x53, 0xf6, 0x0a, 0xc6, 0x24, 0x6c, 0x82, 0xfd, 0x4a, 0x6b, 0xd9, 0xba, 0x68,
	0xa3, 0x2d, 0xf9, 0xda, 0xfe, 0x95, 0xb9, 0xbd, 0x2a, 0xe8, 0x62, 0xd3, 0x16, 0x5d, 0x93, 0x3e,
	0x21, 0x28, 0xed, 0x55, 0x9d, 0x27, 0xd9, 0x1b, 0x98, 0xfa, 0xad, 0xd4, 0x59, 0xdb, 0xdb, 0xff,
	0x53, 0xf7, 0xf4, 0xbb, 0x8f, 0xdf, 0x2f, 0x13, 0x45, 0x69, 0x79, 0x88, 0xa5, 0xce, 0xb9, 0xd1,
	0x99, 0xb0, 0xca, 0xe5, 0xe8, 0x52, 0xee, 0x0c, 0x4a, 0x75, 0x52, 0x4d, 0xdf, 0xdc, 0xe9, 0xd2,
	0x4a, 0xe4, 0x89, 0xe6, 0xc2, 0x28, 0x5e, 0xad, 0xb9, 0x43, 0x5b, 0x29, 0x89, 0x3f, 0x72, 0x51,
	0x88, 0x04, 0x0f, 0x63, 0x6f, 0xbc, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x2f, 0xe1, 0xff,
	0x21, 0x03, 0x00, 0x00,
}