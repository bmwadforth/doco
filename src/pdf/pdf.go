package pdf

import (
	"bytes"
	"doco/src/file"
	"doco/src/to_be_named"
)

type Doco struct {
	Header *file.Header
	Body *file.Body
	Xref *file.Xref
	Trailer *file.Trailer
}

func New(v file.Version) (*Doco, error) {
	doco := Doco{
		Header:  &file.Header{Version: v},
		Body:    &file.Body{},
		Xref:    &file.Xref{},
		Trailer: &file.Trailer{
			Size:       0,
			RootObject: to_be_named.ObjectReference{
				ObjectNumber:  0,
				VersionNumber: 0,
			},
			XrefStart:  0,
		},
	}

	return &doco, nil
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