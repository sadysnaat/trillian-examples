// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package configpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	HubMultiConfig
	HubBackendSet
	HubBackend
	HubConfigSet
	HubConfig
	TrackedLog
*/
package configpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keyspb "github.com/google/trillian/crypto/keyspb"
import sigpb "github.com/google/trillian/crypto/sigpb"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HubMultiConfig wraps up a HubBackendSet and corresponding HubConfigSet so
// that they can easily be parsed as a single proto.
type HubMultiConfig struct {
	// The set of backends that this configuration will use to send requests to.
	// The names of the backends in the HubBackendSet must all be distinct.
	Backends *HubBackendSet `protobuf:"bytes,1,opt,name=backends" json:"backends,omitempty"`
	// The set oubs that will use the above backends. All the protos in this
	// HubConfigSet must set a valid hub_backend_name for the config to be usable.
	HubConfigs *HubConfigSet `protobuf:"bytes,2,opt,name=hub_configs,json=hubConfigs" json:"hub_configs,omitempty"`
}

func (m *HubMultiConfig) Reset()                    { *m = HubMultiConfig{} }
func (m *HubMultiConfig) String() string            { return proto.CompactTextString(m) }
func (*HubMultiConfig) ProtoMessage()               {}
func (*HubMultiConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HubMultiConfig) GetBackends() *HubBackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

func (m *HubMultiConfig) GetHubConfigs() *HubConfigSet {
	if m != nil {
		return m.HubConfigs
	}
	return nil
}

// HubBackendSet describes the set of Trillian backends that support a gossip hub.
type HubBackendSet struct {
	Backend []*HubBackend `protobuf:"bytes,1,rep,name=backend" json:"backend,omitempty"`
}

func (m *HubBackendSet) Reset()                    { *m = HubBackendSet{} }
func (m *HubBackendSet) String() string            { return proto.CompactTextString(m) }
func (*HubBackendSet) ProtoMessage()               {}
func (*HubBackendSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HubBackendSet) GetBackend() []*HubBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

// HubBackend describes a single Trillian backend.
type HubBackend struct {
	// name defines the name of the hub backend for use in HubConfig messages and must be unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// backend_spec defines the RPC endpoint that clients should use to send requests
	// to this hub backend. These should be in the same format as rpcBackendFlag in the
	// hublog main and must not be an empty string.
	BackendSpec string `protobuf:"bytes,2,opt,name=backend_spec,json=backendSpec" json:"backend_spec,omitempty"`
}

func (m *HubBackend) Reset()                    { *m = HubBackend{} }
func (m *HubBackend) String() string            { return proto.CompactTextString(m) }
func (*HubBackend) ProtoMessage()               {}
func (*HubBackend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HubBackend) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HubBackend) GetBackendSpec() string {
	if m != nil {
		return m.BackendSpec
	}
	return ""
}

// HubConfigSet describes the different hub instances that this instance will run.
type HubConfigSet struct {
	Config []*HubConfig `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
}

func (m *HubConfigSet) Reset()                    { *m = HubConfigSet{} }
func (m *HubConfigSet) String() string            { return proto.CompactTextString(m) }
func (*HubConfigSet) ProtoMessage()               {}
func (*HubConfigSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HubConfigSet) GetConfig() []*HubConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// HubConfig describes the configuration options for a hub instance.
type HubConfig struct {
	LogId      int64                `protobuf:"varint,1,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	Prefix     string               `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	SourceLog  []*TrackedLog        `protobuf:"bytes,3,rep,name=source_log,json=sourceLog" json:"source_log,omitempty"`
	PrivateKey *google_protobuf.Any `protobuf:"bytes,4,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// The public key is included for the convenience of test tools (and obviously
	// should match the private key above); it is not used by the hub.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,5,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	// hub_backend_name if set indicates which backend serves this hub. The name must be
	// one of those defined in the HubBackendSet.
	HubBackendName string `protobuf:"bytes,11,opt,name=hub_backend_name,json=hubBackendName" json:"hub_backend_name,omitempty"`
}

func (m *HubConfig) Reset()                    { *m = HubConfig{} }
func (m *HubConfig) String() string            { return proto.CompactTextString(m) }
func (*HubConfig) ProtoMessage()               {}
func (*HubConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HubConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *HubConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *HubConfig) GetSourceLog() []*TrackedLog {
	if m != nil {
		return m.SourceLog
	}
	return nil
}

func (m *HubConfig) GetPrivateKey() *google_protobuf.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *HubConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *HubConfig) GetHubBackendName() string {
	if m != nil {
		return m.HubBackendName
	}
	return ""
}

// TrackedLog describes a source Log that this gossip hub will accept signed tree heads from.
type TrackedLog struct {
	Name          string                              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url           string                              `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	PublicKey     *keyspb.PublicKey                   `protobuf:"bytes,3,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	HashAlgorithm sigpb.DigitallySigned_HashAlgorithm `protobuf:"varint,4,opt,name=hash_algorithm,json=hashAlgorithm,enum=sigpb.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
}

func (m *TrackedLog) Reset()                    { *m = TrackedLog{} }
func (m *TrackedLog) String() string            { return proto.CompactTextString(m) }
func (*TrackedLog) ProtoMessage()               {}
func (*TrackedLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TrackedLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TrackedLog) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *TrackedLog) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *TrackedLog) GetHashAlgorithm() sigpb.DigitallySigned_HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return sigpb.DigitallySigned_NONE
}

func init() {
	proto.RegisterType((*HubMultiConfig)(nil), "configpb.HubMultiConfig")
	proto.RegisterType((*HubBackendSet)(nil), "configpb.HubBackendSet")
	proto.RegisterType((*HubBackend)(nil), "configpb.HubBackend")
	proto.RegisterType((*HubConfigSet)(nil), "configpb.HubConfigSet")
	proto.RegisterType((*HubConfig)(nil), "configpb.HubConfig")
	proto.RegisterType((*TrackedLog)(nil), "configpb.TrackedLog")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0x75, 0x2b, 0xeb, 0x4b, 0x57, 0x0d, 0x33, 0x46, 0xd9, 0x69, 0x44, 0x1c, 0x2a,
	0x21, 0x39, 0xa8, 0xd5, 0xc4, 0x81, 0x03, 0x1a, 0xe3, 0x50, 0xd4, 0x81, 0x90, 0xcb, 0x3d, 0xb2,
	0x13, 0xd7, 0xb1, 0xea, 0xc6, 0x56, 0x7e, 0x20, 0x72, 0xe1, 0xc0, 0x7f, 0xc4, 0x7f, 0x88, 0x62,
	0x3b, 0xdd, 0x0f, 0x55, 0x82, 0x4b, 0xe2, 0x67, 0x7f, 0xbf, 0xcf, 0x9f, 0xf7, 0x9e, 0x61, 0x94,
	0xe8, 0x7c, 0x2d, 0x05, 0x36, 0x85, 0xae, 0x34, 0x3a, 0x76, 0x91, 0x61, 0x17, 0x57, 0x42, 0x56,
	0x59, 0xcd, 0x70, 0xa2, 0xb7, 0x91, 0xd0, 0x5a, 0x28, 0x1e, 0x55, 0x85, 0x54, 0x4a, 0xd2, 0x3c,
	0x4a, 0x8a, 0xc6, 0x54, 0x3a, 0xda, 0xf0, 0xa6, 0x34, 0xcc, 0xff, 0x5c, 0x82, 0x8b, 0xf9, 0xbf,
	0x6d, 0x65, 0x9b, 0xdf, 0x7d, 0xbd, 0xe9, 0xa5, 0x57, 0xda, 0x88, 0xd5, 0xeb, 0x88, 0xe6, 0x8d,
	0x3b, 0x0a, 0x7f, 0xc1, 0x78, 0x51, 0xb3, 0x2f, 0xb5, 0xaa, 0xe4, 0x8d, 0x45, 0x43, 0x73, 0x38,
	0x66, 0x34, 0xd9, 0xf0, 0x3c, 0x2d, 0x27, 0xbd, 0xcb, 0xde, 0x34, 0x98, 0xbd, 0xc0, 0x1d, 0x35,
	0x5e, 0xd4, 0xec, 0xa3, 0x3b, 0x5c, 0xf1, 0x8a, 0xec, 0x84, 0xe8, 0x1d, 0x04, 0x59, 0xcd, 0x62,
	0xa7, 0x2b, 0x27, 0x07, 0xd6, 0x77, 0xfe, 0xc0, 0xe7, 0xd2, 0xb7, 0x36, 0xc8, 0xba, 0xa8, 0x0c,
	0x3f, 0xc0, 0xc9, 0x83, 0x9c, 0x08, 0xc3, 0x13, 0x9f, 0x75, 0xd2, 0xbb, 0xec, 0x4f, 0x83, 0xd9,
	0xd9, 0xbe, 0xdb, 0x49, 0x27, 0x0a, 0x6f, 0x00, 0xee, 0xb6, 0x11, 0x82, 0xc3, 0x9c, 0x6e, 0xb9,
	0x05, 0x1f, 0x12, 0xbb, 0x46, 0xaf, 0x60, 0xe4, 0xc5, 0x71, 0x69, 0x78, 0x62, 0xe1, 0x86, 0x24,
	0xf0, 0x7b, 0x2b, 0xc3, 0x93, 0xf0, 0x3d, 0x8c, 0xee, 0x13, 0xa2, 0x37, 0x30, 0x70, 0x97, 0x7a,
	0x86, 0x67, 0x7b, 0x2a, 0x21, 0x5e, 0x12, 0xfe, 0x3e, 0x80, 0xe1, 0x6e, 0x17, 0x3d, 0x87, 0x81,
	0xd2, 0x22, 0x96, 0xa9, 0x65, 0xe8, 0x93, 0x23, 0xa5, 0xc5, 0xe7, 0x14, 0x9d, 0xc3, 0xc0, 0x14,
	0x7c, 0x2d, 0x7f, 0xfa, 0xeb, 0x7d, 0x84, 0xe6, 0x00, 0xa5, 0xae, 0x8b, 0x84, 0xc7, 0x4a, 0x8b,
	0x49, 0xff, 0x71, 0xc5, 0xdf, 0x8b, 0x96, 0x32, 0xbd, 0xd5, 0x82, 0x0c, 0x9d, 0xee, 0x56, 0x0b,
	0x74, 0x05, 0x81, 0x29, 0xe4, 0x0f, 0x5a, 0xf1, 0x78, 0xc3, 0x9b, 0xc9, 0xa1, 0xed, 0xf6, 0x19,
	0x76, 0x53, 0xc6, 0xdd, 0x94, 0xf1, 0x75, 0xde, 0x10, 0xf0, 0xc2, 0x25, 0x6f, 0xd0, 0x5b, 0x00,
	0x53, 0x33, 0x25, 0x13, 0xeb, 0x3a, 0xb2, 0xae, 0xa7, 0xd8, 0x3f, 0xaf, 0x6f, 0xf6, 0x64, 0xc9,
	0x1b, 0x32, 0x34, 0xdd, 0x12, 0x4d, 0xe1, 0xb4, 0x1d, 0x6b, 0xd7, 0x3e, 0xdb, 0xda, 0xc0, 0xf2,
	0x8f, 0xb3, 0x5d, 0xd3, 0xbf, 0xd2, 0x2d, 0x0f, 0xff, 0xf4, 0x00, 0xee, 0x60, 0xf7, 0xce, 0xe1,
	0x14, 0xfa, 0x75, 0xa1, 0x7c, 0xfd, 0xed, 0xf2, 0x11, 0x50, 0xff, 0x3f, 0x80, 0x96, 0x30, 0xce,
	0x68, 0x99, 0xc5, 0x54, 0x09, 0x5d, 0xc8, 0x2a, 0xdb, 0xda, 0xe2, 0xc7, 0xb3, 0xd7, 0xd8, 0xbd,
	0xf7, 0x4f, 0x52, 0xc8, 0x8a, 0x2a, 0xd5, 0xac, 0xa4, 0xc8, 0x79, 0x8a, 0x17, 0xb4, 0xcc, 0xae,
	0x3b, 0x2d, 0x39, 0xc9, 0xee, 0x87, 0x6c, 0x60, 0x3b, 0x35, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x86, 0x15, 0x75, 0x02, 0xa3, 0x03, 0x00, 0x00,
}