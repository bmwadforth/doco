package Pages

import "bytes"

type PageType uint

const (
	A3 PageType = 0
	A4 PageType = 1
)

type Font struct {
	Name     string
	FontName string
}

type Dimension struct {
	Height uint
	Width  uint
}

type Margin struct {
	Top    uint
	Right  uint
	Bottom uint
	Left   uint
}

type PageTree struct {
	Children *[]Page
}

type Page struct {
	PageNumber uint
	Font       Font
	PageType   PageType
	Margin     Margin
	Body       *bytes.Buffer
	Parent     *PageTree
	Errors     []error
}