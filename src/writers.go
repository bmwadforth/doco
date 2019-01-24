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
		buff.Write([]byte(fmt.Sprintf("%d %d obj\n", obj.ObjectNumber, obj.GenerationNumber)))
		buff.Write([]byte("<<\n"))
		if obj.Data != nil {
			buff.Write([]byte(fmt.Sprintf("%s\n", obj.Data.(string))))
		} else {
			for _, el := range obj.Dictionary {
				for key, val := range el {
					buff.Write([]byte(fmt.Sprintf("%s %s\n", key, val)))
				}
			}
		}
		buff.Write([]byte(">>\n"))
		buff.Write([]byte("endobj\n"))
	}
	d.buffer.Write(buff.Bytes())
}

func (d *Doco) writeCrossRef() {
	d.buildCrossRef()

	buff := bytes.NewBuffer(make([]byte, 0))
	buff.Write([]byte("xref\n"))
	buff.Write([]byte(fmt.Sprintf("%d %d\n", d.CrossReference.FirstObject, d.CrossReference.Count)))

	for _, ref := range d.CrossReference.References {
		buff.Write([]byte(fmt.Sprintf("%010d %05d %s \n", ref.GenerationNumber, ref.ByteOffset, string(ref.RefFlag))))
	}

	d.buffer.Write(buff.Bytes())
}


func (d *Doco) writeTrailer() {
	d.buildTrailer()

	buff := bytes.NewBuffer(make([]byte, 0))
	buff.Write([]byte("trailer\n"))
	buff.Write([]byte("<<\n"))
	buff.Write([]byte(fmt.Sprintf("/Size %d\n", d.Trailer.Size)))
	buff.Write([]byte(fmt.Sprintf("/Root %s\n", GenerateIndirectReference(d.Trailer.Root))))
	buff.Write([]byte(">>\n"))
	buff.Write([]byte("startxref\n"))
	buff.Write([]byte(fmt.Sprintf("%d\n", d.currentPosition)))
	buff.Write([]byte("%%EOF"))

	d.buffer.Write(buff.Bytes())
}

