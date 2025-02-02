// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: vmapi.proto

package vmproto

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

type ExecCommandStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args    []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ExecCommandStreamRequest) Reset() {
	*x = ExecCommandStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandStreamRequest) ProtoMessage() {}

func (x *ExecCommandStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandStreamRequest.ProtoReflect.Descriptor instead.
func (*ExecCommandStreamRequest) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{0}
}

func (x *ExecCommandStreamRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExecCommandStreamRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type ExecCommandStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*ExecCommandStreamResponse_Output
	//	*ExecCommandStreamResponse_Log
	Content isExecCommandStreamResponse_Content `protobuf_oneof:"content"`
}

func (x *ExecCommandStreamResponse) Reset() {
	*x = ExecCommandStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandStreamResponse) ProtoMessage() {}

func (x *ExecCommandStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandStreamResponse.ProtoReflect.Descriptor instead.
func (*ExecCommandStreamResponse) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{1}
}

func (m *ExecCommandStreamResponse) GetContent() isExecCommandStreamResponse_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *ExecCommandStreamResponse) GetOutput() []byte {
	if x, ok := x.GetContent().(*ExecCommandStreamResponse_Output); ok {
		return x.Output
	}
	return nil
}

func (x *ExecCommandStreamResponse) GetLog() *Log {
	if x, ok := x.GetContent().(*ExecCommandStreamResponse_Log); ok {
		return x.Log
	}
	return nil
}

type isExecCommandStreamResponse_Content interface {
	isExecCommandStreamResponse_Content()
}

type ExecCommandStreamResponse_Output struct {
	Output []byte `protobuf:"bytes,1,opt,name=output,proto3,oneof"`
}

type ExecCommandStreamResponse_Log struct {
	Log *Log `protobuf:"bytes,2,opt,name=log,proto3,oneof"`
}

func (*ExecCommandStreamResponse_Output) isExecCommandStreamResponse_Content() {}

func (*ExecCommandStreamResponse_Log) isExecCommandStreamResponse_Content() {}

type ExecCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args    []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ExecCommandRequest) Reset() {
	*x = ExecCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandRequest) ProtoMessage() {}

func (x *ExecCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecCommandRequest) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{2}
}

func (x *ExecCommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExecCommandRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type ExecCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ExecCommandResponse) Reset() {
	*x = ExecCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandResponse) ProtoMessage() {}

func (x *ExecCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecCommandResponse) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{3}
}

func (x *ExecCommandResponse) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

type WriteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Content  []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WriteFileRequest) Reset() {
	*x = WriteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileRequest) ProtoMessage() {}

func (x *WriteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileRequest.ProtoReflect.Descriptor instead.
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{4}
}

func (x *WriteFileRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *WriteFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *WriteFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type WriteFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteFileResponse) Reset() {
	*x = WriteFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileResponse) ProtoMessage() {}

func (x *WriteFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileResponse.ProtoReflect.Descriptor instead.
func (*WriteFileResponse) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{5}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vmapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_vmapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_vmapi_proto_rawDescGZIP(), []int{6}
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_vmapi_proto protoreflect.FileDescriptor

var file_vmapi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x6d, 0x61, 0x70, 0x69, 0x22, 0x48, 0x0a, 0x18, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x60,
	0x0a, 0x19, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x00,
	0x52, 0x03, 0x6c, 0x6f, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x42, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x64, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xe5, 0x01, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x58, 0x0a, 0x11, 0x45, 0x78, 0x65, 0x63, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x76,
	0x6d, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x44, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x12, 0x19, 0x2e, 0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x6d,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x76, 0x6d, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x6e, 0x73, 0x63, 0x68, 0x6c, 0x75, 0x65, 0x74,
	0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x6d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vmapi_proto_rawDescOnce sync.Once
	file_vmapi_proto_rawDescData = file_vmapi_proto_rawDesc
)

func file_vmapi_proto_rawDescGZIP() []byte {
	file_vmapi_proto_rawDescOnce.Do(func() {
		file_vmapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_vmapi_proto_rawDescData)
	})
	return file_vmapi_proto_rawDescData
}

var file_vmapi_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vmapi_proto_goTypes = []interface{}{
	(*ExecCommandStreamRequest)(nil),  // 0: vmapi.ExecCommandStreamRequest
	(*ExecCommandStreamResponse)(nil), // 1: vmapi.ExecCommandStreamResponse
	(*ExecCommandRequest)(nil),        // 2: vmapi.ExecCommandRequest
	(*ExecCommandResponse)(nil),       // 3: vmapi.ExecCommandResponse
	(*WriteFileRequest)(nil),          // 4: vmapi.WriteFileRequest
	(*WriteFileResponse)(nil),         // 5: vmapi.WriteFileResponse
	(*Log)(nil),                       // 6: vmapi.Log
}
var file_vmapi_proto_depIdxs = []int32{
	6, // 0: vmapi.ExecCommandStreamResponse.log:type_name -> vmapi.Log
	0, // 1: vmapi.API.ExecCommandStream:input_type -> vmapi.ExecCommandStreamRequest
	2, // 2: vmapi.API.ExecCommand:input_type -> vmapi.ExecCommandRequest
	4, // 3: vmapi.API.WriteFile:input_type -> vmapi.WriteFileRequest
	1, // 4: vmapi.API.ExecCommandStream:output_type -> vmapi.ExecCommandStreamResponse
	3, // 5: vmapi.API.ExecCommand:output_type -> vmapi.ExecCommandResponse
	5, // 6: vmapi.API.WriteFile:output_type -> vmapi.WriteFileResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vmapi_proto_init() }
func file_vmapi_proto_init() {
	if File_vmapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vmapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandStreamRequest); i {
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
		file_vmapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandStreamResponse); i {
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
		file_vmapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandRequest); i {
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
		file_vmapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandResponse); i {
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
		file_vmapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileRequest); i {
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
		file_vmapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileResponse); i {
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
		file_vmapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
	file_vmapi_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExecCommandStreamResponse_Output)(nil),
		(*ExecCommandStreamResponse_Log)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vmapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vmapi_proto_goTypes,
		DependencyIndexes: file_vmapi_proto_depIdxs,
		MessageInfos:      file_vmapi_proto_msgTypes,
	}.Build()
	File_vmapi_proto = out.File
	file_vmapi_proto_rawDesc = nil
	file_vmapi_proto_goTypes = nil
	file_vmapi_proto_depIdxs = nil
}
