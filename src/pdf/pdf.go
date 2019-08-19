package pdf

import "bytes"

type Doco struct {
	Header *Header
	Body *Body
	Xref *Xref
	Trailer *Trailer
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