// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: address.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Province      string                 `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City          string                 `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	District      string                 `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Address       string                 `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	SignerName    string                 `protobuf:"bytes,7,opt,name=signerName,proto3" json:"signerName,omitempty"`
	SignerMobile  string                 `protobuf:"bytes,8,opt,name=signerMobile,proto3" json:"signerMobile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	mi := &file_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *AddressRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddressRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddressRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *AddressRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressRequest) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *AddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressRequest) GetSignerName() string {
	if x != nil {
		return x.SignerName
	}
	return ""
}

func (x *AddressRequest) GetSignerMobile() string {
	if x != nil {
		return x.SignerMobile
	}
	return ""
}

type AddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Province      string                 `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City          string                 `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	District      string                 `protobuf:"bytes,5,opt,name=district,proto3" json:"district,omitempty"`
	Address       string                 `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	SignerName    string                 `protobuf:"bytes,7,opt,name=signerName,proto3" json:"signerName,omitempty"`
	SignerMobile  string                 `protobuf:"bytes,8,opt,name=signerMobile,proto3" json:"signerMobile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressResponse) Reset() {
	*x = AddressResponse{}
	mi := &file_address_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResponse) ProtoMessage() {}

func (x *AddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResponse.ProtoReflect.Descriptor instead.
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{1}
}

func (x *AddressResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddressResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddressResponse) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *AddressResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressResponse) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *AddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressResponse) GetSignerName() string {
	if x != nil {
		return x.SignerName
	}
	return ""
}

func (x *AddressResponse) GetSignerMobile() string {
	if x != nil {
		return x.SignerMobile
	}
	return ""
}

type AddressListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Total         int32                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data          []*AddressResponse     `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressListResponse) Reset() {
	*x = AddressListResponse{}
	mi := &file_address_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressListResponse) ProtoMessage() {}

func (x *AddressListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressListResponse.ProtoReflect.Descriptor instead.
func (*AddressListResponse) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AddressListResponse) GetData() []*AddressResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0xe3, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x51, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xea, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0f,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData []byte
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_address_proto_rawDesc), len(file_address_proto_rawDesc)))
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_address_proto_goTypes = []any{
	(*AddressRequest)(nil),      // 0: AddressRequest
	(*AddressResponse)(nil),     // 1: AddressResponse
	(*AddressListResponse)(nil), // 2: AddressListResponse
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_address_proto_depIdxs = []int32{
	1, // 0: AddressListResponse.data:type_name -> AddressResponse
	0, // 1: Address.GetAddressList:input_type -> AddressRequest
	0, // 2: Address.CreateAddress:input_type -> AddressRequest
	0, // 3: Address.DeleteAddress:input_type -> AddressRequest
	0, // 4: Address.UpdateAddress:input_type -> AddressRequest
	2, // 5: Address.GetAddressList:output_type -> AddressListResponse
	1, // 6: Address.CreateAddress:output_type -> AddressResponse
	3, // 7: Address.DeleteAddress:output_type -> google.protobuf.Empty
	3, // 8: Address.UpdateAddress:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_address_proto_rawDesc), len(file_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
