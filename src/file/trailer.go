package file

import (
	"bytes"
	"doco/src/pdf"
	"fmt"
)

type Trailer struct {
	Size       uint32
	RootObject pdf.ObjectReference
	XrefStart  uint32
}

func (t *Trailer) ToBytes() []byte {
	trailerBytes := bytes.NewBufferString("trailer")
	trailerBytes.Write([]byte(fmt.Sprintf("<< /Size %s", string(t.Size))))
	trailerBytes.Write([]byte(fmt.Sprintf("/Root %s", t.RootObject.Format())))
	trailerBytes.Write([]byte(fmt.Sprintf(">>")))
	trailerBytes.Write([]byte(fmt.Sprintf("startxref")))
	trailerBytes.Write([]byte(fmt.Sprintf("%s", string(t.XrefStart))))
	trailerBytes.Write([]byte(fmt.Sprintf("%s%sEOF", string(PercentSign), string(PercentSign))))
	return trailerBytes.Bytes()
}
