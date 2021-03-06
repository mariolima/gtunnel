// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.16.0
// source: gTunnel.proto

package gtunnel

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{0}
}

type EndpointControlMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation   int32  `protobuf:"varint,1,opt,name=operation,proto3" json:"operation,omitempty"`
	EndpointID  string `protobuf:"bytes,2,opt,name=endpointID,proto3" json:"endpointID,omitempty"`
	TunnelID    string `protobuf:"bytes,3,opt,name=tunnelID,proto3" json:"tunnelID,omitempty"`
	ErrorStatus int32  `protobuf:"varint,4,opt,name=errorStatus,proto3" json:"errorStatus,omitempty"`
	LocalIp     uint32 `protobuf:"varint,5,opt,name=localIp,proto3" json:"localIp,omitempty"`
	LocalPort   uint32 `protobuf:"varint,6,opt,name=localPort,proto3" json:"localPort,omitempty"`
	RemoteIP    uint32 `protobuf:"varint,7,opt,name=remoteIP,proto3" json:"remoteIP,omitempty"`
	RemotePort  uint32 `protobuf:"varint,8,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
	Cmd         string `protobuf:"bytes,9,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Sc          []byte `protobuf:"bytes,10,opt,name=sc,proto3" json:"sc,omitempty"`
	CmdResult   []byte `protobuf:"bytes,11,opt,name=cmdResult,proto3" json:"cmdResult,omitempty"`
}

func (x *EndpointControlMessage) Reset() {
	*x = EndpointControlMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointControlMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointControlMessage) ProtoMessage() {}

func (x *EndpointControlMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointControlMessage.ProtoReflect.Descriptor instead.
func (*EndpointControlMessage) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointControlMessage) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *EndpointControlMessage) GetEndpointID() string {
	if x != nil {
		return x.EndpointID
	}
	return ""
}

func (x *EndpointControlMessage) GetTunnelID() string {
	if x != nil {
		return x.TunnelID
	}
	return ""
}

func (x *EndpointControlMessage) GetErrorStatus() int32 {
	if x != nil {
		return x.ErrorStatus
	}
	return 0
}

func (x *EndpointControlMessage) GetLocalIp() uint32 {
	if x != nil {
		return x.LocalIp
	}
	return 0
}

func (x *EndpointControlMessage) GetLocalPort() uint32 {
	if x != nil {
		return x.LocalPort
	}
	return 0
}

func (x *EndpointControlMessage) GetRemoteIP() uint32 {
	if x != nil {
		return x.RemoteIP
	}
	return 0
}

func (x *EndpointControlMessage) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

func (x *EndpointControlMessage) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *EndpointControlMessage) GetSc() []byte {
	if x != nil {
		return x.Sc
	}
	return nil
}

func (x *EndpointControlMessage) GetCmdResult() []byte {
	if x != nil {
		return x.CmdResult
	}
	return nil
}

type EndpointTaskMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointID string `protobuf:"bytes,1,opt,name=endpointID,proto3" json:"endpointID,omitempty"`
	Status     int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	TaskId     string `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
	TaskType   int32  `protobuf:"varint,4,opt,name=taskType,proto3" json:"taskType,omitempty"`
	Cmd        string `protobuf:"bytes,5,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Data       []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	Chunk      int32  `protobuf:"varint,7,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Size       int32  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *EndpointTaskMessage) Reset() {
	*x = EndpointTaskMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointTaskMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointTaskMessage) ProtoMessage() {}

func (x *EndpointTaskMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointTaskMessage.ProtoReflect.Descriptor instead.
func (*EndpointTaskMessage) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointTaskMessage) GetEndpointID() string {
	if x != nil {
		return x.EndpointID
	}
	return ""
}

func (x *EndpointTaskMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *EndpointTaskMessage) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *EndpointTaskMessage) GetTaskType() int32 {
	if x != nil {
		return x.TaskType
	}
	return 0
}

func (x *EndpointTaskMessage) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *EndpointTaskMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EndpointTaskMessage) GetChunk() int32 {
	if x != nil {
		return x.Chunk
	}
	return 0
}

func (x *EndpointTaskMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type EndpointTaskFileTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk int32  `protobuf:"varint,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Size  int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EndpointTaskFileTransfer) Reset() {
	*x = EndpointTaskFileTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointTaskFileTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointTaskFileTransfer) ProtoMessage() {}

func (x *EndpointTaskFileTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointTaskFileTransfer.ProtoReflect.Descriptor instead.
func (*EndpointTaskFileTransfer) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointTaskFileTransfer) GetChunk() int32 {
	if x != nil {
		return x.Chunk
	}
	return 0
}

func (x *EndpointTaskFileTransfer) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *EndpointTaskFileTransfer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EndpointTaskCMD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd    string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Result []byte `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *EndpointTaskCMD) Reset() {
	*x = EndpointTaskCMD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointTaskCMD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointTaskCMD) ProtoMessage() {}

func (x *EndpointTaskCMD) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointTaskCMD.ProtoReflect.Descriptor instead.
func (*EndpointTaskCMD) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointTaskCMD) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *EndpointTaskCMD) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type TunnelControlMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation    int32  `protobuf:"varint,1,opt,name=operation,proto3" json:"operation,omitempty"`
	ErrorStatus  int32  `protobuf:"varint,2,opt,name=errorStatus,proto3" json:"errorStatus,omitempty"`
	EndpointID   string `protobuf:"bytes,3,opt,name=endpointID,proto3" json:"endpointID,omitempty"`
	TunnelID     string `protobuf:"bytes,4,opt,name=tunnelID,proto3" json:"tunnelID,omitempty"`
	ConnectionID int32  `protobuf:"varint,5,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
}

func (x *TunnelControlMessage) Reset() {
	*x = TunnelControlMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelControlMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelControlMessage) ProtoMessage() {}

func (x *TunnelControlMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelControlMessage.ProtoReflect.Descriptor instead.
func (*TunnelControlMessage) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{5}
}

func (x *TunnelControlMessage) GetOperation() int32 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *TunnelControlMessage) GetErrorStatus() int32 {
	if x != nil {
		return x.ErrorStatus
	}
	return 0
}

func (x *TunnelControlMessage) GetEndpointID() string {
	if x != nil {
		return x.EndpointID
	}
	return ""
}

func (x *TunnelControlMessage) GetTunnelID() string {
	if x != nil {
		return x.TunnelID
	}
	return ""
}

func (x *TunnelControlMessage) GetConnectionID() int32 {
	if x != nil {
		return x.ConnectionID
	}
	return 0
}

type BytesMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointID   string `protobuf:"bytes,1,opt,name=endpointID,proto3" json:"endpointID,omitempty"`
	TunnelID     string `protobuf:"bytes,2,opt,name=tunnelID,proto3" json:"tunnelID,omitempty"`
	ConnectionID int32  `protobuf:"varint,3,opt,name=connectionID,proto3" json:"connectionID,omitempty"`
	Content      []byte `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *BytesMessage) Reset() {
	*x = BytesMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gTunnel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesMessage) ProtoMessage() {}

func (x *BytesMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gTunnel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesMessage.ProtoReflect.Descriptor instead.
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return file_gTunnel_proto_rawDescGZIP(), []int{6}
}

func (x *BytesMessage) GetEndpointID() string {
	if x != nil {
		return x.EndpointID
	}
	return ""
}

func (x *BytesMessage) GetTunnelID() string {
	if x != nil {
		return x.TunnelID
	}
	return ""
}

func (x *BytesMessage) GetConnectionID() int32 {
	if x != nil {
		return x.ConnectionID
	}
	return 0
}

func (x *BytesMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_gTunnel_proto protoreflect.FileDescriptor

var file_gTunnel_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0xc8, 0x02, 0x0a, 0x16, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x49, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x73, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x63, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xd1, 0x01, 0x0a,
	0x13, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x58, 0x0a, 0x18, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3b, 0x0a, 0x0f, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x4d, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x14, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0x88, 0x01, 0x0a, 0x0c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xfb, 0x02, 0x0a, 0x07,
	0x47, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x63, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x67, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4c, 0x0a,
	0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x15,
	0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x14, 0x0a, 0x07, 0x67, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x07, 0x47, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gTunnel_proto_rawDescOnce sync.Once
	file_gTunnel_proto_rawDescData = file_gTunnel_proto_rawDesc
)

func file_gTunnel_proto_rawDescGZIP() []byte {
	file_gTunnel_proto_rawDescOnce.Do(func() {
		file_gTunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_gTunnel_proto_rawDescData)
	})
	return file_gTunnel_proto_rawDescData
}

var file_gTunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gTunnel_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: gtunnel.empty
	(*EndpointControlMessage)(nil),   // 1: gtunnel.EndpointControlMessage
	(*EndpointTaskMessage)(nil),      // 2: gtunnel.EndpointTaskMessage
	(*EndpointTaskFileTransfer)(nil), // 3: gtunnel.EndpointTaskFileTransfer
	(*EndpointTaskCMD)(nil),          // 4: gtunnel.EndpointTaskCMD
	(*TunnelControlMessage)(nil),     // 5: gtunnel.TunnelControlMessage
	(*BytesMessage)(nil),             // 6: gtunnel.BytesMessage
}
var file_gTunnel_proto_depIdxs = []int32{
	1, // 0: gtunnel.GTunnel.CreateEndpointControlStream:input_type -> gtunnel.EndpointControlMessage
	5, // 1: gtunnel.GTunnel.CreateTunnelControlStream:input_type -> gtunnel.TunnelControlMessage
	6, // 2: gtunnel.GTunnel.CreateConnectionStream:input_type -> gtunnel.BytesMessage
	2, // 3: gtunnel.GTunnel.CreateEndpointTaskStream:input_type -> gtunnel.EndpointTaskMessage
	1, // 4: gtunnel.GTunnel.CreateEndpointControlStream:output_type -> gtunnel.EndpointControlMessage
	5, // 5: gtunnel.GTunnel.CreateTunnelControlStream:output_type -> gtunnel.TunnelControlMessage
	6, // 6: gtunnel.GTunnel.CreateConnectionStream:output_type -> gtunnel.BytesMessage
	2, // 7: gtunnel.GTunnel.CreateEndpointTaskStream:output_type -> gtunnel.EndpointTaskMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gTunnel_proto_init() }
func file_gTunnel_proto_init() {
	if File_gTunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gTunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_gTunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointControlMessage); i {
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
		file_gTunnel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointTaskMessage); i {
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
		file_gTunnel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointTaskFileTransfer); i {
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
		file_gTunnel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointTaskCMD); i {
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
		file_gTunnel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelControlMessage); i {
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
		file_gTunnel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesMessage); i {
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
			RawDescriptor: file_gTunnel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gTunnel_proto_goTypes,
		DependencyIndexes: file_gTunnel_proto_depIdxs,
		MessageInfos:      file_gTunnel_proto_msgTypes,
	}.Build()
	File_gTunnel_proto = out.File
	file_gTunnel_proto_rawDesc = nil
	file_gTunnel_proto_goTypes = nil
	file_gTunnel_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GTunnelClient is the client API for GTunnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GTunnelClient interface {
	// Gets a stream of control messages from the server
	CreateEndpointControlStream(ctx context.Context, in *EndpointControlMessage, opts ...grpc.CallOption) (GTunnel_CreateEndpointControlStreamClient, error)
	CreateTunnelControlStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateTunnelControlStreamClient, error)
	// Bidirectional stream for data
	CreateConnectionStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateConnectionStreamClient, error)
	CreateEndpointTaskStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateEndpointTaskStreamClient, error)
}

type gTunnelClient struct {
	cc grpc.ClientConnInterface
}

func NewGTunnelClient(cc grpc.ClientConnInterface) GTunnelClient {
	return &gTunnelClient{cc}
}

func (c *gTunnelClient) CreateEndpointControlStream(ctx context.Context, in *EndpointControlMessage, opts ...grpc.CallOption) (GTunnel_CreateEndpointControlStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GTunnel_serviceDesc.Streams[0], "/gtunnel.GTunnel/CreateEndpointControlStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gTunnelCreateEndpointControlStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GTunnel_CreateEndpointControlStreamClient interface {
	Recv() (*EndpointControlMessage, error)
	grpc.ClientStream
}

type gTunnelCreateEndpointControlStreamClient struct {
	grpc.ClientStream
}

func (x *gTunnelCreateEndpointControlStreamClient) Recv() (*EndpointControlMessage, error) {
	m := new(EndpointControlMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gTunnelClient) CreateTunnelControlStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateTunnelControlStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GTunnel_serviceDesc.Streams[1], "/gtunnel.GTunnel/CreateTunnelControlStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gTunnelCreateTunnelControlStreamClient{stream}
	return x, nil
}

type GTunnel_CreateTunnelControlStreamClient interface {
	Send(*TunnelControlMessage) error
	Recv() (*TunnelControlMessage, error)
	grpc.ClientStream
}

type gTunnelCreateTunnelControlStreamClient struct {
	grpc.ClientStream
}

func (x *gTunnelCreateTunnelControlStreamClient) Send(m *TunnelControlMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gTunnelCreateTunnelControlStreamClient) Recv() (*TunnelControlMessage, error) {
	m := new(TunnelControlMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gTunnelClient) CreateConnectionStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateConnectionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GTunnel_serviceDesc.Streams[2], "/gtunnel.GTunnel/CreateConnectionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gTunnelCreateConnectionStreamClient{stream}
	return x, nil
}

type GTunnel_CreateConnectionStreamClient interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ClientStream
}

type gTunnelCreateConnectionStreamClient struct {
	grpc.ClientStream
}

func (x *gTunnelCreateConnectionStreamClient) Send(m *BytesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gTunnelCreateConnectionStreamClient) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gTunnelClient) CreateEndpointTaskStream(ctx context.Context, opts ...grpc.CallOption) (GTunnel_CreateEndpointTaskStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GTunnel_serviceDesc.Streams[3], "/gtunnel.GTunnel/CreateEndpointTaskStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gTunnelCreateEndpointTaskStreamClient{stream}
	return x, nil
}

type GTunnel_CreateEndpointTaskStreamClient interface {
	Send(*EndpointTaskMessage) error
	Recv() (*EndpointTaskMessage, error)
	grpc.ClientStream
}

type gTunnelCreateEndpointTaskStreamClient struct {
	grpc.ClientStream
}

func (x *gTunnelCreateEndpointTaskStreamClient) Send(m *EndpointTaskMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gTunnelCreateEndpointTaskStreamClient) Recv() (*EndpointTaskMessage, error) {
	m := new(EndpointTaskMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GTunnelServer is the server API for GTunnel service.
type GTunnelServer interface {
	// Gets a stream of control messages from the server
	CreateEndpointControlStream(*EndpointControlMessage, GTunnel_CreateEndpointControlStreamServer) error
	CreateTunnelControlStream(GTunnel_CreateTunnelControlStreamServer) error
	// Bidirectional stream for data
	CreateConnectionStream(GTunnel_CreateConnectionStreamServer) error
	CreateEndpointTaskStream(GTunnel_CreateEndpointTaskStreamServer) error
}

// UnimplementedGTunnelServer can be embedded to have forward compatible implementations.
type UnimplementedGTunnelServer struct {
}

func (*UnimplementedGTunnelServer) CreateEndpointControlStream(*EndpointControlMessage, GTunnel_CreateEndpointControlStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEndpointControlStream not implemented")
}
func (*UnimplementedGTunnelServer) CreateTunnelControlStream(GTunnel_CreateTunnelControlStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateTunnelControlStream not implemented")
}
func (*UnimplementedGTunnelServer) CreateConnectionStream(GTunnel_CreateConnectionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateConnectionStream not implemented")
}
func (*UnimplementedGTunnelServer) CreateEndpointTaskStream(GTunnel_CreateEndpointTaskStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateEndpointTaskStream not implemented")
}

func RegisterGTunnelServer(s *grpc.Server, srv GTunnelServer) {
	s.RegisterService(&_GTunnel_serviceDesc, srv)
}

func _GTunnel_CreateEndpointControlStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EndpointControlMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GTunnelServer).CreateEndpointControlStream(m, &gTunnelCreateEndpointControlStreamServer{stream})
}

type GTunnel_CreateEndpointControlStreamServer interface {
	Send(*EndpointControlMessage) error
	grpc.ServerStream
}

type gTunnelCreateEndpointControlStreamServer struct {
	grpc.ServerStream
}

func (x *gTunnelCreateEndpointControlStreamServer) Send(m *EndpointControlMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _GTunnel_CreateTunnelControlStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GTunnelServer).CreateTunnelControlStream(&gTunnelCreateTunnelControlStreamServer{stream})
}

type GTunnel_CreateTunnelControlStreamServer interface {
	Send(*TunnelControlMessage) error
	Recv() (*TunnelControlMessage, error)
	grpc.ServerStream
}

type gTunnelCreateTunnelControlStreamServer struct {
	grpc.ServerStream
}

func (x *gTunnelCreateTunnelControlStreamServer) Send(m *TunnelControlMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gTunnelCreateTunnelControlStreamServer) Recv() (*TunnelControlMessage, error) {
	m := new(TunnelControlMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GTunnel_CreateConnectionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GTunnelServer).CreateConnectionStream(&gTunnelCreateConnectionStreamServer{stream})
}

type GTunnel_CreateConnectionStreamServer interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ServerStream
}

type gTunnelCreateConnectionStreamServer struct {
	grpc.ServerStream
}

func (x *gTunnelCreateConnectionStreamServer) Send(m *BytesMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gTunnelCreateConnectionStreamServer) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GTunnel_CreateEndpointTaskStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GTunnelServer).CreateEndpointTaskStream(&gTunnelCreateEndpointTaskStreamServer{stream})
}

type GTunnel_CreateEndpointTaskStreamServer interface {
	Send(*EndpointTaskMessage) error
	Recv() (*EndpointTaskMessage, error)
	grpc.ServerStream
}

type gTunnelCreateEndpointTaskStreamServer struct {
	grpc.ServerStream
}

func (x *gTunnelCreateEndpointTaskStreamServer) Send(m *EndpointTaskMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gTunnelCreateEndpointTaskStreamServer) Recv() (*EndpointTaskMessage, error) {
	m := new(EndpointTaskMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GTunnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gtunnel.GTunnel",
	HandlerType: (*GTunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateEndpointControlStream",
			Handler:       _GTunnel_CreateEndpointControlStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateTunnelControlStream",
			Handler:       _GTunnel_CreateTunnelControlStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateConnectionStream",
			Handler:       _GTunnel_CreateConnectionStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateEndpointTaskStream",
			Handler:       _GTunnel_CreateEndpointTaskStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gTunnel.proto",
}
