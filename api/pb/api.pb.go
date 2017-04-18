// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Empty
	KeySelection
	Address
	Height
	Balances
	Key
	BoolResponse
	NetParams
	TransactionList
	Tx
	Txid
	FeeLevelSelection
	FeePerByte
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyPurpose int32

const (
	KeyPurpose_INTERNAL KeyPurpose = 0
	KeyPurpose_EXTERNAL KeyPurpose = 1
)

var KeyPurpose_name = map[int32]string{
	0: "INTERNAL",
	1: "EXTERNAL",
}
var KeyPurpose_value = map[string]int32{
	"INTERNAL": 0,
	"EXTERNAL": 1,
}

func (x KeyPurpose) String() string {
	return proto.EnumName(KeyPurpose_name, int32(x))
}
func (KeyPurpose) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FeeLevel int32

const (
	FeeLevel_ECONOMIC FeeLevel = 0
	FeeLevel_NORMAL   FeeLevel = 1
	FeeLevel_PRIORITY FeeLevel = 2
)

var FeeLevel_name = map[int32]string{
	0: "ECONOMIC",
	1: "NORMAL",
	2: "PRIORITY",
}
var FeeLevel_value = map[string]int32{
	"ECONOMIC": 0,
	"NORMAL":   1,
	"PRIORITY": 2,
}

func (x FeeLevel) String() string {
	return proto.EnumName(FeeLevel_name, int32(x))
}
func (FeeLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type KeySelection struct {
	Purpose KeyPurpose `protobuf:"varint,1,opt,name=purpose,enum=pb.KeyPurpose" json:"purpose,omitempty"`
}

func (m *KeySelection) Reset()                    { *m = KeySelection{} }
func (m *KeySelection) String() string            { return proto.CompactTextString(m) }
func (*KeySelection) ProtoMessage()               {}
func (*KeySelection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeySelection) GetPurpose() KeyPurpose {
	if m != nil {
		return m.Purpose
	}
	return KeyPurpose_INTERNAL
}

type Address struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Address) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type Height struct {
	Height uint32 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *Height) Reset()                    { *m = Height{} }
func (m *Height) String() string            { return proto.CompactTextString(m) }
func (*Height) ProtoMessage()               {}
func (*Height) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Height) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Balances struct {
	Confirmed   uint64 `protobuf:"varint,1,opt,name=confirmed" json:"confirmed,omitempty"`
	Unconfirmed uint64 `protobuf:"varint,2,opt,name=unconfirmed" json:"unconfirmed,omitempty"`
}

func (m *Balances) Reset()                    { *m = Balances{} }
func (m *Balances) String() string            { return proto.CompactTextString(m) }
func (*Balances) ProtoMessage()               {}
func (*Balances) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Balances) GetConfirmed() uint64 {
	if m != nil {
		return m.Confirmed
	}
	return 0
}

func (m *Balances) GetUnconfirmed() uint64 {
	if m != nil {
		return m.Unconfirmed
	}
	return 0
}

type Key struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type BoolResponse struct {
	Bool bool `protobuf:"varint,1,opt,name=bool" json:"bool,omitempty"`
}

func (m *BoolResponse) Reset()                    { *m = BoolResponse{} }
func (m *BoolResponse) String() string            { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()               {}
func (*BoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BoolResponse) GetBool() bool {
	if m != nil {
		return m.Bool
	}
	return false
}

type NetParams struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NetParams) Reset()                    { *m = NetParams{} }
func (m *NetParams) String() string            { return proto.CompactTextString(m) }
func (*NetParams) ProtoMessage()               {}
func (*NetParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NetParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TransactionList struct {
	Transactions []*Tx `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *TransactionList) Reset()                    { *m = TransactionList{} }
func (m *TransactionList) String() string            { return proto.CompactTextString(m) }
func (*TransactionList) ProtoMessage()               {}
func (*TransactionList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TransactionList) GetTransactions() []*Tx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Tx struct {
	Txid      string                     `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
	Value     int64                      `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Height    int32                      `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	WatchOnly bool                       `protobuf:"varint,5,opt,name=watchOnly" json:"watchOnly,omitempty"`
	Raw       []byte                     `protobuf:"bytes,6,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Tx) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Tx) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Tx) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Tx) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Tx) GetWatchOnly() bool {
	if m != nil {
		return m.WatchOnly
	}
	return false
}

func (m *Tx) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

type Txid struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *Txid) Reset()                    { *m = Txid{} }
func (m *Txid) String() string            { return proto.CompactTextString(m) }
func (*Txid) ProtoMessage()               {}
func (*Txid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Txid) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type FeeLevelSelection struct {
	FeeLevel FeeLevel `protobuf:"varint,1,opt,name=feeLevel,enum=pb.FeeLevel" json:"feeLevel,omitempty"`
}

func (m *FeeLevelSelection) Reset()                    { *m = FeeLevelSelection{} }
func (m *FeeLevelSelection) String() string            { return proto.CompactTextString(m) }
func (*FeeLevelSelection) ProtoMessage()               {}
func (*FeeLevelSelection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FeeLevelSelection) GetFeeLevel() FeeLevel {
	if m != nil {
		return m.FeeLevel
	}
	return FeeLevel_ECONOMIC
}

type FeePerByte struct {
	Fee uint64 `protobuf:"varint,1,opt,name=fee" json:"fee,omitempty"`
}

func (m *FeePerByte) Reset()                    { *m = FeePerByte{} }
func (m *FeePerByte) String() string            { return proto.CompactTextString(m) }
func (*FeePerByte) ProtoMessage()               {}
func (*FeePerByte) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FeePerByte) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*KeySelection)(nil), "pb.KeySelection")
	proto.RegisterType((*Address)(nil), "pb.Address")
	proto.RegisterType((*Height)(nil), "pb.Height")
	proto.RegisterType((*Balances)(nil), "pb.Balances")
	proto.RegisterType((*Key)(nil), "pb.Key")
	proto.RegisterType((*BoolResponse)(nil), "pb.BoolResponse")
	proto.RegisterType((*NetParams)(nil), "pb.NetParams")
	proto.RegisterType((*TransactionList)(nil), "pb.TransactionList")
	proto.RegisterType((*Tx)(nil), "pb.Tx")
	proto.RegisterType((*Txid)(nil), "pb.Txid")
	proto.RegisterType((*FeeLevelSelection)(nil), "pb.FeeLevelSelection")
	proto.RegisterType((*FeePerByte)(nil), "pb.FeePerByte")
	proto.RegisterEnum("pb.KeyPurpose", KeyPurpose_name, KeyPurpose_value)
	proto.RegisterEnum("pb.FeeLevel", FeeLevel_name, FeeLevel_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	CurrentAddress(ctx context.Context, in *KeySelection, opts ...grpc.CallOption) (*Address, error)
	NewAddress(ctx context.Context, in *KeySelection, opts ...grpc.CallOption) (*Address, error)
	ChainTip(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Height, error)
	Balance(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Balances, error)
	MasterPrivateKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error)
	MasterPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error)
	HasKey(ctx context.Context, in *Address, opts ...grpc.CallOption) (*BoolResponse, error)
	Params(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetParams, error)
	Transactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TransactionList, error)
	GetTransaction(ctx context.Context, in *Txid, opts ...grpc.CallOption) (*Tx, error)
	GetFeePerByte(ctx context.Context, in *FeeLevelSelection, opts ...grpc.CallOption) (*FeePerByte, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.API/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CurrentAddress(ctx context.Context, in *KeySelection, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := grpc.Invoke(ctx, "/pb.API/CurrentAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) NewAddress(ctx context.Context, in *KeySelection, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := grpc.Invoke(ctx, "/pb.API/NewAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ChainTip(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Height, error) {
	out := new(Height)
	err := grpc.Invoke(ctx, "/pb.API/ChainTip", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Balance(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Balances, error) {
	out := new(Balances)
	err := grpc.Invoke(ctx, "/pb.API/Balance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) MasterPrivateKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := grpc.Invoke(ctx, "/pb.API/MasterPrivateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) MasterPublicKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Key, error) {
	out := new(Key)
	err := grpc.Invoke(ctx, "/pb.API/MasterPublicKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) HasKey(ctx context.Context, in *Address, opts ...grpc.CallOption) (*BoolResponse, error) {
	out := new(BoolResponse)
	err := grpc.Invoke(ctx, "/pb.API/HasKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Params(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetParams, error) {
	out := new(NetParams)
	err := grpc.Invoke(ctx, "/pb.API/Params", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Transactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TransactionList, error) {
	out := new(TransactionList)
	err := grpc.Invoke(ctx, "/pb.API/Transactions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTransaction(ctx context.Context, in *Txid, opts ...grpc.CallOption) (*Tx, error) {
	out := new(Tx)
	err := grpc.Invoke(ctx, "/pb.API/GetTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetFeePerByte(ctx context.Context, in *FeeLevelSelection, opts ...grpc.CallOption) (*FeePerByte, error) {
	out := new(FeePerByte)
	err := grpc.Invoke(ctx, "/pb.API/GetFeePerByte", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Stop(context.Context, *Empty) (*Empty, error)
	CurrentAddress(context.Context, *KeySelection) (*Address, error)
	NewAddress(context.Context, *KeySelection) (*Address, error)
	ChainTip(context.Context, *Empty) (*Height, error)
	Balance(context.Context, *Empty) (*Balances, error)
	MasterPrivateKey(context.Context, *Empty) (*Key, error)
	MasterPublicKey(context.Context, *Empty) (*Key, error)
	HasKey(context.Context, *Address) (*BoolResponse, error)
	Params(context.Context, *Empty) (*NetParams, error)
	Transactions(context.Context, *Empty) (*TransactionList, error)
	GetTransaction(context.Context, *Txid) (*Tx, error)
	GetFeePerByte(context.Context, *FeeLevelSelection) (*FeePerByte, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CurrentAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySelection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CurrentAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/CurrentAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CurrentAddress(ctx, req.(*KeySelection))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySelection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).NewAddress(ctx, req.(*KeySelection))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ChainTip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ChainTip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/ChainTip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ChainTip(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Balance(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_MasterPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).MasterPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/MasterPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).MasterPrivateKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_MasterPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).MasterPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/MasterPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).MasterPublicKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_HasKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).HasKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/HasKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).HasKey(ctx, req.(*Address))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Params(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Transactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Transactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/Transactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Transactions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Txid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTransaction(ctx, req.(*Txid))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetFeePerByte_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeLevelSelection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetFeePerByte(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.API/GetFeePerByte",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetFeePerByte(ctx, req.(*FeeLevelSelection))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _API_Stop_Handler,
		},
		{
			MethodName: "CurrentAddress",
			Handler:    _API_CurrentAddress_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _API_NewAddress_Handler,
		},
		{
			MethodName: "ChainTip",
			Handler:    _API_ChainTip_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _API_Balance_Handler,
		},
		{
			MethodName: "MasterPrivateKey",
			Handler:    _API_MasterPrivateKey_Handler,
		},
		{
			MethodName: "MasterPublicKey",
			Handler:    _API_MasterPublicKey_Handler,
		},
		{
			MethodName: "HasKey",
			Handler:    _API_HasKey_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _API_Params_Handler,
		},
		{
			MethodName: "Transactions",
			Handler:    _API_Transactions_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _API_GetTransaction_Handler,
		},
		{
			MethodName: "GetFeePerByte",
			Handler:    _API_GetFeePerByte_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xda, 0x4c,
	0x14, 0xb5, 0xf9, 0x31, 0x70, 0x21, 0x84, 0x6f, 0xbe, 0xfe, 0x20, 0xd4, 0x26, 0x68, 0x5a, 0xa9,
	0x6e, 0xa4, 0x92, 0x96, 0x6e, 0xb2, 0xc9, 0x22, 0x41, 0xf9, 0xa1, 0x49, 0x00, 0x4d, 0xbc, 0x68,
	0x97, 0x03, 0x5c, 0x82, 0x55, 0x63, 0x5b, 0xf6, 0x90, 0xc4, 0x6f, 0xd0, 0xc7, 0xe9, 0x23, 0x56,
	0x33, 0x8c, 0xb1, 0x59, 0x54, 0xea, 0xee, 0xcc, 0x39, 0xc7, 0xe3, 0x7b, 0xef, 0xdc, 0x03, 0x35,
	0x1e, 0xba, 0xbd, 0x30, 0x0a, 0x44, 0x40, 0x0a, 0xe1, 0xb4, 0x73, 0xf8, 0x10, 0x04, 0x0f, 0x1e,
	0x1e, 0x2b, 0x66, 0xba, 0x5e, 0x1c, 0x0b, 0x77, 0x85, 0xb1, 0xe0, 0xab, 0x70, 0x63, 0xa2, 0x15,
	0x28, 0x5f, 0xac, 0x42, 0x91, 0xd0, 0x13, 0x68, 0xdc, 0x60, 0x72, 0x8f, 0x1e, 0xce, 0x84, 0x1b,
	0xf8, 0xc4, 0x86, 0x4a, 0xb8, 0x8e, 0xc2, 0x20, 0xc6, 0xb6, 0xd9, 0x35, 0xed, 0x66, 0xbf, 0xd9,
	0x0b, 0xa7, 0xbd, 0x1b, 0x4c, 0x26, 0x1b, 0x96, 0xa5, 0x32, 0x7d, 0x0b, 0x95, 0xb3, 0xf9, 0x3c,
	0xc2, 0x38, 0x26, 0x04, 0x4a, 0x7c, 0x3e, 0x8f, 0xd4, 0x17, 0x35, 0xa6, 0x30, 0xed, 0x82, 0x75,
	0x8d, 0xee, 0xc3, 0x52, 0x90, 0x57, 0x60, 0x2d, 0x15, 0x52, 0xfa, 0x1e, 0xd3, 0x27, 0xfa, 0x0d,
	0xaa, 0xe7, 0xdc, 0xe3, 0xfe, 0x0c, 0x63, 0xf2, 0x06, 0x6a, 0xb3, 0xc0, 0x5f, 0xb8, 0xd1, 0x0a,
	0xe7, 0xca, 0x56, 0x62, 0x19, 0x41, 0xba, 0x50, 0x5f, 0xfb, 0x99, 0x5e, 0x50, 0x7a, 0x9e, 0xa2,
	0xaf, 0xa1, 0x78, 0x83, 0x09, 0x69, 0x41, 0xf1, 0x27, 0x26, 0xba, 0x0e, 0x09, 0x29, 0x85, 0xc6,
	0x79, 0x10, 0x78, 0x0c, 0xe3, 0x30, 0xf0, 0x63, 0x94, 0xa5, 0x4e, 0x83, 0xc0, 0x53, 0x96, 0x2a,
	0x53, 0x98, 0x1e, 0x42, 0x6d, 0x84, 0x62, 0xc2, 0x23, 0xbe, 0x52, 0xbd, 0xf8, 0x7c, 0x85, 0x69,
	0x2f, 0x12, 0xd3, 0x53, 0xd8, 0x77, 0x22, 0xee, 0xc7, 0x5c, 0xcd, 0xe8, 0xd6, 0x8d, 0x05, 0x39,
	0x82, 0x86, 0xc8, 0xa8, 0xb8, 0x6d, 0x76, 0x8b, 0x76, 0xbd, 0x6f, 0xc9, 0x61, 0x39, 0xcf, 0x6c,
	0x47, 0xa3, 0xbf, 0x4d, 0x28, 0x38, 0xcf, 0xf2, 0x66, 0xf1, 0xec, 0xce, 0xd3, 0x9b, 0x25, 0x26,
	0x2f, 0xa0, 0xfc, 0xc8, 0xbd, 0x35, 0xaa, 0x9e, 0x8a, 0x6c, 0x73, 0xc8, 0x4d, 0xac, 0xd8, 0x35,
	0xed, 0x72, 0x3a, 0x31, 0x72, 0x02, 0xb5, 0xed, 0x43, 0xb6, 0x4b, 0x5d, 0xd3, 0xae, 0xf7, 0x3b,
	0xbd, 0xcd, 0x53, 0xf7, 0xd2, 0xa7, 0xee, 0x39, 0xa9, 0x83, 0x65, 0x66, 0x39, 0xdf, 0x27, 0x2e,
	0x66, 0xcb, 0xb1, 0xef, 0x25, 0xed, 0xb2, 0xea, 0x3d, 0x23, 0xe4, 0xd8, 0x22, 0xfe, 0xd4, 0xb6,
	0xba, 0xa6, 0xdd, 0x60, 0x12, 0xd2, 0x0e, 0x94, 0x1c, 0x59, 0x1f, 0x81, 0xd2, 0x92, 0xc7, 0xcb,
	0xb4, 0x66, 0x89, 0xe9, 0x29, 0xfc, 0x77, 0x89, 0x78, 0x8b, 0x8f, 0xe8, 0xe5, 0xf7, 0xa6, 0xba,
	0xd0, 0xa4, 0x5e, 0x9c, 0x86, 0x9c, 0x45, 0x6a, 0x64, 0x5b, 0x95, 0x1e, 0x00, 0x5c, 0x22, 0x4e,
	0x30, 0x3a, 0x4f, 0x04, 0xca, 0x5f, 0x2f, 0x10, 0xf5, 0x93, 0x4b, 0x78, 0x64, 0x03, 0x64, 0xeb,
	0x46, 0x1a, 0x50, 0x1d, 0x8e, 0x9c, 0x0b, 0x36, 0x3a, 0xbb, 0x6d, 0x19, 0xf2, 0x74, 0xf1, 0x5d,
	0x9f, 0xcc, 0xa3, 0x3e, 0x54, 0xd3, 0xfb, 0x95, 0x32, 0x18, 0x8f, 0xc6, 0x77, 0xc3, 0x41, 0xcb,
	0x20, 0x00, 0xd6, 0x68, 0xcc, 0xee, 0xa4, 0x4b, 0x2a, 0x13, 0x36, 0x1c, 0xb3, 0xa1, 0xf3, 0xa3,
	0x55, 0xe8, 0xff, 0x2a, 0x41, 0xf1, 0x6c, 0x32, 0x24, 0x07, 0x50, 0xba, 0x17, 0x41, 0x48, 0x6a,
	0xb2, 0x4a, 0x15, 0x85, 0x4e, 0x06, 0xa9, 0x41, 0xbe, 0x40, 0x73, 0xb0, 0x8e, 0x22, 0xf4, 0x45,
	0xba, 0xe4, 0x2d, 0x1d, 0x84, 0x6d, 0xcf, 0x9d, 0xba, 0x64, 0xb4, 0x4c, 0x0d, 0xf2, 0x09, 0x60,
	0x84, 0x4f, 0xff, 0x6c, 0x7f, 0x07, 0xd5, 0xc1, 0x92, 0xbb, 0xbe, 0xe3, 0xee, 0x54, 0x01, 0x12,
	0x6e, 0x92, 0x43, 0x0d, 0xf2, 0x1e, 0x2a, 0x3a, 0x23, 0x79, 0x8f, 0x1a, 0x6d, 0x9a, 0x1d, 0x6a,
	0x10, 0x1b, 0x5a, 0x77, 0x3c, 0x16, 0x18, 0x4d, 0x22, 0xf7, 0x91, 0x0b, 0x94, 0x51, 0xc8, 0xd9,
	0x2b, 0xba, 0x14, 0x6a, 0x90, 0x0f, 0xb0, 0xaf, 0x9d, 0xeb, 0xa9, 0xe7, 0xce, 0xfe, 0x6e, 0xfc,
	0x08, 0xd6, 0x35, 0x8f, 0xa5, 0x9e, 0x2f, 0xbb, 0xa3, 0xba, 0xca, 0x07, 0x4a, 0xd5, 0x68, 0xe9,
	0xec, 0xe4, 0xae, 0xda, 0x93, 0x70, 0x9b, 0x2a, 0x6a, 0x90, 0xcf, 0xd0, 0xc8, 0x65, 0x68, 0xc7,
	0xfb, 0xbf, 0x4a, 0xcd, 0x6e, 0xc0, 0xd4, 0xbd, 0xcd, 0x2b, 0x14, 0x39, 0x9e, 0x54, 0x37, 0xf1,
	0x72, 0xe7, 0x1d, 0x1d, 0x34, 0x6a, 0x90, 0x13, 0xd8, 0xbb, 0x42, 0x91, 0xdb, 0xa8, 0x97, 0xf9,
	0xbd, 0xcb, 0xa6, 0xdf, 0xd4, 0xb4, 0xb6, 0x51, 0x63, 0x6a, 0xa9, 0xc8, 0x7c, 0xfd, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x89, 0xc1, 0xc2, 0x5c, 0x3c, 0x05, 0x00, 0x00,
}
