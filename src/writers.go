package src

import (
	"bytes"
	"fmt"
)

func (d *Doco) writeHeader() {
	d.buildHeader()
	d.buffer.Write([]byte(d.Header))
}

func (d *Doco) writeBody() {
	d.buildBody()

	buff := bytes.NewBuffer(make([]byte, 0))
	for _, obj := range d.Body.Objects {
		buff.Write([]byte(fmt.Sprintf("%d d obj%x", obj.ObjectNumber, obj.GenerationNumber, LF)))
		buff.Write([]byte(fmt.Sprintf("<<%x", LF)))
		for _, el := range obj.Dictionary {
			for key, val := range el {
				buff.Write([]byte(fmt.Sprintf("%s %s%x", key, val, LF)))
			}
		}
		buff.Write([]byte(fmt.Sprintf(">>%x", LF)))
		buff.Write([]byte(fmt.Sprintf("endobj%x", LF)))
	}
	d.buffer.Write(buff.Bytes())
}

func (d *Doco) writeCrossRef() {
	//Foreach Indirect Object, Write Line With The Byte Offset Of That Object
	//Foreach Cross-reference section

}

func (d *Doco) writeTrailer() {
	d.buildTrailer()

	buff := bytes.NewBuffer(make([]byte, 0))
	buff.Write([]byte(fmt.Sprintf("trailer%x", LF)))
	buff.Write([]byte(fmt.Sprintf("<<%x", LF)))
	buff.Write([]byte(fmt.Sprintf("/Size %d%x", d.Trailer.Size, LF)))
	buff.Write([]byte(fmt.Sprintf("/Root %s%x", GenerateIndirectReference(d.Trailer.Root), LF)))
	buff.Write([]byte(fmt.Sprintf(">>%x", LF)))
	buff.Write([]byte(fmt.Sprintf("startxref%x", LF)))
	buff.Write([]byte(fmt.Sprintf("%d%x", d.lastCrOffset, LF)))
	buff.Write([]byte("%%EOF"))

	d.buffer.Write(buff.Bytes())
}

