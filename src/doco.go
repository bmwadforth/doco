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
		PaperSize: size,
		PageCount: 1,
		Version:   "1.7",
		Pages:     []Page{},
		buffer:    bytes.NewBuffer(make([]byte, 0)),
	}

	initialPageTree := PageTree{
		Object: DocumentObject{
			ObjectType:       pageTree,
			ObjectNumber:     1,
			GenerationNumber: 0,
			Dictionary:       []map[string]string{},
		},
		parent: nil,
		kids:   nil,
		count:  0,
	}

	initialPage := Page{Object: DocumentObject{
		ObjectType:       page,
		ObjectNumber:     2,
		GenerationNumber: 0,
	},
		parent: &initialPageTree,
	}

	pages := make([]Page, 1)
	pages = append(pages, initialPage)
	initialPageTree.kids = &pages

	pgeTreeAddErr := doco.addPageTree(initialPageTree)
	if pgeTreeAddErr != nil {
		panic(fmt.Sprintf("%v", pgeTreeAddErr))
	}

	pgeAddErr := doco.addPage(initialPage)
	if pgeAddErr != nil {
		panic(fmt.Sprintf("%v", pgeAddErr))
	}

	doco.buildHeader()

	docCatalog := DocumentCatalog{}
	err := doco.addDocumentCatalog(docCatalog)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	//add doc catalog
	//add page tree
	//add single page
	//set page dimensions
	//set font
	//

	return &doco
}

//Public Receivers
func (d *Doco) AddPage() error {
	err := d.addPage()
	if err != nil {

	}
	return nil
}

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
