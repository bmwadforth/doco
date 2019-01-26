package src

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

func New(size PaperSize) *Doco {
	return newDoco(size)
}

func newDoco(size PaperSize) *Doco {
	doco := Doco{
		Meta: DocumentMeta{
			PaperSize: size,
			Version:   "1.7",
			Unit:      Millimeters,
		},
		PageTrees: []PageTree{},
		Pages:     []Page{},
		Errors:    []DocumentError{},
		buffer:    bytes.NewBuffer(make([]byte, 0)),
	}

	switch size {
	case A0:
		doco.Meta.Dimensions = DocumentDimensions{Height: 1189, Width: 841}
	case A1:
		doco.Meta.Dimensions = DocumentDimensions{Height: 841, Width: 594}
	case A2:
		doco.Meta.Dimensions = DocumentDimensions{Height: 594, Width: 420}
	case A3:
		doco.Meta.Dimensions = DocumentDimensions{Height: 420, Width: 297}
	case A4:
		doco.Meta.Dimensions = DocumentDimensions{Height: 297, Width: 210}
	}

	//TODO:
	//page tree should have method for appending pages
	initialPageTree := PageTree{
		Object: DocumentObject{
			ObjectType:       pageTree,
			ObjectNumber:     2,
			GenerationNumber: 0,
			Dictionary:       []map[string]string{},
		},
		parent: nil,
		pages:  &[]Page{},
	}

	//TODO:
	//page struct should have method for associating/disassociating page with pagetree
	initialPage := Page{
		Object: DocumentObject{
			ObjectType:       page,
			ObjectNumber:     3,
			GenerationNumber: 0,
		},
		mediaBox: Rectangle{
			upperRightX: StdUnitToPoint(float32(doco.Meta.Dimensions.Width), doco.Meta.Unit),
			upperRightY: StdUnitToPoint(float32(doco.Meta.Dimensions.Height), doco.Meta.Unit),
		},
		parent: &initialPageTree,
	}

	*initialPageTree.pages = append(*initialPageTree.pages, initialPage)

	pgeTreeAddErr := doco.addPageTree(initialPageTree)
	if pgeTreeAddErr != nil {
		panic(fmt.Sprintf("%v", pgeTreeAddErr))
	}

	pgeAddErr := doco.addPage(initialPage)
	if pgeAddErr != nil {
		panic(fmt.Sprintf("%v", pgeAddErr))
	}

	docCatalog := DocumentCatalog{
		Object: DocumentObject{
			ObjectType:       catalog,
			ObjectNumber:     1,
			GenerationNumber: 0,
		},
		Pages:      &doco.Pages,
		PageLayout: Single,
	}
	docCatalogErr := doco.addDocumentCatalog(docCatalog)
	if docCatalogErr != nil {
		panic(fmt.Sprintf("%v", docCatalogErr))
	}

	//set font
	//

	return &doco
}

//Public Receivers
func (d *Doco) Output() string {
	return d.buffer.String()
}

func (d *Doco) Build() error {
	if len(d.Errors) > 1 {
		return errors.New(fmt.Sprintf("unable to build, %d error(s) exist", len(d.Errors)))
	}

	d.writeHeader()
	d.writeBody()
	d.writeCrossRef()
	d.writeTrailer()
	return nil
}

func (d *Doco) Save(fileName string) error {
	err := d.Build()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable To Build PDF: %v", err))
	}

	writeErr := ioutil.WriteFile(fileName, d.buffer.Bytes(), 0777)
	if writeErr != nil {
		return errors.New(fmt.Sprintf("Unable To Save PDF: %v", err))
	}

	return nil
}

func (d *Doco) SetError(err error) uint {
	if len(d.Errors) > 1 {
		newErrorId := d.Errors[len(d.Errors)-1].Id + 1
		d.Errors = append(d.Errors, DocumentError{Id: newErrorId, Error: err})
		return newErrorId
	}
	d.Errors = append(d.Errors, DocumentError{Id: 1, Error: err})
	return 1
}

func (d *Doco) ClearError(idToRemove uint) {
	newErrors := make([]DocumentError, 0)
	for _, err := range d.Errors {
		if err.Id != idToRemove {
			newErrors = append(newErrors, err)
		}
	}
	d.Errors = newErrors
}
