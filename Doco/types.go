package Doco

import "bytes"

type DocoPageType uint

const (
	A3 DocoPageType = 0
	A4 DocoPageType = 1
)

type DocoFont struct {
	Name string
	FontName string
}

type DocoDimension struct {
	Height uint
	Width uint
}

type DocoMargin struct {
	Top uint
	Right uint
	Bottom uint
	Left uint
}

type DocoMeta struct {
	PdfVersion float32
}

type Doco struct {
	Meta DocoMeta
	Pages []DocoPageTree
	CurrentPage *DocoPage
	BufferPosition uint
}

type DocoPageTree struct {
	Children *[]DocoPage
}

type DocoPage struct {
	PageNumber uint
	Font       DocoFont
	PageType DocoPageType
	Margin     DocoMargin
	Body *bytes.Buffer
	Parent *DocoPageTree
	Errors []error
}



//Stream Types
