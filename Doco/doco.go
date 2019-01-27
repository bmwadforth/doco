package Doco

import (
	"bytes"
	"io/ioutil"
)

func New(size PaperSize) DocoInstance {
	return newDoco(size)
}

func newDoco(size PaperSize) *Doco {
	doco := &Doco{
		Meta:DocoMeta{
			Version:"1.7",
			Unit:Milimeters,
		},
		PageTrees: []DocoPageTree{},
		Pages: []DocoPage{},
		buffer:bytes.NewBuffer(make([]byte, 0)),
	}

	switch size {
	case A4:
		doco.Meta.Dimensions = DocoDimensions{Width:210, Height:297}
	}

	initialFont := DocoFont{
		BaseFont:"Helvetica",
	}

	initialPageTree := DocoPageTree{
		Parent:nil,
		Pages:&[]DocoPage{},
	}

	initialPage := DocoPage{
		Parent:&initialPageTree,
		Resources:&DocoPageResources{
			Font:&initialFont,
		},
	}
	*initialPageTree.Pages = append(*initialPageTree.Pages, initialPage)

	initialDocCatalog := DocoCatalog{
		RootPageTree:&initialPageTree,
	}

	doco.addCatalog(initialDocCatalog)
	doco.addPageTree(initialPageTree)
	doco.addPage(initialPage)

	return doco
}

//
func (d *Doco) SetMargin(margin DocoMargin) {

}

func (d *Doco) Save(path string) error {
	err := ioutil.WriteFile(path, d.buffer.Bytes(), 0777)
	if err != nil {
		return  err
	}
	return nil
}

func (d *Doco) Output() string {
	return d.buffer.String()
}

func (d *Doco) WriteText(text string) error {


	return nil
}