package src

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

func New(size PageSize) *Doco {
	return newDoco(size)
}

func newDoco(size PageSize) *Doco {
	doco := Doco{PageSize: size, PageCount:1}
	doco.buffer = bytes.NewBuffer(make([]byte, 0))
	return &doco
}


//Public Receivers
func (d *Doco) Build() error {
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

func (d *Doco) Output() string {
	return d.buffer.String()
}