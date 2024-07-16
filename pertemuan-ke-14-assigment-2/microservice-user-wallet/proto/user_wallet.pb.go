// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/user_wallet.proto

package user_wallet

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type HistoryTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserIdFrom      int32                  `protobuf:"varint,2,opt,name=user_id_from,json=userIdFrom,proto3" json:"user_id_from,omitempty"`
	UserIdTo        int32                  `protobuf:"varint,3,opt,name=user_id_to,json=userIdTo,proto3" json:"user_id_to,omitempty"`
	TypeTransaction string                 `protobuf:"bytes,4,opt,name=type_transaction,json=typeTransaction,proto3" json:"type_transaction,omitempty"`
	TypeCredit      string                 `protobuf:"bytes,5,opt,name=type_credit,json=typeCredit,proto3" json:"type_credit,omitempty"`
	Total           float32                `protobuf:"fixed32,6,opt,name=total,proto3" json:"total,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *HistoryTransaction) Reset() {
	*x = HistoryTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTransaction) ProtoMessage() {}

func (x *HistoryTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTransaction.ProtoReflect.Descriptor instead.
func (*HistoryTransaction) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryTransaction) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HistoryTransaction) GetUserIdFrom() int32 {
	if x != nil {
		return x.UserIdFrom
	}
	return 0
}

func (x *HistoryTransaction) GetUserIdTo() int32 {
	if x != nil {
		return x.UserIdTo
	}
	return 0
}

func (x *HistoryTransaction) GetTypeTransaction() string {
	if x != nil {
		return x.TypeTransaction
	}
	return ""
}

func (x *HistoryTransaction) GetTypeCredit() string {
	if x != nil {
		return x.TypeCredit
	}
	return ""
}

func (x *HistoryTransaction) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *HistoryTransaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type TopupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TopupRequest) Reset() {
	*x = TopupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopupRequest) ProtoMessage() {}

func (x *TopupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopupRequest.ProtoReflect.Descriptor instead.
func (*TopupRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *TopupRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TopupRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TopupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	History *HistoryTransaction `protobuf:"bytes,1,opt,name=history,proto3" json:"history,omitempty"`
}

func (x *TopupResponse) Reset() {
	*x = TopupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopupResponse) ProtoMessage() {}

func (x *TopupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopupResponse.ProtoReflect.Descriptor instead.
func (*TopupResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *TopupResponse) GetHistory() *HistoryTransaction {
	if x != nil {
		return x.History
	}
	return nil
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From   int32   `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To     int32   `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *TransferRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TransferRequest) GetFrom() int32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *TransferRequest) GetTo() int32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *TransferRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	History1 *HistoryTransaction `protobuf:"bytes,1,opt,name=history1,proto3" json:"history1,omitempty"`
	History2 *HistoryTransaction `protobuf:"bytes,2,opt,name=history2,proto3" json:"history2,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *TransferResponse) GetHistory1() *HistoryTransaction {
	if x != nil {
		return x.History1
	}
	return nil
}

func (x *TransferResponse) GetHistory2() *HistoryTransaction {
	if x != nil {
		return x.History2
	}
	return nil
}

type GetUserBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // Menambahkan field id
}

func (x *GetUserBalanceRequest) Reset() {
	*x = GetUserBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceRequest) ProtoMessage() {}

func (x *GetUserBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetUserBalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserBalanceRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Saldo     float32                `protobuf:"fixed32,2,opt,name=saldo,proto3" json:"saldo,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetUserBalanceResponse) Reset() {
	*x = GetUserBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBalanceResponse) ProtoMessage() {}

func (x *GetUserBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetUserBalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserBalanceResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetUserBalanceResponse) GetSaldo() float32 {
	if x != nil {
		return x.Saldo
	}
	return 0
}

func (x *GetUserBalanceResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetTransactionHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTransactionHistoryRequest) Reset() {
	*x = GetTransactionHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionHistoryRequest) ProtoMessage() {}

func (x *GetTransactionHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *GetTransactionHistoryRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetTransactionHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	History []*HistoryTransaction `protobuf:"bytes,1,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *GetTransactionHistoryResponse) Reset() {
	*x = GetTransactionHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionHistoryResponse) ProtoMessage() {}

func (x *GetTransactionHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *GetTransactionHistoryResponse) GetHistory() []*HistoryTransaction {
	if x != nil {
		return x.History
	}
	return nil
}

var File_proto_user_wallet_proto protoreflect.FileDescriptor

var file_proto_user_wallet_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x12, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1c, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x36, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x5b, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x5d, 0x0a, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x08, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x31, 0x12, 0x4c,
	0x0a, 0x08, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x32, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c,
	0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x64, 0x6f, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x37, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x32, 0xe4, 0x04, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x75, 0x70, 0x12,
	0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x70, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b,
	0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x75, 0x70, 0x12, 0x7f, 0x0a, 0x08, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x9f, 0x01, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xb6,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x1d, 0x5a, 0x1b, 0x70, 0x72, 0x61, 0x69, 0x73,
	0x69, 0x6e, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_wallet_proto_rawDescOnce sync.Once
	file_proto_user_wallet_proto_rawDescData = file_proto_user_wallet_proto_rawDesc
)

func file_proto_user_wallet_proto_rawDescGZIP() []byte {
	file_proto_user_wallet_proto_rawDescOnce.Do(func() {
		file_proto_user_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_wallet_proto_rawDescData)
	})
	return file_proto_user_wallet_proto_rawDescData
}

var file_proto_user_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_user_wallet_proto_goTypes = []any{
	(*HistoryTransaction)(nil),            // 0: proto.user_wallet_service.v1.HistoryTransaction
	(*TopupRequest)(nil),                  // 1: proto.user_wallet_service.v1.TopupRequest
	(*TopupResponse)(nil),                 // 2: proto.user_wallet_service.v1.TopupResponse
	(*TransferRequest)(nil),               // 3: proto.user_wallet_service.v1.TransferRequest
	(*TransferResponse)(nil),              // 4: proto.user_wallet_service.v1.TransferResponse
	(*GetUserBalanceRequest)(nil),         // 5: proto.user_wallet_service.v1.GetUserBalanceRequest
	(*GetUserBalanceResponse)(nil),        // 6: proto.user_wallet_service.v1.GetUserBalanceResponse
	(*GetTransactionHistoryRequest)(nil),  // 7: proto.user_wallet_service.v1.GetTransactionHistoryRequest
	(*GetTransactionHistoryResponse)(nil), // 8: proto.user_wallet_service.v1.GetTransactionHistoryResponse
	(*timestamppb.Timestamp)(nil),         // 9: google.protobuf.Timestamp
}
var file_proto_user_wallet_proto_depIdxs = []int32{
	9,  // 0: proto.user_wallet_service.v1.HistoryTransaction.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: proto.user_wallet_service.v1.TopupResponse.history:type_name -> proto.user_wallet_service.v1.HistoryTransaction
	0,  // 2: proto.user_wallet_service.v1.TransferResponse.history1:type_name -> proto.user_wallet_service.v1.HistoryTransaction
	0,  // 3: proto.user_wallet_service.v1.TransferResponse.history2:type_name -> proto.user_wallet_service.v1.HistoryTransaction
	9,  // 4: proto.user_wallet_service.v1.GetUserBalanceResponse.created_at:type_name -> google.protobuf.Timestamp
	0,  // 5: proto.user_wallet_service.v1.GetTransactionHistoryResponse.history:type_name -> proto.user_wallet_service.v1.HistoryTransaction
	1,  // 6: proto.user_wallet_service.v1.UserWalletService.Topup:input_type -> proto.user_wallet_service.v1.TopupRequest
	3,  // 7: proto.user_wallet_service.v1.UserWalletService.Transfer:input_type -> proto.user_wallet_service.v1.TransferRequest
	5,  // 8: proto.user_wallet_service.v1.UserWalletService.GetUserBalance:input_type -> proto.user_wallet_service.v1.GetUserBalanceRequest
	7,  // 9: proto.user_wallet_service.v1.UserWalletService.GetTransactionHistory:input_type -> proto.user_wallet_service.v1.GetTransactionHistoryRequest
	2,  // 10: proto.user_wallet_service.v1.UserWalletService.Topup:output_type -> proto.user_wallet_service.v1.TopupResponse
	4,  // 11: proto.user_wallet_service.v1.UserWalletService.Transfer:output_type -> proto.user_wallet_service.v1.TransferResponse
	6,  // 12: proto.user_wallet_service.v1.UserWalletService.GetUserBalance:output_type -> proto.user_wallet_service.v1.GetUserBalanceResponse
	8,  // 13: proto.user_wallet_service.v1.UserWalletService.GetTransactionHistory:output_type -> proto.user_wallet_service.v1.GetTransactionHistoryResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_user_wallet_proto_init() }
func file_proto_user_wallet_proto_init() {
	if File_proto_user_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_wallet_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HistoryTransaction); i {
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
		file_proto_user_wallet_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TopupRequest); i {
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
		file_proto_user_wallet_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TopupResponse); i {
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
		file_proto_user_wallet_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TransferRequest); i {
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
		file_proto_user_wallet_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TransferResponse); i {
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
		file_proto_user_wallet_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserBalanceRequest); i {
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
		file_proto_user_wallet_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserBalanceResponse); i {
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
		file_proto_user_wallet_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetTransactionHistoryRequest); i {
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
		file_proto_user_wallet_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetTransactionHistoryResponse); i {
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
			RawDescriptor: file_proto_user_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_wallet_proto_goTypes,
		DependencyIndexes: file_proto_user_wallet_proto_depIdxs,
		MessageInfos:      file_proto_user_wallet_proto_msgTypes,
	}.Build()
	File_proto_user_wallet_proto = out.File
	file_proto_user_wallet_proto_rawDesc = nil
	file_proto_user_wallet_proto_goTypes = nil
	file_proto_user_wallet_proto_depIdxs = nil
}
