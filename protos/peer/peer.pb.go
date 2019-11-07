// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/peer.proto

package peer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HelloRequest struct {
	Ordinal              uint32   `protobuf:"varint,1,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type NextResourceRequest struct {
	Ordinal              uint32   `protobuf:"varint,1,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	ResourceType         uint32   `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Readiness            []byte   `protobuf:"bytes,3,opt,name=readiness,proto3" json:"readiness,omitempty"`
	ReadyForEqual        bool     `protobuf:"varint,4,opt,name=ready_for_equal,json=readyForEqual,proto3" json:"ready_for_equal,omitempty"`
	ReadinessMax         bool     `protobuf:"varint,5,opt,name=readiness_max,json=readinessMax,proto3" json:"readiness_max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextResourceRequest) Reset()         { *m = NextResourceRequest{} }
func (m *NextResourceRequest) String() string { return proto.CompactTextString(m) }
func (*NextResourceRequest) ProtoMessage()    {}
func (*NextResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{1}
}

func (m *NextResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextResourceRequest.Unmarshal(m, b)
}
func (m *NextResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextResourceRequest.Marshal(b, m, deterministic)
}
func (m *NextResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextResourceRequest.Merge(m, src)
}
func (m *NextResourceRequest) XXX_Size() int {
	return xxx_messageInfo_NextResourceRequest.Size(m)
}
func (m *NextResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NextResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NextResourceRequest proto.InternalMessageInfo

func (m *NextResourceRequest) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *NextResourceRequest) GetResourceType() uint32 {
	if m != nil {
		return m.ResourceType
	}
	return 0
}

func (m *NextResourceRequest) GetReadiness() []byte {
	if m != nil {
		return m.Readiness
	}
	return nil
}

func (m *NextResourceRequest) GetReadyForEqual() bool {
	if m != nil {
		return m.ReadyForEqual
	}
	return false
}

func (m *NextResourceRequest) GetReadinessMax() bool {
	if m != nil {
		return m.ReadinessMax
	}
	return false
}

type NextResourceResponse struct {
	Has                  bool     `protobuf:"varint,1,opt,name=has,proto3" json:"has,omitempty"`
	Last                 bool     `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
	ResourceId           []byte   `protobuf:"bytes,3,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NextResourceResponse) Reset()         { *m = NextResourceResponse{} }
func (m *NextResourceResponse) String() string { return proto.CompactTextString(m) }
func (*NextResourceResponse) ProtoMessage()    {}
func (*NextResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{2}
}

func (m *NextResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NextResourceResponse.Unmarshal(m, b)
}
func (m *NextResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NextResourceResponse.Marshal(b, m, deterministic)
}
func (m *NextResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextResourceResponse.Merge(m, src)
}
func (m *NextResourceResponse) XXX_Size() int {
	return xxx_messageInfo_NextResourceResponse.Size(m)
}
func (m *NextResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NextResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NextResourceResponse proto.InternalMessageInfo

func (m *NextResourceResponse) GetHas() bool {
	if m != nil {
		return m.Has
	}
	return false
}

func (m *NextResourceResponse) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

func (m *NextResourceResponse) GetResourceId() []byte {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

type ReadinessRequest struct {
	Ordinal uint32 `protobuf:"varint,1,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	// would prefer to use a map, but []byte is not a supported key type, and string must be in UTF-8 format (cannot be arbitrary bytes)
	Readiness            []*Readiness `protobuf:"bytes,2,rep,name=readiness,proto3" json:"readiness,omitempty"`
	ShuttingDown         bool         `protobuf:"varint,3,opt,name=shutting_down,json=shuttingDown,proto3" json:"shutting_down,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadinessRequest) Reset()         { *m = ReadinessRequest{} }
func (m *ReadinessRequest) String() string { return proto.CompactTextString(m) }
func (*ReadinessRequest) ProtoMessage()    {}
func (*ReadinessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{3}
}

func (m *ReadinessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadinessRequest.Unmarshal(m, b)
}
func (m *ReadinessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadinessRequest.Marshal(b, m, deterministic)
}
func (m *ReadinessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadinessRequest.Merge(m, src)
}
func (m *ReadinessRequest) XXX_Size() int {
	return xxx_messageInfo_ReadinessRequest.Size(m)
}
func (m *ReadinessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadinessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadinessRequest proto.InternalMessageInfo

func (m *ReadinessRequest) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *ReadinessRequest) GetReadiness() []*Readiness {
	if m != nil {
		return m.Readiness
	}
	return nil
}

func (m *ReadinessRequest) GetShuttingDown() bool {
	if m != nil {
		return m.ShuttingDown
	}
	return false
}

type Readiness struct {
	ResourceType         uint32   `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Readiness            []byte   `protobuf:"bytes,2,opt,name=readiness,proto3" json:"readiness,omitempty"`
	ReadyForEqual        bool     `protobuf:"varint,3,opt,name=ready_for_equal,json=readyForEqual,proto3" json:"ready_for_equal,omitempty"`
	Max                  bool     `protobuf:"varint,4,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Readiness) Reset()         { *m = Readiness{} }
func (m *Readiness) String() string { return proto.CompactTextString(m) }
func (*Readiness) ProtoMessage()    {}
func (*Readiness) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{4}
}

func (m *Readiness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Readiness.Unmarshal(m, b)
}
func (m *Readiness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Readiness.Marshal(b, m, deterministic)
}
func (m *Readiness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Readiness.Merge(m, src)
}
func (m *Readiness) XXX_Size() int {
	return xxx_messageInfo_Readiness.Size(m)
}
func (m *Readiness) XXX_DiscardUnknown() {
	xxx_messageInfo_Readiness.DiscardUnknown(m)
}

var xxx_messageInfo_Readiness proto.InternalMessageInfo

func (m *Readiness) GetResourceType() uint32 {
	if m != nil {
		return m.ResourceType
	}
	return 0
}

func (m *Readiness) GetReadiness() []byte {
	if m != nil {
		return m.Readiness
	}
	return nil
}

func (m *Readiness) GetReadyForEqual() bool {
	if m != nil {
		return m.ReadyForEqual
	}
	return false
}

func (m *Readiness) GetMax() bool {
	if m != nil {
		return m.Max
	}
	return false
}

type StatsRequest struct {
	ResourceStats        map[uint32]*NodeStats `protobuf:"bytes,1,rep,name=resource_stats,json=resourceStats,proto3" json:"resource_stats,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{5}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}
func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}
func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}
func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}
func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

func (m *StatsRequest) GetResourceStats() map[uint32]*NodeStats {
	if m != nil {
		return m.ResourceStats
	}
	return nil
}

type NodeStats struct {
	Stats                []*NodeStat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{6}
}

func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (m *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(m, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetStats() []*NodeStat {
	if m != nil {
		return m.Stats
	}
	return nil
}

type NodeStat struct {
	Ordinal              uint32   `protobuf:"varint,1,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStat) Reset()         { *m = NodeStat{} }
func (m *NodeStat) String() string { return proto.CompactTextString(m) }
func (*NodeStat) ProtoMessage()    {}
func (*NodeStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{7}
}

func (m *NodeStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStat.Unmarshal(m, b)
}
func (m *NodeStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStat.Marshal(b, m, deterministic)
}
func (m *NodeStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStat.Merge(m, src)
}
func (m *NodeStat) XXX_Size() int {
	return xxx_messageInfo_NodeStat.Size(m)
}
func (m *NodeStat) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStat.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStat proto.InternalMessageInfo

func (m *NodeStat) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *NodeStat) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type HandoffRequest struct {
	ResourceType         uint32   `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId           []byte   `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Ordinal              uint32   `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	ResourceCount        uint32   `protobuf:"varint,4,opt,name=resource_count,json=resourceCount,proto3" json:"resource_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandoffRequest) Reset()         { *m = HandoffRequest{} }
func (m *HandoffRequest) String() string { return proto.CompactTextString(m) }
func (*HandoffRequest) ProtoMessage()    {}
func (*HandoffRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c302117fbb08ad42, []int{8}
}

func (m *HandoffRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandoffRequest.Unmarshal(m, b)
}
func (m *HandoffRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandoffRequest.Marshal(b, m, deterministic)
}
func (m *HandoffRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandoffRequest.Merge(m, src)
}
func (m *HandoffRequest) XXX_Size() int {
	return xxx_messageInfo_HandoffRequest.Size(m)
}
func (m *HandoffRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandoffRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandoffRequest proto.InternalMessageInfo

func (m *HandoffRequest) GetResourceType() uint32 {
	if m != nil {
		return m.ResourceType
	}
	return 0
}

func (m *HandoffRequest) GetResourceId() []byte {
	if m != nil {
		return m.ResourceId
	}
	return nil
}

func (m *HandoffRequest) GetOrdinal() uint32 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func (m *HandoffRequest) GetResourceCount() uint32 {
	if m != nil {
		return m.ResourceCount
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "peer.HelloRequest")
	proto.RegisterType((*NextResourceRequest)(nil), "peer.NextResourceRequest")
	proto.RegisterType((*NextResourceResponse)(nil), "peer.NextResourceResponse")
	proto.RegisterType((*ReadinessRequest)(nil), "peer.ReadinessRequest")
	proto.RegisterType((*Readiness)(nil), "peer.Readiness")
	proto.RegisterType((*StatsRequest)(nil), "peer.StatsRequest")
	proto.RegisterMapType((map[uint32]*NodeStats)(nil), "peer.StatsRequest.ResourceStatsEntry")
	proto.RegisterType((*NodeStats)(nil), "peer.NodeStats")
	proto.RegisterType((*NodeStat)(nil), "peer.NodeStat")
	proto.RegisterType((*HandoffRequest)(nil), "peer.HandoffRequest")
}

func init() { proto.RegisterFile("peer/peer.proto", fileDescriptor_c302117fbb08ad42) }

var fileDescriptor_c302117fbb08ad42 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xee, 0xda, 0xc9, 0xaf, 0xc9, 0xc4, 0xf9, 0xa3, 0xfd, 0x45, 0x95, 0x09, 0x48, 0x44, 0x2e,
	0x41, 0xb9, 0xe0, 0x88, 0x22, 0x24, 0x28, 0xc7, 0x12, 0x28, 0x12, 0x54, 0xb0, 0xc0, 0x11, 0x45,
	0xdb, 0x7a, 0x93, 0x46, 0xb8, 0x5e, 0xd7, 0x5e, 0xd3, 0xf8, 0xc6, 0x91, 0x37, 0xe0, 0x31, 0x78,
	0x04, 0x0e, 0xbc, 0x18, 0xda, 0x5d, 0xdb, 0x71, 0x92, 0xd2, 0xf4, 0x62, 0xcd, 0x7e, 0x33, 0xe3,
	0xfd, 0xe6, 0x9b, 0xd9, 0x81, 0x76, 0xc8, 0x58, 0x34, 0x92, 0x1f, 0x37, 0x8c, 0xb8, 0xe0, 0xb8,
	0x22, 0xed, 0xde, 0xdd, 0x19, 0xe7, 0x33, 0x9f, 0x8d, 0x14, 0x76, 0x9a, 0x4c, 0x47, 0xec, 0x22,
	0x14, 0xa9, 0x0e, 0x71, 0x86, 0x60, 0x1d, 0x33, 0xdf, 0xe7, 0x84, 0x5d, 0x26, 0x2c, 0x16, 0xd8,
	0x86, 0x5d, 0x1e, 0x79, 0xf3, 0x80, 0xfa, 0x36, 0xea, 0xa3, 0x61, 0x93, 0xe4, 0x47, 0xe7, 0x37,
	0x82, 0xff, 0x4f, 0xd8, 0x42, 0x10, 0x16, 0xf3, 0x24, 0x3a, 0x63, 0x5b, 0x33, 0xf0, 0x3e, 0x34,
	0xa3, 0x2c, 0x78, 0x22, 0xd2, 0x90, 0xd9, 0x86, 0xf2, 0x5b, 0x39, 0xf8, 0x29, 0x0d, 0x19, 0xbe,
	0x07, 0xf5, 0x88, 0x51, 0x6f, 0x1e, 0xb0, 0x38, 0xb6, 0xcd, 0x3e, 0x1a, 0x5a, 0x64, 0x09, 0xe0,
	0x87, 0xd0, 0x96, 0x87, 0x74, 0x32, 0xe5, 0xd1, 0x84, 0x5d, 0x26, 0xd4, 0xb7, 0x2b, 0x7d, 0x34,
	0xac, 0x91, 0xa6, 0x82, 0x5f, 0xf1, 0x68, 0x2c, 0x41, 0x7d, 0x55, 0x96, 0x34, 0xb9, 0xa0, 0x0b,
	0xbb, 0xaa, 0xa2, 0xac, 0x02, 0x7c, 0x47, 0x17, 0xce, 0x17, 0xe8, 0xae, 0x16, 0x10, 0x87, 0x3c,
	0x88, 0x19, 0xee, 0x80, 0x79, 0x4e, 0x63, 0xc5, 0xbe, 0x46, 0xa4, 0x89, 0x31, 0x54, 0x7c, 0x1a,
	0x0b, 0x45, 0xb8, 0x46, 0x94, 0x8d, 0xef, 0x43, 0xa3, 0xa8, 0x66, 0xee, 0x65, 0x54, 0x21, 0x87,
	0xde, 0x78, 0xce, 0x77, 0x04, 0x1d, 0x92, 0xdf, 0xb7, 0x5d, 0x9d, 0x47, 0xe5, 0xc2, 0x8d, 0xbe,
	0x39, 0x6c, 0x1c, 0xb4, 0x5d, 0xd5, 0xbc, 0xe5, 0x4f, 0x4a, 0x4a, 0xec, 0x43, 0x33, 0x3e, 0x4f,
	0x84, 0x98, 0x07, 0xb3, 0x89, 0xc7, 0xaf, 0x02, 0x45, 0xa0, 0x46, 0xac, 0x1c, 0x7c, 0xc9, 0xaf,
	0x02, 0xe7, 0x07, 0x82, 0x3a, 0x29, 0xa7, 0xac, 0xea, 0x8f, 0xb6, 0xe9, 0x6f, 0xdc, 0x42, 0x7f,
	0xf3, 0x3a, 0xfd, 0x3b, 0x60, 0x4a, 0xd5, 0x75, 0x6f, 0xa4, 0xe9, 0xfc, 0x42, 0x60, 0x7d, 0x14,
	0x54, 0x14, 0x4a, 0xbc, 0x85, 0x56, 0xc1, 0x26, 0x96, 0x0e, 0x1b, 0xa9, 0xa2, 0x07, 0xba, 0xe8,
	0x72, 0xac, 0x9b, 0xb7, 0x48, 0x81, 0xe3, 0x40, 0x44, 0x29, 0x29, 0x4a, 0x51, 0x58, 0xef, 0x03,
	0xe0, 0xcd, 0x20, 0x49, 0xe3, 0x2b, 0x4b, 0xb3, 0x3a, 0xa5, 0x89, 0x07, 0x50, 0xfd, 0x46, 0xfd,
	0x44, 0xcf, 0x5e, 0xa1, 0xf0, 0x09, 0xf7, 0x74, 0x1a, 0xd1, 0xde, 0x43, 0xe3, 0x19, 0x72, 0x1e,
	0x43, 0xbd, 0xc0, 0xf1, 0x03, 0xa8, 0x96, 0x49, 0xb6, 0x56, 0xf3, 0x88, 0x76, 0x3a, 0x87, 0x50,
	0xcb, 0xa1, 0x1b, 0x3a, 0xdd, 0x85, 0xea, 0x19, 0x4f, 0x02, 0x91, 0xcd, 0xbf, 0x3e, 0x38, 0x3f,
	0x11, 0xb4, 0x8e, 0x69, 0xe0, 0xf1, 0xe9, 0x34, 0x97, 0xe8, 0x56, 0x0d, 0x5b, 0x9b, 0x43, 0x63,
	0x7d, 0x0e, 0xcb, 0x44, 0xcc, 0x55, 0x22, 0x83, 0x52, 0x0b, 0x34, 0xa3, 0x8a, 0x0a, 0x28, 0x6e,
	0x3d, 0x92, 0xe0, 0xc1, 0x1f, 0x03, 0x2a, 0xef, 0x19, 0x8b, 0xf0, 0x53, 0xa8, 0xaa, 0xe5, 0x80,
	0xb1, 0x2e, 0xbf, 0xbc, 0x29, 0x7a, 0x7b, 0xae, 0xde, 0x2b, 0x6e, 0xbe, 0x57, 0xdc, 0xb1, 0xdc,
	0x2b, 0xce, 0x0e, 0x7e, 0x0d, 0x56, 0xf9, 0x9d, 0xe1, 0x3b, 0x99, 0x78, 0x9b, 0xcb, 0xa3, 0xd7,
	0xbb, 0xce, 0xa5, 0x9f, 0xa5, 0xb3, 0x83, 0x8f, 0xa0, 0xfd, 0x39, 0xf4, 0xa8, 0x60, 0xcb, 0x99,
	0xde, 0x5b, 0x7f, 0x22, 0x5b, 0xd9, 0xbc, 0x80, 0x86, 0xfe, 0x89, 0x6e, 0x2c, 0xde, 0x1c, 0xb7,
	0x1b, 0x92, 0x9f, 0xc3, 0x6e, 0xd6, 0x23, 0xdc, 0xcd, 0x34, 0x58, 0x69, 0xd9, 0xbf, 0x53, 0x4f,
	0xff, 0x53, 0xc8, 0x93, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe4, 0x2d, 0xd8, 0x96, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	NextResource(ctx context.Context, in *NextResourceRequest, opts ...grpc.CallOption) (*NextResourceResponse, error)
	UpdateReadiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Handoff(ctx context.Context, in *HandoffRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type peerClient struct {
	cc *grpc.ClientConn
}

func NewPeerClient(cc *grpc.ClientConn) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peer.Peer/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) NextResource(ctx context.Context, in *NextResourceRequest, opts ...grpc.CallOption) (*NextResourceResponse, error) {
	out := new(NextResourceResponse)
	err := c.cc.Invoke(ctx, "/peer.Peer/NextResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) UpdateReadiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peer.Peer/UpdateReadiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) UpdateStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peer.Peer/UpdateStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) Handoff(ctx context.Context, in *HandoffRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peer.Peer/Handoff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
type PeerServer interface {
	Hello(context.Context, *HelloRequest) (*empty.Empty, error)
	NextResource(context.Context, *NextResourceRequest) (*NextResourceResponse, error)
	UpdateReadiness(context.Context, *ReadinessRequest) (*empty.Empty, error)
	UpdateStats(context.Context, *StatsRequest) (*empty.Empty, error)
	Handoff(context.Context, *HandoffRequest) (*empty.Empty, error)
}

// UnimplementedPeerServer can be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (*UnimplementedPeerServer) Hello(ctx context.Context, req *HelloRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedPeerServer) NextResource(ctx context.Context, req *NextResourceRequest) (*NextResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextResource not implemented")
}
func (*UnimplementedPeerServer) UpdateReadiness(ctx context.Context, req *ReadinessRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReadiness not implemented")
}
func (*UnimplementedPeerServer) UpdateStats(ctx context.Context, req *StatsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStats not implemented")
}
func (*UnimplementedPeerServer) Handoff(ctx context.Context, req *HandoffRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handoff not implemented")
}

func RegisterPeerServer(s *grpc.Server, srv PeerServer) {
	s.RegisterService(&_Peer_serviceDesc, srv)
}

func _Peer_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Peer/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_NextResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).NextResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Peer/NextResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).NextResource(ctx, req.(*NextResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_UpdateReadiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).UpdateReadiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Peer/UpdateReadiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).UpdateReadiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_UpdateStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).UpdateStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Peer/UpdateStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).UpdateStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_Handoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandoffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Handoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.Peer/Handoff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Handoff(ctx, req.(*HandoffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Peer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peer.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Peer_Hello_Handler,
		},
		{
			MethodName: "NextResource",
			Handler:    _Peer_NextResource_Handler,
		},
		{
			MethodName: "UpdateReadiness",
			Handler:    _Peer_UpdateReadiness_Handler,
		},
		{
			MethodName: "UpdateStats",
			Handler:    _Peer_UpdateStats_Handler,
		},
		{
			MethodName: "Handoff",
			Handler:    _Peer_Handoff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/peer.proto",
}
