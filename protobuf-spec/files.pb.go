// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: files.proto

package protobuf_spec

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

type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Length   string `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	Sha256   string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{0}
}

func (x *FileMeta) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileMeta) GetLength() string {
	if x != nil {
		return x.Length
	}
	return ""
}

func (x *FileMeta) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string    `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Meta     *FileMeta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	Data     []byte    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *File) GetMeta() *FileMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *File) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirectoryName string   `protobuf:"bytes,1,opt,name=directory_name,json=directoryName,proto3" json:"directory_name,omitempty"`
	File          []string `protobuf:"bytes,2,rep,name=File,proto3" json:"File,omitempty"`
}

func (x *Directory) Reset() {
	*x = Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{2}
}

func (x *Directory) GetDirectoryName() string {
	if x != nil {
		return x.DirectoryName
	}
	return ""
}

func (x *Directory) GetFile() []string {
	if x != nil {
		return x.File
	}
	return nil
}

// CreateFileSpace
type CreateFileSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
}

func (x *CreateFileSpaceRequest) Reset() {
	*x = CreateFileSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileSpaceRequest) ProtoMessage() {}

func (x *CreateFileSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileSpaceRequest.ProtoReflect.Descriptor instead.
func (*CreateFileSpaceRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFileSpaceRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

type CreateFileSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *CreateFileSpaceResponse) Reset() {
	*x = CreateFileSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileSpaceResponse) ProtoMessage() {}

func (x *CreateFileSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileSpaceResponse.ProtoReflect.Descriptor instead.
func (*CreateFileSpaceResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFileSpaceResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

// CreateDirectory
type CreateDirectoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSpace string `protobuf:"bytes,1,opt,name=file_space,json=fileSpace,proto3" json:"file_space,omitempty"`
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *CreateDirectoryRequest) Reset() {
	*x = CreateDirectoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryRequest) ProtoMessage() {}

func (x *CreateDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryRequest.ProtoReflect.Descriptor instead.
func (*CreateDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{5}
}

func (x *CreateDirectoryRequest) GetFileSpace() string {
	if x != nil {
		return x.FileSpace
	}
	return ""
}

func (x *CreateDirectoryRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type CreateDirectoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *CreateDirectoryResponse) Reset() {
	*x = CreateDirectoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDirectoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDirectoryResponse) ProtoMessage() {}

func (x *CreateDirectoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDirectoryResponse.ProtoReflect.Descriptor instead.
func (*CreateDirectoryResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDirectoryResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

// CreateFile
type CreateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileSpace string `protobuf:"bytes,1,opt,name=file_space,json=fileSpace,proto3" json:"file_space,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateFileRequest) Reset() {
	*x = CreateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileRequest) ProtoMessage() {}

func (x *CreateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileRequest.ProtoReflect.Descriptor instead.
func (*CreateFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{7}
}

func (x *CreateFileRequest) GetFileSpace() string {
	if x != nil {
		return x.FileSpace
	}
	return ""
}

func (x *CreateFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *CreateFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *CreateFileResponse) Reset() {
	*x = CreateFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileResponse) ProtoMessage() {}

func (x *CreateFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileResponse.ProtoReflect.Descriptor instead.
func (*CreateFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{8}
}

func (x *CreateFileResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

// CheckFileSpaceExists
type FetchFileSpaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
}

func (x *FetchFileSpaceRequest) Reset() {
	*x = FetchFileSpaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFileSpaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFileSpaceRequest) ProtoMessage() {}

func (x *FetchFileSpaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFileSpaceRequest.ProtoReflect.Descriptor instead.
func (*FetchFileSpaceRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{9}
}

func (x *FetchFileSpaceRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

type FetchFileSpaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist         bool   `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	PrevFileSpace string `protobuf:"bytes,2,opt,name=prev_file_space,json=prevFileSpace,proto3" json:"prev_file_space,omitempty"`
}

func (x *FetchFileSpaceResponse) Reset() {
	*x = FetchFileSpaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFileSpaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFileSpaceResponse) ProtoMessage() {}

func (x *FetchFileSpaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFileSpaceResponse.ProtoReflect.Descriptor instead.
func (*FetchFileSpaceResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{10}
}

func (x *FetchFileSpaceResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *FetchFileSpaceResponse) GetPrevFileSpace() string {
	if x != nil {
		return x.PrevFileSpace
	}
	return ""
}

type FetchFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpaceName string `protobuf:"bytes,1,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *FetchFileRequest) Reset() {
	*x = FetchFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFileRequest) ProtoMessage() {}

func (x *FetchFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFileRequest.ProtoReflect.Descriptor instead.
func (*FetchFileRequest) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{11}
}

func (x *FetchFileRequest) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *FetchFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type FetchFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	File   *File  `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FetchFileResponse) Reset() {
	*x = FetchFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_files_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchFileResponse) ProtoMessage() {}

func (x *FetchFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_files_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchFileResponse.ProtoReflect.Descriptor instead.
func (*FetchFileResponse) Descriptor() ([]byte, []int) {
	return file_files_proto_rawDescGZIP(), []int{12}
}

func (x *FetchFileResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_success
}

func (x *FetchFileResponse) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

var File_files_proto protoreflect.FileDescriptor

var file_files_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x22, 0x55, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x09, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x37, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3a,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x63, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x35, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56,
	0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4f, 0x0a, 0x11, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xfe, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x11, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_files_proto_rawDescOnce sync.Once
	file_files_proto_rawDescData = file_files_proto_rawDesc
)

func file_files_proto_rawDescGZIP() []byte {
	file_files_proto_rawDescOnce.Do(func() {
		file_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_files_proto_rawDescData)
	})
	return file_files_proto_rawDescData
}

var file_files_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_files_proto_goTypes = []interface{}{
	(*FileMeta)(nil),                // 0: FileMeta
	(*File)(nil),                    // 1: File
	(*Directory)(nil),               // 2: Directory
	(*CreateFileSpaceRequest)(nil),  // 3: CreateFileSpaceRequest
	(*CreateFileSpaceResponse)(nil), // 4: CreateFileSpaceResponse
	(*CreateDirectoryRequest)(nil),  // 5: CreateDirectoryRequest
	(*CreateDirectoryResponse)(nil), // 6: CreateDirectoryResponse
	(*CreateFileRequest)(nil),       // 7: CreateFileRequest
	(*CreateFileResponse)(nil),      // 8: CreateFileResponse
	(*FetchFileSpaceRequest)(nil),   // 9: FetchFileSpaceRequest
	(*FetchFileSpaceResponse)(nil),  // 10: FetchFileSpaceResponse
	(*FetchFileRequest)(nil),        // 11: FetchFileRequest
	(*FetchFileResponse)(nil),       // 12: FetchFileResponse
	(Status)(0),                     // 13: Status
}
var file_files_proto_depIdxs = []int32{
	0,  // 0: File.meta:type_name -> FileMeta
	13, // 1: CreateFileSpaceResponse.status:type_name -> Status
	13, // 2: CreateDirectoryResponse.status:type_name -> Status
	13, // 3: CreateFileResponse.status:type_name -> Status
	13, // 4: FetchFileResponse.status:type_name -> Status
	1,  // 5: FetchFileResponse.file:type_name -> File
	3,  // 6: Files.CreateFileSpace:input_type -> CreateFileSpaceRequest
	5,  // 7: Files.CreateDirectory:input_type -> CreateDirectoryRequest
	7,  // 8: Files.CreateFile:input_type -> CreateFileRequest
	11, // 9: Files.FetchFile:input_type -> FetchFileRequest
	4,  // 10: Files.CreateFileSpace:output_type -> CreateFileSpaceResponse
	6,  // 11: Files.CreateDirectory:output_type -> CreateDirectoryResponse
	8,  // 12: Files.CreateFile:output_type -> CreateFileResponse
	12, // 13: Files.FetchFile:output_type -> FetchFileResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_files_proto_init() }
func file_files_proto_init() {
	if File_files_proto != nil {
		return
	}
	file_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_files_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directory); i {
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
		file_files_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileSpaceRequest); i {
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
		file_files_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileSpaceResponse); i {
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
		file_files_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryRequest); i {
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
		file_files_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDirectoryResponse); i {
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
		file_files_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileRequest); i {
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
		file_files_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileResponse); i {
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
		file_files_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFileSpaceRequest); i {
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
		file_files_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFileSpaceResponse); i {
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
		file_files_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFileRequest); i {
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
		file_files_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchFileResponse); i {
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
			RawDescriptor: file_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_files_proto_goTypes,
		DependencyIndexes: file_files_proto_depIdxs,
		MessageInfos:      file_files_proto_msgTypes,
	}.Build()
	File_files_proto = out.File
	file_files_proto_rawDesc = nil
	file_files_proto_goTypes = nil
	file_files_proto_depIdxs = nil
}