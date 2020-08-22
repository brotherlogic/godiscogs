// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: godiscogs.proto

package godiscogs

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

type SaleState int32

const (
	SaleState_NOT_FOR_SALE SaleState = 0
	SaleState_FOR_SALE     SaleState = 1
	SaleState_SOLD         SaleState = 2
	SaleState_EXPIRED      SaleState = 3
)

// Enum value maps for SaleState.
var (
	SaleState_name = map[int32]string{
		0: "NOT_FOR_SALE",
		1: "FOR_SALE",
		2: "SOLD",
		3: "EXPIRED",
	}
	SaleState_value = map[string]int32{
		"NOT_FOR_SALE": 0,
		"FOR_SALE":     1,
		"SOLD":         2,
		"EXPIRED":      3,
	}
)

func (x SaleState) Enum() *SaleState {
	p := new(SaleState)
	*p = x
	return p
}

func (x SaleState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SaleState) Descriptor() protoreflect.EnumDescriptor {
	return file_godiscogs_proto_enumTypes[0].Descriptor()
}

func (SaleState) Type() protoreflect.EnumType {
	return &file_godiscogs_proto_enumTypes[0]
}

func (x SaleState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SaleState.Descriptor instead.
func (SaleState) EnumDescriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{0}
}

type Track_TrackType int32

const (
	Track_UNKNOWN Track_TrackType = 0
	Track_TRACK   Track_TrackType = 1
	Track_HEADING Track_TrackType = 2
)

// Enum value maps for Track_TrackType.
var (
	Track_TrackType_name = map[int32]string{
		0: "UNKNOWN",
		1: "TRACK",
		2: "HEADING",
	}
	Track_TrackType_value = map[string]int32{
		"UNKNOWN": 0,
		"TRACK":   1,
		"HEADING": 2,
	}
)

func (x Track_TrackType) Enum() *Track_TrackType {
	p := new(Track_TrackType)
	*p = x
	return p
}

func (x Track_TrackType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Track_TrackType) Descriptor() protoreflect.EnumDescriptor {
	return file_godiscogs_proto_enumTypes[1].Descriptor()
}

func (Track_TrackType) Type() protoreflect.EnumType {
	return &file_godiscogs_proto_enumTypes[1]
}

func (x Track_TrackType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Track_TrackType.Descriptor instead.
func (Track_TrackType) EnumDescriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{8, 0}
}

type ForSale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the record
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of the sale
	SaleId int32 `protobuf:"varint,2,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	// The current price
	SalePrice int32 `protobuf:"varint,3,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
}

func (x *ForSale) Reset() {
	*x = ForSale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForSale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForSale) ProtoMessage() {}

func (x *ForSale) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForSale.ProtoReflect.Descriptor instead.
func (*ForSale) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{0}
}

func (x *ForSale) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ForSale) GetSaleId() int32 {
	if x != nil {
		return x.SaleId
	}
	return 0
}

func (x *ForSale) GetSalePrice() int32 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the label
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The catalogue number
	Catno string `protobuf:"bytes,2,opt,name=catno,proto3" json:"catno,omitempty"`
	// The id of the label
	Id int32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{1}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetCatno() string {
	if x != nil {
		return x.Catno
	}
	return ""
}

func (x *Label) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//The id number of the folder
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the folder
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{2}
}

func (x *Folder) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the artist
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{3}
}

func (x *Artist) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uri to the image
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The type of image
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{4}
}

func (x *Image) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Image) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The descriptions of the Format
	Descriptions []string `protobuf:"bytes,1,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
	// The name of the Format
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The number of the format
	Qty string `protobuf:"bytes,3,opt,name=qty,proto3" json:"qty,omitempty"`
	// Text associated with the Format
	Text string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{5}
}

func (x *Format) GetDescriptions() []string {
	if x != nil {
		return x.Descriptions
	}
	return nil
}

func (x *Format) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Format) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *Format) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id number of the release
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The title of the release
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Artists associated with the release
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists,proto3" json:"artists,omitempty"`
	// The folder in which the record is stored
	FolderId int32 `protobuf:"varint,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Images associated with the release
	Images []*Image `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`
	// The instance id of this release
	InstanceId int32 `protobuf:"varint,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// The labels connected to this release
	Labels []*Label `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty"`
	// The number of discs in the release
	FormatQuantity int32 `protobuf:"varint,8,opt,name=format_quantity,json=formatQuantity,proto3" json:"format_quantity,omitempty"`
	// The rating given to this release
	Rating int32 `protobuf:"varint,9,opt,name=rating,proto3" json:"rating,omitempty"`
	// The earliest release date of this record
	EarliestReleaseDate int64 `protobuf:"varint,10,opt,name=earliest_release_date,json=earliestReleaseDate,proto3" json:"earliest_release_date,omitempty"`
	// The master ID of this release
	MasterId int32 `protobuf:"varint,11,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	// The release date of this release
	Released string `protobuf:"bytes,12,opt,name=released,proto3" json:"released,omitempty"`
	// The formats of the release
	Formats []*Format `protobuf:"bytes,13,rep,name=formats,proto3" json:"formats,omitempty"`
	// Is this a gatefold?
	Gatefold bool `protobuf:"varint,14,opt,name=gatefold,proto3" json:"gatefold,omitempty"`
	// Is this a boxset
	Boxset bool `protobuf:"varint,15,opt,name=boxset,proto3" json:"boxset,omitempty"`
	// The tracks for this release
	Tracklist       []*Track `protobuf:"bytes,16,rep,name=tracklist,proto3" json:"tracklist,omitempty"`
	InstanceNotes   []*Note  `protobuf:"bytes,19,rep,name=instance_notes,json=instanceNotes,proto3" json:"instance_notes,omitempty"`
	RecordCondition string   `protobuf:"bytes,17,opt,name=record_condition,json=recordCondition,proto3" json:"record_condition,omitempty"`
	SleeveCondition string   `protobuf:"bytes,18,opt,name=sleeve_condition,json=sleeveCondition,proto3" json:"sleeve_condition,omitempty"`
	DigitalVersions []int32  `protobuf:"varint,20,rep,packed,name=digital_versions,json=digitalVersions,proto3" json:"digital_versions,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{6}
}

func (x *Release) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Release) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Release) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *Release) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *Release) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Release) GetInstanceId() int32 {
	if x != nil {
		return x.InstanceId
	}
	return 0
}

func (x *Release) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Release) GetFormatQuantity() int32 {
	if x != nil {
		return x.FormatQuantity
	}
	return 0
}

func (x *Release) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Release) GetEarliestReleaseDate() int64 {
	if x != nil {
		return x.EarliestReleaseDate
	}
	return 0
}

func (x *Release) GetMasterId() int32 {
	if x != nil {
		return x.MasterId
	}
	return 0
}

func (x *Release) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *Release) GetFormats() []*Format {
	if x != nil {
		return x.Formats
	}
	return nil
}

func (x *Release) GetGatefold() bool {
	if x != nil {
		return x.Gatefold
	}
	return false
}

func (x *Release) GetBoxset() bool {
	if x != nil {
		return x.Boxset
	}
	return false
}

func (x *Release) GetTracklist() []*Track {
	if x != nil {
		return x.Tracklist
	}
	return nil
}

func (x *Release) GetInstanceNotes() []*Note {
	if x != nil {
		return x.InstanceNotes
	}
	return nil
}

func (x *Release) GetRecordCondition() string {
	if x != nil {
		return x.RecordCondition
	}
	return ""
}

func (x *Release) GetSleeveCondition() string {
	if x != nil {
		return x.SleeveCondition
	}
	return ""
}

func (x *Release) GetDigitalVersions() []int32 {
	if x != nil {
		return x.DigitalVersions
	}
	return nil
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldId int32  `protobuf:"varint,1,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{7}
}

func (x *Note) GetFieldId() int32 {
	if x != nil {
		return x.FieldId
	}
	return 0
}

func (x *Note) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position  string          `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Title     string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type_     string          `protobuf:"bytes,3,opt,name=type_,json=type,proto3" json:"type_,omitempty"`
	TrackType Track_TrackType `protobuf:"varint,4,opt,name=track_type,json=trackType,proto3,enum=godiscogs.Track_TrackType" json:"track_type,omitempty"`
	SubTracks []*Track        `protobuf:"bytes,5,rep,name=sub_tracks,json=subTracks,proto3" json:"sub_tracks,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_godiscogs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_godiscogs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_godiscogs_proto_rawDescGZIP(), []int{8}
}

func (x *Track) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Track) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Track) GetType_() string {
	if x != nil {
		return x.Type_
	}
	return ""
}

func (x *Track) GetTrackType() Track_TrackType {
	if x != nil {
		return x.TrackType
	}
	return Track_UNKNOWN
}

func (x *Track) GetSubTracks() []*Track {
	if x != nil {
		return x.SubTracks
	}
	return nil
}

var File_godiscogs_proto protoreflect.FileDescriptor

var file_godiscogs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x22, 0x51, 0x0a, 0x07,
	0x46, 0x6f, 0x72, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x41, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x61, 0x74, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x74,
	0x6e, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2c, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2c, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2d,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x66, 0x0a,
	0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xe6, 0x05, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x67, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67,
	0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x61, 0x72, 0x6c, 0x69,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18,
	0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67,
	0x73, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x66, 0x6f, 0x6c, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x67, 0x61, 0x74, 0x65, 0x66, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x6f, 0x78, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62,
	0x6f, 0x78, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x67, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x6c, 0x65, 0x65,
	0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x73, 0x6c, 0x65, 0x65, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0f, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x67, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x73, 0x75, 0x62, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x73, 0x22, 0x30, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x42, 0x0a, 0x09, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53,
	0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x41, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x67, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_godiscogs_proto_rawDescOnce sync.Once
	file_godiscogs_proto_rawDescData = file_godiscogs_proto_rawDesc
)

func file_godiscogs_proto_rawDescGZIP() []byte {
	file_godiscogs_proto_rawDescOnce.Do(func() {
		file_godiscogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_godiscogs_proto_rawDescData)
	})
	return file_godiscogs_proto_rawDescData
}

var file_godiscogs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_godiscogs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_godiscogs_proto_goTypes = []interface{}{
	(SaleState)(0),       // 0: godiscogs.SaleState
	(Track_TrackType)(0), // 1: godiscogs.Track.TrackType
	(*ForSale)(nil),      // 2: godiscogs.ForSale
	(*Label)(nil),        // 3: godiscogs.Label
	(*Folder)(nil),       // 4: godiscogs.Folder
	(*Artist)(nil),       // 5: godiscogs.Artist
	(*Image)(nil),        // 6: godiscogs.Image
	(*Format)(nil),       // 7: godiscogs.Format
	(*Release)(nil),      // 8: godiscogs.Release
	(*Note)(nil),         // 9: godiscogs.Note
	(*Track)(nil),        // 10: godiscogs.Track
}
var file_godiscogs_proto_depIdxs = []int32{
	5,  // 0: godiscogs.Release.artists:type_name -> godiscogs.Artist
	6,  // 1: godiscogs.Release.images:type_name -> godiscogs.Image
	3,  // 2: godiscogs.Release.labels:type_name -> godiscogs.Label
	7,  // 3: godiscogs.Release.formats:type_name -> godiscogs.Format
	10, // 4: godiscogs.Release.tracklist:type_name -> godiscogs.Track
	9,  // 5: godiscogs.Release.instance_notes:type_name -> godiscogs.Note
	1,  // 6: godiscogs.Track.track_type:type_name -> godiscogs.Track.TrackType
	10, // 7: godiscogs.Track.sub_tracks:type_name -> godiscogs.Track
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_godiscogs_proto_init() }
func file_godiscogs_proto_init() {
	if File_godiscogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_godiscogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForSale); i {
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
		file_godiscogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_godiscogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
		file_godiscogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
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
		file_godiscogs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_godiscogs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
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
		file_godiscogs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
		file_godiscogs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_godiscogs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
			RawDescriptor: file_godiscogs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_godiscogs_proto_goTypes,
		DependencyIndexes: file_godiscogs_proto_depIdxs,
		EnumInfos:         file_godiscogs_proto_enumTypes,
		MessageInfos:      file_godiscogs_proto_msgTypes,
	}.Build()
	File_godiscogs_proto = out.File
	file_godiscogs_proto_rawDesc = nil
	file_godiscogs_proto_goTypes = nil
	file_godiscogs_proto_depIdxs = nil
}
