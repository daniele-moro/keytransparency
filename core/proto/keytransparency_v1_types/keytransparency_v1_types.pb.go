// Code generated by protoc-gen-go.
// source: core/proto/keytransparency_v1_types/keytransparency_v1_types.proto
// DO NOT EDIT!

/*
Package keytransparency_v1_types is a generated protocol buffer package.

It is generated from these files:
	core/proto/keytransparency_v1_types/keytransparency_v1_types.proto

It has these top-level messages:
	Committed
	Profile
	EntryUpdate
	Entry
	PublicKey
	KeyValue
	SignedKV
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
	HkpLookupRequest
	HttpResponse
*/
package keytransparency_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ctmap "github.com/google/key-transparency/core/proto/ctmap"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Profile contains data hidden behind the cryptographic commitment.
type Profile struct {
	// Keys is a map of application IDs to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the serialized Profile protobuf.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign MapHeads with.
type PublicKey struct {
	// Key formats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,json=rsaVerifyingSha2562048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	// key contains the map entry key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value contains the map entry value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// key_value is a serialized KeyValue.
	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	// signatures on keyvalue. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SignedKV) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse merkle tree at end_epoch.
	LeafProof *ctmap.GetLeafResponse `protobuf:"bytes,5,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smh contains the signed map head for the sparse merkle tree.
	// smh is also stored in the append only log.
	Smh *ctmap.SignedMapHead `protobuf:"bytes,6,opt,name=smh" json:"smh,omitempty"`
	// smh_sct is the SCT showing that smh was submitted to CT logs.
	// TODO: Support storing smh in multiple logs.
	SmhSct []byte `protobuf:"bytes,7,opt,name=smh_sct,json=smhSct,proto3" json:"smh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *ctmap.GetLeafResponse {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmh() *ctmap.SignedMapHead {
	if m != nil {
		return m.Smh
	}
	return nil
}

// Get a list of historical values for a user.
type ListEntryHistoryRequest struct {
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epcoh.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// A paginated history of values for a user.
type ListEntryHistoryResponse struct {
	// values represents the list of values this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// next_start is the next page token to query for pagination.
	// next_start is 0 when there are no more results to fetch.
	NextStart int64 `protobuf:"varint,2,opt,name=next_start,json=nextStart" json:"next_start,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

// Update a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the new account to be registered.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// entry_update contains the user submitted update(s).
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	// proof contains a proof that the update has been included in the tree.
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()                    { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()               {}
func (*HkpLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()                    { *m = HttpResponse{} }
func (m *HttpResponse) String() string            { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()               {}
func (*HttpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.types.Committed")
	proto.RegisterType((*Profile)(nil), "keytransparency.v1.types.Profile")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.types.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.types.Entry")
	proto.RegisterType((*PublicKey)(nil), "keytransparency.v1.types.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.types.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.types.SignedKV")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.types.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.types.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.types.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.types.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.types.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.types.UpdateEntryResponse")
	proto.RegisterType((*HkpLookupRequest)(nil), "keytransparency.v1.types.HkpLookupRequest")
	proto.RegisterType((*HttpResponse)(nil), "keytransparency.v1.types.HttpResponse")
}

func init() {
	proto.RegisterFile("core/proto/keytransparency_v1_types/keytransparency_v1_types.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x4d, 0xf3, 0xe3, 0x93, 0x6a, 0x5b, 0x0d, 0x55, 0xd7, 0x14, 0x81, 0x16, 0x23, 0xd0,
	0x6a, 0x11, 0x2e, 0x35, 0x14, 0x96, 0x45, 0x68, 0xa1, 0x68, 0x45, 0xd0, 0x16, 0xa9, 0x72, 0xa0,
	0x5c, 0x5a, 0x53, 0x7b, 0x92, 0x58, 0x4d, 0x3c, 0xc3, 0xcc, 0x38, 0x6c, 0x90, 0xb8, 0xe2, 0x96,
	0xc7, 0xe0, 0x99, 0x78, 0x1a, 0x2e, 0x38, 0x33, 0x63, 0xa7, 0x6e, 0xd9, 0x50, 0x90, 0xb8, 0x69,
	0xe7, 0xfc, 0x7c, 0xe7, 0xe7, 0x3b, 0x27, 0xc7, 0x70, 0x9a, 0x71, 0xc9, 0x8e, 0x84, 0xe4, 0x9a,
	0x1f, 0x5d, 0xb1, 0x95, 0x96, 0xb4, 0x54, 0x82, 0x4a, 0x56, 0x66, 0xab, 0x74, 0x79, 0x9c, 0xea,
	0x95, 0x60, 0x6a, 0xa3, 0x21, 0xb2, 0x38, 0x12, 0xdc, 0xb2, 0x47, 0xcb, 0xe3, 0xc8, 0xda, 0x0f,
	0x9f, 0x4e, 0x0b, 0x3d, 0xab, 0x2e, 0xa3, 0x8c, 0x2f, 0x8e, 0xa6, 0x9c, 0x4f, 0xe7, 0xcc, 0xc4,
	0x7a, 0xbf, 0xed, 0x7c, 0xd4, 0x2a, 0x20, 0xd3, 0x0b, 0x2a, 0xdc, 0x5f, 0x17, 0x3a, 0x3c, 0x06,
	0xff, 0x2b, 0xbe, 0x58, 0x14, 0x5a, 0xb3, 0x9c, 0xec, 0x41, 0x07, 0xd1, 0x81, 0xf7, 0xc0, 0x7b,
	0xb8, 0x93, 0x98, 0x27, 0x21, 0xb0, 0x9d, 0x53, 0x4d, 0x83, 0x2d, 0xab, 0xb2, 0xef, 0xf0, 0x57,
	0x0f, 0xfa, 0xe7, 0x92, 0x4f, 0x8a, 0x39, 0x23, 0x4f, 0x61, 0x1b, 0xdd, 0x14, 0x42, 0x3a, 0x0f,
	0x87, 0xf1, 0x7b, 0xd1, 0xa6, 0x42, 0xa3, 0x1a, 0x10, 0x3d, 0x47, 0xef, 0x67, 0xa5, 0x96, 0xab,
	0xc4, 0x02, 0x0f, 0x3f, 0x01, 0x7f, 0xad, 0x6a, 0xe7, 0xf7, 0x5d, 0xfe, 0x7d, 0xe8, 0x2e, 0xe9,
	0xbc, 0x62, 0x75, 0x01, 0x4e, 0x78, 0xb2, 0xf5, 0xd8, 0x0b, 0x7f, 0xf3, 0x60, 0x68, 0x51, 0xdf,
	0x0b, 0xac, 0x8a, 0x91, 0x27, 0xd0, 0xab, 0xec, 0xcb, 0xba, 0x0e, 0xe3, 0x70, 0x73, 0x2d, 0xe3,
	0x62, 0x5a, 0xb2, 0xfc, 0xf9, 0x45, 0x52, 0x23, 0xc8, 0x97, 0xe0, 0x67, 0x0d, 0x09, 0x41, 0xc7,
	0xc2, 0xdf, 0xde, 0x0c, 0x5f, 0xf3, 0x95, 0x5c, 0xa3, 0xc2, 0x0a, 0xba, 0xae, 0x87, 0x37, 0x01,
	0x9c, 0x76, 0xc1, 0x4a, 0x5d, 0x53, 0xd9, 0xd2, 0x90, 0x33, 0xd8, 0xa5, 0x95, 0x9e, 0x71, 0x59,
	0xfc, 0xcc, 0xf2, 0xd4, 0x92, 0xb7, 0x65, 0xc9, 0xfb, 0x87, 0x8c, 0xe7, 0xd5, 0xe5, 0xbc, 0xc8,
	0x90, 0xa7, 0xe4, 0xde, 0x35, 0xd6, 0xd0, 0x16, 0xfe, 0xee, 0x81, 0xbf, 0xb6, 0x92, 0x43, 0xe8,
	0xb3, 0x3c, 0x3e, 0x39, 0x39, 0xfe, 0xd4, 0x25, 0x1e, 0xbd, 0x92, 0x34, 0x0a, 0xf2, 0x19, 0xbc,
	0x26, 0x15, 0x4d, 0x97, 0x4c, 0x16, 0x93, 0x55, 0x51, 0x4e, 0x53, 0x35, 0xa3, 0xf1, 0xc9, 0xc7,
	0x69, 0xfc, 0xc1, 0x47, 0x8f, 0x1d, 0xbb, 0xe8, 0x7d, 0x80, 0x2e, 0x17, 0x8d, 0xc7, 0xd8, 0x3a,
	0x18, 0x3b, 0x89, 0x61, 0x9f, 0x65, 0xf9, 0x0d, 0xb8, 0x40, 0x9b, 0xe5, 0xca, 0xe0, 0x88, 0xb5,
	0xae, 0x91, 0xe7, 0x68, 0x3b, 0x05, 0x18, 0x60, 0x43, 0x76, 0x8f, 0xc3, 0x18, 0x06, 0x58, 0xdf,
	0x85, 0x19, 0xde, 0x4b, 0x96, 0xec, 0xa5, 0x43, 0x0e, 0xff, 0xf0, 0x60, 0xd0, 0x4c, 0x8a, 0xbc,
	0x0e, 0xbe, 0x09, 0xe6, 0xdc, 0x1c, 0xd4, 0x44, 0x77, 0x11, 0x13, 0x00, 0x85, 0x8e, 0x54, 0x57,
	0x92, 0x35, 0x6c, 0xc6, 0x77, 0x8f, 0xdf, 0x3e, 0x1c, 0xc8, 0x6d, 0x64, 0x2b, 0x0a, 0x52, 0x39,
	0x10, 0x92, 0x2d, 0x0b, 0x5e, 0x29, 0xd7, 0x65, 0xb2, 0x96, 0x0f, 0x3f, 0x87, 0xdd, 0x5b, 0xd0,
	0x76, 0x53, 0xbd, 0xbb, 0x36, 0xf7, 0x11, 0xec, 0x7e, 0xcd, 0xb4, 0x4b, 0xc9, 0x7e, 0xac, 0x98,
	0xd2, 0xe4, 0x3e, 0xf4, 0x2b, 0xc5, 0x64, 0x5a, 0xe4, 0xf5, 0xf2, 0xf7, 0x8c, 0xf8, 0x4d, 0x1e,
	0xfe, 0xe9, 0xc1, 0xde, 0xb5, 0xb3, 0x12, 0xbc, 0x54, 0x96, 0xc1, 0xa5, 0x9c, 0x34, 0x0c, 0xe2,
	0xd3, 0xd0, 0x83, 0xff, 0x52, 0xfc, 0x49, 0xf3, 0x49, 0x9d, 0x70, 0x80, 0x8a, 0x73, 0x23, 0xff,
	0x0f, 0xdb, 0x4d, 0x4e, 0x00, 0xe6, 0x8c, 0x36, 0x09, 0xba, 0x36, 0xc6, 0x41, 0xe4, 0xee, 0x08,
	0x96, 0x77, 0x86, 0xb6, 0xa6, 0xba, 0xc4, 0x37, 0x9e, 0x2e, 0xf3, 0xbb, 0xd0, 0x51, 0x8b, 0x59,
	0xd0, 0xb3, 0xfe, 0xfb, 0xb5, 0xbf, 0xa3, 0xff, 0x5b, 0x2a, 0x46, 0x8c, 0xe6, 0x89, 0x71, 0x30,
	0xed, 0xe3, 0xbf, 0x54, 0x65, 0x3a, 0xe8, 0xdb, 0xe2, 0x7b, 0x28, 0x8e, 0x33, 0x1d, 0x32, 0xb8,
	0x7f, 0x56, 0x28, 0xd7, 0xfe, 0x08, 0x1f, 0xfc, 0x6e, 0xca, 0x0c, 0xf1, 0x4a, 0x53, 0xa9, 0x2d,
	0x0f, 0x9d, 0xc4, 0x09, 0x86, 0x21, 0x41, 0xa7, 0x2c, 0x55, 0xf8, 0xd3, 0xb1, 0x24, 0x74, 0x71,
	0xa0, 0xa8, 0x18, 0xa3, 0x1c, 0xfe, 0x02, 0xc1, 0xdf, 0xd3, 0xd4, 0x64, 0x9f, 0x42, 0xcf, 0x8e,
	0xae, 0xb9, 0x71, 0x8f, 0x36, 0x53, 0x77, 0x7b, 0x50, 0x49, 0x8d, 0x24, 0x6f, 0x00, 0x94, 0xec,
	0x85, 0x4e, 0xdb, 0x75, 0xf9, 0x46, 0x33, 0x36, 0x8a, 0xf0, 0x27, 0x20, 0xee, 0x88, 0xfd, 0xab,
	0x9d, 0x20, 0x23, 0xd8, 0x61, 0xc6, 0x31, 0xbd, 0x71, 0xef, 0xde, 0xd9, 0x5c, 0x57, 0xeb, 0x4c,
	0x26, 0x43, 0x76, 0x2d, 0x84, 0x3f, 0xc0, 0xab, 0x37, 0x12, 0xd7, 0x2d, 0x7f, 0x01, 0x5d, 0x37,
	0x68, 0xcf, 0x46, 0xfe, 0x2f, 0x1d, 0x3b, 0xa0, 0x39, 0xce, 0x7b, 0xa3, 0x2b, 0x71, 0xc6, 0xf9,
	0x55, 0x25, 0x9a, 0x86, 0xee, 0xc1, 0x16, 0x17, 0x75, 0x2f, 0xf8, 0x22, 0x07, 0xd0, 0x53, 0x8c,
	0xca, 0x6c, 0x66, 0x3b, 0xc0, 0xfe, 0x9c, 0x44, 0x02, 0xe8, 0x73, 0xa1, 0x0b, 0x8c, 0x67, 0x07,
	0xe5, 0x27, 0x8d, 0x68, 0x46, 0xcb, 0x5e, 0x50, 0xdc, 0x92, 0x6d, 0xab, 0x77, 0x02, 0x79, 0x00,
	0xc3, 0x09, 0x1e, 0x1d, 0x26, 0x85, 0x2c, 0xf0, 0xe4, 0x76, 0xad, 0xad, 0xad, 0x0a, 0x9f, 0xc1,
	0xce, 0x48, 0x6b, 0xb1, 0x6e, 0xf0, 0x2d, 0xd8, 0xc9, 0x78, 0xa9, 0x91, 0x0a, 0x7b, 0x9e, 0xea,
	0x9a, 0x86, 0xb5, 0xee, 0x3b, 0x54, 0x99, 0x0f, 0xdf, 0x25, 0xcf, 0x57, 0xcd, 0x87, 0xcf, 0xbc,
	0x2f, 0x7b, 0xf6, 0x93, 0xf9, 0xe1, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x00, 0x1c, 0xe4, 0x1e,
	0xd3, 0x07, 0x00, 0x00,
}