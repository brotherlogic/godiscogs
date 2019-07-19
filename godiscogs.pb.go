// Code generated by protoc-gen-go. DO NOT EDIT.
// source: godiscogs.proto

package godiscogs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SaleState int32

const (
	SaleState_NOT_FOR_SALE SaleState = 0
	SaleState_FOR_SALE     SaleState = 1
	SaleState_SOLD         SaleState = 2
)

var SaleState_name = map[int32]string{
	0: "NOT_FOR_SALE",
	1: "FOR_SALE",
	2: "SOLD",
}

var SaleState_value = map[string]int32{
	"NOT_FOR_SALE": 0,
	"FOR_SALE":     1,
	"SOLD":         2,
}

func (x SaleState) String() string {
	return proto.EnumName(SaleState_name, int32(x))
}

func (SaleState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

type Track_TrackType int32

const (
	Track_UNKNOWN Track_TrackType = 0
	Track_TRACK   Track_TrackType = 1
	Track_HEADING Track_TrackType = 2
)

var Track_TrackType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TRACK",
	2: "HEADING",
}

var Track_TrackType_value = map[string]int32{
	"UNKNOWN": 0,
	"TRACK":   1,
	"HEADING": 2,
}

func (x Track_TrackType) String() string {
	return proto.EnumName(Track_TrackType_name, int32(x))
}

func (Track_TrackType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{7, 0}
}

type Label struct {
	// The name of the label
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The catalogue number
	Catno string `protobuf:"bytes,2,opt,name=catno,proto3" json:"catno,omitempty"`
	// The id of the label
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetCatno() string {
	if m != nil {
		return m.Catno
	}
	return ""
}

func (m *Label) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Folder struct {
	//The id number of the folder
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the folder
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Folder) Reset()         { *m = Folder{} }
func (m *Folder) String() string { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()    {}
func (*Folder) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{1}
}

func (m *Folder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Folder.Unmarshal(m, b)
}
func (m *Folder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Folder.Marshal(b, m, deterministic)
}
func (m *Folder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Folder.Merge(m, src)
}
func (m *Folder) XXX_Size() int {
	return xxx_messageInfo_Folder.Size(m)
}
func (m *Folder) XXX_DiscardUnknown() {
	xxx_messageInfo_Folder.DiscardUnknown(m)
}

var xxx_messageInfo_Folder proto.InternalMessageInfo

func (m *Folder) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Folder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Artist struct {
	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//The name of the artist
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artist) Reset()         { *m = Artist{} }
func (m *Artist) String() string { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()    {}
func (*Artist) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{2}
}

func (m *Artist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artist.Unmarshal(m, b)
}
func (m *Artist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artist.Marshal(b, m, deterministic)
}
func (m *Artist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artist.Merge(m, src)
}
func (m *Artist) XXX_Size() int {
	return xxx_messageInfo_Artist.Size(m)
}
func (m *Artist) XXX_DiscardUnknown() {
	xxx_messageInfo_Artist.DiscardUnknown(m)
}

var xxx_messageInfo_Artist proto.InternalMessageInfo

func (m *Artist) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Artist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Image struct {
	// The uri to the image
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// The type of image
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Image) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Format struct {
	// The descriptions of the Format
	Descriptions []string `protobuf:"bytes,1,rep,name=descriptions,proto3" json:"descriptions,omitempty"`
	// The name of the Format
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The number of the format
	Qty string `protobuf:"bytes,3,opt,name=qty,proto3" json:"qty,omitempty"`
	// Text associated with the Format
	Text                 string   `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Format) Reset()         { *m = Format{} }
func (m *Format) String() string { return proto.CompactTextString(m) }
func (*Format) ProtoMessage()    {}
func (*Format) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{4}
}

func (m *Format) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Format.Unmarshal(m, b)
}
func (m *Format) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Format.Marshal(b, m, deterministic)
}
func (m *Format) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Format.Merge(m, src)
}
func (m *Format) XXX_Size() int {
	return xxx_messageInfo_Format.Size(m)
}
func (m *Format) XXX_DiscardUnknown() {
	xxx_messageInfo_Format.DiscardUnknown(m)
}

var xxx_messageInfo_Format proto.InternalMessageInfo

func (m *Format) GetDescriptions() []string {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

func (m *Format) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Format) GetQty() string {
	if m != nil {
		return m.Qty
	}
	return ""
}

func (m *Format) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Release struct {
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
	Tracklist            []*Track `protobuf:"bytes,16,rep,name=tracklist,proto3" json:"tracklist,omitempty"`
	InstanceNotes        []*Note  `protobuf:"bytes,19,rep,name=instance_notes,json=instanceNotes,proto3" json:"instance_notes,omitempty"`
	RecordCondition      string   `protobuf:"bytes,17,opt,name=record_condition,json=recordCondition,proto3" json:"record_condition,omitempty"`
	SleeveCondition      string   `protobuf:"bytes,18,opt,name=sleeve_condition,json=sleeveCondition,proto3" json:"sleeve_condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Release) Reset()         { *m = Release{} }
func (m *Release) String() string { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()    {}
func (*Release) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{5}
}

func (m *Release) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Release.Unmarshal(m, b)
}
func (m *Release) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Release.Marshal(b, m, deterministic)
}
func (m *Release) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Release.Merge(m, src)
}
func (m *Release) XXX_Size() int {
	return xxx_messageInfo_Release.Size(m)
}
func (m *Release) XXX_DiscardUnknown() {
	xxx_messageInfo_Release.DiscardUnknown(m)
}

var xxx_messageInfo_Release proto.InternalMessageInfo

func (m *Release) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Release) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Release) GetArtists() []*Artist {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *Release) GetFolderId() int32 {
	if m != nil {
		return m.FolderId
	}
	return 0
}

func (m *Release) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Release) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *Release) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Release) GetFormatQuantity() int32 {
	if m != nil {
		return m.FormatQuantity
	}
	return 0
}

func (m *Release) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Release) GetEarliestReleaseDate() int64 {
	if m != nil {
		return m.EarliestReleaseDate
	}
	return 0
}

func (m *Release) GetMasterId() int32 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *Release) GetReleased() string {
	if m != nil {
		return m.Released
	}
	return ""
}

func (m *Release) GetFormats() []*Format {
	if m != nil {
		return m.Formats
	}
	return nil
}

func (m *Release) GetGatefold() bool {
	if m != nil {
		return m.Gatefold
	}
	return false
}

func (m *Release) GetBoxset() bool {
	if m != nil {
		return m.Boxset
	}
	return false
}

func (m *Release) GetTracklist() []*Track {
	if m != nil {
		return m.Tracklist
	}
	return nil
}

func (m *Release) GetInstanceNotes() []*Note {
	if m != nil {
		return m.InstanceNotes
	}
	return nil
}

func (m *Release) GetRecordCondition() string {
	if m != nil {
		return m.RecordCondition
	}
	return ""
}

func (m *Release) GetSleeveCondition() string {
	if m != nil {
		return m.SleeveCondition
	}
	return ""
}

type Note struct {
	FieldId              int32    `protobuf:"varint,1,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Note) Reset()         { *m = Note{} }
func (m *Note) String() string { return proto.CompactTextString(m) }
func (*Note) ProtoMessage()    {}
func (*Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{6}
}

func (m *Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Note.Unmarshal(m, b)
}
func (m *Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Note.Marshal(b, m, deterministic)
}
func (m *Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Note.Merge(m, src)
}
func (m *Note) XXX_Size() int {
	return xxx_messageInfo_Note.Size(m)
}
func (m *Note) XXX_DiscardUnknown() {
	xxx_messageInfo_Note.DiscardUnknown(m)
}

var xxx_messageInfo_Note proto.InternalMessageInfo

func (m *Note) GetFieldId() int32 {
	if m != nil {
		return m.FieldId
	}
	return 0
}

func (m *Note) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Track struct {
	Position             string          `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Title                string          `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type_                string          `protobuf:"bytes,3,opt,name=type_,json=type,proto3" json:"type_,omitempty"`
	TrackType            Track_TrackType `protobuf:"varint,4,opt,name=track_type,json=trackType,proto3,enum=godiscogs.Track_TrackType" json:"track_type,omitempty"`
	SubTracks            []*Track        `protobuf:"bytes,5,rep,name=sub_tracks,json=subTracks,proto3" json:"sub_tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Track) Reset()         { *m = Track{} }
func (m *Track) String() string { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()    {}
func (*Track) Descriptor() ([]byte, []int) {
	return fileDescriptor_579be35d7413271f, []int{7}
}

func (m *Track) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Track.Unmarshal(m, b)
}
func (m *Track) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Track.Marshal(b, m, deterministic)
}
func (m *Track) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Track.Merge(m, src)
}
func (m *Track) XXX_Size() int {
	return xxx_messageInfo_Track.Size(m)
}
func (m *Track) XXX_DiscardUnknown() {
	xxx_messageInfo_Track.DiscardUnknown(m)
}

var xxx_messageInfo_Track proto.InternalMessageInfo

func (m *Track) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Track) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Track) GetType_() string {
	if m != nil {
		return m.Type_
	}
	return ""
}

func (m *Track) GetTrackType() Track_TrackType {
	if m != nil {
		return m.TrackType
	}
	return Track_UNKNOWN
}

func (m *Track) GetSubTracks() []*Track {
	if m != nil {
		return m.SubTracks
	}
	return nil
}

func init() {
	proto.RegisterEnum("godiscogs.SaleState", SaleState_name, SaleState_value)
	proto.RegisterEnum("godiscogs.Track_TrackType", Track_TrackType_name, Track_TrackType_value)
	proto.RegisterType((*Label)(nil), "godiscogs.Label")
	proto.RegisterType((*Folder)(nil), "godiscogs.Folder")
	proto.RegisterType((*Artist)(nil), "godiscogs.Artist")
	proto.RegisterType((*Image)(nil), "godiscogs.Image")
	proto.RegisterType((*Format)(nil), "godiscogs.Format")
	proto.RegisterType((*Release)(nil), "godiscogs.Release")
	proto.RegisterType((*Note)(nil), "godiscogs.Note")
	proto.RegisterType((*Track)(nil), "godiscogs.Track")
}

func init() { proto.RegisterFile("godiscogs.proto", fileDescriptor_579be35d7413271f) }

var fileDescriptor_579be35d7413271f = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xda, 0x4a,
	0x10, 0x8d, 0x01, 0x03, 0x1e, 0x08, 0x38, 0x9b, 0xdc, 0xab, 0xbd, 0xb9, 0x0f, 0x17, 0xf9, 0xe5,
	0xd2, 0xaf, 0xb4, 0x4a, 0xd5, 0x56, 0x7d, 0x44, 0xf9, 0x68, 0x51, 0x22, 0x50, 0x17, 0xaa, 0x3e,
	0x5a, 0x0b, 0x5e, 0xd0, 0xaa, 0xc6, 0x26, 0xde, 0x25, 0x0a, 0xbf, 0xad, 0x7f, 0xa9, 0x3f, 0xa2,
	0xda, 0x59, 0xdb, 0x49, 0x4b, 0x2a, 0xf5, 0x05, 0xed, 0x9c, 0x39, 0x3e, 0xb3, 0x73, 0x76, 0x06,
	0xe8, 0x2e, 0xd3, 0x48, 0xaa, 0x79, 0xba, 0x54, 0x27, 0xeb, 0x2c, 0xd5, 0x29, 0xf1, 0x4a, 0x20,
	0x18, 0x80, 0x7b, 0xcd, 0x67, 0x22, 0x26, 0x04, 0x6a, 0x09, 0x5f, 0x09, 0xea, 0xf4, 0x9c, 0xbe,
	0xc7, 0xf0, 0x4c, 0x8e, 0xc0, 0x9d, 0x73, 0x9d, 0xa4, 0xb4, 0x82, 0xa0, 0x0d, 0x48, 0x07, 0x2a,
	0x32, 0xa2, 0xd5, 0x9e, 0xd3, 0x77, 0x59, 0x45, 0x46, 0xc1, 0x73, 0xa8, 0x5f, 0xa6, 0x71, 0x24,
	0xb2, 0x3c, 0xe3, 0x14, 0x99, 0x52, 0xb3, 0x72, 0xaf, 0x69, 0xd8, 0x83, 0x4c, 0x4b, 0xa5, 0xff,
	0x88, 0xfd, 0x02, 0xdc, 0xe1, 0x8a, 0x2f, 0x05, 0xf1, 0xa1, 0xba, 0xc9, 0x64, 0x7e, 0x3b, 0x73,
	0x34, 0x74, 0xbd, 0x5d, 0x97, 0x74, 0x73, 0x0e, 0x16, 0xe6, 0x2a, 0xd9, 0x8a, 0x6b, 0x12, 0x40,
	0x3b, 0x12, 0x6a, 0x9e, 0xc9, 0xb5, 0x96, 0x69, 0xa2, 0xa8, 0xd3, 0xab, 0xf6, 0x3d, 0xf6, 0x13,
	0xf6, 0x58, 0x41, 0x53, 0xe7, 0x46, 0x6f, 0xb1, 0x3b, 0x8f, 0x99, 0x23, 0xd6, 0x11, 0x77, 0x9a,
	0xd6, 0xf2, 0x3a, 0xe2, 0x4e, 0x07, 0xdf, 0x5c, 0x68, 0x30, 0x11, 0x0b, 0xae, 0xc4, 0x4e, 0x1b,
	0x47, 0xe0, 0x6a, 0xa9, 0xe3, 0x42, 0xd6, 0x06, 0xe4, 0x19, 0x34, 0x38, 0xb6, 0xad, 0x68, 0xb5,
	0x57, 0xed, 0xb7, 0x4e, 0x0f, 0x4e, 0xee, 0x5f, 0xc5, 0x1a, 0xc2, 0x0a, 0x06, 0xf9, 0x17, 0xbc,
	0x05, 0x3a, 0x1a, 0xca, 0x08, 0xeb, 0xba, 0xac, 0x69, 0x81, 0x61, 0x44, 0xfa, 0x50, 0x97, 0xc6,
	0x12, 0x45, 0x5d, 0x14, 0xf2, 0x1f, 0x08, 0xa1, 0x57, 0x2c, 0xcf, 0x93, 0xff, 0xa0, 0x25, 0x13,
	0xa5, 0x79, 0x32, 0x17, 0x46, 0xa8, 0x8e, 0x42, 0x50, 0x40, 0x56, 0x2a, 0x36, 0x8f, 0xaf, 0x68,
	0x63, 0x47, 0x0a, 0xa7, 0x82, 0xe5, 0x79, 0xf2, 0x3f, 0x74, 0x17, 0x68, 0x6c, 0x78, 0xb3, 0xe1,
	0x89, 0x96, 0x7a, 0x4b, 0x9b, 0x28, 0xd7, 0xb1, 0xf0, 0xa7, 0x1c, 0x25, 0x7f, 0x43, 0x3d, 0xe3,
	0x5a, 0x26, 0x4b, 0xea, 0x61, 0x3e, 0x8f, 0xc8, 0x29, 0xfc, 0x25, 0x78, 0x16, 0x4b, 0xa1, 0x74,
	0x98, 0x59, 0xe7, 0xc2, 0x88, 0x6b, 0x41, 0xa1, 0xe7, 0xf4, 0xab, 0xec, 0xb0, 0x48, 0xe6, 0xae,
	0x9e, 0x73, 0x2d, 0x8c, 0x0d, 0x2b, 0xae, 0xb4, 0xb5, 0xa1, 0x65, 0x6d, 0xb0, 0xc0, 0x30, 0x22,
	0xc7, 0xd0, 0xcc, 0x75, 0x22, 0xda, 0x46, 0xa7, 0xcb, 0xd8, 0x98, 0x6d, 0xaf, 0xa5, 0xe8, 0xfe,
	0x8e, 0xd9, 0x76, 0x40, 0x58, 0xc1, 0x30, 0x42, 0x4b, 0xae, 0x85, 0xf1, 0x97, 0x76, 0x7a, 0x4e,
	0xbf, 0xc9, 0xca, 0xd8, 0x74, 0x33, 0x4b, 0xef, 0x94, 0xd0, 0xb4, 0x8b, 0x99, 0x3c, 0x22, 0x27,
	0xe0, 0xe9, 0x8c, 0xcf, 0xbf, 0xc6, 0x52, 0x69, 0xea, 0xef, 0x78, 0x37, 0x35, 0x39, 0x76, 0x4f,
	0x21, 0x6f, 0xa1, 0x53, 0xbe, 0x44, 0x92, 0x6a, 0xa1, 0xe8, 0x21, 0x7e, 0xd4, 0x7d, 0xf0, 0xd1,
	0x28, 0xd5, 0x82, 0xed, 0x17, 0x34, 0x13, 0x29, 0xf2, 0x04, 0xfc, 0x4c, 0xcc, 0xd3, 0x2c, 0x0a,
	0xe7, 0x69, 0x12, 0x49, 0x33, 0xb6, 0xf4, 0x00, 0x9b, 0xed, 0x5a, 0xfc, 0xac, 0x80, 0x0d, 0x55,
	0xc5, 0x42, 0xdc, 0x8a, 0x07, 0x54, 0x62, 0xa9, 0x16, 0x2f, 0xa9, 0xc1, 0x3b, 0xa8, 0x19, 0x79,
	0xf2, 0x0f, 0x34, 0x17, 0x52, 0xc4, 0x51, 0x58, 0xce, 0x6f, 0x03, 0xe3, 0x21, 0x0e, 0xf1, 0x2d,
	0x8f, 0x37, 0xe5, 0x10, 0x63, 0x10, 0x7c, 0x77, 0xc0, 0xc5, 0xde, 0x8c, 0x69, 0xeb, 0x54, 0xd9,
	0x2a, 0x76, 0x27, 0xcb, 0xf8, 0x37, 0x0b, 0x70, 0x08, 0xae, 0x59, 0xd1, 0x30, 0x5f, 0x2d, 0xdc,
	0x57, 0xf2, 0x1e, 0x00, 0x4d, 0x0a, 0x71, 0x93, 0xcd, 0xa4, 0x77, 0x4e, 0x8f, 0x7f, 0x35, 0xd2,
	0xfe, 0x4e, 0xb7, 0x6b, 0x91, 0x5b, 0x6a, 0x8e, 0xe4, 0x25, 0x80, 0xda, 0xcc, 0x42, 0x04, 0x1e,
	0x5b, 0x85, 0xfc, 0x0d, 0xd4, 0x66, 0x86, 0x27, 0x15, 0xbc, 0x02, 0xaf, 0x14, 0x22, 0x2d, 0x68,
	0x7c, 0x1e, 0x5d, 0x8d, 0xc6, 0x5f, 0x46, 0xfe, 0x1e, 0xf1, 0xc0, 0x9d, 0xb2, 0xc1, 0xd9, 0x95,
	0xef, 0x18, 0xfc, 0xe3, 0xc5, 0xe0, 0x7c, 0x38, 0xfa, 0xe0, 0x57, 0x9e, 0xbe, 0x01, 0x6f, 0xc2,
	0x63, 0x31, 0xd1, 0x66, 0x18, 0x7d, 0x68, 0x8f, 0xc6, 0xd3, 0xf0, 0x72, 0xcc, 0xc2, 0xc9, 0xe0,
	0xfa, 0xc2, 0xdf, 0x23, 0x6d, 0x68, 0x96, 0x91, 0x43, 0x9a, 0x50, 0x9b, 0x8c, 0xaf, 0xcf, 0xfd,
	0xca, 0xac, 0x8e, 0x7f, 0xb2, 0xaf, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xac, 0xfb, 0xdf,
	0x77, 0x05, 0x00, 0x00,
}
