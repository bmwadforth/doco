package Doco

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func New(size PaperSize) Instance {
	return newDoco(size)
}

func newDoco(size PaperSize) *Core {
	doco := &Core{
		Meta: Meta{
			Version: "1.7",
			Unit:    UnitMillimeters,
		},
		Pages: []Page{},
	}

	switch size {
	case A4:
		doco.Meta.Dimensions = Dimensions{Width: 210, Height:297}
	case A3:
		doco.Meta.Dimensions = Dimensions{Width: 297, Height:420}
	}

	initialFont := Font{
		BaseFont:"Helvetica",
	}

	initialPageTree := PageTree{
		Parent:nil,
		Pages:&[]Page{},
	}

	initialPage := Page{
		Parent:&initialPageTree,
		Resources:&PageResources{
			Font:&initialFont,
		},
		Contents:&PageContents{},
		Core:doco,
	}
	*initialPageTree.Pages = append(*initialPageTree.Pages, initialPage)

	initialDocCatalog := Catalog{
		RootPageTree:&initialPageTree,
	}

	doco.addCatalog(initialDocCatalog)
	doco.addPage(initialPage)

	return doco
}

//
func (d *Core) SetMargin(margin Margin) {

}

func (d *Core) Save(path string) error {
	if len(d.Errors) != 0 {
		return errors.New(fmt.Sprintf("%d Errors Exist. Please Fix Them Before Saving PDF", len(d.Errors)))
	}

	raw := Raw{}
	raw.buildFrom(*d)

	raw.writeHeader()
	raw.writeBody()
	raw.writeXRef()
	raw.writeTrailer()

	err := ioutil.WriteFile(path, raw.Buffer.Bytes(), 0777)
	if err != nil {
		return  err
	}
	return nil
}

func (d *Core) Output() string {
	raw := Raw{}
	raw.buildFrom(*d)

	raw.writeHeader()
	raw.writeBody()
	raw.writeXRef()
	raw.writeTrailer()
	return raw.Buffer.String()
}

func (d *Core) Write(contentType ContentType, data interface{}) {
	switch contentType {
	case Text:

	case Image:

	}
}