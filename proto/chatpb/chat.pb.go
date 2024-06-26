// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/chatpb/chat.proto

package chatpb

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

type SentSingleIssueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	IssueId  int32  `protobuf:"varint,3,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *SentSingleIssueReq) Reset() {
	*x = SentSingleIssueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentSingleIssueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentSingleIssueReq) ProtoMessage() {}

func (x *SentSingleIssueReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentSingleIssueReq.ProtoReflect.Descriptor instead.
func (*SentSingleIssueReq) Descriptor() ([]byte, []int) {
	return file_proto_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *SentSingleIssueReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SentSingleIssueReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SentSingleIssueReq) GetIssueId() int32 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

type SentSingleIssueResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SentSingleIssueResp) Reset() {
	*x = SentSingleIssueResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatpb_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentSingleIssueResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentSingleIssueResp) ProtoMessage() {}

func (x *SentSingleIssueResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatpb_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentSingleIssueResp.ProtoReflect.Descriptor instead.
func (*SentSingleIssueResp) Descriptor() ([]byte, []int) {
	return file_proto_chatpb_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SentSingleIssueResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SentSingleIssueResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SentSingleIssueResp) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type SentProjectIssueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectName string `protobuf:"bytes,3,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *SentProjectIssueReq) Reset() {
	*x = SentProjectIssueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatpb_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentProjectIssueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentProjectIssueReq) ProtoMessage() {}

func (x *SentProjectIssueReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatpb_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentProjectIssueReq.ProtoReflect.Descriptor instead.
func (*SentProjectIssueReq) Descriptor() ([]byte, []int) {
	return file_proto_chatpb_chat_proto_rawDescGZIP(), []int{2}
}

func (x *SentProjectIssueReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SentProjectIssueReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SentProjectIssueReq) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type SentProjectIssueResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    map[string]string `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SentProjectIssueResp) Reset() {
	*x = SentProjectIssueResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chatpb_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentProjectIssueResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentProjectIssueResp) ProtoMessage() {}

func (x *SentProjectIssueResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chatpb_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentProjectIssueResp.ProtoReflect.Descriptor instead.
func (*SentProjectIssueResp) Descriptor() ([]byte, []int) {
	return file_proto_chatpb_chat_proto_rawDescGZIP(), []int{3}
}

func (x *SentProjectIssueResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SentProjectIssueResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SentProjectIssueResp) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_chatpb_chat_proto protoreflect.FileDescriptor

var file_proto_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x22, 0x61, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6a,
	0x0a, 0x13, 0x53, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x14, 0x53,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa5, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x4c, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chatpb_chat_proto_rawDescOnce sync.Once
	file_proto_chatpb_chat_proto_rawDescData = file_proto_chatpb_chat_proto_rawDesc
)

func file_proto_chatpb_chat_proto_rawDescGZIP() []byte {
	file_proto_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_proto_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chatpb_chat_proto_rawDescData)
	})
	return file_proto_chatpb_chat_proto_rawDescData
}

var file_proto_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_chatpb_chat_proto_goTypes = []interface{}{
	(*SentSingleIssueReq)(nil),   // 0: chatpb.SentSingleIssueReq
	(*SentSingleIssueResp)(nil),  // 1: chatpb.SentSingleIssueResp
	(*SentProjectIssueReq)(nil),  // 2: chatpb.SentProjectIssueReq
	(*SentProjectIssueResp)(nil), // 3: chatpb.SentProjectIssueResp
	nil,                          // 4: chatpb.SentSingleIssueResp.DataEntry
	nil,                          // 5: chatpb.SentProjectIssueResp.DataEntry
}
var file_proto_chatpb_chat_proto_depIdxs = []int32{
	4, // 0: chatpb.SentSingleIssueResp.data:type_name -> chatpb.SentSingleIssueResp.DataEntry
	5, // 1: chatpb.SentProjectIssueResp.data:type_name -> chatpb.SentProjectIssueResp.DataEntry
	0, // 2: chatpb.Chat.SentSingleIssue:input_type -> chatpb.SentSingleIssueReq
	2, // 3: chatpb.Chat.SentProjectIssue:input_type -> chatpb.SentProjectIssueReq
	1, // 4: chatpb.Chat.SentSingleIssue:output_type -> chatpb.SentSingleIssueResp
	3, // 5: chatpb.Chat.SentProjectIssue:output_type -> chatpb.SentProjectIssueResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_chatpb_chat_proto_init() }
func file_proto_chatpb_chat_proto_init() {
	if File_proto_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentSingleIssueReq); i {
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
		file_proto_chatpb_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentSingleIssueResp); i {
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
		file_proto_chatpb_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentProjectIssueReq); i {
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
		file_proto_chatpb_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentProjectIssueResp); i {
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
			RawDescriptor: file_proto_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_proto_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_proto_chatpb_chat_proto_msgTypes,
	}.Build()
	File_proto_chatpb_chat_proto = out.File
	file_proto_chatpb_chat_proto_rawDesc = nil
	file_proto_chatpb_chat_proto_goTypes = nil
	file_proto_chatpb_chat_proto_depIdxs = nil
}
