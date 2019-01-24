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
	doco := Doco{pageSize:size, pageCount:1}

	//Standard Unit Defaults To Millimeters
	switch size {
	case A0:
		doco.dimensions = DocumentDimensions{DocumentWidth: 841, DocumentHeight: 1189}
	case A1:
		doco.dimensions = DocumentDimensions{DocumentWidth: 594, DocumentHeight: 841}
	case A2:
		doco.dimensions = DocumentDimensions{DocumentWidth: 420, DocumentHeight: 594}
	case A3:
		doco.dimensions = DocumentDimensions{DocumentWidth: 297, DocumentHeight: 420}
	case A4:
		doco.dimensions = DocumentDimensions{DocumentWidth: 210, DocumentHeight: 297}
	}

	buff := make([]byte, 64)
	doco.buffer = bytes.NewBuffer(buff)

	return &doco
}


//Receivers
func (d *Doco) GetDimensions() (width DocumentWidth, height DocumentHeight) {
	return d.dimensions.DocumentWidth, d.dimensions.DocumentHeight
}

func writeHeader() []byte {
	//If PDF has binary data in it
	return []byte(fmt.Sprintf("%x%x%x%x%x", `%PDF-1.7`, LF, PERCENT, `\0xB5\0xB5\0xB5\0xB5`, LF))
}

func writeBody() []byte {
	//If Binary Data
	return []byte(fmt.Sprintf("%x", `<5361792048656c6c6f20746f204d79204c6974746c6520467269656e64>`))
}

func writeCrossRefTable() []byte {
	//Foreach Indirect Object, Write Line With The Byte Offset Of That Object
	//Foreach Cross-reference section
	crossReferenceTable := make([]byte, 1)
	for i := 0; i < 1; i++ {
		crossReferenceTable = append(crossReferenceTable, fmt.Sprintf("%x", LF))
	}
	return []byte(fmt.Sprintf("%x", ``))
}

func writeTrailer() []byte {
	return []byte(fmt.Sprintf("%x", `%%EOF`))
}

func (d *Doco) Build() error {
	//See incremental updates
	d.buffer.Write(writeHeader())
	d.buffer.Write(writeBody())
	d.buffer.Write(writeCrossRefTable())
	d.buffer.Write(writeTrailer())
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