// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: grpc.proto

package grpc_package

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

type CurrencyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Amount        float64                `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`                               // Сумма для конвертации
	FromCurrency  string                 `protobuf:"bytes,2,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"` // Исходная валюта (например, "USD")
	ToCurrency    string                 `protobuf:"bytes,3,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`       // Целевая валюта (например, "EUR")
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CurrencyRequest) Reset() {
	*x = CurrencyRequest{}
	mi := &file_grpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyRequest) ProtoMessage() {}

func (x *CurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CurrencyRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *CurrencyRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

type CurrencyResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ConvertedAmount float64                `protobuf:"fixed64,1,opt,name=converted_amount,json=convertedAmount,proto3" json:"converted_amount,omitempty"` // Конвертированная сумма
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CurrencyResponse) Reset() {
	*x = CurrencyResponse{}
	mi := &file_grpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyResponse) ProtoMessage() {}

func (x *CurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyResponse.ProtoReflect.Descriptor instead.
func (*CurrencyResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyResponse) GetConvertedAmount() float64 {
	if x != nil {
		return x.ConvertedAmount
	}
	return 0
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x3d, 0x0a, 0x10, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x65, 0x0a, 0x11, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12,
	0x50, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x26, 0x5a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x3b, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_goTypes = []any{
	(*CurrencyRequest)(nil),  // 0: grpc_package.CurrencyRequest
	(*CurrencyResponse)(nil), // 1: grpc_package.CurrencyResponse
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: grpc_package.CurrencyConverter.ConvertCurrency:input_type -> grpc_package.CurrencyRequest
	1, // 1: grpc_package.CurrencyConverter.ConvertCurrency:output_type -> grpc_package.CurrencyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}
