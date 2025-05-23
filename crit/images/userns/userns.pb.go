// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: userns.proto

package userns

import (
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

type UidGidExtent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	First         *uint32                `protobuf:"varint,1,req,name=first" json:"first,omitempty"`
	LowerFirst    *uint32                `protobuf:"varint,2,req,name=lower_first,json=lowerFirst" json:"lower_first,omitempty"`
	Count         *uint32                `protobuf:"varint,3,req,name=count" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UidGidExtent) Reset() {
	*x = UidGidExtent{}
	mi := &file_userns_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UidGidExtent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UidGidExtent) ProtoMessage() {}

func (x *UidGidExtent) ProtoReflect() protoreflect.Message {
	mi := &file_userns_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UidGidExtent.ProtoReflect.Descriptor instead.
func (*UidGidExtent) Descriptor() ([]byte, []int) {
	return file_userns_proto_rawDescGZIP(), []int{0}
}

func (x *UidGidExtent) GetFirst() uint32 {
	if x != nil && x.First != nil {
		return *x.First
	}
	return 0
}

func (x *UidGidExtent) GetLowerFirst() uint32 {
	if x != nil && x.LowerFirst != nil {
		return *x.LowerFirst
	}
	return 0
}

func (x *UidGidExtent) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type UsernsEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UidMap        []*UidGidExtent        `protobuf:"bytes,1,rep,name=uid_map,json=uidMap" json:"uid_map,omitempty"`
	GidMap        []*UidGidExtent        `protobuf:"bytes,2,rep,name=gid_map,json=gidMap" json:"gid_map,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsernsEntry) Reset() {
	*x = UsernsEntry{}
	mi := &file_userns_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsernsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsernsEntry) ProtoMessage() {}

func (x *UsernsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_userns_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsernsEntry.ProtoReflect.Descriptor instead.
func (*UsernsEntry) Descriptor() ([]byte, []int) {
	return file_userns_proto_rawDescGZIP(), []int{1}
}

func (x *UsernsEntry) GetUidMap() []*UidGidExtent {
	if x != nil {
		return x.UidMap
	}
	return nil
}

func (x *UsernsEntry) GetGidMap() []*UidGidExtent {
	if x != nil {
		return x.GidMap
	}
	return nil
}

var File_userns_proto protoreflect.FileDescriptor

const file_userns_proto_rawDesc = "" +
	"\n" +
	"\fuserns.proto\"]\n" +
	"\x0euid_gid_extent\x12\x14\n" +
	"\x05first\x18\x01 \x02(\rR\x05first\x12\x1f\n" +
	"\vlower_first\x18\x02 \x02(\rR\n" +
	"lowerFirst\x12\x14\n" +
	"\x05count\x18\x03 \x02(\rR\x05count\"b\n" +
	"\fuserns_entry\x12(\n" +
	"\auid_map\x18\x01 \x03(\v2\x0f.uid_gid_extentR\x06uidMap\x12(\n" +
	"\agid_map\x18\x02 \x03(\v2\x0f.uid_gid_extentR\x06gidMap"

var (
	file_userns_proto_rawDescOnce sync.Once
	file_userns_proto_rawDescData []byte
)

func file_userns_proto_rawDescGZIP() []byte {
	file_userns_proto_rawDescOnce.Do(func() {
		file_userns_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_userns_proto_rawDesc), len(file_userns_proto_rawDesc)))
	})
	return file_userns_proto_rawDescData
}

var file_userns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_userns_proto_goTypes = []any{
	(*UidGidExtent)(nil), // 0: uid_gid_extent
	(*UsernsEntry)(nil),  // 1: userns_entry
}
var file_userns_proto_depIdxs = []int32{
	0, // 0: userns_entry.uid_map:type_name -> uid_gid_extent
	0, // 1: userns_entry.gid_map:type_name -> uid_gid_extent
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userns_proto_init() }
func file_userns_proto_init() {
	if File_userns_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_userns_proto_rawDesc), len(file_userns_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userns_proto_goTypes,
		DependencyIndexes: file_userns_proto_depIdxs,
		MessageInfos:      file_userns_proto_msgTypes,
	}.Build()
	File_userns_proto = out.File
	file_userns_proto_goTypes = nil
	file_userns_proto_depIdxs = nil
}
