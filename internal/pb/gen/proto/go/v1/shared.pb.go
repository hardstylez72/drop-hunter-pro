// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: v1/shared.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskKind int32

const (
	TaskKind_TKSimple TaskKind = 0
	TaskKind_TKSwap   TaskKind = 1
	TaskKind_TKBridge TaskKind = 2
)

// Enum value maps for TaskKind.
var (
	TaskKind_name = map[int32]string{
		0: "TKSimple",
		1: "TKSwap",
		2: "TKBridge",
	}
	TaskKind_value = map[string]int32{
		"TKSimple": 0,
		"TKSwap":   1,
		"TKBridge": 2,
	}
)

func (x TaskKind) Enum() *TaskKind {
	p := new(TaskKind)
	*p = x
	return p
}

func (x TaskKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskKind) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[0].Descriptor()
}

func (TaskKind) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[0]
}

func (x TaskKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskKind.Descriptor instead.
func (TaskKind) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{0}
}

type Network int32

const (
	Network_ARBITRUM         Network = 0
	Network_OPTIMISM         Network = 1
	Network_BinanaceBNB      Network = 2
	Network_Etherium         Network = 3
	Network_POLIGON          Network = 4
	Network_AVALANCHE        Network = 5
	Network_GOERLIETH        Network = 6
	Network_ZKSYNCERA        Network = 7
	Network_ZKSYNCERATESTNET Network = 8 //deprecated
	Network_ZKSYNCLITE       Network = 9
	Network_StarkNet         Network = 10
	Network_Meter            Network = 11
	Network_Tenet            Network = 12
	Network_Canto            Network = 13
	Network_ArbitrumNova     Network = 14
	Network_PolygonZKEVM     Network = 15
	Network_Fantom           Network = 16
	Network_Base             Network = 17
	Network_opBNB            Network = 18
	Network_Linea            Network = 19
	Network_Zora             Network = 20
	Network_Core             Network = 21
)

// Enum value maps for Network.
var (
	Network_name = map[int32]string{
		0:  "ARBITRUM",
		1:  "OPTIMISM",
		2:  "BinanaceBNB",
		3:  "Etherium",
		4:  "POLIGON",
		5:  "AVALANCHE",
		6:  "GOERLIETH",
		7:  "ZKSYNCERA",
		8:  "ZKSYNCERATESTNET",
		9:  "ZKSYNCLITE",
		10: "StarkNet",
		11: "Meter",
		12: "Tenet",
		13: "Canto",
		14: "ArbitrumNova",
		15: "PolygonZKEVM",
		16: "Fantom",
		17: "Base",
		18: "opBNB",
		19: "Linea",
		20: "Zora",
		21: "Core",
	}
	Network_value = map[string]int32{
		"ARBITRUM":         0,
		"OPTIMISM":         1,
		"BinanaceBNB":      2,
		"Etherium":         3,
		"POLIGON":          4,
		"AVALANCHE":        5,
		"GOERLIETH":        6,
		"ZKSYNCERA":        7,
		"ZKSYNCERATESTNET": 8,
		"ZKSYNCLITE":       9,
		"StarkNet":         10,
		"Meter":            11,
		"Tenet":            12,
		"Canto":            13,
		"ArbitrumNova":     14,
		"PolygonZKEVM":     15,
		"Fantom":           16,
		"Base":             17,
		"opBNB":            18,
		"Linea":            19,
		"Zora":             20,
		"Core":             21,
	}
)

func (x Network) Enum() *Network {
	p := new(Network)
	*p = x
	return p
}

func (x Network) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Network) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[1].Descriptor()
}

func (Network) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[1]
}

func (x Network) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Network.Descriptor instead.
func (Network) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{1}
}

type Token int32

const (
	Token_USDT        Token = 0
	Token_ETH         Token = 1
	Token_USDC        Token = 2
	Token_STG         Token = 3
	Token_BNB         Token = 4
	Token_MATIC       Token = 5
	Token_AVAX        Token = 6
	Token_veSTG       Token = 7
	Token_WETH        Token = 8
	Token_LUSD        Token = 9
	Token_LSD         Token = 10
	Token_MUTE        Token = 11
	Token_MAV         Token = 12
	Token_SPACE       Token = 13
	Token_VC          Token = 14
	Token_IZI         Token = 15
	Token_USDCBridged Token = 16
	Token_BUSD        Token = 17
	Token_USDp        Token = 18
	Token_CORE        Token = 19
)

// Enum value maps for Token.
var (
	Token_name = map[int32]string{
		0:  "USDT",
		1:  "ETH",
		2:  "USDC",
		3:  "STG",
		4:  "BNB",
		5:  "MATIC",
		6:  "AVAX",
		7:  "veSTG",
		8:  "WETH",
		9:  "LUSD",
		10: "LSD",
		11: "MUTE",
		12: "MAV",
		13: "SPACE",
		14: "VC",
		15: "IZI",
		16: "USDCBridged",
		17: "BUSD",
		18: "USDp",
		19: "CORE",
	}
	Token_value = map[string]int32{
		"USDT":        0,
		"ETH":         1,
		"USDC":        2,
		"STG":         3,
		"BNB":         4,
		"MATIC":       5,
		"AVAX":        6,
		"veSTG":       7,
		"WETH":        8,
		"LUSD":        9,
		"LSD":         10,
		"MUTE":        11,
		"MAV":         12,
		"SPACE":       13,
		"VC":          14,
		"IZI":         15,
		"USDCBridged": 16,
		"BUSD":        17,
		"USDp":        18,
		"CORE":        19,
	}
)

func (x Token) Enum() *Token {
	p := new(Token)
	*p = x
	return p
}

func (x Token) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Token) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[2].Descriptor()
}

func (Token) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[2]
}

func (x Token) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Token.Descriptor instead.
func (Token) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{2}
}

type ProcessStatus int32

const (
	ProcessStatus_StatusReady   ProcessStatus = 0
	ProcessStatus_StatusRunning ProcessStatus = 1
	ProcessStatus_StatusError   ProcessStatus = 2
	ProcessStatus_StatusDone    ProcessStatus = 3
	ProcessStatus_StatusStop    ProcessStatus = 4 // delete
	ProcessStatus_StatusRetry   ProcessStatus = 5
)

// Enum value maps for ProcessStatus.
var (
	ProcessStatus_name = map[int32]string{
		0: "StatusReady",
		1: "StatusRunning",
		2: "StatusError",
		3: "StatusDone",
		4: "StatusStop",
		5: "StatusRetry",
	}
	ProcessStatus_value = map[string]int32{
		"StatusReady":   0,
		"StatusRunning": 1,
		"StatusError":   2,
		"StatusDone":    3,
		"StatusStop":    4,
		"StatusRetry":   5,
	}
)

func (x ProcessStatus) Enum() *ProcessStatus {
	p := new(ProcessStatus)
	*p = x
	return p
}

func (x ProcessStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[3].Descriptor()
}

func (ProcessStatus) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[3]
}

func (x ProcessStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessStatus.Descriptor instead.
func (ProcessStatus) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{3}
}

type ProfileAccountType int32

const (
	ProfileAccountType_AccountDiscord ProfileAccountType = 0
	ProfileAccountType_AccountOkex    ProfileAccountType = 1
	ProfileAccountType_AccountEmail   ProfileAccountType = 2
)

// Enum value maps for ProfileAccountType.
var (
	ProfileAccountType_name = map[int32]string{
		0: "AccountDiscord",
		1: "AccountOkex",
		2: "AccountEmail",
	}
	ProfileAccountType_value = map[string]int32{
		"AccountDiscord": 0,
		"AccountOkex":    1,
		"AccountEmail":   2,
	}
)

func (x ProfileAccountType) Enum() *ProfileAccountType {
	p := new(ProfileAccountType)
	*p = x
	return p
}

func (x ProfileAccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileAccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[4].Descriptor()
}

func (ProfileAccountType) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[4]
}

func (x ProfileAccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileAccountType.Descriptor instead.
func (ProfileAccountType) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{4}
}

type AmountType int32

const (
	AmountType_All     AmountType = 0
	AmountType_Percent AmountType = 1
	AmountType_Value   AmountType = 2
)

// Enum value maps for AmountType.
var (
	AmountType_name = map[int32]string{
		0: "All",
		1: "Percent",
		2: "Value",
	}
	AmountType_value = map[string]int32{
		"All":     0,
		"Percent": 1,
		"Value":   2,
	}
)

func (x AmountType) Enum() *AmountType {
	p := new(AmountType)
	*p = x
	return p
}

func (x AmountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AmountType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_shared_proto_enumTypes[5].Descriptor()
}

func (AmountType) Type() protoreflect.EnumType {
	return &file_v1_shared_proto_enumTypes[5]
}

func (x AmountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AmountType.Descriptor instead.
func (AmountType) EnumDescriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{5}
}

type AmUni struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gwei    string  `protobuf:"bytes,1,opt,name=gwei,proto3" json:"gwei,omitempty"`
	Eth     string  `protobuf:"bytes,2,opt,name=eth,proto3" json:"eth,omitempty"`
	Usd     string  `protobuf:"bytes,3,opt,name=usd,proto3" json:"usd,omitempty"`
	Network Network `protobuf:"varint,4,opt,name=network,proto3,enum=shared.Network" json:"network,omitempty"`
	Wei     string  `protobuf:"bytes,5,opt,name=wei,proto3" json:"wei,omitempty"`
}

func (x *AmUni) Reset() {
	*x = AmUni{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmUni) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmUni) ProtoMessage() {}

func (x *AmUni) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmUni.ProtoReflect.Descriptor instead.
func (*AmUni) Descriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{0}
}

func (x *AmUni) GetGwei() string {
	if x != nil {
		return x.Gwei
	}
	return ""
}

func (x *AmUni) GetEth() string {
	if x != nil {
		return x.Eth
	}
	return ""
}

func (x *AmUni) GetUsd() string {
	if x != nil {
		return x.Usd
	}
	return ""
}

func (x *AmUni) GetNetwork() Network {
	if x != nil {
		return x.Network
	}
	return Network_ARBITRUM
}

func (x *AmUni) GetWei() string {
	if x != nil {
		return x.Wei
	}
	return ""
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Amount_SendAll
	//	*Amount_SendPercent
	//	*Amount_SendAmount
	//	*Amount_SendValue
	//	*Amount_PercRange
	Kind         isAmount_Kind `protobuf_oneof:"kind"`
	Send         *AmUni        `protobuf:"bytes,7,opt,name=send,proto3,oneof" json:"send,omitempty"`
	Balance      *AmUni        `protobuf:"bytes,8,opt,name=balance,proto3,oneof" json:"balance,omitempty"`
	GasEstimated *AmUni        `protobuf:"bytes,9,opt,name=gas_estimated,json=gasEstimated,proto3,oneof" json:"gas_estimated,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{1}
}

func (m *Amount) GetKind() isAmount_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Amount) GetSendAll() bool {
	if x, ok := x.GetKind().(*Amount_SendAll); ok {
		return x.SendAll
	}
	return false
}

func (x *Amount) GetSendPercent() float32 {
	if x, ok := x.GetKind().(*Amount_SendPercent); ok {
		return x.SendPercent
	}
	return 0
}

func (x *Amount) GetSendAmount() float32 {
	if x, ok := x.GetKind().(*Amount_SendAmount); ok {
		return x.SendAmount
	}
	return 0
}

func (x *Amount) GetSendValue() string {
	if x, ok := x.GetKind().(*Amount_SendValue); ok {
		return x.SendValue
	}
	return ""
}

func (x *Amount) GetPercRange() *PercentRange {
	if x, ok := x.GetKind().(*Amount_PercRange); ok {
		return x.PercRange
	}
	return nil
}

func (x *Amount) GetSend() *AmUni {
	if x != nil {
		return x.Send
	}
	return nil
}

func (x *Amount) GetBalance() *AmUni {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *Amount) GetGasEstimated() *AmUni {
	if x != nil {
		return x.GasEstimated
	}
	return nil
}

type isAmount_Kind interface {
	isAmount_Kind()
}

type Amount_SendAll struct {
	SendAll bool `protobuf:"varint,4,opt,name=send_all,json=sendAll,proto3,oneof"`
}

type Amount_SendPercent struct {
	SendPercent float32 `protobuf:"fixed32,5,opt,name=send_percent,json=sendPercent,proto3,oneof"`
}

type Amount_SendAmount struct {
	SendAmount float32 `protobuf:"fixed32,6,opt,name=send_amount,json=sendAmount,proto3,oneof"` //deprecated
}

type Amount_SendValue struct {
	SendValue string `protobuf:"bytes,10,opt,name=send_value,json=sendValue,proto3,oneof"`
}

type Amount_PercRange struct {
	PercRange *PercentRange `protobuf:"bytes,11,opt,name=perc_range,json=percRange,proto3,oneof"`
}

func (*Amount_SendAll) isAmount_Kind() {}

func (*Amount_SendPercent) isAmount_Kind() {}

func (*Amount_SendAmount) isAmount_Kind() {}

func (*Amount_SendValue) isAmount_Kind() {}

func (*Amount_PercRange) isAmount_Kind() {}

type PercentRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min int64 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max int64 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *PercentRange) Reset() {
	*x = PercentRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PercentRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PercentRange) ProtoMessage() {}

func (x *PercentRange) ProtoReflect() protoreflect.Message {
	mi := &file_v1_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PercentRange.ProtoReflect.Descriptor instead.
func (*PercentRange) Descriptor() ([]byte, []int) {
	return file_v1_shared_proto_rawDescGZIP(), []int{2}
}

func (x *PercentRange) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PercentRange) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_v1_shared_proto protoreflect.FileDescriptor

var file_v1_shared_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x05, 0x41, 0x6d, 0x55, 0x6e,
	0x69, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x77, 0x65, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x67, 0x77, 0x65, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x73, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x73, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x65, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x77, 0x65, 0x69, 0x22, 0x91, 0x03, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x23,
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x09, 0x70, 0x65, 0x72, 0x63, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x41, 0x6d, 0x55, 0x6e, 0x69, 0x48, 0x01, 0x52, 0x04, 0x73,
	0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x41, 0x6d, 0x55, 0x6e, 0x69, 0x48, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x61, 0x73, 0x5f, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x41, 0x6d, 0x55, 0x6e, 0x69, 0x48, 0x03, 0x52, 0x0c, 0x67, 0x61,
	0x73, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x0c, 0x92,
	0x41, 0x09, 0x0a, 0x07, 0xd2, 0x01, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x67, 0x61, 0x73,
	0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x22, 0x32, 0x0a, 0x0c, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x2a, 0x32,
	0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4b,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x4b, 0x53, 0x77,
	0x61, 0x70, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4b, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x10, 0x02, 0x2a, 0xb7, 0x02, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x52, 0x42, 0x49, 0x54, 0x52, 0x55, 0x4d, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4f, 0x50, 0x54, 0x49, 0x4d, 0x49, 0x53, 0x4d, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x69,
	0x6e, 0x61, 0x6e, 0x61, 0x63, 0x65, 0x42, 0x4e, 0x42, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x69, 0x75, 0x6d, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x4f, 0x4c,
	0x49, 0x47, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x4c, 0x41, 0x4e,
	0x43, 0x48, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x4f, 0x45, 0x52, 0x4c, 0x49, 0x45,
	0x54, 0x48, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x5a, 0x4b, 0x53, 0x59, 0x4e, 0x43, 0x45, 0x52,
	0x41, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x5a, 0x4b, 0x53, 0x59, 0x4e, 0x43, 0x45, 0x52, 0x41,
	0x54, 0x45, 0x53, 0x54, 0x4e, 0x45, 0x54, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x5a, 0x4b, 0x53,
	0x59, 0x4e, 0x43, 0x4c, 0x49, 0x54, 0x45, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x72, 0x6b, 0x4e, 0x65, 0x74, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x65, 0x74, 0x65, 0x72,
	0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x65, 0x6e, 0x65, 0x74, 0x10, 0x0c, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x61, 0x6e, 0x74, 0x6f, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x72, 0x62, 0x69,
	0x74, 0x72, 0x75, 0x6d, 0x4e, 0x6f, 0x76, 0x61, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x6f,
	0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x5a, 0x4b, 0x45, 0x56, 0x4d, 0x10, 0x0f, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x70, 0x42, 0x4e, 0x42, 0x10, 0x12, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x10, 0x13, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x6f, 0x72, 0x61,
	0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x65, 0x10, 0x15, 0x2a, 0xd1, 0x01, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x44, 0x54, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x44,
	0x43, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x54, 0x47, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x4e, 0x42, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10, 0x05,
	0x12, 0x08, 0x0a, 0x04, 0x41, 0x56, 0x41, 0x58, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x76, 0x65,
	0x53, 0x54, 0x47, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x54, 0x48, 0x10, 0x08, 0x12,
	0x08, 0x0a, 0x04, 0x4c, 0x55, 0x53, 0x44, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x53, 0x44,
	0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x55, 0x54, 0x45, 0x10, 0x0b, 0x12, 0x07, 0x0a, 0x03,
	0x4d, 0x41, 0x56, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x0d,
	0x12, 0x06, 0x0a, 0x02, 0x56, 0x43, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x5a, 0x49, 0x10,
	0x0f, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x44, 0x43, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x64,
	0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x44, 0x10, 0x11, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x53, 0x44, 0x70, 0x10, 0x12, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x13,
	0x2a, 0x75, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x44, 0x6f, 0x6e, 0x65, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x74, 0x6f, 0x70, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x10, 0x05, 0x2a, 0x4b, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x6b, 0x65, 0x78,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x10, 0x02, 0x2a, 0x2d, 0x0a, 0x0a, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_shared_proto_rawDescOnce sync.Once
	file_v1_shared_proto_rawDescData = file_v1_shared_proto_rawDesc
)

func file_v1_shared_proto_rawDescGZIP() []byte {
	file_v1_shared_proto_rawDescOnce.Do(func() {
		file_v1_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_shared_proto_rawDescData)
	})
	return file_v1_shared_proto_rawDescData
}

var file_v1_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_v1_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_shared_proto_goTypes = []interface{}{
	(TaskKind)(0),           // 0: shared.TaskKind
	(Network)(0),            // 1: shared.Network
	(Token)(0),              // 2: shared.Token
	(ProcessStatus)(0),      // 3: shared.ProcessStatus
	(ProfileAccountType)(0), // 4: shared.ProfileAccountType
	(AmountType)(0),         // 5: shared.AmountType
	(*AmUni)(nil),           // 6: shared.AmUni
	(*Amount)(nil),          // 7: shared.Amount
	(*PercentRange)(nil),    // 8: shared.PercentRange
}
var file_v1_shared_proto_depIdxs = []int32{
	1, // 0: shared.AmUni.network:type_name -> shared.Network
	8, // 1: shared.Amount.perc_range:type_name -> shared.PercentRange
	6, // 2: shared.Amount.send:type_name -> shared.AmUni
	6, // 3: shared.Amount.balance:type_name -> shared.AmUni
	6, // 4: shared.Amount.gas_estimated:type_name -> shared.AmUni
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_shared_proto_init() }
func file_v1_shared_proto_init() {
	if File_v1_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmUni); i {
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
		file_v1_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_v1_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PercentRange); i {
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
	file_v1_shared_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Amount_SendAll)(nil),
		(*Amount_SendPercent)(nil),
		(*Amount_SendAmount)(nil),
		(*Amount_SendValue)(nil),
		(*Amount_PercRange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_shared_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_shared_proto_goTypes,
		DependencyIndexes: file_v1_shared_proto_depIdxs,
		EnumInfos:         file_v1_shared_proto_enumTypes,
		MessageInfos:      file_v1_shared_proto_msgTypes,
	}.Build()
	File_v1_shared_proto = out.File
	file_v1_shared_proto_rawDesc = nil
	file_v1_shared_proto_goTypes = nil
	file_v1_shared_proto_depIdxs = nil
}
