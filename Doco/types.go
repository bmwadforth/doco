package Doco

type DocoFont struct {
	Name string
	FontName string
}

type DocoDimensions struct {
	Height uint
	Width uint
}

type DocoMeta struct {
	Font DocoFont
	Dimensions DocoDimensions
	PdfVersion float32
}

type Doco struct {
	Meta DocoMeta
	Pages []DocoPageTree
	Errors []error
}

type DocoPageTree struct {
	Children *[]DocoPage
}

type DocoPage struct {
	PageNumber uint
	Parent *DocoPageTree
}