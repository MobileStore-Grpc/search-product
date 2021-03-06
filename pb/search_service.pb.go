// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: search_service.proto

package pb

import (
	pb "github.com/MobileStore-Grpc/product/pb"
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

type SearchMobileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *SearchMobileRequest) Reset() {
	*x = SearchMobileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMobileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMobileRequest) ProtoMessage() {}

func (x *SearchMobileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMobileRequest.ProtoReflect.Descriptor instead.
func (*SearchMobileRequest) Descriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchMobileRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type SearchMobileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile *pb.Mobile `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *SearchMobileResponse) Reset() {
	*x = SearchMobileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMobileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMobileResponse) ProtoMessage() {}

func (x *SearchMobileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMobileResponse.ProtoReflect.Descriptor instead.
func (*SearchMobileResponse) Descriptor() ([]byte, []int) {
	return file_search_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMobileResponse) GetMobile() *pb.Mobile {
	if x != nil {
		return x.Mobile
	}
	return nil
}

var File_search_service_proto protoreflect.FileDescriptor

var file_search_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x3a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x32, 0x6f, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x47, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_service_proto_rawDescOnce sync.Once
	file_search_service_proto_rawDescData = file_search_service_proto_rawDesc
)

func file_search_service_proto_rawDescGZIP() []byte {
	file_search_service_proto_rawDescOnce.Do(func() {
		file_search_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_service_proto_rawDescData)
	})
	return file_search_service_proto_rawDescData
}

var file_search_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_search_service_proto_goTypes = []interface{}{
	(*SearchMobileRequest)(nil),  // 0: pb.SearchMobileRequest
	(*SearchMobileResponse)(nil), // 1: pb.SearchMobileResponse
	(*Filter)(nil),               // 2: pb.Filter
	(*pb.Mobile)(nil),            // 3: pb.Mobile
}
var file_search_service_proto_depIdxs = []int32{
	2, // 0: pb.SearchMobileRequest.filter:type_name -> pb.Filter
	3, // 1: pb.SearchMobileResponse.mobile:type_name -> pb.Mobile
	0, // 2: pb.SearchService.SearchMobile:input_type -> pb.SearchMobileRequest
	1, // 3: pb.SearchService.SearchMobile:output_type -> pb.SearchMobileResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_search_service_proto_init() }
func file_search_service_proto_init() {
	if File_search_service_proto != nil {
		return
	}
	file_filter_message_proto_init()
	file_google_api_annotations_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_search_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMobileRequest); i {
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
		file_search_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMobileResponse); i {
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
			RawDescriptor: file_search_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_service_proto_goTypes,
		DependencyIndexes: file_search_service_proto_depIdxs,
		MessageInfos:      file_search_service_proto_msgTypes,
	}.Build()
	File_search_service_proto = out.File
	file_search_service_proto_rawDesc = nil
	file_search_service_proto_goTypes = nil
	file_search_service_proto_depIdxs = nil
}
