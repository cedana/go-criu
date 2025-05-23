// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.28.0
// source: packet-sock.proto

package packet_sock

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

type PacketMclist struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Index         *uint32                `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	Type          *uint32                `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	Addr          []byte                 `protobuf:"bytes,3,req,name=addr" json:"addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PacketMclist) Reset() {
	*x = PacketMclist{}
	mi := &file_packet_sock_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketMclist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketMclist) ProtoMessage() {}

func (x *PacketMclist) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketMclist.ProtoReflect.Descriptor instead.
func (*PacketMclist) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{0}
}

func (x *PacketMclist) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *PacketMclist) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *PacketMclist) GetAddr() []byte {
	if x != nil {
		return x.Addr
	}
	return nil
}

type PacketRing struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlockSize     *uint32                `protobuf:"varint,1,req,name=block_size,json=blockSize" json:"block_size,omitempty"`
	BlockNr       *uint32                `protobuf:"varint,2,req,name=block_nr,json=blockNr" json:"block_nr,omitempty"`
	FrameSize     *uint32                `protobuf:"varint,3,req,name=frame_size,json=frameSize" json:"frame_size,omitempty"`
	FrameNr       *uint32                `protobuf:"varint,4,req,name=frame_nr,json=frameNr" json:"frame_nr,omitempty"`
	RetireTmo     *uint32                `protobuf:"varint,5,req,name=retire_tmo,json=retireTmo" json:"retire_tmo,omitempty"`
	SizeofPriv    *uint32                `protobuf:"varint,6,req,name=sizeof_priv,json=sizeofPriv" json:"sizeof_priv,omitempty"`
	Features      *uint32                `protobuf:"varint,7,req,name=features" json:"features,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PacketRing) Reset() {
	*x = PacketRing{}
	mi := &file_packet_sock_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketRing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketRing) ProtoMessage() {}

func (x *PacketRing) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketRing.ProtoReflect.Descriptor instead.
func (*PacketRing) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{1}
}

func (x *PacketRing) GetBlockSize() uint32 {
	if x != nil && x.BlockSize != nil {
		return *x.BlockSize
	}
	return 0
}

func (x *PacketRing) GetBlockNr() uint32 {
	if x != nil && x.BlockNr != nil {
		return *x.BlockNr
	}
	return 0
}

func (x *PacketRing) GetFrameSize() uint32 {
	if x != nil && x.FrameSize != nil {
		return *x.FrameSize
	}
	return 0
}

func (x *PacketRing) GetFrameNr() uint32 {
	if x != nil && x.FrameNr != nil {
		return *x.FrameNr
	}
	return 0
}

func (x *PacketRing) GetRetireTmo() uint32 {
	if x != nil && x.RetireTmo != nil {
		return *x.RetireTmo
	}
	return 0
}

func (x *PacketRing) GetSizeofPriv() uint32 {
	if x != nil && x.SizeofPriv != nil {
		return *x.SizeofPriv
	}
	return 0
}

func (x *PacketRing) GetFeatures() uint32 {
	if x != nil && x.Features != nil {
		return *x.Features
	}
	return 0
}

type PacketSockEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *uint32                `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type          *uint32                `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	Protocol      *uint32                `protobuf:"varint,3,req,name=protocol" json:"protocol,omitempty"`
	Flags         *uint32                `protobuf:"varint,4,req,name=flags" json:"flags,omitempty"`
	Ifindex       *uint32                `protobuf:"varint,5,req,name=ifindex" json:"ifindex,omitempty"`
	Fown          *fown.FownEntry        `protobuf:"bytes,6,req,name=fown" json:"fown,omitempty"`
	Opts          *sk_opts.SkOptsEntry   `protobuf:"bytes,7,req,name=opts" json:"opts,omitempty"`
	Version       *uint32                `protobuf:"varint,8,req,name=version" json:"version,omitempty"`
	Reserve       *uint32                `protobuf:"varint,9,req,name=reserve" json:"reserve,omitempty"`
	AuxData       *bool                  `protobuf:"varint,10,req,name=aux_data,json=auxData" json:"aux_data,omitempty"`
	OrigDev       *bool                  `protobuf:"varint,11,req,name=orig_dev,json=origDev" json:"orig_dev,omitempty"`
	VnetHdr       *bool                  `protobuf:"varint,12,req,name=vnet_hdr,json=vnetHdr" json:"vnet_hdr,omitempty"`
	Loss          *bool                  `protobuf:"varint,13,req,name=loss" json:"loss,omitempty"`
	Timestamp     *uint32                `protobuf:"varint,14,req,name=timestamp" json:"timestamp,omitempty"`
	CopyThresh    *uint32                `protobuf:"varint,15,req,name=copy_thresh,json=copyThresh" json:"copy_thresh,omitempty"`
	Mclist        []*PacketMclist        `protobuf:"bytes,16,rep,name=mclist" json:"mclist,omitempty"`
	Fanout        *uint32                `protobuf:"varint,17,opt,name=fanout,def=4294967295" json:"fanout,omitempty"`
	RxRing        *PacketRing            `protobuf:"bytes,18,opt,name=rx_ring,json=rxRing" json:"rx_ring,omitempty"`
	TxRing        *PacketRing            `protobuf:"bytes,19,opt,name=tx_ring,json=txRing" json:"tx_ring,omitempty"`
	NsId          *uint32                `protobuf:"varint,20,opt,name=ns_id,json=nsId" json:"ns_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// Default values for PacketSockEntry fields.
const (
	Default_PacketSockEntry_Fanout = uint32(4294967295)
)

func (x *PacketSockEntry) Reset() {
	*x = PacketSockEntry{}
	mi := &file_packet_sock_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PacketSockEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketSockEntry) ProtoMessage() {}

func (x *PacketSockEntry) ProtoReflect() protoreflect.Message {
	mi := &file_packet_sock_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketSockEntry.ProtoReflect.Descriptor instead.
func (*PacketSockEntry) Descriptor() ([]byte, []int) {
	return file_packet_sock_proto_rawDescGZIP(), []int{2}
}

func (x *PacketSockEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PacketSockEntry) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *PacketSockEntry) GetProtocol() uint32 {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return 0
}

func (x *PacketSockEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *PacketSockEntry) GetIfindex() uint32 {
	if x != nil && x.Ifindex != nil {
		return *x.Ifindex
	}
	return 0
}

func (x *PacketSockEntry) GetFown() *fown.FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *PacketSockEntry) GetOpts() *sk_opts.SkOptsEntry {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *PacketSockEntry) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *PacketSockEntry) GetReserve() uint32 {
	if x != nil && x.Reserve != nil {
		return *x.Reserve
	}
	return 0
}

func (x *PacketSockEntry) GetAuxData() bool {
	if x != nil && x.AuxData != nil {
		return *x.AuxData
	}
	return false
}

func (x *PacketSockEntry) GetOrigDev() bool {
	if x != nil && x.OrigDev != nil {
		return *x.OrigDev
	}
	return false
}

func (x *PacketSockEntry) GetVnetHdr() bool {
	if x != nil && x.VnetHdr != nil {
		return *x.VnetHdr
	}
	return false
}

func (x *PacketSockEntry) GetLoss() bool {
	if x != nil && x.Loss != nil {
		return *x.Loss
	}
	return false
}

func (x *PacketSockEntry) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *PacketSockEntry) GetCopyThresh() uint32 {
	if x != nil && x.CopyThresh != nil {
		return *x.CopyThresh
	}
	return 0
}

func (x *PacketSockEntry) GetMclist() []*PacketMclist {
	if x != nil {
		return x.Mclist
	}
	return nil
}

func (x *PacketSockEntry) GetFanout() uint32 {
	if x != nil && x.Fanout != nil {
		return *x.Fanout
	}
	return Default_PacketSockEntry_Fanout
}

func (x *PacketSockEntry) GetRxRing() *PacketRing {
	if x != nil {
		return x.RxRing
	}
	return nil
}

func (x *PacketSockEntry) GetTxRing() *PacketRing {
	if x != nil {
		return x.TxRing
	}
	return nil
}

func (x *PacketSockEntry) GetNsId() uint32 {
	if x != nil && x.NsId != nil {
		return *x.NsId
	}
	return 0
}

var File_packet_sock_proto protoreflect.FileDescriptor

const file_packet_sock_proto_rawDesc = "" +
	"\n" +
	"\x11packet-sock.proto\x1a\n" +
	"opts.proto\x1a\n" +
	"fown.proto\x1a\rsk-opts.proto\"M\n" +
	"\rpacket_mclist\x12\x14\n" +
	"\x05index\x18\x01 \x02(\rR\x05index\x12\x12\n" +
	"\x04type\x18\x02 \x02(\rR\x04type\x12\x12\n" +
	"\x04addr\x18\x03 \x02(\fR\x04addr\"\xdd\x01\n" +
	"\vpacket_ring\x12\x1d\n" +
	"\n" +
	"block_size\x18\x01 \x02(\rR\tblockSize\x12\x19\n" +
	"\bblock_nr\x18\x02 \x02(\rR\ablockNr\x12\x1d\n" +
	"\n" +
	"frame_size\x18\x03 \x02(\rR\tframeSize\x12\x19\n" +
	"\bframe_nr\x18\x04 \x02(\rR\aframeNr\x12\x1d\n" +
	"\n" +
	"retire_tmo\x18\x05 \x02(\rR\tretireTmo\x12\x1f\n" +
	"\vsizeof_priv\x18\x06 \x02(\rR\n" +
	"sizeofPriv\x12\x1a\n" +
	"\bfeatures\x18\a \x02(\rR\bfeatures\"\xd6\x04\n" +
	"\x11packet_sock_entry\x12\x0e\n" +
	"\x02id\x18\x01 \x02(\rR\x02id\x12\x12\n" +
	"\x04type\x18\x02 \x02(\rR\x04type\x12\x1a\n" +
	"\bprotocol\x18\x03 \x02(\rR\bprotocol\x12\x1b\n" +
	"\x05flags\x18\x04 \x02(\rB\x05\x92~\x02\b\x01R\x05flags\x12\x18\n" +
	"\aifindex\x18\x05 \x02(\rR\aifindex\x12\x1f\n" +
	"\x04fown\x18\x06 \x02(\v2\v.fown_entryR\x04fown\x12\"\n" +
	"\x04opts\x18\a \x02(\v2\x0e.sk_opts_entryR\x04opts\x12\x18\n" +
	"\aversion\x18\b \x02(\rR\aversion\x12\x18\n" +
	"\areserve\x18\t \x02(\rR\areserve\x12\x19\n" +
	"\baux_data\x18\n" +
	" \x02(\bR\aauxData\x12\x19\n" +
	"\borig_dev\x18\v \x02(\bR\aorigDev\x12\x19\n" +
	"\bvnet_hdr\x18\f \x02(\bR\avnetHdr\x12\x12\n" +
	"\x04loss\x18\r \x02(\bR\x04loss\x12\x1c\n" +
	"\ttimestamp\x18\x0e \x02(\rR\ttimestamp\x12\x1f\n" +
	"\vcopy_thresh\x18\x0f \x02(\rR\n" +
	"copyThresh\x12&\n" +
	"\x06mclist\x18\x10 \x03(\v2\x0e.packet_mclistR\x06mclist\x12\"\n" +
	"\x06fanout\x18\x11 \x01(\r:\n" +
	"4294967295R\x06fanout\x12%\n" +
	"\arx_ring\x18\x12 \x01(\v2\f.packet_ringR\x06rxRing\x12%\n" +
	"\atx_ring\x18\x13 \x01(\v2\f.packet_ringR\x06txRing\x12\x13\n" +
	"\x05ns_id\x18\x14 \x01(\rR\x04nsId"

var (
	file_packet_sock_proto_rawDescOnce sync.Once
	file_packet_sock_proto_rawDescData []byte
)

func file_packet_sock_proto_rawDescGZIP() []byte {
	file_packet_sock_proto_rawDescOnce.Do(func() {
		file_packet_sock_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_packet_sock_proto_rawDesc), len(file_packet_sock_proto_rawDesc)))
	})
	return file_packet_sock_proto_rawDescData
}

var file_packet_sock_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packet_sock_proto_goTypes = []any{
	(*PacketMclist)(nil),        // 0: packet_mclist
	(*PacketRing)(nil),          // 1: packet_ring
	(*PacketSockEntry)(nil),     // 2: packet_sock_entry
	(*fown.FownEntry)(nil),      // 3: fown_entry
	(*sk_opts.SkOptsEntry)(nil), // 4: sk_opts_entry
}
var file_packet_sock_proto_depIdxs = []int32{
	3, // 0: packet_sock_entry.fown:type_name -> fown_entry
	4, // 1: packet_sock_entry.opts:type_name -> sk_opts_entry
	0, // 2: packet_sock_entry.mclist:type_name -> packet_mclist
	1, // 3: packet_sock_entry.rx_ring:type_name -> packet_ring
	1, // 4: packet_sock_entry.tx_ring:type_name -> packet_ring
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_packet_sock_proto_init() }
func file_packet_sock_proto_init() {
	if File_packet_sock_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_packet_sock_proto_rawDesc), len(file_packet_sock_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_packet_sock_proto_goTypes,
		DependencyIndexes: file_packet_sock_proto_depIdxs,
		MessageInfos:      file_packet_sock_proto_msgTypes,
	}.Build()
	File_packet_sock_proto = out.File
	file_packet_sock_proto_goTypes = nil
	file_packet_sock_proto_depIdxs = nil
}
