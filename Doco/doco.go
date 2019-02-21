package Doco

import "bytes"

func New(pageType DocoPageType) DocoSpecimen {
	docoMeta := DocoMeta{PdfVersion:1.7,}

	docoPageTree := DocoPageTree{Children: &[]DocoPage{}}
	docoPage := DocoPage{
		PageNumber:1,
		Parent:&docoPageTree,
		PageType:pageType,
		Font:DocoFont{},
		Margin:DocoMargin{0, 0, 0, 0},
		Body:bytes.NewBuffer(make([]byte, 0)),
		Errors:nil,
	}

	*docoPageTree.Children = append(*docoPageTree.Children, docoPage)

	doco := Doco{
		Meta:docoMeta,
		Pages:[]DocoPageTree{docoPageTree},
		CurrentPage:&docoPage,
		BufferPosition:0,
	}

	return &doco
}

