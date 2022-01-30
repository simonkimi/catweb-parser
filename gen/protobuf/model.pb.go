// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: model.proto

package protobuf

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

type ImageRpcModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string  `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CacheKey string  `protobuf:"bytes,2,opt,name=cacheKey,proto3" json:"cacheKey,omitempty"`
	Width    float64 `protobuf:"fixed64,3,opt,name=width,proto3" json:"width,omitempty"`
	Height   float64 `protobuf:"fixed64,4,opt,name=height,proto3" json:"height,omitempty"`
	ImgX     float64 `protobuf:"fixed64,5,opt,name=imgX,proto3" json:"imgX,omitempty"`
	ImgY     float64 `protobuf:"fixed64,6,opt,name=imgY,proto3" json:"imgY,omitempty"`
	Target   string  `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *ImageRpcModel) Reset() {
	*x = ImageRpcModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageRpcModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageRpcModel) ProtoMessage() {}

func (x *ImageRpcModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageRpcModel.ProtoReflect.Descriptor instead.
func (*ImageRpcModel) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *ImageRpcModel) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageRpcModel) GetCacheKey() string {
	if x != nil {
		return x.CacheKey
	}
	return ""
}

func (x *ImageRpcModel) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ImageRpcModel) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageRpcModel) GetImgX() float64 {
	if x != nil {
		return x.ImgX
	}
	return 0
}

func (x *ImageRpcModel) GetImgY() float64 {
	if x != nil {
		return x.ImgY
	}
	return 0
}

func (x *ImageRpcModel) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type ColorRpcModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	R int32 `protobuf:"varint,2,opt,name=r,proto3" json:"r,omitempty"`
	G int32 `protobuf:"varint,3,opt,name=g,proto3" json:"g,omitempty"`
	B int32 `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *ColorRpcModel) Reset() {
	*x = ColorRpcModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColorRpcModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColorRpcModel) ProtoMessage() {}

func (x *ColorRpcModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColorRpcModel.ProtoReflect.Descriptor instead.
func (*ColorRpcModel) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *ColorRpcModel) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *ColorRpcModel) GetR() int32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *ColorRpcModel) GetG() int32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *ColorRpcModel) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type ListRpcModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListRpcModel_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Env   map[string]string    `protobuf:"bytes,10,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListRpcModel) Reset() {
	*x = ListRpcModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpcModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpcModel) ProtoMessage() {}

func (x *ListRpcModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpcModel.ProtoReflect.Descriptor instead.
func (*ListRpcModel) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *ListRpcModel) GetItems() []*ListRpcModel_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRpcModel) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type DetailRpcModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string                    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle     string                    `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Language     string                    `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	ImageCount   int32                     `protobuf:"varint,4,opt,name=image_count,json=imageCount,proto3" json:"image_count,omitempty"`
	UploadTime   string                    `protobuf:"bytes,5,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
	CountPrePage int32                     `protobuf:"varint,6,opt,name=count_pre_page,json=countPrePage,proto3" json:"count_pre_page,omitempty"`
	Description  string                    `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Star         float64                   `protobuf:"fixed64,8,opt,name=star,proto3" json:"star,omitempty"`
	Tag          *DetailRpcModel_Tag       `protobuf:"bytes,20,opt,name=tag,proto3" json:"tag,omitempty"`
	Badges       []*DetailRpcModel_Badge   `protobuf:"bytes,21,rep,name=badges,proto3" json:"badges,omitempty"`
	Comments     []*DetailRpcModel_Comment `protobuf:"bytes,22,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *DetailRpcModel) Reset() {
	*x = DetailRpcModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRpcModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRpcModel) ProtoMessage() {}

func (x *DetailRpcModel) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRpcModel.ProtoReflect.Descriptor instead.
func (*DetailRpcModel) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *DetailRpcModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DetailRpcModel) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *DetailRpcModel) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *DetailRpcModel) GetImageCount() int32 {
	if x != nil {
		return x.ImageCount
	}
	return 0
}

func (x *DetailRpcModel) GetUploadTime() string {
	if x != nil {
		return x.UploadTime
	}
	return ""
}

func (x *DetailRpcModel) GetCountPrePage() int32 {
	if x != nil {
		return x.CountPrePage
	}
	return 0
}

func (x *DetailRpcModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DetailRpcModel) GetStar() float64 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *DetailRpcModel) GetTag() *DetailRpcModel_Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *DetailRpcModel) GetBadges() []*DetailRpcModel_Badge {
	if x != nil {
		return x.Badges
	}
	return nil
}

func (x *DetailRpcModel) GetComments() []*DetailRpcModel_Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type ListRpcModel_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string         `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Color *ColorRpcModel `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *ListRpcModel_Tag) Reset() {
	*x = ListRpcModel_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpcModel_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpcModel_Tag) ProtoMessage() {}

func (x *ListRpcModel_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpcModel_Tag.ProtoReflect.Descriptor instead.
func (*ListRpcModel_Tag) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListRpcModel_Tag) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ListRpcModel_Tag) GetColor() *ColorRpcModel {
	if x != nil {
		return x.Color
	}
	return nil
}

type ListRpcModel_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string              `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle   string              `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	UploadTime string              `protobuf:"bytes,3,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
	Star       float64             `protobuf:"fixed64,4,opt,name=star,proto3" json:"star,omitempty"`
	ImgCount   int32               `protobuf:"varint,5,opt,name=img_count,json=imgCount,proto3" json:"img_count,omitempty"`
	PreviewImg *ImageRpcModel      `protobuf:"bytes,6,opt,name=preview_img,json=previewImg,proto3" json:"preview_img,omitempty"`
	Tag        *ListRpcModel_Tag   `protobuf:"bytes,10,opt,name=tag,proto3" json:"tag,omitempty"`
	Badges     []*ListRpcModel_Tag `protobuf:"bytes,11,rep,name=badges,proto3" json:"badges,omitempty"`
	Paper      string              `protobuf:"bytes,12,opt,name=paper,proto3" json:"paper,omitempty"`
	NextPage   string              `protobuf:"bytes,20,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	IdCode     string              `protobuf:"bytes,21,opt,name=id_code,json=idCode,proto3" json:"id_code,omitempty"`
	Env        map[string]string   `protobuf:"bytes,30,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListRpcModel_Item) Reset() {
	*x = ListRpcModel_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpcModel_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpcModel_Item) ProtoMessage() {}

func (x *ListRpcModel_Item) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpcModel_Item.ProtoReflect.Descriptor instead.
func (*ListRpcModel_Item) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ListRpcModel_Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListRpcModel_Item) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *ListRpcModel_Item) GetUploadTime() string {
	if x != nil {
		return x.UploadTime
	}
	return ""
}

func (x *ListRpcModel_Item) GetStar() float64 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *ListRpcModel_Item) GetImgCount() int32 {
	if x != nil {
		return x.ImgCount
	}
	return 0
}

func (x *ListRpcModel_Item) GetPreviewImg() *ImageRpcModel {
	if x != nil {
		return x.PreviewImg
	}
	return nil
}

func (x *ListRpcModel_Item) GetTag() *ListRpcModel_Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *ListRpcModel_Item) GetBadges() []*ListRpcModel_Tag {
	if x != nil {
		return x.Badges
	}
	return nil
}

func (x *ListRpcModel_Item) GetPaper() string {
	if x != nil {
		return x.Paper
	}
	return ""
}

func (x *ListRpcModel_Item) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

func (x *ListRpcModel_Item) GetIdCode() string {
	if x != nil {
		return x.IdCode
	}
	return ""
}

func (x *ListRpcModel_Item) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type DetailRpcModel_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string            `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Color *ColorRpcModel    `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Env   map[string]string `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DetailRpcModel_Tag) Reset() {
	*x = DetailRpcModel_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRpcModel_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRpcModel_Tag) ProtoMessage() {}

func (x *DetailRpcModel_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRpcModel_Tag.ProtoReflect.Descriptor instead.
func (*DetailRpcModel_Tag) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DetailRpcModel_Tag) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *DetailRpcModel_Tag) GetColor() *ColorRpcModel {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *DetailRpcModel_Tag) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type DetailRpcModel_Badge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *DetailRpcModel_Badge) Reset() {
	*x = DetailRpcModel_Badge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRpcModel_Badge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRpcModel_Badge) ProtoMessage() {}

func (x *DetailRpcModel_Badge) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRpcModel_Badge.ProtoReflect.Descriptor instead.
func (*DetailRpcModel_Badge) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3, 1}
}

func (x *DetailRpcModel_Badge) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *DetailRpcModel_Badge) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type DetailRpcModel_Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Time     string `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Score    string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *DetailRpcModel_Comment) Reset() {
	*x = DetailRpcModel_Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRpcModel_Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRpcModel_Comment) ProtoMessage() {}

func (x *DetailRpcModel_Comment) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRpcModel_Comment.ProtoReflect.Descriptor instead.
func (*DetailRpcModel_Comment) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3, 2}
}

func (x *DetailRpcModel_Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DetailRpcModel_Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *DetailRpcModel_Comment) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *DetailRpcModel_Comment) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01,
	0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6d, 0x67, 0x58, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x69, 0x6d, 0x67, 0x58, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x6d, 0x67, 0x59, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x69,
	0x6d, 0x67, 0x59, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x47, 0x0a, 0x0d, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x62, 0x22, 0x9c, 0x05, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x63,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x28, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x76, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x1a, 0x3f, 0x0a, 0x03, 0x54, 0x61, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x70, 0x63, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x1a, 0xbe, 0x03, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x6d, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6d, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0a, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6d, 0x67, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x29, 0x0a,
	0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x70, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x1e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03,
	0x65, 0x6e, 0x76, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x36, 0x0a, 0x08, 0x45,
	0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xd5, 0x05, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x70,
	0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x74,
	0x61, 0x72, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65,
	0x52, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xa7, 0x01,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72,
	0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x2e, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x70, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x61,
	0x67, 0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x1a,
	0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x05, 0x42, 0x61, 0x64, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x1a, 0x69, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_model_proto_goTypes = []interface{}{
	(*ImageRpcModel)(nil),          // 0: ImageRpcModel
	(*ColorRpcModel)(nil),          // 1: ColorRpcModel
	(*ListRpcModel)(nil),           // 2: ListRpcModel
	(*DetailRpcModel)(nil),         // 3: DetailRpcModel
	(*ListRpcModel_Tag)(nil),       // 4: ListRpcModel.Tag
	(*ListRpcModel_Item)(nil),      // 5: ListRpcModel.Item
	nil,                            // 6: ListRpcModel.EnvEntry
	nil,                            // 7: ListRpcModel.Item.EnvEntry
	(*DetailRpcModel_Tag)(nil),     // 8: DetailRpcModel.Tag
	(*DetailRpcModel_Badge)(nil),   // 9: DetailRpcModel.Badge
	(*DetailRpcModel_Comment)(nil), // 10: DetailRpcModel.Comment
	nil,                            // 11: DetailRpcModel.Tag.EnvEntry
}
var file_model_proto_depIdxs = []int32{
	5,  // 0: ListRpcModel.items:type_name -> ListRpcModel.Item
	6,  // 1: ListRpcModel.env:type_name -> ListRpcModel.EnvEntry
	8,  // 2: DetailRpcModel.tag:type_name -> DetailRpcModel.Tag
	9,  // 3: DetailRpcModel.badges:type_name -> DetailRpcModel.Badge
	10, // 4: DetailRpcModel.comments:type_name -> DetailRpcModel.Comment
	1,  // 5: ListRpcModel.Tag.color:type_name -> ColorRpcModel
	0,  // 6: ListRpcModel.Item.preview_img:type_name -> ImageRpcModel
	4,  // 7: ListRpcModel.Item.tag:type_name -> ListRpcModel.Tag
	4,  // 8: ListRpcModel.Item.badges:type_name -> ListRpcModel.Tag
	7,  // 9: ListRpcModel.Item.env:type_name -> ListRpcModel.Item.EnvEntry
	1,  // 10: DetailRpcModel.Tag.color:type_name -> ColorRpcModel
	11, // 11: DetailRpcModel.Tag.env:type_name -> DetailRpcModel.Tag.EnvEntry
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageRpcModel); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColorRpcModel); i {
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
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpcModel); i {
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
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRpcModel); i {
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
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpcModel_Tag); i {
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
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpcModel_Item); i {
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
		file_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRpcModel_Tag); i {
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
		file_model_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRpcModel_Badge); i {
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
		file_model_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRpcModel_Comment); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
