// Code generated by protoc-gen-go.
// source: keymaster.proto
// DO NOT EDIT!

/*
Package keymaster is a generated protocol buffer package.

It is generated from these files:
	keymaster.proto

It has these top-level messages:
	Metadata
	SigningKey
	VerifyingKey
	KeySet
*/
package keymaster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KeyStatus defines a key status.
type SigningKey_KeyStatus int32

const (
	SigningKey_ACTIVE     SigningKey_KeyStatus = 0
	SigningKey_INACTIVE   SigningKey_KeyStatus = 1
	SigningKey_DEPRECATED SigningKey_KeyStatus = 2
)

var SigningKey_KeyStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
	2: "DEPRECATED",
}
var SigningKey_KeyStatus_value = map[string]int32{
	"ACTIVE":     0,
	"INACTIVE":   1,
	"DEPRECATED": 2,
}

func (x SigningKey_KeyStatus) String() string {
	return proto.EnumName(SigningKey_KeyStatus_name, int32(x))
}
func (SigningKey_KeyStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// KeyStatus defines a key status.
type VerifyingKey_KeyStatus int32

const (
	VerifyingKey_ACTIVE     VerifyingKey_KeyStatus = 0
	VerifyingKey_DEPRECATED VerifyingKey_KeyStatus = 1
)

var VerifyingKey_KeyStatus_name = map[int32]string{
	0: "ACTIVE",
	1: "DEPRECATED",
}
var VerifyingKey_KeyStatus_value = map[string]int32{
	"ACTIVE":     0,
	"DEPRECATED": 1,
}

func (x VerifyingKey_KeyStatus) String() string {
	return proto.EnumName(VerifyingKey_KeyStatus_name, int32(x))
}
func (VerifyingKey_KeyStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Metadata struct {
	// key_id represents a key identifier.
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	// added_at determines the time this key has been added to the key set.
	AddedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=added_at,json=addedAt" json:"added_at,omitempty"`
	// description contains an arbitrary text describing the key.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Metadata) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *Metadata) GetAddedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.AddedAt
	}
	return nil
}

func (m *Metadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// SigningKey represents a private key.
type SigningKey struct {
	// metadata contains information about this key..
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// key_material contains the key material in PEM format.
	KeyMaterial []byte `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
	// status determines the status of this key, e.g., active, deprecated, etc.
	Status SigningKey_KeyStatus `protobuf:"varint,3,opt,name=status,enum=keymaster.SigningKey_KeyStatus" json:"status,omitempty"`
}

func (m *SigningKey) Reset()                    { *m = SigningKey{} }
func (m *SigningKey) String() string            { return proto.CompactTextString(m) }
func (*SigningKey) ProtoMessage()               {}
func (*SigningKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SigningKey) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SigningKey) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

func (m *SigningKey) GetStatus() SigningKey_KeyStatus {
	if m != nil {
		return m.Status
	}
	return SigningKey_ACTIVE
}

// VerifyingKey represents a public key.
type VerifyingKey struct {
	// metadata contains information about this key..
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// key_material contains the key material in PEM format.
	KeyMaterial []byte `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
	// status determines the status of this key, e.g., active, deprecated, etc.
	Status VerifyingKey_KeyStatus `protobuf:"varint,3,opt,name=status,enum=keymaster.VerifyingKey_KeyStatus" json:"status,omitempty"`
}

func (m *VerifyingKey) Reset()                    { *m = VerifyingKey{} }
func (m *VerifyingKey) String() string            { return proto.CompactTextString(m) }
func (*VerifyingKey) ProtoMessage()               {}
func (*VerifyingKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *VerifyingKey) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VerifyingKey) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

func (m *VerifyingKey) GetStatus() VerifyingKey_KeyStatus {
	if m != nil {
		return m.Status
	}
	return VerifyingKey_ACTIVE
}

// KeySet contains a set of public and private keys.
type KeySet struct {
	// signing_keys holds a map of private keys keyed by the ID of their
	// corresponding public keys.
	SigningKeys map[string]*SigningKey `protobuf:"bytes,1,rep,name=signing_keys,json=signingKeys" json:"signing_keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// verifying_keys holds a map of public keys keyed by their IDs.
	VerifyingKeys map[string]*VerifyingKey `protobuf:"bytes,2,rep,name=verifying_keys,json=verifyingKeys" json:"verifying_keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *KeySet) Reset()                    { *m = KeySet{} }
func (m *KeySet) String() string            { return proto.CompactTextString(m) }
func (*KeySet) ProtoMessage()               {}
func (*KeySet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *KeySet) GetSigningKeys() map[string]*SigningKey {
	if m != nil {
		return m.SigningKeys
	}
	return nil
}

func (m *KeySet) GetVerifyingKeys() map[string]*VerifyingKey {
	if m != nil {
		return m.VerifyingKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "keymaster.Metadata")
	proto.RegisterType((*SigningKey)(nil), "keymaster.SigningKey")
	proto.RegisterType((*VerifyingKey)(nil), "keymaster.VerifyingKey")
	proto.RegisterType((*KeySet)(nil), "keymaster.KeySet")
	proto.RegisterEnum("keymaster.SigningKey_KeyStatus", SigningKey_KeyStatus_name, SigningKey_KeyStatus_value)
	proto.RegisterEnum("keymaster.VerifyingKey_KeyStatus", VerifyingKey_KeyStatus_name, VerifyingKey_KeyStatus_value)
}

func init() { proto.RegisterFile("keymaster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x8b, 0xda, 0x40,
	0x1c, 0xc5, 0x3b, 0x91, 0x4d, 0xf5, 0x9b, 0xd4, 0xca, 0x94, 0xa5, 0xe2, 0x65, 0xdd, 0x50, 0xa8,
	0x50, 0x1a, 0xc1, 0xb2, 0xf4, 0xc7, 0x4d, 0x76, 0x73, 0x10, 0xd9, 0x52, 0x66, 0xed, 0x42, 0x4f,
	0x32, 0xdb, 0x7c, 0x57, 0x86, 0x98, 0x44, 0x32, 0xa3, 0x30, 0xfd, 0x03, 0x7b, 0x2c, 0xfd, 0x93,
	0x8a, 0x93, 0x44, 0x67, 0x77, 0xf5, 0xd8, 0x9b, 0x3e, 0x1f, 0x9f, 0xf7, 0xde, 0xcc, 0x08, 0x2f,
	0x13, 0xd4, 0x29, 0x97, 0x0a, 0x8b, 0x70, 0x55, 0xe4, 0x2a, 0xa7, 0xad, 0x9d, 0xd0, 0x3b, 0x5b,
	0xe4, 0xf9, 0x62, 0x89, 0x43, 0xf3, 0xc3, 0xdd, 0xfa, 0x7e, 0xa8, 0x44, 0x8a, 0x52, 0xf1, 0x74,
	0x55, 0x7a, 0x83, 0x5f, 0xd0, 0xbc, 0x46, 0xc5, 0x63, 0xae, 0x38, 0x3d, 0x05, 0x37, 0x41, 0x3d,
	0x17, 0x71, 0x97, 0xf4, 0xc9, 0xa0, 0xc5, 0x4e, 0x12, 0xd4, 0x93, 0x98, 0x5e, 0x40, 0x93, 0xc7,
	0x31, 0xc6, 0x73, 0xae, 0xba, 0x4e, 0x9f, 0x0c, 0xbc, 0x51, 0x2f, 0x2c, 0xb1, 0x61, 0x8d, 0x0d,
	0x67, 0x35, 0x96, 0x3d, 0x37, 0xde, 0xb1, 0xa2, 0x7d, 0xf0, 0x62, 0x94, 0x3f, 0x0b, 0xb1, 0x52,
	0x22, 0xcf, 0xba, 0x0d, 0x83, 0xb4, 0xa5, 0xe0, 0x2f, 0x01, 0xb8, 0x11, 0x8b, 0x4c, 0x64, 0x8b,
	0x29, 0x6a, 0x3a, 0x84, 0x66, 0x5a, 0x55, 0x31, 0x05, 0xbc, 0xd1, 0xab, 0x70, 0x3f, 0xad, 0x6e,
	0xc9, 0x76, 0x26, 0x7a, 0x0e, 0xfe, 0xb6, 0x6f, 0xca, 0x15, 0x16, 0x82, 0x2f, 0x4d, 0x39, 0x9f,
	0x79, 0x09, 0xea, 0xeb, 0x4a, 0xa2, 0x1f, 0xc1, 0x95, 0x8a, 0xab, 0xb5, 0x34, 0xf9, 0xed, 0xd1,
	0x99, 0x45, 0xdc, 0x47, 0x87, 0x53, 0xd4, 0x37, 0xc6, 0xc6, 0x2a, 0x7b, 0x70, 0x01, 0xad, 0x9d,
	0x48, 0x01, 0xdc, 0xf1, 0xe5, 0x6c, 0x72, 0x1b, 0x75, 0x9e, 0x51, 0x1f, 0x9a, 0x93, 0xaf, 0xd5,
	0x37, 0x42, 0xdb, 0x00, 0x57, 0xd1, 0x37, 0x16, 0x5d, 0x8e, 0x67, 0xd1, 0x55, 0xc7, 0x09, 0x7e,
	0x13, 0xf0, 0x6f, 0xb1, 0x10, 0xf7, 0xfa, 0x3f, 0x8e, 0xfa, 0xfc, 0x68, 0xd4, 0xb9, 0x45, 0xb4,
	0xc3, 0x0f, 0xcc, 0x7a, 0x7b, 0x6c, 0xd6, 0xc3, 0x21, 0x24, 0xf8, 0xe3, 0x80, 0xbb, 0x75, 0xa2,
	0xa2, 0x11, 0xf8, 0xb2, 0x3c, 0xaa, 0x79, 0x82, 0x5a, 0x76, 0x49, 0xbf, 0x31, 0xf0, 0x46, 0x81,
	0x15, 0x5a, 0x1a, 0xad, 0x03, 0x95, 0x51, 0xa6, 0x0a, 0xcd, 0x3c, 0xb9, 0x57, 0xe8, 0x14, 0xda,
	0x9b, 0xba, 0x5c, 0x09, 0x72, 0x0c, 0xe8, 0xcd, 0x53, 0x90, 0x3d, 0xa2, 0x42, 0xbd, 0xd8, 0xd8,
	0x5a, 0xef, 0x3b, 0x74, 0x1e, 0xa7, 0xd1, 0x0e, 0x34, 0x12, 0xd4, 0xd5, 0xdb, 0xdd, 0x7e, 0xa4,
	0xef, 0xe0, 0x64, 0xc3, 0x97, 0x6b, 0xac, 0x9e, 0xed, 0xe9, 0xc1, 0xcb, 0x67, 0xa5, 0xe7, 0x8b,
	0xf3, 0x89, 0xf4, 0x7e, 0x00, 0x7d, 0x9a, 0x7d, 0x00, 0xfc, 0xfe, 0x21, 0xf8, 0xf5, 0x91, 0x0b,
	0xb0, 0xd0, 0x77, 0xae, 0xf9, 0xaf, 0x7c, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x9e, 0x98,
	0x95, 0xae, 0x03, 0x00, 0x00,
}
