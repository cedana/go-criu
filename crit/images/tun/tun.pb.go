// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: tun.proto

package tun

import (
	_ "github.com/cedana/go-criu/v7/crit/images/opts"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TunfileEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Netdev        *string                `protobuf:"bytes,2,opt,name=netdev" json:"netdev,omitempty"`
	Detached      *bool                  `protobuf:"varint,3,opt,name=detached" json:"detached,omitempty"`
	NsId          *uint32                `protobuf:"varint,4,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunfileEntry) Reset() {
	*x = TunfileEntry{}
	mi := &file_tun_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunfileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunfileEntry) ProtoMessage() {}

func (x *TunfileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tun_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunfileEntry.ProtoReflect.Descriptor instead.
func (*TunfileEntry) Descriptor() ([]byte, []int) {
	return file_tun_proto_rawDescGZIP(), []int{0}
}

func (x *TunfileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *TunfileEntry) GetNetdev() string {
	if x != nil && x.Netdev != nil {
		return *x.Netdev
	}
	return ""
}

func (x *TunfileEntry) GetDetached() bool {
	if x != nil && x.Detached != nil {
		return *x.Detached
	}
	return false
}

func (x *TunfileEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

type TunLinkEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Flags         *uint32                `protobuf:"varint,1,req,name=flags" json:"flags,omitempty"`
	Owner         *int32                 `protobuf:"varint,2,req,name=owner" json:"owner,omitempty"`
	Group         *int32                 `protobuf:"varint,3,req,name=group" json:"group,omitempty"`
	Vnethdr       *uint32                `protobuf:"varint,4,req,name=vnethdr" json:"vnethdr,omitempty"`
	Sndbuf        *uint32                `protobuf:"varint,5,req,name=sndbuf" json:"sndbuf,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunLinkEntry) Reset() {
	*x = TunLinkEntry{}
	mi := &file_tun_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunLinkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunLinkEntry) ProtoMessage() {}

func (x *TunLinkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tun_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunLinkEntry.ProtoReflect.Descriptor instead.
func (*TunLinkEntry) Descriptor() ([]byte, []int) {
	return file_tun_proto_rawDescGZIP(), []int{1}
}

func (x *TunLinkEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *TunLinkEntry) GetOwner() int32 {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return 0
}

func (x *TunLinkEntry) GetGroup() int32 {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return 0
}

func (x *TunLinkEntry) GetVnethdr() uint32 {
	if x != nil && x.Vnethdr != nil {
		return *x.Vnethdr
	}
	return 0
}

func (x *TunLinkEntry) GetSndbuf() uint32 {
	if x != nil && x.Sndbuf != nil {
		return *x.Sndbuf
	}
	return 0
}

var File_tun_proto protoreflect.FileDescriptor

const file_tun_proto_rawDesc = "" +
	"\n" +
	"\ttun.proto\x1a\n" +
	"opts.proto\"h\n" +
	"\rtunfile_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x16\n" +
	"\x06netdev\x18\x02 \x01(\tR\x06netdev\x12\x1a\n" +
	"\bdetached\x18\x03 \x01(\bR\bdetached\x12\x13\n" +
	"\x05ns_id\x18\x04 \x01(\rR\x04nsId\"\x8b\x01\n" +
	"\x0etun_link_entry\x12\x1b\n" +
	"\x05flags\x18\x01 \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x14\n" +
	"\x05owner\x18\x02 \x02(\x05R\x05owner\x12\x14\n" +
	"\x05group\x18\x03 \x02(\x05R\x05group\x12\x18\n" +
	"\avnethdr\x18\x04 \x02(\rR\avnethdr\x12\x16\n" +
	"\x06sndbuf\x18\x05 \x02(\rR\x06sndbuf"

var (
	file_tun_proto_rawDescOnce sync.Once
	file_tun_proto_rawDescData []byte
)

func file_tun_proto_rawDescGZIP() []byte {
	file_tun_proto_rawDescOnce.Do(func() {
		file_tun_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tun_proto_rawDesc), len(file_tun_proto_rawDesc)))
	})
	return file_tun_proto_rawDescData
}

var file_tun_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tun_proto_goTypes = []any{
	(*TunfileEntry)(nil), // 0: tunfile_entry
	(*TunLinkEntry)(nil), // 1: tun_link_entry
}
var file_tun_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tun_proto_init() }
func file_tun_proto_init() {
	if File_tun_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tun_proto_rawDesc), len(file_tun_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tun_proto_goTypes,
		DependencyIndexes: file_tun_proto_depIdxs,
		MessageInfos:      file_tun_proto_msgTypes,
	}.Build()
	File_tun_proto = out.File
	file_tun_proto_goTypes = nil
	file_tun_proto_depIdxs = nil
}
