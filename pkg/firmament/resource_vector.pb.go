/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_vector.proto

package firmament

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ResourceVector stores all the resources which will be take in account during scheduling.
type ResourceVector struct {
	// cpu_cores is the cpu millicores of node..
	CpuCores float32 `protobuf:"fixed32,1,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	RamBw    uint64  `protobuf:"varint,2,opt,name=ram_bw,json=ramBw,proto3" json:"ram_bw,omitempty"`
	// ram_cap is the memory capacity of node.
	RamCap uint64 `protobuf:"varint,3,opt,name=ram_cap,json=ramCap,proto3" json:"ram_cap,omitempty"`
	DiskBw uint64 `protobuf:"varint,4,opt,name=disk_bw,json=diskBw,proto3" json:"disk_bw,omitempty"`
	// disk_cap is the disk capacity of node.
	DiskCap uint64 `protobuf:"varint,5,opt,name=disk_cap,json=diskCap,proto3" json:"disk_cap,omitempty"`
	// net_tx_bw is transmit network packets in KB.
	NetTxBw uint64 `protobuf:"varint,6,opt,name=net_tx_bw,json=netTxBw,proto3" json:"net_tx_bw,omitempty"`
	// net_tx_bw is receive network packets in KB.
	NetRxBw uint64 `protobuf:"varint,7,opt,name=net_rx_bw,json=netRxBw,proto3" json:"net_rx_bw,omitempty"`
	// ephemeral storage
	EphemeralCap         uint64   `protobuf:"varint,8,opt,name=ephemeral_cap,json=ephemeralCap,proto3" json:"ephemeral_cap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceVector) Reset()         { *m = ResourceVector{} }
func (m *ResourceVector) String() string { return proto.CompactTextString(m) }
func (*ResourceVector) ProtoMessage()    {}
func (*ResourceVector) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd2f68a0615029fb, []int{0}
}

func (m *ResourceVector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceVector.Unmarshal(m, b)
}
func (m *ResourceVector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceVector.Marshal(b, m, deterministic)
}
func (m *ResourceVector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceVector.Merge(m, src)
}
func (m *ResourceVector) XXX_Size() int {
	return xxx_messageInfo_ResourceVector.Size(m)
}
func (m *ResourceVector) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceVector.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceVector proto.InternalMessageInfo

func (m *ResourceVector) GetCpuCores() float32 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *ResourceVector) GetRamBw() uint64 {
	if m != nil {
		return m.RamBw
	}
	return 0
}

func (m *ResourceVector) GetRamCap() uint64 {
	if m != nil {
		return m.RamCap
	}
	return 0
}

func (m *ResourceVector) GetDiskBw() uint64 {
	if m != nil {
		return m.DiskBw
	}
	return 0
}

func (m *ResourceVector) GetDiskCap() uint64 {
	if m != nil {
		return m.DiskCap
	}
	return 0
}

func (m *ResourceVector) GetNetTxBw() uint64 {
	if m != nil {
		return m.NetTxBw
	}
	return 0
}

func (m *ResourceVector) GetNetRxBw() uint64 {
	if m != nil {
		return m.NetRxBw
	}
	return 0
}

func (m *ResourceVector) GetEphemeralCap() uint64 {
	if m != nil {
		return m.EphemeralCap
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourceVector)(nil), "firmament.ResourceVector")
}

func init() { proto.RegisterFile("resource_vector.proto", fileDescriptor_dd2f68a0615029fb) }

var fileDescriptor_dd2f68a0615029fb = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0xb9, 0xb4, 0x69, 0x72, 0x02, 0x06, 0x4b, 0x15, 0x01, 0x96, 0x0a, 0x96, 0x4e,
	0x2c, 0xbc, 0x41, 0xf2, 0x06, 0x16, 0x62, 0xb5, 0x5c, 0x73, 0x88, 0x0a, 0x1c, 0x5b, 0x17, 0x07,
	0xf7, 0xa9, 0x79, 0x06, 0x74, 0x17, 0xc8, 0x78, 0xff, 0xf7, 0xff, 0x1e, 0x0c, 0x3b, 0xc2, 0x31,
	0x4e, 0xe4, 0xd1, 0x7e, 0xa3, 0xcf, 0x91, 0x9e, 0x12, 0xc5, 0x1c, 0x75, 0xf3, 0x7e, 0xa2, 0xe0,
	0x02, 0x0e, 0xf9, 0xe1, 0x47, 0xc1, 0xb5, 0xf9, 0x2b, 0xbd, 0x4a, 0x47, 0xdf, 0x43, 0xe3, 0xd3,
	0x64, 0x7d, 0x24, 0x1c, 0x5b, 0xb5, 0x57, 0x87, 0x95, 0xa9, 0x7d, 0x9a, 0x7a, 0xbe, 0xf5, 0x0e,
	0x2a, 0x72, 0xc1, 0x1e, 0x4b, 0xbb, 0xda, 0xab, 0xc3, 0xda, 0x6c, 0xc8, 0x85, 0xae, 0xe8, 0x1b,
	0xd8, 0x72, 0xec, 0x5d, 0x6a, 0x2f, 0x24, 0xe7, 0x56, 0xef, 0x12, 0xc3, 0xdb, 0x69, 0xfc, 0xe4,
	0xc1, 0x7a, 0x06, 0x3e, 0xbb, 0xa2, 0x6f, 0xa1, 0x16, 0xe0, 0xc9, 0x46, 0x44, 0x8a, 0xbc, 0xb9,
	0x83, 0x66, 0xc0, 0x6c, 0xf3, 0x99, 0x57, 0xd5, 0x6c, 0x03, 0xe6, 0x97, 0x73, 0x57, 0xfe, 0x8d,
	0xc4, 0xb6, 0x8b, 0x19, 0xb6, 0x47, 0xb8, 0xc2, 0xf4, 0x81, 0x01, 0xc9, 0x7d, 0xc9, 0xbb, 0xb5,
	0xf8, 0xe5, 0x12, 0xf6, 0x2e, 0x1d, 0x2b, 0xf9, 0x82, 0xe7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x04, 0xf4, 0xe3, 0xc8, 0x1b, 0x01, 0x00, 0x00,
}
