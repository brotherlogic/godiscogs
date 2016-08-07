// Code generated by protoc-gen-go.
// source: godiscogs.proto
// DO NOT EDIT!

/*
Package godiscogs is a generated protocol buffer package.

It is generated from these files:
	godiscogs.proto

It has these top-level messages:
	Label
	Folder
	Artist
	Image
	Release
*/
package godiscogs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Label struct {
	// The name of the label
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Folder struct {
	// The id number of the folder
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The name of the folder
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Folder) Reset()                    { *m = Folder{} }
func (m *Folder) String() string            { return proto.CompactTextString(m) }
func (*Folder) ProtoMessage()               {}
func (*Folder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Artist struct {
	// The id number of the artist
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The name of the artist
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Image struct {
	// The uri to the image
	Uri string `protobuf:"bytes,1,opt,name=uri" json:"uri,omitempty"`
	// The type of image
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Release struct {
	// The id number of the release
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The title of the release
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// Artists associated with the release
	Artists []*Artist `protobuf:"bytes,3,rep,name=artists" json:"artists,omitempty"`
	// The folder in which the record is stored
	FolderId int32 `protobuf:"varint,4,opt,name=folder_id,json=folderId" json:"folder_id,omitempty"`
	// Images associated with the release
	Images []*Image `protobuf:"bytes,5,rep,name=images" json:"images,omitempty"`
	// The instance id of this release
	InstanceId int32 `protobuf:"varint,6,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
	// The labels connected to this release
	Labels []*Label `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty"`
	// The number of discs in the release
	FormatQuantity int32 `protobuf:"varint,8,opt,name=format_quantity,json=formatQuantity" json:"format_quantity,omitempty"`
}

func (m *Release) Reset()                    { *m = Release{} }
func (m *Release) String() string            { return proto.CompactTextString(m) }
func (*Release) ProtoMessage()               {}
func (*Release) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Release) GetArtists() []*Artist {
	if m != nil {
		return m.Artists
	}
	return nil
}

func (m *Release) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Release) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "godiscogs.Label")
	proto.RegisterType((*Folder)(nil), "godiscogs.Folder")
	proto.RegisterType((*Artist)(nil), "godiscogs.Artist")
	proto.RegisterType((*Image)(nil), "godiscogs.Image")
	proto.RegisterType((*Release)(nil), "godiscogs.Release")
}

func init() { proto.RegisterFile("godiscogs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x03, 0xa5, 0x2d, 0x1d, 0x12, 0xc0, 0x89, 0x87, 0x26, 0x1c, 0x24, 0xbd, 0xd8, 0x44,
	0xe5, 0xa0, 0x4f, 0xe0, 0xc5, 0x84, 0xc4, 0x8b, 0xfb, 0x02, 0x64, 0xa1, 0x4b, 0xb3, 0x49, 0xdb,
	0xc5, 0xdd, 0xe5, 0xc0, 0x93, 0xf8, 0xba, 0x76, 0x67, 0xdb, 0xaa, 0xe1, 0xe2, 0xed, 0xeb, 0x37,
	0x5f, 0x7f, 0x3b, 0x7f, 0x60, 0x51, 0xaa, 0x42, 0x9a, 0x83, 0x2a, 0xcd, 0xe6, 0xa4, 0x95, 0x55,
	0x98, 0x0c, 0x46, 0xb6, 0x82, 0xf0, 0x9d, 0xef, 0x45, 0x85, 0x08, 0x93, 0x86, 0xd7, 0x22, 0x1d,
	0xad, 0x47, 0x79, 0xc2, 0x48, 0x67, 0x8f, 0x10, 0xbd, 0xa9, 0xaa, 0x10, 0x1a, 0xe7, 0x30, 0x96,
	0x05, 0xd5, 0x42, 0xd6, 0xaa, 0x21, 0x3d, 0xfe, 0x9b, 0x7e, 0xd5, 0x56, 0x1a, 0xfb, 0xaf, 0xf4,
	0x13, 0x84, 0xdb, 0x9a, 0x97, 0x02, 0x97, 0x10, 0x9c, 0xb5, 0xec, 0xde, 0x75, 0xd2, 0xc5, 0xed,
	0xe5, 0x34, 0xc4, 0x9d, 0xce, 0xbe, 0xc6, 0x10, 0x33, 0x51, 0x09, 0x6e, 0xc4, 0x15, 0xfe, 0x16,
	0x42, 0x2b, 0x6d, 0xd5, 0xff, 0xe0, 0x3f, 0xf0, 0x01, 0x62, 0x4e, 0xed, 0x98, 0x34, 0x58, 0x07,
	0xf9, 0xec, 0xf9, 0x66, 0xf3, 0xb3, 0x07, 0xdf, 0x28, 0xeb, 0x13, 0xb8, 0x82, 0xe4, 0x48, 0x93,
	0xee, 0x5a, 0xf2, 0x84, 0xc8, 0x53, 0x6f, 0x6c, 0x0b, 0xcc, 0x21, 0x92, 0xae, 0x55, 0x93, 0x86,
	0x04, 0x5a, 0xfe, 0x02, 0xd1, 0x0c, 0xac, 0xab, 0xe3, 0x1d, 0xcc, 0x64, 0x63, 0x2c, 0x6f, 0x0e,
	0xc2, 0x81, 0x22, 0x02, 0x41, 0x6f, 0x79, 0x54, 0xe5, 0xd6, 0x6d, 0xd2, 0xf8, 0x0a, 0x45, 0x77,
	0x60, 0x5d, 0x1d, 0xef, 0x61, 0x71, 0x54, 0xba, 0xe6, 0x76, 0xf7, 0x79, 0xe6, 0x4d, 0x3b, 0xd2,
	0x25, 0x9d, 0x12, 0x6e, 0xee, 0xed, 0x8f, 0xce, 0xdd, 0x47, 0x74, 0xd3, 0x97, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xff, 0xe5, 0xc0, 0x66, 0xe6, 0x01, 0x00, 0x00,
}
