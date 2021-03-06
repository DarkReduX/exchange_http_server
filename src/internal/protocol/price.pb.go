// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: src/internal/protocol/price.proto

package protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DonateValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val   float32 `protobuf:"fixed32,1,opt,name=val,proto3" json:"val,omitempty"`
	Token string  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DonateValue) Reset() {
	*x = DonateValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DonateValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DonateValue) ProtoMessage() {}

func (x *DonateValue) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DonateValue.ProtoReflect.Descriptor instead.
func (*DonateValue) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{0}
}

func (x *DonateValue) GetVal() float32 {
	if x != nil {
		return x.Val
	}
	return 0
}

func (x *DonateValue) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance float32 `protobuf:"fixed32,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{1}
}

func (x *Balance) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []byte  `protobuf:"bytes,1,opt,name=positions,proto3" json:"positions,omitempty"`
	Balance   float32 `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetPositions() []byte {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *UserData) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   int64   `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Symbol string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Bid    float32 `protobuf:"fixed32,3,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask    float32 `protobuf:"fixed32,4,opt,name=ask,proto3" json:"ask,omitempty"`
	IsBay  bool    `protobuf:"varint,5,opt,name=isBay,proto3" json:"isBay,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{6}
}

func (x *Position) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *Position) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Position) GetBid() float32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *Position) GetAsk() float32 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *Position) GetIsBay() bool {
	if x != nil {
		return x.IsBay
	}
	return false
}

type PositionCloseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol     string  `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Id         int64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Token      string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Isbay      bool    `protobuf:"varint,4,opt,name=isbay,proto3" json:"isbay,omitempty"`
	ClosePrice float32 `protobuf:"fixed32,5,opt,name=closePrice,proto3" json:"closePrice,omitempty"`
}

func (x *PositionCloseReq) Reset() {
	*x = PositionCloseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionCloseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionCloseReq) ProtoMessage() {}

func (x *PositionCloseReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionCloseReq.ProtoReflect.Descriptor instead.
func (*PositionCloseReq) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{7}
}

func (x *PositionCloseReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PositionCloseReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PositionCloseReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PositionCloseReq) GetIsbay() bool {
	if x != nil {
		return x.Isbay
	}
	return false
}

func (x *PositionCloseReq) GetClosePrice() float32 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

type PositionOpenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   int64   `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Symbol string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Bid    float32 `protobuf:"fixed32,3,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask    float32 `protobuf:"fixed32,4,opt,name=ask,proto3" json:"ask,omitempty"`
	IsBay  bool    `protobuf:"varint,5,opt,name=isBay,proto3" json:"isBay,omitempty"`
	Token  string  `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PositionOpenReq) Reset() {
	*x = PositionOpenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionOpenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionOpenReq) ProtoMessage() {}

func (x *PositionOpenReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionOpenReq.ProtoReflect.Descriptor instead.
func (*PositionOpenReq) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{8}
}

func (x *PositionOpenReq) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *PositionOpenReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PositionOpenReq) GetBid() float32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *PositionOpenReq) GetAsk() float32 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *PositionOpenReq) GetIsBay() bool {
	if x != nil {
		return x.IsBay
	}
	return false
}

func (x *PositionOpenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{9}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   int64   `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Symbol string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Bid    float32 `protobuf:"fixed32,3,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask    float32 `protobuf:"fixed32,4,opt,name=ask,proto3" json:"ask,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{10}
}

func (x *Price) GetUuid() int64 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *Price) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Price) GetBid() float32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *Price) GetAsk() float32 {
	if x != nil {
		return x.Ask
	}
	return 0
}

type Conn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Conn) Reset() {
	*x = Conn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_internal_protocol_price_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conn) ProtoMessage() {}

func (x *Conn) ProtoReflect() protoreflect.Message {
	mi := &file_src_internal_protocol_price_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conn.ProtoReflect.Descriptor instead.
func (*Conn) Descriptor() ([]byte, []int) {
	return file_src_internal_protocol_price_proto_rawDescGZIP(), []int{11}
}

func (x *Conn) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_src_internal_protocol_price_proto protoreflect.FileDescriptor

var file_src_internal_protocol_price_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x35, 0x0a,
	0x0b, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1d, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3c, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3b, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x70, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x62, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x61, 0x79, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x73, 0x62, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x62,
	0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x42, 0x61, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x61, 0x73,
	0x6b, 0x22, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x3e, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x30, 0x01, 0x32, 0xa7, 0x03, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x6e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x2d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x44, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_internal_protocol_price_proto_rawDescOnce sync.Once
	file_src_internal_protocol_price_proto_rawDescData = file_src_internal_protocol_price_proto_rawDesc
)

func file_src_internal_protocol_price_proto_rawDescGZIP() []byte {
	file_src_internal_protocol_price_proto_rawDescOnce.Do(func() {
		file_src_internal_protocol_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_internal_protocol_price_proto_rawDescData)
	})
	return file_src_internal_protocol_price_proto_rawDescData
}

var file_src_internal_protocol_price_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_src_internal_protocol_price_proto_goTypes = []interface{}{
	(*DonateValue)(nil),      // 0: protocol.DonateValue
	(*Balance)(nil),          // 1: protocol.Balance
	(*UserData)(nil),         // 2: protocol.UserData
	(*Token)(nil),            // 3: protocol.Token
	(*LoginReq)(nil),         // 4: protocol.LoginReq
	(*LoginResp)(nil),        // 5: protocol.LoginResp
	(*Position)(nil),         // 6: protocol.Position
	(*PositionCloseReq)(nil), // 7: protocol.PositionCloseReq
	(*PositionOpenReq)(nil),  // 8: protocol.PositionOpenReq
	(*Response)(nil),         // 9: protocol.Response
	(*Price)(nil),            // 10: protocol.Price
	(*Conn)(nil),             // 11: protocol.Conn
}
var file_src_internal_protocol_price_proto_depIdxs = []int32{
	11, // 0: protocol.PriceService.SendPrice:input_type -> protocol.Conn
	7,  // 1: protocol.PositionService.SendClosePositionRequest:input_type -> protocol.PositionCloseReq
	8,  // 2: protocol.PositionService.SendOpenPositionRequest:input_type -> protocol.PositionOpenReq
	4,  // 3: protocol.PositionService.Login:input_type -> protocol.LoginReq
	3,  // 4: protocol.PositionService.Logout:input_type -> protocol.Token
	3,  // 5: protocol.PositionService.GetUserData:input_type -> protocol.Token
	3,  // 6: protocol.PositionService.GetUserBalance:input_type -> protocol.Token
	0,  // 7: protocol.PositionService.Donate:input_type -> protocol.DonateValue
	10, // 8: protocol.PriceService.SendPrice:output_type -> protocol.Price
	9,  // 9: protocol.PositionService.SendClosePositionRequest:output_type -> protocol.Response
	9,  // 10: protocol.PositionService.SendOpenPositionRequest:output_type -> protocol.Response
	5,  // 11: protocol.PositionService.Login:output_type -> protocol.LoginResp
	9,  // 12: protocol.PositionService.Logout:output_type -> protocol.Response
	2,  // 13: protocol.PositionService.GetUserData:output_type -> protocol.UserData
	1,  // 14: protocol.PositionService.GetUserBalance:output_type -> protocol.Balance
	9,  // 15: protocol.PositionService.Donate:output_type -> protocol.Response
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_src_internal_protocol_price_proto_init() }
func file_src_internal_protocol_price_proto_init() {
	if File_src_internal_protocol_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_internal_protocol_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DonateValue); i {
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
		file_src_internal_protocol_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_src_internal_protocol_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_src_internal_protocol_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_src_internal_protocol_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_src_internal_protocol_price_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_src_internal_protocol_price_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_src_internal_protocol_price_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionCloseReq); i {
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
		file_src_internal_protocol_price_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionOpenReq); i {
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
		file_src_internal_protocol_price_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_src_internal_protocol_price_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_src_internal_protocol_price_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conn); i {
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
			RawDescriptor: file_src_internal_protocol_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_src_internal_protocol_price_proto_goTypes,
		DependencyIndexes: file_src_internal_protocol_price_proto_depIdxs,
		MessageInfos:      file_src_internal_protocol_price_proto_msgTypes,
	}.Build()
	File_src_internal_protocol_price_proto = out.File
	file_src_internal_protocol_price_proto_rawDesc = nil
	file_src_internal_protocol_price_proto_goTypes = nil
	file_src_internal_protocol_price_proto_depIdxs = nil
}
