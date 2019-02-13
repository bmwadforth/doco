package Doco

import "bytes"
//Interfaces
type Instance interface {
	SetMargin(margin Margin)
	Output() string
	Save(path string) error
	Write(contentType ContentType, data interface{})
}

type Margin struct {
	Top uint
	Right uint
	Bottom uint
	Left uint
}

//Custom Types
type ObjectType uint
type ContentType uint
type PaperSize uint
type Unit uint

type Core struct {
	Meta      Meta
	Catalog   Catalog
	Pages     []Page
	Errors []error
}

type Meta struct {
	Version    string
	Dimensions Dimensions
	Unit       Unit
}

type Dimensions struct {
	Width float32
	Height float32
}

type Catalog struct {
	RootPageTree *PageTree
}

//TypePage Structs
type PageTree struct {
	Parent *PageTree
	Pages *[]Page
}

type Page struct {
	Parent *PageTree
	Resources *PageResources
	Contents *PageContents
	*Core
}

type PageResources struct {
	Font *Font
}

type Font struct {
	BaseFont string
}

type PageContents struct {
	Type   ContentType
	Data   interface{}
	Length uint
}





//Raw Structs
//Used When Writing Data To Buffer For Raw File Structure
type RawHeader struct {
	Version string
	FileHasBinary bool
}

type RawBodyObject struct {
	ObjectType ObjectType
	ObjectNumber uint
	GenerationNumber uint
	Data interface{}
	Offset uint
}

type RawBody struct {
	Objects *[]RawBodyObject
}

type RawXrefObject struct {
	RefToBodyObject *RawBodyObject
	InUse bool
	Offset uint
}

type RawXref struct {
	Objects *[]RawXrefObject
}

type RawTrailer struct {
	Size uint
	FirstBodyObject *RawBodyObject
	LastXref *RawXrefObject
}

type Raw struct {
	Header  RawHeader
	Body    RawBody
	Xref    RawXref
	Trailer RawTrailer
	Buffer *bytes.Buffer
}