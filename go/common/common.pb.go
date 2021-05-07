// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
}

var Modulation_value = map[string]int32{
	"LORA": 0,
	"FSK":  1,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
	// ISM2400 (LoRaWAN 2.4 GHz)
	Region_ISM2400 Region = 11
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
	11: "ISM2400",
}

var Region_value = map[string]int32{
	"EU868":   0,
	"US915":   2,
	"CN779":   3,
	"EU433":   4,
	"AU915":   5,
	"CN470":   6,
	"AS923":   7,
	"KR920":   8,
	"IN865":   9,
	"RU864":   10,
	"ISM2400": 11,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type MType int32

const (
	// JoinRequest.
	MType_JoinRequest MType = 0
	// JoinAccept.
	MType_JoinAccept MType = 1
	// UnconfirmedDataUp.
	MType_UnconfirmedDataUp MType = 2
	// UnconfirmedDataDown.
	MType_UnconfirmedDataDown MType = 3
	// ConfirmedDataUp.
	MType_ConfirmedDataUp MType = 4
	// ConfirmedDataDown.
	MType_ConfirmedDataDown MType = 5
	// RejoinRequest.
	MType_RejoinRequest MType = 6
	// Proprietary.
	MType_Proprietary MType = 7
)

var MType_name = map[int32]string{
	0: "JoinRequest",
	1: "JoinAccept",
	2: "UnconfirmedDataUp",
	3: "UnconfirmedDataDown",
	4: "ConfirmedDataUp",
	5: "ConfirmedDataDown",
	6: "RejoinRequest",
	7: "Proprietary",
}

var MType_value = map[string]int32{
	"JoinRequest":         0,
	"JoinAccept":          1,
	"UnconfirmedDataUp":   2,
	"UnconfirmedDataDown": 3,
	"ConfirmedDataUp":     4,
	"ConfirmedDataDown":   5,
	"RejoinRequest":       6,
	"Proprietary":         7,
}

func (x MType) String() string {
	return proto.EnumName(MType_name, int32(x))
}

func (MType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver (TDOA).
	LocationSource_GEO_RESOLVER_TDOA LocationSource = 3
	// Geo resolver (RSSI).
	LocationSource_GEO_RESOLVER_RSSI LocationSource = 4
	// Geo resolver (GNSS).
	LocationSource_GEO_RESOLVER_GNSS LocationSource = 5
	// Geo resolver (WIFI).
	LocationSource_GEO_RESOLVER_WIFI LocationSource = 6
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER_TDOA",
	4: "GEO_RESOLVER_RSSI",
	5: "GEO_RESOLVER_GNSS",
	6: "GEO_RESOLVER_WIFI",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":           0,
	"GPS":               1,
	"CONFIG":            2,
	"GEO_RESOLVER_TDOA": 3,
	"GEO_RESOLVER_RSSI": 4,
	"GEO_RESOLVER_GNSS": 5,
	"GEO_RESOLVER_WIFI": 6,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.MType", MType_name, MType_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x5f, 0x53, 0xda, 0x4e,
	0x14, 0x75, 0xf9, 0x13, 0xe0, 0xf2, 0x53, 0xd7, 0x75, 0x7e, 0x95, 0x69, 0x3b, 0x53, 0xc6, 0x27,
	0x86, 0x99, 0x02, 0x05, 0x54, 0x78, 0x44, 0x44, 0x86, 0x82, 0x09, 0xb3, 0x31, 0x75, 0xda, 0x17,
	0x66, 0x59, 0xb7, 0x98, 0x26, 0x66, 0xd3, 0x90, 0xd8, 0xc9, 0x73, 0xdf, 0xfb, 0x39, 0xfa, 0xd0,
	0x0f, 0xd9, 0xd9, 0x04, 0xb5, 0xea, 0xd3, 0x9e, 0x7b, 0xce, 0xb9, 0x7f, 0x36, 0xd9, 0x0b, 0xfb,
	0x5c, 0xde, 0xde, 0x4a, 0xaf, 0x99, 0x1e, 0x0d, 0x3f, 0x90, 0xa1, 0x24, 0x5a, 0x1a, 0x1d, 0x0e,
	0xa1, 0x3c, 0x15, 0xf1, 0xc8, 0xbb, 0x13, 0xae, 0xf4, 0x05, 0x79, 0x03, 0x25, 0x47, 0x38, 0x0b,
	0x97, 0x2d, 0x85, 0x5b, 0x41, 0x55, 0x54, 0x2b, 0xd1, 0xa2, 0x23, 0x9c, 0x99, 0x8a, 0xc9, 0x01,
	0x14, 0x98, 0x58, 0x2f, 0x1c, 0x11, 0x57, 0x32, 0x55, 0x54, 0xfb, 0x8f, 0x6a, 0x4c, 0xac, 0xa7,
	0x22, 0x3e, 0xfc, 0x83, 0xa0, 0x38, 0x93, 0x9c, 0x85, 0xb6, 0xf4, 0xc8, 0x6b, 0x28, 0xba, 0x2c,
	0xb4, 0xc3, 0xe8, 0x5a, 0x24, 0x15, 0x10, 0x7d, 0x88, 0xc9, 0x5b, 0x28, 0xb9, 0xd2, 0x5b, 0xa5,
	0x62, 0x26, 0x11, 0x1f, 0x09, 0x95, 0xc9, 0xdc, 0x4d, 0x66, 0x36, 0xcd, 0xbc, 0x8f, 0x49, 0x03,
	0xb4, 0xb5, 0x8c, 0x02, 0x2e, 0x2a, 0xb9, 0x2a, 0xaa, 0xed, 0xb4, 0x5f, 0x35, 0x36, 0xd7, 0xb9,
	0xef, 0x6b, 0x26, 0x2a, 0xdd, 0xb8, 0x92, 0x5a, 0x9c, 0x47, 0x01, 0xe3, 0x71, 0x25, 0x5f, 0x45,
	0xb5, 0x6d, 0xfa, 0x10, 0xd7, 0xdf, 0x01, 0x5c, 0xc8, 0xeb, 0xc8, 0x4d, 0xe7, 0x2d, 0x42, 0x6e,
	0x66, 0xd0, 0x01, 0xde, 0x22, 0x05, 0xc8, 0x9e, 0x9b, 0x53, 0x8c, 0xea, 0x3f, 0x11, 0x68, 0x54,
	0xac, 0x94, 0x5a, 0x82, 0xfc, 0xc8, 0xea, 0x1d, 0xf7, 0xf0, 0x96, 0x82, 0x96, 0xd9, 0xff, 0x70,
	0x84, 0x33, 0x0a, 0x0e, 0xf5, 0x93, 0x93, 0x3e, 0xce, 0xa6, 0x86, 0x6e, 0xa7, 0x83, 0x73, 0x0a,
	0x0e, 0x2c, 0x65, 0xc8, 0xa7, 0x86, 0xee, 0x49, 0x0b, 0x6b, 0x09, 0x6b, 0xf6, 0xdb, 0x1d, 0x5c,
	0x50, 0x70, 0x4a, 0xfb, 0xed, 0x16, 0x2e, 0x2a, 0x38, 0xd1, 0x7b, 0xc7, 0x47, 0xb8, 0xa4, 0x20,
	0xb5, 0x7a, 0xc7, 0x5d, 0x0c, 0xa4, 0x0c, 0x85, 0x89, 0x79, 0xd1, 0xee, 0xb6, 0x5a, 0xb8, 0x5c,
	0xff, 0x8d, 0x20, 0x7f, 0x71, 0x19, 0xfb, 0x82, 0xec, 0x42, 0xf9, 0xa3, 0xb4, 0x3d, 0x2a, 0xbe,
	0x47, 0x62, 0x1d, 0xe2, 0x2d, 0xb2, 0x03, 0xa0, 0x88, 0x01, 0xe7, 0xc2, 0x0f, 0x31, 0x22, 0xff,
	0xc3, 0x9e, 0xe5, 0x71, 0xe9, 0x7d, 0xb5, 0x83, 0x5b, 0x71, 0x7d, 0xc6, 0x42, 0x66, 0xf9, 0x38,
	0x43, 0x0e, 0x60, 0xff, 0x19, 0x7d, 0x26, 0x7f, 0x78, 0x38, 0x4b, 0xf6, 0x61, 0x77, 0xf8, 0xcc,
	0x9d, 0x53, 0x45, 0x86, 0x2f, 0xbc, 0x79, 0xb2, 0x07, 0xdb, 0x54, 0x7c, 0xfb, 0xa7, 0xbd, 0xa6,
	0xe6, 0x99, 0x07, 0xd2, 0x0f, 0x6c, 0x11, 0xb2, 0x20, 0xc6, 0x85, 0xfa, 0x2f, 0x04, 0x3b, 0x4f,
	0x7f, 0x84, 0xba, 0x8a, 0xa5, 0x4f, 0x75, 0xe3, 0x4a, 0x4f, 0xbf, 0xec, 0x78, 0x6e, 0x62, 0x44,
	0x00, 0xb4, 0xa1, 0xa1, 0x9f, 0x4f, 0xc6, 0x38, 0xa3, 0xfa, 0x8d, 0x47, 0xc6, 0x82, 0x8e, 0x4c,
	0x63, 0xf6, 0x69, 0x44, 0x17, 0x97, 0x67, 0xc6, 0x00, 0x67, 0x5f, 0xd0, 0xd4, 0x34, 0x27, 0xe9,
	0x74, 0x4f, 0xe8, 0xb1, 0x6e, 0x9a, 0x38, 0xff, 0x82, 0xbe, 0x9a, 0x9c, 0x4f, 0xb0, 0x76, 0xfa,
	0x19, 0x2a, 0xb6, 0x6c, 0xf0, 0x1b, 0x3b, 0xf0, 0xd7, 0x21, 0xe3, 0x4e, 0x83, 0xf9, 0xf6, 0xe6,
	0xc5, 0x9c, 0x96, 0x87, 0xc9, 0x39, 0x57, 0x7b, 0x30, 0x47, 0x5f, 0x1a, 0x2b, 0x3b, 0xbc, 0x89,
	0x96, 0x4a, 0x6d, 0x2e, 0x03, 0xc9, 0x19, 0x0b, 0x9a, 0x8f, 0x89, 0xef, 0x99, 0x6f, 0x37, 0x57,
	0xb2, 0x79, 0xd7, 0xd9, 0xec, 0xcf, 0x52, 0x4b, 0x16, 0xa8, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xb1, 0x20, 0x81, 0x72, 0x57, 0x03, 0x00, 0x00,
}
