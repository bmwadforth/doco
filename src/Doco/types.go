package Doco

import (
	. "doco/src/Pages"
)

type Meta struct {
	PdfVersion float32
}

type Document struct {
	Meta           Meta
	Pages          []PageTree
	CurrentPage    *Page
	BufferPosition uint
}

//Stream Types
