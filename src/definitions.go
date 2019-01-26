package src

import (
	"bytes"
	"fmt"
	"time"
)

//Interfaces
type IDoco interface {
	Build() error
	Save(fileName string) error
}

//Custom Types

//The Type Of Object.
//Example: /Page, /Pages, /Font, etc.
type ObjectType uint

//A Typical Dictionary Type With A Key And Value Pair
//Example: /Content 1 0 R
type Dictionary map[string]interface{}

//The PaperSize, i.e A4
type PaperSize uint

//A Standard Unit, Can Be Millimeters, Pixels, etc.
type StandardUnit uint

//The Document Width In Standard Unit (mm, px)
type DocumentWidth float32
//The Document Height In Standard Unit (mm, px)
type DocumentHeight float32

//Flag For Cross-Reference Section (F/N)
type Flag rune

//White-space Chars
type WhiteSpaceChars rune

//Delimiter Chars
type DelimiterChars rune

//Document Header
//Example: %PDF-1.7....
type DocumentHeader string

//Structs
//Core Doco Structs

type DocumentLayout uint


//A Struct Representing The Document Dimensions
type DocumentDimensions struct {
	Height DocumentHeight
	Width DocumentWidth
}

//A Struct Representing An Indirect Reference
type IndirectReference struct {
	ObjectNumber     uint
	GenerationNumber uint
}

type DocumentError struct {
	Id uint
	Error error
}

type DocumentMeta struct {
	Version string
	PaperSize PaperSize
	Unit StandardUnit
	Dimensions DocumentDimensions
}

//The Main Doco Struct. Represents a PDF Document
type Doco struct {
	//Public Members
	Meta DocumentMeta
	DocumentCatalog DocumentCatalog
	PageTrees []PageTree
	Pages      []Page
	Errors []DocumentError

	//Private Members
	header         DocumentHeader
	body           DocumentBody
	crossReference DocumentCrossReferenceTable
	trailer        DocumentTrailer
	buffer *bytes.Buffer
}

//Other Structs

func (doc *DocumentObject) genIndirectRef() string{
	return fmt.Sprintf("%d %d R", doc.ObjectNumber, doc.GenerationNumber)
}

//A Document Object.
type DocumentObject struct {
	ObjectType ObjectType
	ObjectNumber     uint
	GenerationNumber uint
	Dictionary       []map[string]string
	Data             interface{}
	ByteOffset       uint
}

//Document Catalog Object
type DocumentCatalog struct {
	Object DocumentObject
	Pages *[]Page
	PageLayout DocumentLayout
}

//A Page Tree
//Sits Above Pages
type PageTree struct {
	Object DocumentObject
	parent *PageTree
	pages  *[]Page
}

//A Page
func (p *Page) addResource(resource Resource) error {
	//check if old resource exists, if it does, merge argument with old one
	p.resources = resource
	return nil
}

type Page struct {
	Object DocumentObject
	parent         *PageTree
	lastModified   time.Time
	resources      Resource
	mediaBox       Rectangle
	cropBox        Rectangle
	bleedBox       Rectangle
	trimBox        Rectangle
	artBox         Rectangle
	boxColorInfo   Dictionary
	contents       interface{}
	rotate         uint
	group          Dictionary
	thumb          interface{}
	b              interface{}
	dur            uint
	trans          Dictionary
	annots         interface{}
	aa             Dictionary
	meta           interface{}
	placeInfo      Dictionary
	structParents  uint
	id             []byte
	pz             uint
	separationInfo Dictionary
	tabs interface{}
	templateInstantiated interface{}
	presSteps Dictionary
	userUnits uint
	vp Dictionary
}

//A Resource Dictionary
//Example: Pulled Into Page Struct, i.e. To Define Font
type Resource struct {
	extGState Dictionary
	colorSpace Dictionary
	pattern Dictionary
	shading Dictionary
	xObject Dictionary
	font Dictionary
	procSet []string
	properties Dictionary
}

//Represents Rectangle With Various Points
type Rectangle struct {
	lowerLeftY float32
	lowerLeftX float32
	upperRightX float32
	upperRightY float32
}








/*
type DocumentHeader struct {
	Version uint
}*/

//Redo how these structs are structured



//The Document Body
type DocumentBody struct {
	Count uint
	Objects []DocumentObject
}

//An Item In The Document Cross-Reference Table
type DocumentCrossRefItem struct {
	ByteOffset uint
	GenerationNumber uint
	RefFlag Flag
}

//The Document Cross-Reference Table
type DocumentCrossReferenceTable struct {
	FirstObject uint
	Count uint
	References []DocumentCrossRefItem
}

//The Document Trailer
type DocumentTrailer struct {
	Size uint
	Root IndirectReference
}