// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/deviceconfiguration.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"
	protocol "github.com/syncthing/syncthing/lib/protocol"
	_ "github.com/syncthing/syncthing/proto/ext"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DeviceConfiguration struct {
	DeviceID                 github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"deviceID" xml:"id,attr"`
	Name                     string                                               `protobuf:"bytes,2,opt,name=name,proto3" json:"name" xml:"name,attr,omitempty"`
	Addresses                []string                                             `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses" xml:"address,omitempty" default:"dynamic"`
	Compression              protocol.Compression                                 `protobuf:"varint,4,opt,name=compression,proto3,enum=protocol.Compression" json:"compression" xml:"compression,attr"`
	CertName                 string                                               `protobuf:"bytes,5,opt,name=cert_name,json=certName,proto3" json:"certName" xml:"certName,attr,omitempty"`
	Introducer               bool                                                 `protobuf:"varint,6,opt,name=introducer,proto3" json:"introducer" xml:"introducer,attr"`
	SkipIntroductionRemovals bool                                                 `protobuf:"varint,7,opt,name=skip_introduction_removals,json=skipIntroductionRemovals,proto3" json:"skipIntroductionRemovals" xml:"skipIntroductionRemovals,attr"`
	IntroducedBy             github_com_syncthing_syncthing_lib_protocol.DeviceID `protobuf:"bytes,8,opt,name=introduced_by,json=introducedBy,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.DeviceID" json:"introducedBy" xml:"introducedBy,attr"`
	Paused                   bool                                                 `protobuf:"varint,9,opt,name=paused,proto3" json:"paused" xml:"paused"`
	AllowedNetworks          []string                                             `protobuf:"bytes,10,rep,name=allowed_networks,json=allowedNetworks,proto3" json:"allowedNetworks" xml:"allowedNetwork,omitempty"`
	AutoAcceptFolders        bool                                                 `protobuf:"varint,11,opt,name=auto_accept_folders,json=autoAcceptFolders,proto3" json:"autoAcceptFolders" xml:"autoAcceptFolders"`
	MaxSendKbps              int                                                  `protobuf:"varint,12,opt,name=max_send_kbps,json=maxSendKbps,proto3,casttype=int" json:"maxSendKbps" xml:"maxSendKbps"`
	MaxRecvKbps              int                                                  `protobuf:"varint,13,opt,name=max_recv_kbps,json=maxRecvKbps,proto3,casttype=int" json:"maxRecvKbps" xml:"maxRecvKbps"`
	IgnoredFolders           []ObservedFolder                                     `protobuf:"bytes,14,rep,name=ignored_folders,json=ignoredFolders,proto3" json:"ignoredFolders" xml:"ignoredFolder"`
	PendingFolders           []ObservedFolder                                     `protobuf:"bytes,15,rep,name=pending_folders,json=pendingFolders,proto3" json:"pendingFolders" xml:"pendingFolder"`
	MaxRequestKiB            int                                                  `protobuf:"varint,16,opt,name=max_request_kib,json=maxRequestKib,proto3,casttype=int" json:"maxRequestKiB" xml:"maxRequestKiB"`
}

func (m *DeviceConfiguration) Reset()         { *m = DeviceConfiguration{} }
func (m *DeviceConfiguration) String() string { return proto.CompactTextString(m) }
func (*DeviceConfiguration) ProtoMessage()    {}
func (*DeviceConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_744b782bd13071dd, []int{0}
}
func (m *DeviceConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceConfiguration.Merge(m, src)
}
func (m *DeviceConfiguration) XXX_Size() int {
	return m.ProtoSize()
}
func (m *DeviceConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DeviceConfiguration)(nil), "config.DeviceConfiguration")
}

func init() {
	proto.RegisterFile("lib/config/deviceconfiguration.proto", fileDescriptor_744b782bd13071dd)
}

var fileDescriptor_744b782bd13071dd = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x31, 0x6f, 0xdb, 0x46,
	0x18, 0x15, 0xeb, 0xc4, 0xb6, 0xce, 0x96, 0x65, 0xd3, 0x88, 0xc3, 0x18, 0x88, 0x4e, 0x60, 0x35,
	0x28, 0x68, 0x2a, 0x17, 0x6e, 0x27, 0xa3, 0x1d, 0xca, 0x04, 0x45, 0x03, 0xa3, 0x49, 0x7b, 0xdd,
	0xbc, 0xb0, 0x24, 0xef, 0xac, 0x1c, 0x2c, 0xf2, 0x58, 0xf2, 0xa4, 0x48, 0x40, 0x87, 0x8e, 0x1d,
	0x3a, 0x14, 0x59, 0xbb, 0x14, 0x1d, 0x3a, 0xf4, 0x97, 0x18, 0xe8, 0x60, 0x8d, 0x45, 0x87, 0x03,
	0x62, 0x6f, 0x1c, 0x39, 0x66, 0x2a, 0x78, 0x47, 0x51, 0x24, 0x1d, 0x17, 0x06, 0xba, 0xdd, 0xbd,
	0xf7, 0xee, 0xbd, 0xe3, 0xd3, 0x7d, 0x10, 0xe8, 0x8d, 0xa8, 0x7b, 0xe0, 0xb1, 0xe0, 0x94, 0x0e,
	0x0f, 0x30, 0x99, 0x50, 0x8f, 0xa8, 0xcd, 0x38, 0x72, 0x38, 0x65, 0xc1, 0x20, 0x8c, 0x18, 0x67,
	0xfa, 0xaa, 0x02, 0xf7, 0xf7, 0x32, 0xb5, 0x84, 0x3c, 0x36, 0x3a, 0x70, 0x49, 0xa8, 0xf8, 0xfd,
	0x07, 0x25, 0x17, 0xe6, 0xc6, 0x24, 0x9a, 0x10, 0x9c, 0x53, 0x4d, 0x32, 0xe5, 0x6a, 0x69, 0xfe,
	0xb5, 0x05, 0x76, 0x9f, 0xca, 0x8c, 0x27, 0xe5, 0x0c, 0xfd, 0x4f, 0x0d, 0x34, 0x55, 0xb6, 0x4d,
	0xb1, 0xa1, 0x75, 0xb5, 0xfe, 0xa6, 0xf5, 0xb3, 0x76, 0x2e, 0x60, 0xe3, 0x1f, 0x01, 0x3f, 0x19,
	0x52, 0xfe, 0x72, 0xec, 0x0e, 0x3c, 0xe6, 0x1f, 0xc4, 0xb3, 0xc0, 0xe3, 0x2f, 0x69, 0x30, 0x2c,
	0xad, 0xca, 0x37, 0x1a, 0x28, 0xf7, 0x67, 0x4f, 0x2f, 0x05, 0x5c, 0x5f, 0xac, 0x13, 0x01, 0xd7,
	0x71, 0xbe, 0x4e, 0x05, 0x6c, 0x4d, 0xfd, 0xd1, 0x91, 0x49, 0xf1, 0x63, 0x87, 0xf3, 0xc8, 0x4c,
	0x2e, 0x7a, 0x6b, 0xf9, 0x3a, 0xbd, 0xe8, 0x15, 0xba, 0x9f, 0xe6, 0x3d, 0xed, 0xf5, 0xbc, 0x57,
	0x78, 0xa0, 0x05, 0x83, 0xf5, 0xaf, 0xc1, 0x9d, 0xc0, 0xf1, 0x89, 0xf1, 0x5e, 0x57, 0xeb, 0x37,
	0xad, 0x4f, 0x13, 0x01, 0xe5, 0x3e, 0x15, 0xf0, 0x81, 0x74, 0xce, 0x36, 0xd2, 0xef, 0x31, 0xf3,
	0x29, 0x27, 0x7e, 0xc8, 0x67, 0x59, 0xca, 0xee, 0x3b, 0x70, 0x24, 0x4f, 0xea, 0x53, 0xd0, 0x74,
	0x30, 0x8e, 0x48, 0x1c, 0x93, 0xd8, 0x58, 0xe9, 0xae, 0xf4, 0x9b, 0xd6, 0x49, 0x22, 0xe0, 0x12,
	0x4c, 0x05, 0x7c, 0x24, 0xbd, 0x73, 0xa4, 0xe4, 0xdc, 0xc5, 0xe4, 0xd4, 0x19, 0x8f, 0xf8, 0x91,
	0x89, 0x67, 0x81, 0xe3, 0x53, 0x2f, 0xcb, 0xda, 0xb9, 0xa6, 0x7b, 0x7b, 0xd1, 0x5b, 0xcb, 0x05,
	0x68, 0xe9, 0xab, 0x4f, 0xc0, 0x86, 0xc7, 0xfc, 0x30, 0xdb, 0x51, 0x16, 0x18, 0x77, 0xba, 0x5a,
	0x7f, 0xeb, 0xf0, 0xde, 0xa0, 0xa8, 0xf3, 0xc9, 0x92, 0xb4, 0x3e, 0x4b, 0x04, 0x2c, 0xab, 0x53,
	0x01, 0xf7, 0xe4, 0xa5, 0x4a, 0x58, 0xd1, 0xe9, 0x76, 0x1d, 0x44, 0xe5, 0xa3, 0x3a, 0x01, 0x4d,
	0x8f, 0x44, 0xdc, 0x96, 0x45, 0xde, 0x95, 0x45, 0x7e, 0x99, 0xfd, 0x4c, 0x19, 0xf8, 0x5c, 0x95,
	0xf9, 0x50, 0x79, 0xe7, 0xc0, 0x3b, 0x0a, 0xbd, 0x7f, 0x03, 0x87, 0x0a, 0x17, 0xfd, 0x04, 0x00,
	0x1a, 0xf0, 0x88, 0xe1, 0xb1, 0x47, 0x22, 0x63, 0xb5, 0xab, 0xf5, 0xd7, 0xad, 0xa3, 0x44, 0xc0,
	0x12, 0x9a, 0x0a, 0x78, 0x4f, 0x3d, 0x88, 0x02, 0x2a, 0x3e, 0xa2, 0x5d, 0xc3, 0x50, 0xe9, 0x9c,
	0xfe, 0xbb, 0x06, 0xf6, 0xe3, 0x33, 0x1a, 0xda, 0x0b, 0x2c, 0x7b, 0xc9, 0x76, 0x44, 0x7c, 0x36,
	0x71, 0x46, 0xb1, 0xb1, 0x26, 0xc3, 0x70, 0x22, 0xa0, 0x91, 0xa9, 0x9e, 0x95, 0x44, 0x28, 0xd7,
	0xa4, 0x02, 0xbe, 0x2f, 0xa3, 0x6f, 0x12, 0x14, 0x17, 0x79, 0xf8, 0x9f, 0x0a, 0x74, 0x63, 0x82,
	0xfe, 0x87, 0x06, 0x5a, 0xc5, 0x9d, 0xb1, 0xed, 0xce, 0x8c, 0x75, 0x39, 0x5c, 0x3f, 0xfe, 0xaf,
	0xe1, 0x4a, 0x04, 0xdc, 0x5c, 0xba, 0x5a, 0xb3, 0x54, 0xc0, 0xfb, 0xd5, 0x0e, 0xb1, 0x35, 0x2b,
	0x2e, 0xbf, 0x73, 0x0d, 0xcd, 0x86, 0x0b, 0x55, 0x1c, 0xf4, 0x43, 0xb0, 0x1a, 0x3a, 0xe3, 0x98,
	0x60, 0xa3, 0x29, 0x8b, 0xdb, 0x4f, 0x04, 0xcc, 0x91, 0x54, 0xc0, 0x4d, 0xe9, 0xae, 0xb6, 0x26,
	0xca, 0x71, 0xfd, 0x07, 0xb0, 0xed, 0x8c, 0x46, 0xec, 0x15, 0xc1, 0x76, 0x40, 0xf8, 0x2b, 0x16,
	0x9d, 0xc5, 0x06, 0x90, 0xd3, 0xf3, 0x4d, 0x22, 0x60, 0x3b, 0xe7, 0x9e, 0xe7, 0x54, 0x2a, 0x60,
	0x47, 0xcd, 0x50, 0x05, 0xaf, 0xbe, 0x29, 0xe3, 0x26, 0x12, 0xd5, 0xed, 0xf4, 0xef, 0xc0, 0xae,
	0x33, 0xe6, 0xcc, 0x76, 0x3c, 0x8f, 0x84, 0xdc, 0x3e, 0x65, 0x23, 0x4c, 0xa2, 0xd8, 0xd8, 0x90,
	0xd7, 0xff, 0x28, 0x11, 0x70, 0x27, 0xa3, 0x3f, 0x97, 0xec, 0x17, 0x8a, 0x2c, 0x7a, 0xba, 0xc6,
	0x98, 0xe8, 0xba, 0x5a, 0x7f, 0x01, 0x5a, 0xbe, 0x33, 0xb5, 0x63, 0x12, 0x60, 0xfb, 0xcc, 0x0d,
	0x63, 0x63, 0xb3, 0xab, 0xf5, 0xef, 0x5a, 0x1f, 0x64, 0x73, 0xe8, 0x3b, 0xd3, 0x6f, 0x49, 0x80,
	0x8f, 0xdd, 0x30, 0x73, 0xdd, 0x91, 0xae, 0x25, 0xcc, 0x7c, 0x2b, 0xe0, 0x0a, 0x0d, 0x38, 0x2a,
	0x0b, 0x17, 0x86, 0x11, 0xf1, 0x26, 0xca, 0xb0, 0x55, 0x31, 0x44, 0xc4, 0x9b, 0xd4, 0x0d, 0x17,
	0x58, 0xc5, 0x70, 0x01, 0xea, 0x01, 0x68, 0xd3, 0x61, 0xc0, 0x22, 0x82, 0x8b, 0xef, 0xdf, 0xea,
	0xae, 0xf4, 0x37, 0x0e, 0xf7, 0x06, 0xea, 0xbf, 0x60, 0xf0, 0x22, 0xff, 0x2f, 0x50, 0xdf, 0x64,
	0x7d, 0x98, 0x3d, 0xbb, 0x44, 0xc0, 0xad, 0xfc, 0xd8, 0xb2, 0x98, 0x5d, 0xf5, 0x80, 0xca, 0xb0,
	0x89, 0x6a, 0xb2, 0x2c, 0x2f, 0x24, 0x01, 0xa6, 0xc1, 0xb0, 0xc8, 0x6b, 0xdf, 0x2e, 0x2f, 0x3f,
	0x56, 0xcf, 0xab, 0xc0, 0x26, 0xaa, 0xc9, 0xf4, 0x5f, 0x35, 0xd0, 0x56, 0x8d, 0x7d, 0x3f, 0x26,
	0x31, 0xb7, 0xcf, 0xa8, 0x6b, 0x6c, 0xcb, 0xce, 0xe2, 0x4b, 0x01, 0x5b, 0x5f, 0x65, 0x55, 0x48,
	0xe6, 0x98, 0x5a, 0x89, 0x80, 0x2d, 0xbf, 0x0c, 0x14, 0x21, 0x15, 0x74, 0x51, 0x64, 0x72, 0xd1,
	0xab, 0xc9, 0xeb, 0xc0, 0xeb, 0x79, 0xaf, 0x9a, 0x80, 0x2a, 0xbc, 0x6b, 0x1d, 0x9f, 0xbf, 0xe9,
	0x34, 0xe6, 0x6f, 0x3a, 0x8d, 0xf3, 0xcb, 0x8e, 0x36, 0xbf, 0xec, 0x68, 0xbf, 0x5c, 0x75, 0x1a,
	0xbf, 0x5d, 0x75, 0xb4, 0xf9, 0x55, 0xa7, 0xf1, 0xf7, 0x55, 0xa7, 0x71, 0xf2, 0xe8, 0x16, 0xd3,
	0xad, 0x8a, 0x73, 0x57, 0xe5, 0x94, 0x7f, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x35,
	0x78, 0xb9, 0x0f, 0x08, 0x00, 0x00,
}

func (m *DeviceConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxRequestKiB != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxRequestKiB))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.PendingFolders) > 0 {
		for iNdEx := len(m.PendingFolders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingFolders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.IgnoredFolders) > 0 {
		for iNdEx := len(m.IgnoredFolders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IgnoredFolders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if m.MaxRecvKbps != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxRecvKbps))
		i--
		dAtA[i] = 0x68
	}
	if m.MaxSendKbps != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.MaxSendKbps))
		i--
		dAtA[i] = 0x60
	}
	if m.AutoAcceptFolders {
		i--
		if m.AutoAcceptFolders {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.AllowedNetworks) > 0 {
		for iNdEx := len(m.AllowedNetworks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedNetworks[iNdEx])
			copy(dAtA[i:], m.AllowedNetworks[iNdEx])
			i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.AllowedNetworks[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.IntroducedBy.ProtoSize()
		i -= size
		if _, err := m.IntroducedBy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if m.SkipIntroductionRemovals {
		i--
		if m.SkipIntroductionRemovals {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Introducer {
		i--
		if m.Introducer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.CertName) > 0 {
		i -= len(m.CertName)
		copy(dAtA[i:], m.CertName)
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.CertName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Compression != 0 {
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(m.Compression))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.DeviceID.ProtoSize()
		i -= size
		if _, err := m.DeviceID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDeviceconfiguration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDeviceconfiguration(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceconfiguration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceConfiguration) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DeviceID.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.Compression != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.Compression))
	}
	l = len(m.CertName)
	if l > 0 {
		n += 1 + l + sovDeviceconfiguration(uint64(l))
	}
	if m.Introducer {
		n += 2
	}
	if m.SkipIntroductionRemovals {
		n += 2
	}
	l = m.IntroducedBy.ProtoSize()
	n += 1 + l + sovDeviceconfiguration(uint64(l))
	if m.Paused {
		n += 2
	}
	if len(m.AllowedNetworks) > 0 {
		for _, s := range m.AllowedNetworks {
			l = len(s)
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.AutoAcceptFolders {
		n += 2
	}
	if m.MaxSendKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxSendKbps))
	}
	if m.MaxRecvKbps != 0 {
		n += 1 + sovDeviceconfiguration(uint64(m.MaxRecvKbps))
	}
	if len(m.IgnoredFolders) > 0 {
		for _, e := range m.IgnoredFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if len(m.PendingFolders) > 0 {
		for _, e := range m.PendingFolders {
			l = e.ProtoSize()
			n += 1 + l + sovDeviceconfiguration(uint64(l))
		}
	}
	if m.MaxRequestKiB != 0 {
		n += 2 + sovDeviceconfiguration(uint64(m.MaxRequestKiB))
	}
	return n
}

func sovDeviceconfiguration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceconfiguration(x uint64) (n int) {
	return sovDeviceconfiguration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceconfiguration
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DeviceID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compression", wireType)
			}
			m.Compression = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Compression |= protocol.Compression(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Introducer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Introducer = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkipIntroductionRemovals", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SkipIntroductionRemovals = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntroducedBy", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IntroducedBy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedNetworks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedNetworks = append(m.AllowedNetworks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoAcceptFolders", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoAcceptFolders = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSendKbps", wireType)
			}
			m.MaxSendKbps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSendKbps |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRecvKbps", wireType)
			}
			m.MaxRecvKbps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRecvKbps |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredFolders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredFolders = append(m.IgnoredFolders, ObservedFolder{})
			if err := m.IgnoredFolders[len(m.IgnoredFolders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingFolders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PendingFolders = append(m.PendingFolders, ObservedFolder{})
			if err := m.PendingFolders[len(m.PendingFolders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestKiB", wireType)
			}
			m.MaxRequestKiB = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRequestKiB |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceconfiguration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDeviceconfiguration
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDeviceconfiguration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceconfiguration
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeviceconfiguration
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDeviceconfiguration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceconfiguration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceconfiguration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceconfiguration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceconfiguration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceconfiguration = fmt.Errorf("proto: unexpected end of group")
)