// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/books.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string               `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Image       string               `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Pages       uint64               `protobuf:"varint,5,opt,name=pages,proto3" json:"pages,omitempty"`
	Author      string               `protobuf:"bytes,6,opt,name=author,proto3" json:"author,omitempty"`
	Genre       string               `protobuf:"bytes,7,opt,name=genre,proto3" json:"genre,omitempty"`
	Publisher   string               `protobuf:"bytes,8,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Binding     bool                 `protobuf:"varint,9,opt,name=binding,proto3" json:"binding,omitempty"`
	Series      string               `protobuf:"bytes,10,opt,name=series,proto3" json:"series,omitempty"`
	Langugae    string               `protobuf:"bytes,11,opt,name=langugae,proto3" json:"langugae,omitempty"`
	AddedAt     *timestamp.Timestamp `protobuf:"bytes,12,opt,name=added_at,json=addedAt,proto3" json:"added_at,omitempty"`
}

func (x *BookInfo) Reset() {
	*x = BookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfo) ProtoMessage() {}

func (x *BookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfo.ProtoReflect.Descriptor instead.
func (*BookInfo) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{0}
}

func (x *BookInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BookInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BookInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BookInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *BookInfo) GetPages() uint64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *BookInfo) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *BookInfo) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *BookInfo) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *BookInfo) GetBinding() bool {
	if x != nil {
		return x.Binding
	}
	return false
}

func (x *BookInfo) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *BookInfo) GetLangugae() string {
	if x != nil {
		return x.Langugae
	}
	return ""
}

func (x *BookInfo) GetAddedAt() *timestamp.Timestamp {
	if x != nil {
		return x.AddedAt
	}
	return nil
}

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Image       string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Pages       uint64 `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	Author      string `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Genre       string `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Publisher   string `protobuf:"bytes,7,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Binding     bool   `protobuf:"varint,8,opt,name=binding,proto3" json:"binding,omitempty"`
	Series      string `protobuf:"bytes,9,opt,name=series,proto3" json:"series,omitempty"`
	Langugae    string `protobuf:"bytes,10,opt,name=langugae,proto3" json:"langugae,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBookRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CreateBookRequest) GetPages() uint64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *CreateBookRequest) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *CreateBookRequest) GetBinding() bool {
	if x != nil {
		return x.Binding
	}
	return false
}

func (x *CreateBookRequest) GetSeries() string {
	if x != nil {
		return x.Series
	}
	return ""
}

func (x *CreateBookRequest) GetLangugae() string {
	if x != nil {
		return x.Langugae
	}
	return ""
}

type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=bookID,proto3" json:"bookID,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBookResponse) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

type DeleteBookRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
}

func (x *DeleteBookRequestResponse) Reset() {
	*x = DeleteBookRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequestResponse) ProtoMessage() {}

func (x *DeleteBookRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequestResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookRequestResponse) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBookRequestResponse) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

type GetBookRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookID string `protobuf:"bytes,1,opt,name=BookID,proto3" json:"BookID,omitempty"`
}

func (x *GetBookRequset) Reset() {
	*x = GetBookRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequset) ProtoMessage() {}

func (x *GetBookRequset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequset.ProtoReflect.Descriptor instead.
func (*GetBookRequset) Descriptor() ([]byte, []int) {
	return file_proto_books_proto_rawDescGZIP(), []int{4}
}

func (x *GetBookRequset) GetBookID() string {
	if x != nil {
		return x.BookID
	}
	return ""
}

var File_proto_books_proto protoreflect.FileDescriptor

var file_proto_books_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x08, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x67, 0x61, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x67, 0x61, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x67, 0x61, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x67, 0x61, 0x65, 0x22, 0x2c,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x19,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49,
	0x44, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x44, 0x32, 0xf7, 0x01, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x73, 0x65, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_books_proto_rawDescOnce sync.Once
	file_proto_books_proto_rawDescData = file_proto_books_proto_rawDesc
)

func file_proto_books_proto_rawDescGZIP() []byte {
	file_proto_books_proto_rawDescOnce.Do(func() {
		file_proto_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_books_proto_rawDescData)
	})
	return file_proto_books_proto_rawDescData
}

var file_proto_books_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_books_proto_goTypes = []interface{}{
	(*BookInfo)(nil),                  // 0: proto.BookInfo
	(*CreateBookRequest)(nil),         // 1: proto.CreateBookRequest
	(*CreateBookResponse)(nil),        // 2: proto.CreateBookResponse
	(*DeleteBookRequestResponse)(nil), // 3: proto.DeleteBookRequestResponse
	(*GetBookRequset)(nil),            // 4: proto.GetBookRequset
	(*timestamp.Timestamp)(nil),       // 5: google.protobuf.Timestamp
	(*empty.Empty)(nil),               // 6: google.protobuf.Empty
}
var file_proto_books_proto_depIdxs = []int32{
	5, // 0: proto.BookInfo.added_at:type_name -> google.protobuf.Timestamp
	1, // 1: proto.Book.Create:input_type -> proto.CreateBookRequest
	3, // 2: proto.Book.Delete:input_type -> proto.DeleteBookRequestResponse
	6, // 3: proto.Book.GetAll:input_type -> google.protobuf.Empty
	4, // 4: proto.Book.Get:input_type -> proto.GetBookRequset
	2, // 5: proto.Book.Create:output_type -> proto.CreateBookResponse
	3, // 6: proto.Book.Delete:output_type -> proto.DeleteBookRequestResponse
	0, // 7: proto.Book.GetAll:output_type -> proto.BookInfo
	0, // 8: proto.Book.Get:output_type -> proto.BookInfo
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_books_proto_init() }
func file_proto_books_proto_init() {
	if File_proto_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookInfo); i {
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
		file_proto_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_proto_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResponse); i {
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
		file_proto_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequestResponse); i {
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
		file_proto_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequset); i {
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
			RawDescriptor: file_proto_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_books_proto_goTypes,
		DependencyIndexes: file_proto_books_proto_depIdxs,
		MessageInfos:      file_proto_books_proto_msgTypes,
	}.Build()
	File_proto_books_proto = out.File
	file_proto_books_proto_rawDesc = nil
	file_proto_books_proto_goTypes = nil
	file_proto_books_proto_depIdxs = nil
}