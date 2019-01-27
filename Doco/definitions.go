package Doco

import "bytes"

//Interfaces

type DocoMargin struct {
	Top uint
	Right uint
	Bottom uint
	Left uint
}

type DocoInstance interface {
	SetMargin(margin DocoMargin)
	Output() string
	Save(path string) error
	WriteText(text string) error
	//Write
	//Etc
}

//Custom Types
type ObjectType uint
type ContentType uint
type PaperSize uint
type Unit uint

type Doco struct {
	Meta DocoMeta
	Catalog DocoCatalog
	PageTrees []DocoPageTree
	Pages []DocoPage
	buffer *bytes.Buffer
}

type DocoMeta struct {
	Version string
	Dimensions DocoDimensions
	Unit Unit
}

type DocoDimensions struct {
	Width float32
	Height float32
}

type DocoCatalog struct {
	RootPageTree *DocoPageTree
}

//Page Structs
type DocoPageTree struct {
	Parent *DocoPageTree
	Pages *[]DocoPage
}

type DocoPage struct {
	Parent *DocoPageTree
	Resources *DocoPageResources
	Contents *DocoPageContents
}

type DocoPageResources struct {
	Font *DocoFont
}

type DocoFont struct {
	BaseFont string
}

type DocoPageContents struct {
	Type   ContentType
	Data   interface{}
	Length uint
}





//Raw Structs
//Used When Writing Data To Buffer For Raw File Structure
type DocoRawHeader struct {
	Version string
	FileHasBinary bool
}

type DocoRawBodyObject struct {
	ObjectType ObjectType
	ObjectNumber uint
	GenerationNumber uint
	Data interface{}
	Offset uint
}

type DocoRawBody struct {
	Objects *[]DocoRawBodyObject
}

type DocoRawXrefObject struct {
	RefToBodyObject *DocoRawBodyObject
	InUse bool
}

type DocoRawXref struct {
	Objects *[]DocoRawXrefObject
}

type DocoRawTrailer struct {
	Root *DocoRawBody
	Size uint
	LastXref *DocoRawXref
}

type DocoRaw struct {
	Header DocoRawHeader
	Body DocoRawBody
	Xref DocoRawXref
	Trailer DocoRawTrailer
}