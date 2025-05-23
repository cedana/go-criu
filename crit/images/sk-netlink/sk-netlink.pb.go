// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: sk-netlink.proto

package sk_netlink

import (
	fown "github.com/cedana/go-criu/v7/crit/images/fown"
	_ "github.com/cedana/go-criu/v7/crit/images/opts"
	sk_opts "github.com/cedana/go-criu/v7/crit/images/sk-opts"
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

type NetlinkSkEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Ino           *uint32                `protobuf:"varint,2,req,name=ino" json:"ino,omitempty"`
	Protocol      *uint32                `protobuf:"varint,3,req,name=protocol" json:"protocol,omitempty"`
	State         *uint32                `protobuf:"varint,4,req,name=state" json:"state,omitempty"`
	Flags         *uint32                `protobuf:"varint,6,req,name=flags" json:"flags,omitempty"`
	Portid        *uint32                `protobuf:"varint,7,req,name=portid" json:"portid,omitempty"`
	Groups        []uint32               `protobuf:"varint,8,rep,name=groups" json:"groups,omitempty"`
	DstPortid     *uint32                `protobuf:"varint,9,req,name=dst_portid,json=dstPortid" json:"dst_portid,omitempty"`
	DstGroup      *uint32                `protobuf:"varint,10,req,name=dst_group,json=dstGroup" json:"dst_group,omitempty"`
	Fown          *fown.FownEntry        `protobuf:"bytes,11,req,name=fown" json:"fown,omitempty"`
	Opts          *sk_opts.SkOptsEntry   `protobuf:"bytes,12,req,name=opts" json:"opts,omitempty"`
	NsId          *uint32                `protobuf:"varint,13,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetlinkSkEntry) Reset() {
	*x = NetlinkSkEntry{}
	mi := &file_sk_netlink_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetlinkSkEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetlinkSkEntry) ProtoMessage() {}

func (x *NetlinkSkEntry) ProtoReflect() protoreflect.Message {
	mi := &file_sk_netlink_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetlinkSkEntry.ProtoReflect.Descriptor instead.
func (*NetlinkSkEntry) Descriptor() ([]byte, []int) {
	return file_sk_netlink_proto_rawDescGZIP(), []int{0}
}

func (x *NetlinkSkEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *NetlinkSkEntry) GetIno() uint32 {
	if x != nil && x.Ino != nil {
		return *x.Ino
	}
	return 0
}

func (x *NetlinkSkEntry) GetProtocol() uint32 {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return 0
}

func (x *NetlinkSkEntry) GetState() uint32 {
	if x != nil && x.State != nil {
		return *x.State
	}
	return 0
}

func (x *NetlinkSkEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *NetlinkSkEntry) GetPortid() uint32 {
	if x != nil && x.Portid != nil {
		return *x.Portid
	}
	return 0
}

func (x *NetlinkSkEntry) GetGroups() []uint32 {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *NetlinkSkEntry) GetDstPortid() uint32 {
	if x != nil && x.DstPortid != nil {
		return *x.DstPortid
	}
	return 0
}

func (x *NetlinkSkEntry) GetDstGroup() uint32 {
	if x != nil && x.DstGroup != nil {
		return *x.DstGroup
	}
	return 0
}

func (x *NetlinkSkEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *NetlinkSkEntry) GetOpts() *sk_opts.SkOptsEntry {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *NetlinkSkEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

var File_sk_netlink_proto protoreflect.FileDescriptor

const file_sk_netlink_proto_rawDesc = "" +
	"\n" +
	"\x10sk-netlink.proto\x1a\n" +
	"opts.proto\x1a\n" +
	"fown.proto\x1a\rsk-opts.proto\"\xc9\x02\n" +
	"\x10netlink_sk_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x10\n" +
	"\x03ino\x18\x02 \x02(\rR\x03ino\x12\x1a\n" +
	"\bprotocol\x18\x03 \x02(\rR\bprotocol\x12\x14\n" +
	"\x05state\x18\x04 \x02(\rR\x05state\x12\x1b\n" +
	"\x05flags\x18\x06 \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x16\n" +
	"\x06portid\x18\a \x02(\rR\x06portid\x12\x16\n" +
	"\x06groups\x18\b \x03(\rR\x06groups\x12\x1d\n" +
	"\n" +
	"dst_portid\x18\t \x02(\rR\tdstPortid\x12\x1b\n" +
	"\tdst_group\x18\n" +
	" \x02(\rR\bdstGroup\x12\x1f\n" +
	"\x04fown\x18\v \x02(\v2\v.fown_entryR\x04fown\x12\"\n" +
	"\x04opts\x18\f \x02(\v2\x0e.sk_opts_entryR\x04opts\x12\x13\n" +
	"\x05ns_id\x18\r \x01(\rR\x04nsId"

var (
	file_sk_netlink_proto_rawDescOnce sync.Once
	file_sk_netlink_proto_rawDescData []byte
)

func file_sk_netlink_proto_rawDescGZIP() []byte {
	file_sk_netlink_proto_rawDescOnce.Do(func() {
		file_sk_netlink_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sk_netlink_proto_rawDesc), len(file_sk_netlink_proto_rawDesc)))
	})
	return file_sk_netlink_proto_rawDescData
}

var file_sk_netlink_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sk_netlink_proto_goTypes = []any{
	(*NetlinkSkEntry)(nil),      // 0: netlink_sk_entry
	(*fown.FownEntry)(nil),      // 1: fown_entry
	(*sk_opts.SkOptsEntry)(nil), // 2: sk_opts_entry
}
var file_sk_netlink_proto_depIdxs = []int32{
	1, // 0: netlink_sk_entry.fown:type_name -> fown_entry
	2, // 1: netlink_sk_entry.opts:type_name -> sk_opts_entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sk_netlink_proto_init() }
func file_sk_netlink_proto_init() {
	if File_sk_netlink_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sk_netlink_proto_rawDesc), len(file_sk_netlink_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sk_netlink_proto_goTypes,
		DependencyIndexes: file_sk_netlink_proto_depIdxs,
		MessageInfos:      file_sk_netlink_proto_msgTypes,
	}.Build()
	File_sk_netlink_proto = out.File
	file_sk_netlink_proto_goTypes = nil
	file_sk_netlink_proto_depIdxs = nil
}
