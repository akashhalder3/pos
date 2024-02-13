// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: proto/eth/v2/beacon_lightclient.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	github_com_prysmaticlabs_prysm_v4_consensus_types_primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
	_ "github.com/prysmaticlabs/prysm/v4/proto/eth/ext"
	v1 "github.com/prysmaticlabs/prysm/v4/proto/eth/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LightClientBootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header                     *v1.BeaconBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	CurrentSyncCommittee       *SyncCommittee        `protobuf:"bytes,2,opt,name=current_sync_committee,json=currentSyncCommittee,proto3" json:"current_sync_committee,omitempty"`
	CurrentSyncCommitteeBranch [][]byte              `protobuf:"bytes,3,rep,name=current_sync_committee_branch,json=currentSyncCommitteeBranch,proto3" json:"current_sync_committee_branch,omitempty"`
}

func (x *LightClientBootstrap) Reset() {
	*x = LightClientBootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientBootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientBootstrap) ProtoMessage() {}

func (x *LightClientBootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientBootstrap.ProtoReflect.Descriptor instead.
func (*LightClientBootstrap) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{0}
}

func (x *LightClientBootstrap) GetHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *LightClientBootstrap) GetCurrentSyncCommittee() *SyncCommittee {
	if x != nil {
		return x.CurrentSyncCommittee
	}
	return nil
}

func (x *LightClientBootstrap) GetCurrentSyncCommitteeBranch() [][]byte {
	if x != nil {
		return x.CurrentSyncCommitteeBranch
	}
	return nil
}

type LightClientUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestedHeader          *v1.BeaconBlockHeader                                             `protobuf:"bytes,1,opt,name=attested_header,json=attestedHeader,proto3" json:"attested_header,omitempty"`
	NextSyncCommittee       *SyncCommittee                                                    `protobuf:"bytes,2,opt,name=next_sync_committee,json=nextSyncCommittee,proto3" json:"next_sync_committee,omitempty"`
	NextSyncCommitteeBranch [][]byte                                                          `protobuf:"bytes,3,rep,name=next_sync_committee_branch,json=nextSyncCommitteeBranch,proto3" json:"next_sync_committee_branch,omitempty"`
	FinalizedHeader         *v1.BeaconBlockHeader                                             `protobuf:"bytes,4,opt,name=finalized_header,json=finalizedHeader,proto3" json:"finalized_header,omitempty"`
	FinalityBranch          [][]byte                                                          `protobuf:"bytes,5,rep,name=finality_branch,json=finalityBranch,proto3" json:"finality_branch,omitempty"`
	SyncAggregate           *v1.SyncAggregate                                                 `protobuf:"bytes,6,opt,name=sync_aggregate,json=syncAggregate,proto3" json:"sync_aggregate,omitempty"`
	SignatureSlot           github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,7,opt,name=signature_slot,json=signatureSlot,proto3" json:"signature_slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
}

func (x *LightClientUpdate) Reset() {
	*x = LightClientUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientUpdate) ProtoMessage() {}

func (x *LightClientUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientUpdate.ProtoReflect.Descriptor instead.
func (*LightClientUpdate) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{1}
}

func (x *LightClientUpdate) GetAttestedHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.AttestedHeader
	}
	return nil
}

func (x *LightClientUpdate) GetNextSyncCommittee() *SyncCommittee {
	if x != nil {
		return x.NextSyncCommittee
	}
	return nil
}

func (x *LightClientUpdate) GetNextSyncCommitteeBranch() [][]byte {
	if x != nil {
		return x.NextSyncCommitteeBranch
	}
	return nil
}

func (x *LightClientUpdate) GetFinalizedHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.FinalizedHeader
	}
	return nil
}

func (x *LightClientUpdate) GetFinalityBranch() [][]byte {
	if x != nil {
		return x.FinalityBranch
	}
	return nil
}

func (x *LightClientUpdate) GetSyncAggregate() *v1.SyncAggregate {
	if x != nil {
		return x.SyncAggregate
	}
	return nil
}

func (x *LightClientUpdate) GetSignatureSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.SignatureSlot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

type LightClientFinalityUpdateWithVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version Version                    `protobuf:"varint,1,opt,name=version,proto3,enum=ethereum.eth.v2.Version" json:"version,omitempty"`
	Data    *LightClientFinalityUpdate `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LightClientFinalityUpdateWithVersion) Reset() {
	*x = LightClientFinalityUpdateWithVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientFinalityUpdateWithVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientFinalityUpdateWithVersion) ProtoMessage() {}

func (x *LightClientFinalityUpdateWithVersion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientFinalityUpdateWithVersion.ProtoReflect.Descriptor instead.
func (*LightClientFinalityUpdateWithVersion) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{2}
}

func (x *LightClientFinalityUpdateWithVersion) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_PHASE0
}

func (x *LightClientFinalityUpdateWithVersion) GetData() *LightClientFinalityUpdate {
	if x != nil {
		return x.Data
	}
	return nil
}

type LightClientFinalityUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestedHeader  *v1.BeaconBlockHeader                                             `protobuf:"bytes,1,opt,name=attested_header,json=attestedHeader,proto3" json:"attested_header,omitempty"`
	FinalizedHeader *v1.BeaconBlockHeader                                             `protobuf:"bytes,2,opt,name=finalized_header,json=finalizedHeader,proto3" json:"finalized_header,omitempty"`
	FinalityBranch  [][]byte                                                          `protobuf:"bytes,3,rep,name=finality_branch,json=finalityBranch,proto3" json:"finality_branch,omitempty"`
	SyncAggregate   *v1.SyncAggregate                                                 `protobuf:"bytes,4,opt,name=sync_aggregate,json=syncAggregate,proto3" json:"sync_aggregate,omitempty"`
	SignatureSlot   github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,5,opt,name=signature_slot,json=signatureSlot,proto3" json:"signature_slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
}

func (x *LightClientFinalityUpdate) Reset() {
	*x = LightClientFinalityUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientFinalityUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientFinalityUpdate) ProtoMessage() {}

func (x *LightClientFinalityUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientFinalityUpdate.ProtoReflect.Descriptor instead.
func (*LightClientFinalityUpdate) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{3}
}

func (x *LightClientFinalityUpdate) GetAttestedHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.AttestedHeader
	}
	return nil
}

func (x *LightClientFinalityUpdate) GetFinalizedHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.FinalizedHeader
	}
	return nil
}

func (x *LightClientFinalityUpdate) GetFinalityBranch() [][]byte {
	if x != nil {
		return x.FinalityBranch
	}
	return nil
}

func (x *LightClientFinalityUpdate) GetSyncAggregate() *v1.SyncAggregate {
	if x != nil {
		return x.SyncAggregate
	}
	return nil
}

func (x *LightClientFinalityUpdate) GetSignatureSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.SignatureSlot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

type LightClientOptimisticUpdateWithVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version Version                      `protobuf:"varint,1,opt,name=version,proto3,enum=ethereum.eth.v2.Version" json:"version,omitempty"`
	Data    *LightClientOptimisticUpdate `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LightClientOptimisticUpdateWithVersion) Reset() {
	*x = LightClientOptimisticUpdateWithVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientOptimisticUpdateWithVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientOptimisticUpdateWithVersion) ProtoMessage() {}

func (x *LightClientOptimisticUpdateWithVersion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientOptimisticUpdateWithVersion.ProtoReflect.Descriptor instead.
func (*LightClientOptimisticUpdateWithVersion) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{4}
}

func (x *LightClientOptimisticUpdateWithVersion) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_PHASE0
}

func (x *LightClientOptimisticUpdateWithVersion) GetData() *LightClientOptimisticUpdate {
	if x != nil {
		return x.Data
	}
	return nil
}

type LightClientOptimisticUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttestedHeader *v1.BeaconBlockHeader                                             `protobuf:"bytes,1,opt,name=attested_header,json=attestedHeader,proto3" json:"attested_header,omitempty"`
	SyncAggregate  *v1.SyncAggregate                                                 `protobuf:"bytes,2,opt,name=sync_aggregate,json=syncAggregate,proto3" json:"sync_aggregate,omitempty"`
	SignatureSlot  github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot `protobuf:"varint,3,opt,name=signature_slot,json=signatureSlot,proto3" json:"signature_slot,omitempty" cast-type:"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives.Slot"`
}

func (x *LightClientOptimisticUpdate) Reset() {
	*x = LightClientOptimisticUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LightClientOptimisticUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LightClientOptimisticUpdate) ProtoMessage() {}

func (x *LightClientOptimisticUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eth_v2_beacon_lightclient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LightClientOptimisticUpdate.ProtoReflect.Descriptor instead.
func (*LightClientOptimisticUpdate) Descriptor() ([]byte, []int) {
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP(), []int{5}
}

func (x *LightClientOptimisticUpdate) GetAttestedHeader() *v1.BeaconBlockHeader {
	if x != nil {
		return x.AttestedHeader
	}
	return nil
}

func (x *LightClientOptimisticUpdate) GetSyncAggregate() *v1.SyncAggregate {
	if x != nil {
		return x.SyncAggregate
	}
	return nil
}

func (x *LightClientOptimisticUpdate) GetSignatureSlot() github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot {
	if x != nil {
		return x.SignatureSlot
	}
	return github_com_prysmaticlabs_prysm_v4_consensus_types_primitives.Slot(0)
}

var File_proto_eth_v2_beacon_lightclient_proto protoreflect.FileDescriptor

var file_proto_eth_v2_beacon_lightclient_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74,
	0x68, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x3a,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x16, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x52, 0x14, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65,
	0x12, 0x41, 0x0a, 0x1d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x1a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x22, 0x9a, 0x04, 0x0a, 0x11, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x52, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x17, 0x6e, 0x65, 0x78, 0x74,
	0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x12, 0x4d, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x45, 0x0a, 0x0e, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x52, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x45, 0x82, 0xb5, 0x18, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76,
	0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f,
	0x74, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x6c, 0x6f, 0x74,
	0x22, 0x9a, 0x01, 0x0a, 0x24, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x03,
	0x0a, 0x19, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x12, 0x45, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x45, 0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x26, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74,
	0x68, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x02, 0x0a, 0x1b, 0x4c, 0x69, 0x67, 0x68, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79,
	0x6e, 0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x73, 0x79, 0x6e,
	0x63, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x45, 0x82, 0xb5, 0x18, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x42, 0x83, 0x01, 0x0a, 0x13, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32,
	0x42, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x76, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0f, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eth_v2_beacon_lightclient_proto_rawDescOnce sync.Once
	file_proto_eth_v2_beacon_lightclient_proto_rawDescData = file_proto_eth_v2_beacon_lightclient_proto_rawDesc
)

func file_proto_eth_v2_beacon_lightclient_proto_rawDescGZIP() []byte {
	file_proto_eth_v2_beacon_lightclient_proto_rawDescOnce.Do(func() {
		file_proto_eth_v2_beacon_lightclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eth_v2_beacon_lightclient_proto_rawDescData)
	})
	return file_proto_eth_v2_beacon_lightclient_proto_rawDescData
}

var file_proto_eth_v2_beacon_lightclient_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_eth_v2_beacon_lightclient_proto_goTypes = []interface{}{
	(*LightClientBootstrap)(nil),                   // 0: ethereum.eth.v2.LightClientBootstrap
	(*LightClientUpdate)(nil),                      // 1: ethereum.eth.v2.LightClientUpdate
	(*LightClientFinalityUpdateWithVersion)(nil),   // 2: ethereum.eth.v2.LightClientFinalityUpdateWithVersion
	(*LightClientFinalityUpdate)(nil),              // 3: ethereum.eth.v2.LightClientFinalityUpdate
	(*LightClientOptimisticUpdateWithVersion)(nil), // 4: ethereum.eth.v2.LightClientOptimisticUpdateWithVersion
	(*LightClientOptimisticUpdate)(nil),            // 5: ethereum.eth.v2.LightClientOptimisticUpdate
	(*v1.BeaconBlockHeader)(nil),                   // 6: ethereum.eth.v1.BeaconBlockHeader
	(*SyncCommittee)(nil),                          // 7: ethereum.eth.v2.SyncCommittee
	(*v1.SyncAggregate)(nil),                       // 8: ethereum.eth.v1.SyncAggregate
	(Version)(0),                                   // 9: ethereum.eth.v2.Version
}
var file_proto_eth_v2_beacon_lightclient_proto_depIdxs = []int32{
	6,  // 0: ethereum.eth.v2.LightClientBootstrap.header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	7,  // 1: ethereum.eth.v2.LightClientBootstrap.current_sync_committee:type_name -> ethereum.eth.v2.SyncCommittee
	6,  // 2: ethereum.eth.v2.LightClientUpdate.attested_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	7,  // 3: ethereum.eth.v2.LightClientUpdate.next_sync_committee:type_name -> ethereum.eth.v2.SyncCommittee
	6,  // 4: ethereum.eth.v2.LightClientUpdate.finalized_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	8,  // 5: ethereum.eth.v2.LightClientUpdate.sync_aggregate:type_name -> ethereum.eth.v1.SyncAggregate
	9,  // 6: ethereum.eth.v2.LightClientFinalityUpdateWithVersion.version:type_name -> ethereum.eth.v2.Version
	3,  // 7: ethereum.eth.v2.LightClientFinalityUpdateWithVersion.data:type_name -> ethereum.eth.v2.LightClientFinalityUpdate
	6,  // 8: ethereum.eth.v2.LightClientFinalityUpdate.attested_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	6,  // 9: ethereum.eth.v2.LightClientFinalityUpdate.finalized_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	8,  // 10: ethereum.eth.v2.LightClientFinalityUpdate.sync_aggregate:type_name -> ethereum.eth.v1.SyncAggregate
	9,  // 11: ethereum.eth.v2.LightClientOptimisticUpdateWithVersion.version:type_name -> ethereum.eth.v2.Version
	5,  // 12: ethereum.eth.v2.LightClientOptimisticUpdateWithVersion.data:type_name -> ethereum.eth.v2.LightClientOptimisticUpdate
	6,  // 13: ethereum.eth.v2.LightClientOptimisticUpdate.attested_header:type_name -> ethereum.eth.v1.BeaconBlockHeader
	8,  // 14: ethereum.eth.v2.LightClientOptimisticUpdate.sync_aggregate:type_name -> ethereum.eth.v1.SyncAggregate
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_proto_eth_v2_beacon_lightclient_proto_init() }
func file_proto_eth_v2_beacon_lightclient_proto_init() {
	if File_proto_eth_v2_beacon_lightclient_proto != nil {
		return
	}
	file_proto_eth_v2_version_proto_init()
	file_proto_eth_v2_sync_committee_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientBootstrap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientFinalityUpdateWithVersion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientFinalityUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientOptimisticUpdateWithVersion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_eth_v2_beacon_lightclient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LightClientOptimisticUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_eth_v2_beacon_lightclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eth_v2_beacon_lightclient_proto_goTypes,
		DependencyIndexes: file_proto_eth_v2_beacon_lightclient_proto_depIdxs,
		MessageInfos:      file_proto_eth_v2_beacon_lightclient_proto_msgTypes,
	}.Build()
	File_proto_eth_v2_beacon_lightclient_proto = out.File
	file_proto_eth_v2_beacon_lightclient_proto_rawDesc = nil
	file_proto_eth_v2_beacon_lightclient_proto_goTypes = nil
	file_proto_eth_v2_beacon_lightclient_proto_depIdxs = nil
}
