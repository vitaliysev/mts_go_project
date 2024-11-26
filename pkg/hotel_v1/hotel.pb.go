// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: hotel.proto

package hotel_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotel *HotelInfo `protobuf:"bytes,1,opt,name=hotel,proto3" json:"hotel,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *GetInfoResponse) GetHotel() *HotelInfo {
	if x != nil {
		return x.Hotel
	}
	return nil
}

type HotelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Price    int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *HotelInfo) Reset() {
	*x = HotelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelInfo) ProtoMessage() {}

func (x *HotelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelInfo.ProtoReflect.Descriptor instead.
func (*HotelInfo) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *HotelInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HotelInfo) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *HotelInfo) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_hotel_proto protoreflect.FileDescriptor

var file_hotel_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0x51, 0x0a, 0x09, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x49, 0x0a, 0x07,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x56, 0x31, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x4f, 0x55, 0x52, 0x2d, 0x55, 0x53, 0x45, 0x52, 0x2d,
	0x4f, 0x52, 0x2d, 0x4f, 0x52, 0x47, 0x2d, 0x4e, 0x41, 0x4d, 0x45, 0x2f, 0x59, 0x4f, 0x55, 0x52,
	0x2d, 0x52, 0x45, 0x50, 0x4f, 0x2d, 0x4e, 0x41, 0x4d, 0x45, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31, 0x3b, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hotel_proto_rawDescOnce sync.Once
	file_hotel_proto_rawDescData = file_hotel_proto_rawDesc
)

func file_hotel_proto_rawDescGZIP() []byte {
	file_hotel_proto_rawDescOnce.Do(func() {
		file_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(file_hotel_proto_rawDescData)
	})
	return file_hotel_proto_rawDescData
}

var file_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hotel_proto_goTypes = []any{
	(*GetInfoRequest)(nil),  // 0: hotel_v1.GetInfoRequest
	(*GetInfoResponse)(nil), // 1: hotel_v1.GetInfoResponse
	(*HotelInfo)(nil),       // 2: hotel_v1.HotelInfo
}
var file_hotel_proto_depIdxs = []int32{
	2, // 0: hotel_v1.GetInfoResponse.hotel:type_name -> hotel_v1.HotelInfo
	0, // 1: hotel_v1.HotelV1.GetInfo:input_type -> hotel_v1.GetInfoRequest
	1, // 2: hotel_v1.HotelV1.GetInfo:output_type -> hotel_v1.GetInfoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hotel_proto_init() }
func file_hotel_proto_init() {
	if File_hotel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hotel_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoRequest); i {
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
		file_hotel_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetInfoResponse); i {
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
		file_hotel_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HotelInfo); i {
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
			RawDescriptor: file_hotel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hotel_proto_goTypes,
		DependencyIndexes: file_hotel_proto_depIdxs,
		MessageInfos:      file_hotel_proto_msgTypes,
	}.Build()
	File_hotel_proto = out.File
	file_hotel_proto_rawDesc = nil
	file_hotel_proto_goTypes = nil
	file_hotel_proto_depIdxs = nil
}