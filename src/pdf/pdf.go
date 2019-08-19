package pdf

import (
	"bytes"
	"doco/src/file"
)

type Doco struct {
	Header *file.Header
	Body *file.Body
	Xref *file.Xref
	Trailer *file.Trailer
}

func (d *Doco) Build() ([]byte, error) {
	docBytes := bytes.NewBuffer(make([]byte, 256))

	headerBytes := d.Header.ToBytes()
	bodyBytes := d.Body.ToBytes()
	xrefBytes := d.Xref.ToBytes()
	trailerBytes := d.Trailer.ToBytes()

	docBytes.Write(headerBytes)
	docBytes.Write(bodyBytes)
	docBytes.Write(xrefBytes)
	docBytes.Write(trailerBytes)

	return docBytes.Bytes(), nil
}