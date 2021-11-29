// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/databricks/databricks.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host           string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port           string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	Pwd            string `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	HttpPath       string `protobuf:"bytes,4,opt,name=httpPath,proto3" json:"httpPath,omitempty"`
	UserAgentEntry string `protobuf:"bytes,5,opt,name=userAgentEntry,proto3" json:"userAgentEntry,omitempty"`
	Identifier     string `protobuf:"bytes,6,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ConnectionRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ConnectionRequest) GetPwd() string {
	if x != nil {
		return x.Pwd
	}
	return ""
}

func (x *ConnectionRequest) GetHttpPath() string {
	if x != nil {
		return x.HttpPath
	}
	return ""
}

func (x *ConnectionRequest) GetUserAgentEntry() string {
	if x != nil {
		return x.UserAgentEntry
	}
	return ""
}

func (x *ConnectionRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SqlStatement string `protobuf:"bytes,1,opt,name=sqlStatement,proto3" json:"sqlStatement,omitempty"`
	Identifier   string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteRequest) GetSqlStatement() string {
	if x != nil {
		return x.SqlStatement
	}
	return ""
}

func (x *ExecuteRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FetchSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Databases []string `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	Error     string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FetchSchemasResponse) Reset() {
	*x = FetchSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchSchemasResponse) ProtoMessage() {}

func (x *FetchSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchSchemasResponse.ProtoReflect.Descriptor instead.
func (*FetchSchemasResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{4}
}

func (x *FetchSchemasResponse) GetDatabases() []string {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *FetchSchemasResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FetchTablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []string `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
	Error  string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FetchTablesResponse) Reset() {
	*x = FetchTablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTablesResponse) ProtoMessage() {}

func (x *FetchTablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTablesResponse.ProtoReflect.Descriptor instead.
func (*FetchTablesResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{5}
}

func (x *FetchTablesResponse) GetTables() []string {
	if x != nil {
		return x.Tables
	}
	return nil
}

func (x *FetchTablesResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ItemAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColName  string `protobuf:"bytes,1,opt,name=colName,proto3" json:"colName,omitempty"`
	DataType string `protobuf:"bytes,2,opt,name=dataType,proto3" json:"dataType,omitempty"`
}

func (x *ItemAttribute) Reset() {
	*x = ItemAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemAttribute) ProtoMessage() {}

func (x *ItemAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemAttribute.ProtoReflect.Descriptor instead.
func (*ItemAttribute) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{6}
}

func (x *ItemAttribute) GetColName() string {
	if x != nil {
		return x.ColName
	}
	return ""
}

func (x *ItemAttribute) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

type FetchTableAttributesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*ItemAttribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Error      string           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FetchTableAttributesResponse) Reset() {
	*x = FetchTableAttributesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTableAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTableAttributesResponse) ProtoMessage() {}

func (x *FetchTableAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTableAttributesResponse.ProtoReflect.Descriptor instead.
func (*FetchTableAttributesResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{7}
}

func (x *FetchTableAttributesResponse) GetAttributes() []*ItemAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *FetchTableAttributesResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FetchTotalCountInTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FetchTotalCountInTableResponse) Reset() {
	*x = FetchTotalCountInTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_databricks_databricks_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchTotalCountInTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchTotalCountInTableResponse) ProtoMessage() {}

func (x *FetchTotalCountInTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_databricks_databricks_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchTotalCountInTableResponse.ProtoReflect.Descriptor instead.
func (*FetchTotalCountInTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_databricks_databricks_proto_rawDescGZIP(), []int{8}
}

func (x *FetchTotalCountInTableResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *FetchTotalCountInTableResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_databricks_databricks_proto protoreflect.FileDescriptor

var file_proto_databricks_databricks_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63,
	0x6b, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x77, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x77, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74,
	0x74, 0x70, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74,
	0x74, 0x70, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x2a,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x0e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x71, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x71, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x14, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x43, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x45, 0x0a, 0x0d, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x6a, 0x0a, 0x1c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4c, 0x0a,
	0x1e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x84, 0x04, 0x0a, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x72, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x54, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_databricks_databricks_proto_rawDescOnce sync.Once
	file_proto_databricks_databricks_proto_rawDescData = file_proto_databricks_databricks_proto_rawDesc
)

func file_proto_databricks_databricks_proto_rawDescGZIP() []byte {
	file_proto_databricks_databricks_proto_rawDescOnce.Do(func() {
		file_proto_databricks_databricks_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_databricks_databricks_proto_rawDescData)
	})
	return file_proto_databricks_databricks_proto_rawDescData
}

var file_proto_databricks_databricks_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_databricks_databricks_proto_goTypes = []interface{}{
	(*ConnectionRequest)(nil),              // 0: proto.ConnectionRequest
	(*ConnectionResponse)(nil),             // 1: proto.ConnectionResponse
	(*ExecuteRequest)(nil),                 // 2: proto.ExecuteRequest
	(*ExecuteResponse)(nil),                // 3: proto.ExecuteResponse
	(*FetchSchemasResponse)(nil),           // 4: proto.FetchSchemasResponse
	(*FetchTablesResponse)(nil),            // 5: proto.FetchTablesResponse
	(*ItemAttribute)(nil),                  // 6: proto.ItemAttribute
	(*FetchTableAttributesResponse)(nil),   // 7: proto.FetchTableAttributesResponse
	(*FetchTotalCountInTableResponse)(nil), // 8: proto.FetchTotalCountInTableResponse
}
var file_proto_databricks_databricks_proto_depIdxs = []int32{
	6, // 0: proto.FetchTableAttributesResponse.attributes:type_name -> proto.ItemAttribute
	0, // 1: proto.Databricks.Connect:input_type -> proto.ConnectionRequest
	0, // 2: proto.Databricks.Close:input_type -> proto.ConnectionRequest
	2, // 3: proto.Databricks.Execute:input_type -> proto.ExecuteRequest
	2, // 4: proto.Databricks.FetchSchemas:input_type -> proto.ExecuteRequest
	2, // 5: proto.Databricks.FetchTables:input_type -> proto.ExecuteRequest
	2, // 6: proto.Databricks.FetchTableAttributes:input_type -> proto.ExecuteRequest
	2, // 7: proto.Databricks.FetchTotalCountInTable:input_type -> proto.ExecuteRequest
	1, // 8: proto.Databricks.Connect:output_type -> proto.ConnectionResponse
	1, // 9: proto.Databricks.Close:output_type -> proto.ConnectionResponse
	3, // 10: proto.Databricks.Execute:output_type -> proto.ExecuteResponse
	4, // 11: proto.Databricks.FetchSchemas:output_type -> proto.FetchSchemasResponse
	5, // 12: proto.Databricks.FetchTables:output_type -> proto.FetchTablesResponse
	7, // 13: proto.Databricks.FetchTableAttributes:output_type -> proto.FetchTableAttributesResponse
	8, // 14: proto.Databricks.FetchTotalCountInTable:output_type -> proto.FetchTotalCountInTableResponse
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_databricks_databricks_proto_init() }
func file_proto_databricks_databricks_proto_init() {
	if File_proto_databricks_databricks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_databricks_databricks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_proto_databricks_databricks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
		file_proto_databricks_databricks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_proto_databricks_databricks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
		file_proto_databricks_databricks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchSchemasResponse); i {
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
		file_proto_databricks_databricks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTablesResponse); i {
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
		file_proto_databricks_databricks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemAttribute); i {
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
		file_proto_databricks_databricks_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTableAttributesResponse); i {
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
		file_proto_databricks_databricks_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchTotalCountInTableResponse); i {
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
			RawDescriptor: file_proto_databricks_databricks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_databricks_databricks_proto_goTypes,
		DependencyIndexes: file_proto_databricks_databricks_proto_depIdxs,
		MessageInfos:      file_proto_databricks_databricks_proto_msgTypes,
	}.Build()
	File_proto_databricks_databricks_proto = out.File
	file_proto_databricks_databricks_proto_rawDesc = nil
	file_proto_databricks_databricks_proto_goTypes = nil
	file_proto_databricks_databricks_proto_depIdxs = nil
}
