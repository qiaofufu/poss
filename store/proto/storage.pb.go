// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/storage.proto

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

type MkdirParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Mode uint32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *MkdirParam) Reset() {
	*x = MkdirParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MkdirParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkdirParam) ProtoMessage() {}

func (x *MkdirParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkdirParam.ProtoReflect.Descriptor instead.
func (*MkdirParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{0}
}

func (x *MkdirParam) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MkdirParam) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type MkdirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MkdirResponse) Reset() {
	*x = MkdirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MkdirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MkdirResponse) ProtoMessage() {}

func (x *MkdirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MkdirResponse.ProtoReflect.Descriptor instead.
func (*MkdirResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{1}
}

func (x *MkdirResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type RmdirParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *RmdirParam) Reset() {
	*x = RmdirParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RmdirParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RmdirParam) ProtoMessage() {}

func (x *RmdirParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RmdirParam.ProtoReflect.Descriptor instead.
func (*RmdirParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{2}
}

func (x *RmdirParam) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type RmdirResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RmdirResponse) Reset() {
	*x = RmdirResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RmdirResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RmdirResponse) ProtoMessage() {}

func (x *RmdirResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RmdirResponse.ProtoReflect.Descriptor instead.
func (*RmdirResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{3}
}

func (x *RmdirResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type StatParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StatParam) Reset() {
	*x = StatParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatParam) ProtoMessage() {}

func (x *StatParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatParam.ProtoReflect.Descriptor instead.
func (*StatParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{4}
}

func (x *StatParam) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type StatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size    int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Mode    uint32 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	ModTime int64  `protobuf:"varint,4,opt,name=modTime,proto3" json:"modTime,omitempty"`
	IsDir   bool   `protobuf:"varint,5,opt,name=isDir,proto3" json:"isDir,omitempty"`
}

func (x *StatResponse) Reset() {
	*x = StatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponse) ProtoMessage() {}

func (x *StatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponse.ProtoReflect.Descriptor instead.
func (*StatResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{5}
}

func (x *StatResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StatResponse) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *StatResponse) GetModTime() int64 {
	if x != nil {
		return x.ModTime
	}
	return 0
}

func (x *StatResponse) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

type OpenParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flags int32  `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	Mode  uint32 `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *OpenParam) Reset() {
	*x = OpenParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenParam) ProtoMessage() {}

func (x *OpenParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenParam.ProtoReflect.Descriptor instead.
func (*OpenParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{6}
}

func (x *OpenParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OpenParam) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *OpenParam) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

type OpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Fd     int32  `protobuf:"varint,2,opt,name=fd,proto3" json:"fd,omitempty"`
}

func (x *OpenResponse) Reset() {
	*x = OpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenResponse) ProtoMessage() {}

func (x *OpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenResponse.ProtoReflect.Descriptor instead.
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{7}
}

func (x *OpenResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OpenResponse) GetFd() int32 {
	if x != nil {
		return x.Fd
	}
	return 0
}

type WriteParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd   int32  `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Buf  []byte `protobuf:"bytes,2,opt,name=buf,proto3" json:"buf,omitempty"`
	Size int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *WriteParam) Reset() {
	*x = WriteParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteParam) ProtoMessage() {}

func (x *WriteParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteParam.ProtoReflect.Descriptor instead.
func (*WriteParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{8}
}

func (x *WriteParam) GetFd() int32 {
	if x != nil {
		return x.Fd
	}
	return 0
}

func (x *WriteParam) GetBuf() []byte {
	if x != nil {
		return x.Buf
	}
	return nil
}

func (x *WriteParam) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{9}
}

func (x *WriteResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ReadParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd   int32 `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Size int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ReadParam) Reset() {
	*x = ReadParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadParam) ProtoMessage() {}

func (x *ReadParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadParam.ProtoReflect.Descriptor instead.
func (*ReadParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{10}
}

func (x *ReadParam) GetFd() int32 {
	if x != nil {
		return x.Fd
	}
	return 0
}

func (x *ReadParam) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Buf    []byte `protobuf:"bytes,2,opt,name=buf,proto3" json:"buf,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{11}
}

func (x *ReadResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReadResponse) GetBuf() []byte {
	if x != nil {
		return x.Buf
	}
	return nil
}

type DelParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd int32 `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
}

func (x *DelParam) Reset() {
	*x = DelParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelParam) ProtoMessage() {}

func (x *DelParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelParam.ProtoReflect.Descriptor instead.
func (*DelParam) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{12}
}

func (x *DelParam) GetFd() int32 {
	if x != nil {
		return x.Fd
	}
	return 0
}

type DelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DelResponse) Reset() {
	*x = DelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_storage_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResponse) ProtoMessage() {}

func (x *DelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_storage_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResponse.ProtoReflect.Descriptor instead.
func (*DelResponse) Descriptor() ([]byte, []int) {
	return file_proto_storage_proto_rawDescGZIP(), []int{13}
}

func (x *DelResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_proto_storage_proto protoreflect.FileDescriptor

var file_proto_storage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0a,
	0x4d, 0x6b, 0x64, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x22, 0x27, 0x0a, 0x0d, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x52,
	0x6d, 0x64, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x27, 0x0a,
	0x0d, 0x52, 0x6d, 0x64, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x7a, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73,
	0x44, 0x69, 0x72, 0x22, 0x49, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x36,
	0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x66, 0x64, 0x22, 0x42, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x66, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x75, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x62, 0x75, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x27, 0x0a, 0x0d, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x66, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x75, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x62, 0x75, 0x66, 0x22, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x66, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xdc, 0x02, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x6b, 0x64, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x6b, 0x64, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x52, 0x6d, 0x64, 0x69, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x6d, 0x64, 0x69, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6d, 0x64, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x2f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x12, 0x2a, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_storage_proto_rawDescOnce sync.Once
	file_proto_storage_proto_rawDescData = file_proto_storage_proto_rawDesc
)

func file_proto_storage_proto_rawDescGZIP() []byte {
	file_proto_storage_proto_rawDescOnce.Do(func() {
		file_proto_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_storage_proto_rawDescData)
	})
	return file_proto_storage_proto_rawDescData
}

var file_proto_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_storage_proto_goTypes = []interface{}{
	(*MkdirParam)(nil),    // 0: proto.MkdirParam
	(*MkdirResponse)(nil), // 1: proto.MkdirResponse
	(*RmdirParam)(nil),    // 2: proto.RmdirParam
	(*RmdirResponse)(nil), // 3: proto.RmdirResponse
	(*StatParam)(nil),     // 4: proto.StatParam
	(*StatResponse)(nil),  // 5: proto.StatResponse
	(*OpenParam)(nil),     // 6: proto.OpenParam
	(*OpenResponse)(nil),  // 7: proto.OpenResponse
	(*WriteParam)(nil),    // 8: proto.WriteParam
	(*WriteResponse)(nil), // 9: proto.WriteResponse
	(*ReadParam)(nil),     // 10: proto.ReadParam
	(*ReadResponse)(nil),  // 11: proto.ReadResponse
	(*DelParam)(nil),      // 12: proto.DelParam
	(*DelResponse)(nil),   // 13: proto.DelResponse
}
var file_proto_storage_proto_depIdxs = []int32{
	0,  // 0: proto.Storage.Mkdir:input_type -> proto.MkdirParam
	2,  // 1: proto.Storage.Rmdir:input_type -> proto.RmdirParam
	4,  // 2: proto.Storage.Stat:input_type -> proto.StatParam
	6,  // 3: proto.Storage.Open:input_type -> proto.OpenParam
	8,  // 4: proto.Storage.Write:input_type -> proto.WriteParam
	10, // 5: proto.Storage.Read:input_type -> proto.ReadParam
	12, // 6: proto.Storage.Del:input_type -> proto.DelParam
	1,  // 7: proto.Storage.Mkdir:output_type -> proto.MkdirResponse
	3,  // 8: proto.Storage.Rmdir:output_type -> proto.RmdirResponse
	5,  // 9: proto.Storage.Stat:output_type -> proto.StatResponse
	7,  // 10: proto.Storage.Open:output_type -> proto.OpenResponse
	9,  // 11: proto.Storage.Write:output_type -> proto.WriteResponse
	11, // 12: proto.Storage.Read:output_type -> proto.ReadResponse
	13, // 13: proto.Storage.Del:output_type -> proto.DelResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_storage_proto_init() }
func file_proto_storage_proto_init() {
	if File_proto_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MkdirParam); i {
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
		file_proto_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MkdirResponse); i {
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
		file_proto_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RmdirParam); i {
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
		file_proto_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RmdirResponse); i {
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
		file_proto_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatParam); i {
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
		file_proto_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatResponse); i {
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
		file_proto_storage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenParam); i {
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
		file_proto_storage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenResponse); i {
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
		file_proto_storage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteParam); i {
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
		file_proto_storage_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_proto_storage_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadParam); i {
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
		file_proto_storage_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_proto_storage_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelParam); i {
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
		file_proto_storage_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelResponse); i {
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
			RawDescriptor: file_proto_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_storage_proto_goTypes,
		DependencyIndexes: file_proto_storage_proto_depIdxs,
		MessageInfos:      file_proto_storage_proto_msgTypes,
	}.Build()
	File_proto_storage_proto = out.File
	file_proto_storage_proto_rawDesc = nil
	file_proto_storage_proto_goTypes = nil
	file_proto_storage_proto_depIdxs = nil
}
