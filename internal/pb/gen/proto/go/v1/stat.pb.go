// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/stat.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveDays     *int64                 `protobuf:"varint,1,opt,name=active_days,json=activeDays,proto3,oneof" json:"active_days,omitempty"`
	ActiveMonths   *int64                 `protobuf:"varint,2,opt,name=active_months,json=activeMonths,proto3,oneof" json:"active_months,omitempty"`
	LastActivity   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_activity,json=lastActivity,proto3,oneof" json:"last_activity,omitempty"`
	Interactions   []*Interaction         `protobuf:"bytes,4,rep,name=interactions,proto3" json:"interactions,omitempty"`
	TotalUsdAmount float64                `protobuf:"fixed64,5,opt,name=total_usd_amount,json=totalUsdAmount,proto3" json:"total_usd_amount,omitempty"`
	TxCount        int64                  `protobuf:"varint,6,opt,name=tx_count,json=txCount,proto3" json:"tx_count,omitempty"`
	UniqAddress    int64                  `protobuf:"varint,7,opt,name=uniq_address,json=uniqAddress,proto3" json:"uniq_address,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_v1_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_v1_stat_proto_rawDescGZIP(), []int{0}
}

func (x *Stat) GetActiveDays() int64 {
	if x != nil && x.ActiveDays != nil {
		return *x.ActiveDays
	}
	return 0
}

func (x *Stat) GetActiveMonths() int64 {
	if x != nil && x.ActiveMonths != nil {
		return *x.ActiveMonths
	}
	return 0
}

func (x *Stat) GetLastActivity() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActivity
	}
	return nil
}

func (x *Stat) GetInteractions() []*Interaction {
	if x != nil {
		return x.Interactions
	}
	return nil
}

func (x *Stat) GetTotalUsdAmount() float64 {
	if x != nil {
		return x.TotalUsdAmount
	}
	return 0
}

func (x *Stat) GetTxCount() int64 {
	if x != nil {
		return x.TxCount
	}
	return 0
}

func (x *Stat) GetUniqAddress() int64 {
	if x != nil {
		return x.UniqAddress
	}
	return 0
}

type Interaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To          string               `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	ToUrl       string               `protobuf:"bytes,3,opt,name=to_url,json=toUrl,proto3" json:"to_url,omitempty"`
	ServiceName string               `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceUrl  *string              `protobuf:"bytes,5,opt,name=service_url,json=serviceUrl,proto3,oneof" json:"service_url,omitempty"`
	Amounts     []*InteractionAmount `protobuf:"bytes,6,rep,name=amounts,proto3" json:"amounts,omitempty"`
	Txs         int64                `protobuf:"varint,7,opt,name=txs,proto3" json:"txs,omitempty"`
	AmountUsd   float64              `protobuf:"fixed64,8,opt,name=amount_usd,json=amountUsd,proto3" json:"amount_usd,omitempty"`
}

func (x *Interaction) Reset() {
	*x = Interaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interaction) ProtoMessage() {}

func (x *Interaction) ProtoReflect() protoreflect.Message {
	mi := &file_v1_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interaction.ProtoReflect.Descriptor instead.
func (*Interaction) Descriptor() ([]byte, []int) {
	return file_v1_stat_proto_rawDescGZIP(), []int{1}
}

func (x *Interaction) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Interaction) GetToUrl() string {
	if x != nil {
		return x.ToUrl
	}
	return ""
}

func (x *Interaction) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Interaction) GetServiceUrl() string {
	if x != nil && x.ServiceUrl != nil {
		return *x.ServiceUrl
	}
	return ""
}

func (x *Interaction) GetAmounts() []*InteractionAmount {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *Interaction) GetTxs() int64 {
	if x != nil {
		return x.Txs
	}
	return 0
}

func (x *Interaction) GetAmountUsd() float64 {
	if x != nil {
		return x.AmountUsd
	}
	return 0
}

type InteractionAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     *Token  `protobuf:"varint,4,opt,name=token,proto3,enum=shared.Token,oneof" json:"token,omitempty"`
	TokenUrl  string  `protobuf:"bytes,1,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	AmountWei string  `protobuf:"bytes,2,opt,name=amount_wei,json=amountWei,proto3" json:"amount_wei,omitempty"`
	AmountUsd *string `protobuf:"bytes,3,opt,name=amount_usd,json=amountUsd,proto3,oneof" json:"amount_usd,omitempty"`
}

func (x *InteractionAmount) Reset() {
	*x = InteractionAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractionAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractionAmount) ProtoMessage() {}

func (x *InteractionAmount) ProtoReflect() protoreflect.Message {
	mi := &file_v1_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractionAmount.ProtoReflect.Descriptor instead.
func (*InteractionAmount) Descriptor() ([]byte, []int) {
	return file_v1_stat_proto_rawDescGZIP(), []int{2}
}

func (x *InteractionAmount) GetToken() Token {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return Token_USDT
}

func (x *InteractionAmount) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *InteractionAmount) GetAmountWei() string {
	if x != nil {
		return x.AmountWei
	}
	return ""
}

func (x *InteractionAmount) GetAmountUsd() string {
	if x != nil && x.AmountUsd != nil {
		return *x.AmountUsd
	}
	return ""
}

type ZkSyncStatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileId string `protobuf:"bytes,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
}

func (x *ZkSyncStatReq) Reset() {
	*x = ZkSyncStatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_stat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZkSyncStatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZkSyncStatReq) ProtoMessage() {}

func (x *ZkSyncStatReq) ProtoReflect() protoreflect.Message {
	mi := &file_v1_stat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZkSyncStatReq.ProtoReflect.Descriptor instead.
func (*ZkSyncStatReq) Descriptor() ([]byte, []int) {
	return file_v1_stat_proto_rawDescGZIP(), []int{3}
}

func (x *ZkSyncStatReq) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

type ZkSyncStatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat *Stat `protobuf:"bytes,1,opt,name=stat,proto3" json:"stat,omitempty"`
}

func (x *ZkSyncStatRes) Reset() {
	*x = ZkSyncStatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_stat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZkSyncStatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZkSyncStatRes) ProtoMessage() {}

func (x *ZkSyncStatRes) ProtoReflect() protoreflect.Message {
	mi := &file_v1_stat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZkSyncStatRes.ProtoReflect.Descriptor instead.
func (*ZkSyncStatRes) Descriptor() ([]byte, []int) {
	return file_v1_stat_proto_rawDescGZIP(), []int{4}
}

func (x *ZkSyncStatRes) GetStat() *Stat {
	if x != nil {
		return x.Stat
	}
	return nil
}

var File_v1_stat_proto protoreflect.FileDescriptor

var file_v1_stat_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x24, 0x0a,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x61, 0x79, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x73, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x02, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x71, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x71, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61,
	0x79, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0xae, 0x02, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x78, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x64, 0x3a, 0x3b, 0x92, 0x41, 0x38, 0x0a,
	0x36, 0xd2, 0x01, 0x02, 0x74, 0x6f, 0xd2, 0x01, 0x06, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0xd2,
	0x01, 0x08, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xd2, 0x01, 0x07, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0xd2, 0x01, 0x03, 0x74, 0x78, 0x73, 0xd2, 0x01, 0x0a, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0xe1, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x77,
	0x65, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x65, 0x69, 0x12, 0x22, 0x0a, 0x0a, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x73, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x29, 0x92, 0x41, 0x26, 0x0a, 0x24, 0xd2, 0x01,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0xd2, 0x01, 0x0a, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x77, 0x65, 0x69, 0xd2, 0x01, 0x08, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x64, 0x22, 0x42, 0x0a, 0x0d, 0x5a,
	0x6b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x3a, 0x12, 0x92, 0x41, 0x0f,
	0x0a, 0x0d, 0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22,
	0x3d, 0x0a, 0x0d, 0x5a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x1e, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74,
	0x3a, 0x0c, 0x92, 0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x73, 0x74, 0x61, 0x74, 0x32, 0x70,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a,
	0x0a, 0x5a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x12, 0x13, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x2e, 0x5a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x5a, 0x6b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a,
	0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7a, 0x6b, 0x73, 0x79, 0x6e, 0x63,
	0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_stat_proto_rawDescOnce sync.Once
	file_v1_stat_proto_rawDescData = file_v1_stat_proto_rawDesc
)

func file_v1_stat_proto_rawDescGZIP() []byte {
	file_v1_stat_proto_rawDescOnce.Do(func() {
		file_v1_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_stat_proto_rawDescData)
	})
	return file_v1_stat_proto_rawDescData
}

var file_v1_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_stat_proto_goTypes = []interface{}{
	(*Stat)(nil),                  // 0: stat.Stat
	(*Interaction)(nil),           // 1: stat.Interaction
	(*InteractionAmount)(nil),     // 2: stat.InteractionAmount
	(*ZkSyncStatReq)(nil),         // 3: stat.ZkSyncStatReq
	(*ZkSyncStatRes)(nil),         // 4: stat.ZkSyncStatRes
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(Token)(0),                    // 6: shared.Token
}
var file_v1_stat_proto_depIdxs = []int32{
	5, // 0: stat.Stat.last_activity:type_name -> google.protobuf.Timestamp
	1, // 1: stat.Stat.interactions:type_name -> stat.Interaction
	2, // 2: stat.Interaction.amounts:type_name -> stat.InteractionAmount
	6, // 3: stat.InteractionAmount.token:type_name -> shared.Token
	0, // 4: stat.ZkSyncStatRes.stat:type_name -> stat.Stat
	3, // 5: stat.StatService.ZkSyncStat:input_type -> stat.ZkSyncStatReq
	4, // 6: stat.StatService.ZkSyncStat:output_type -> stat.ZkSyncStatRes
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_stat_proto_init() }
func file_v1_stat_proto_init() {
	if File_v1_stat_proto != nil {
		return
	}
	file_v1_shared_proto_init()
	file_v1_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
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
		file_v1_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interaction); i {
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
		file_v1_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractionAmount); i {
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
		file_v1_stat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZkSyncStatReq); i {
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
		file_v1_stat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZkSyncStatRes); i {
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
	file_v1_stat_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_stat_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_v1_stat_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_stat_proto_goTypes,
		DependencyIndexes: file_v1_stat_proto_depIdxs,
		MessageInfos:      file_v1_stat_proto_msgTypes,
	}.Build()
	File_v1_stat_proto = out.File
	file_v1_stat_proto_rawDesc = nil
	file_v1_stat_proto_goTypes = nil
	file_v1_stat_proto_depIdxs = nil
}
