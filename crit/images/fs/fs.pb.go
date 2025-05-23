// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: fs.proto

package fs

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

type FsEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CwdId         *uint32                `protobuf:"varint,1,req,name=cwd_id,json=cwdId" json:"cwd_id,omitempty"`
	RootId        *uint32                `protobuf:"varint,2,req,name=root_id,json=rootId" json:"root_id,omitempty"`
	Umask         *uint32                `protobuf:"varint,3,opt,name=umask" json:"umask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FsEntry) Reset() {
	*x = FsEntry{}
	mi := &file_fs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FsEntry) ProtoMessage() {}

func (x *FsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_fs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FsEntry.ProtoReflect.Descriptor instead.
func (*FsEntry) Descriptor() ([]byte, []int) {
	return file_fs_proto_rawDescGZIP(), []int{0}
}

func (x *FsEntry) GetCwdId() uint32 {
	if x != nil && x.CwdId != nil {
		return *x.CwdId
	}
	return 0
}

func (x *FsEntry) GetRootId() uint32 {
	if x != nil && x.RootId != nil {
		return *x.RootId
	}
	return 0
}

func (x *FsEntry) GetUmask() uint32 {
	if x != nil && x.Umask != nil {
		return *x.Umask
	}
	return 0
}

var File_fs_proto protoreflect.FileDescriptor

const file_fs_proto_rawDesc = "" +
	"\n" +
	"\bfs.proto\"P\n" +
	"\bfs_entry\x12\x15\n" +
	"\x06cwd_id\x18\x01 \x02(\rR\x05cwdId\x12\x17\n" +
	"\aroot_id\x18\x02 \x02(\rR\x06rootId\x12\x14\n" +
	"\x05umask\x18\x03 \x01(\rR\x05umask"

var (
	file_fs_proto_rawDescOnce sync.Once
	file_fs_proto_rawDescData []byte
)

func file_fs_proto_rawDescGZIP() []byte {
	file_fs_proto_rawDescOnce.Do(func() {
		file_fs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fs_proto_rawDesc), len(file_fs_proto_rawDesc)))
	})
	return file_fs_proto_rawDescData
}

var file_fs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fs_proto_goTypes = []any{
	(*FsEntry)(nil), // 0: fs_entry
}
var file_fs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fs_proto_init() }
func file_fs_proto_init() {
	if File_fs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fs_proto_rawDesc), len(file_fs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fs_proto_goTypes,
		DependencyIndexes: file_fs_proto_depIdxs,
		MessageInfos:      file_fs_proto_msgTypes,
	}.Build()
	File_fs_proto = out.File
	file_fs_proto_goTypes = nil
	file_fs_proto_depIdxs = nil
}
