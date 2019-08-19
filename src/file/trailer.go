package file

import (
	"bytes"
	"doco/src/to_be_named"
	"fmt"
)

type Trailer struct {
	Size       uint32
	RootObject to_be_named.ObjectReference
	XrefStart  uint32
}

func (t *Trailer) ToBytes() []byte {
	objectKeys := make(map[string]interface{})
	objectKeys["/Size"] = string(t.Size)
	objectKeys["/Root"] = t.RootObject.Format()

	trailerBytes := bytes.NewBufferString("trailer")
	trailerBytes.Write(to_be_named.WriteObject(objectKeys))
	trailerBytes.Write([]byte(fmt.Sprintf("startxref")))
	trailerBytes.Write([]byte(fmt.Sprintf("%s", string(t.XrefStart))))
	trailerBytes.Write([]byte(fmt.Sprintf("%s%sEOF", string(PercentSign), string(PercentSign))))
	return trailerBytes.Bytes()
}
