// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: proto/calc.proto

package proto

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

type PrimeDecomposeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeDecomposeRequest) Reset() {
	*x = PrimeDecomposeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecomposeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecomposeRequest) ProtoMessage() {}

func (x *PrimeDecomposeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecomposeRequest.ProtoReflect.Descriptor instead.
func (*PrimeDecomposeRequest) Descriptor() ([]byte, []int) {
	return file_proto_calc_proto_rawDescGZIP(), []int{0}
}

func (x *PrimeDecomposeRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeDecomposeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeDecomposeResponse) Reset() {
	*x = PrimeDecomposeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecomposeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecomposeResponse) ProtoMessage() {}

func (x *PrimeDecomposeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecomposeResponse.ProtoReflect.Descriptor instead.
func (*PrimeDecomposeResponse) Descriptor() ([]byte, []int) {
	return file_proto_calc_proto_rawDescGZIP(), []int{1}
}

func (x *PrimeDecomposeResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_proto_calc_proto_rawDescGZIP(), []int{2}
}

func (x *AverageRequest) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AverageRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number float64 `protobuf:"fixed64,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRespond) Reset() {
	*x = AverageRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRespond) ProtoMessage() {}

func (x *AverageRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRespond.ProtoReflect.Descriptor instead.
func (*AverageRespond) Descriptor() ([]byte, []int) {
	return file_proto_calc_proto_rawDescGZIP(), []int{3}
}

func (x *AverageRespond) GetNumber() float64 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_proto_calc_proto protoreflect.FileDescriptor

var file_proto_calc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x16, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0e,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0xa7, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a, 0x11, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x28, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_calc_proto_rawDescOnce sync.Once
	file_proto_calc_proto_rawDescData = file_proto_calc_proto_rawDesc
)

func file_proto_calc_proto_rawDescGZIP() []byte {
	file_proto_calc_proto_rawDescOnce.Do(func() {
		file_proto_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calc_proto_rawDescData)
	})
	return file_proto_calc_proto_rawDescData
}

var file_proto_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_calc_proto_goTypes = []interface{}{
	(*PrimeDecomposeRequest)(nil),  // 0: proto.PrimeDecomposeRequest
	(*PrimeDecomposeResponse)(nil), // 1: proto.PrimeDecomposeResponse
	(*AverageRequest)(nil),         // 2: proto.AverageRequest
	(*AverageRespond)(nil),         // 3: proto.AverageRespond
}
var file_proto_calc_proto_depIdxs = []int32{
	0, // 0: proto.CalcService.PrimeDecomposition:input_type -> proto.PrimeDecomposeRequest
	2, // 1: proto.CalcService.AverageCalculator:input_type -> proto.AverageRequest
	1, // 2: proto.CalcService.PrimeDecomposition:output_type -> proto.PrimeDecomposeResponse
	3, // 3: proto.CalcService.AverageCalculator:output_type -> proto.AverageRespond
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_calc_proto_init() }
func file_proto_calc_proto_init() {
	if File_proto_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecomposeRequest); i {
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
		file_proto_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecomposeResponse); i {
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
		file_proto_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_proto_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRespond); i {
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
			RawDescriptor: file_proto_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calc_proto_goTypes,
		DependencyIndexes: file_proto_calc_proto_depIdxs,
		MessageInfos:      file_proto_calc_proto_msgTypes,
	}.Build()
	File_proto_calc_proto = out.File
	file_proto_calc_proto_rawDesc = nil
	file_proto_calc_proto_goTypes = nil
	file_proto_calc_proto_depIdxs = nil
}
