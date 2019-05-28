package Doco

import (
	"bytes"
	. "doco/src/Pages"
)

func New(pageType PageType) IDocument {
	docoMeta := Meta{PdfVersion: 1.7,}

	docoPageTree := PageTree{Children: &[]Page{}}
	docoPage := Page{
		PageNumber: 1,
		Parent:     &docoPageTree,
		PageType:   pageType,
		Font:       Font{},
		Margin:     Margin{},
		Body:       bytes.NewBuffer(make([]byte, 0)),
		Errors:     nil,
	}

	*docoPageTree.Children = append(*docoPageTree.Children, docoPage)

	doco := Document{
		Meta:docoMeta,
		Pages:[]PageTree{docoPageTree},
		CurrentPage:&docoPage,
		BufferPosition:0,
	}

	return &doco
}

