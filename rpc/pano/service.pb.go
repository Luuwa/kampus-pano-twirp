// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc/pano/service.proto

package pano

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

// Id of Post.
type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pano_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pano_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_rpc_pano_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// Post is users post made in pano.kampus
type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Site    string `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
	Slug    string `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Content string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	UserID  string `protobuf:"bytes,6,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pano_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pano_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_rpc_pano_service_proto_rawDescGZIP(), []int{1}
}

func (x *Post) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Post) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

var File_rpc_pano_service_proto protoreflect.FileDescriptor

var file_rpc_pano_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x6e, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2e, 0x70, 0x61, 0x6e, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x32, 0x41, 0x0a, 0x04, 0x50, 0x61, 0x6e, 0x6f, 0x12, 0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x6e,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x70, 0x61, 0x6e, 0x6f, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x6e, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_pano_service_proto_rawDescOnce sync.Once
	file_rpc_pano_service_proto_rawDescData = file_rpc_pano_service_proto_rawDesc
)

func file_rpc_pano_service_proto_rawDescGZIP() []byte {
	file_rpc_pano_service_proto_rawDescOnce.Do(func() {
		file_rpc_pano_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_pano_service_proto_rawDescData)
	})
	return file_rpc_pano_service_proto_rawDescData
}

var file_rpc_pano_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_pano_service_proto_goTypes = []interface{}{
	(*GetPostRequest)(nil), // 0: kampus.pano.GetPostRequest
	(*Post)(nil),           // 1: kampus.pano.Post
}
var file_rpc_pano_service_proto_depIdxs = []int32{
	0, // 0: kampus.pano.Pano.GetPost:input_type -> kampus.pano.GetPostRequest
	1, // 1: kampus.pano.Pano.GetPost:output_type -> kampus.pano.Post
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_pano_service_proto_init() }
func file_rpc_pano_service_proto_init() {
	if File_rpc_pano_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_pano_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_rpc_pano_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
			RawDescriptor: file_rpc_pano_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_pano_service_proto_goTypes,
		DependencyIndexes: file_rpc_pano_service_proto_depIdxs,
		MessageInfos:      file_rpc_pano_service_proto_msgTypes,
	}.Build()
	File_rpc_pano_service_proto = out.File
	file_rpc_pano_service_proto_rawDesc = nil
	file_rpc_pano_service_proto_goTypes = nil
	file_rpc_pano_service_proto_depIdxs = nil
}