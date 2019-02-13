package Doco

import (
	"bytes"
	"fmt"
)

func (r *Raw) buildFrom(core Core) {
	r.Header = RawHeader{}
	r.Body = RawBody{}
	r.Xref = RawXref{}
	r.Trailer = RawTrailer{}
	r.Buffer = bytes.NewBuffer(make([]byte, 0))

	//Build Header
	r.Header.Version = core.Meta.Version
	//Look through core struct for binary data (i.e. text, images, etc, if true, set this flag)
	r.Header.FileHasBinary = true

	//Build Body
	//Write Catalog Object
	r.Body.Objects = &[]RawBodyObject{}
	*r.Body.Objects = append(*r.Body.Objects, RawBodyObject{
		ObjectType:TypeCatalog,
		ObjectNumber:1,
		GenerationNumber:0,
		Data:fmt.Sprintf("/Pages %s\n", "2 0 R"),
	})

	//Check parents in each page, if they all reference same parent, create single, predictable page tree,
	//else create balanced tree structure
	//Write PageTree Object
	*r.Body.Objects = append(*r.Body.Objects, RawBodyObject{
		ObjectType:TypePageTree,
		ObjectNumber:2,
		GenerationNumber:0,
		Data:fmt.Sprintf("/Kids [3 0 R]\n/Count 1\n"),
	})

	objectNum := 3
	for _, page := range core.Pages {
		*r.Body.Objects = append(*r.Body.Objects, RawBodyObject{
			ObjectType:TypePage,
			ObjectNumber:uint(objectNum),
			GenerationNumber:0,
			Data:page.generateObj(),
		})
		objectNum = objectNum + 1
	}

	*r.Body.Objects = append(*r.Body.Objects, RawBodyObject{
		ObjectType:TypeStreamObject,
		ObjectNumber:10,
		GenerationNumber:0,
		Data:fmt.Sprintf("/Length 43 >>\nstream\nBT\n/F1 12 Tf\n10 832 Td\n(Hello world!) Tj\n0 -12 TD\n(Hello world!) Tj\nET\nendstream\n"),
	})

	//Write Font Object
	*r.Body.Objects = append(*r.Body.Objects, RawBodyObject{
		ObjectType:TypeFont,
		ObjectNumber:1000,
		GenerationNumber:0,
		Data:fmt.Sprintf("/Subtype /Type1\n/BaseFont /Helvetica"),
	})


	r.Xref.Objects = &[]RawXrefObject{}
	for _, item := range *r.Body.Objects {
		*r.Xref.Objects = append(*r.Xref.Objects, RawXrefObject{InUse:true, RefToBodyObject:&item})
	}

	//Build Trailer
	r.Trailer.Size = uint(len(*r.Body.Objects))
	lastXref := *r.Xref.Objects
	r.Trailer.LastXref = &lastXref[len(*r.Xref.Objects) - 1]
	bodyObj := *r.Body.Objects
	r.Trailer.FirstBodyObject = &bodyObj[0]
}

func (r *Raw) writeHeader(){
	headerString := ""
	if r.Header.FileHasBinary {
		headerString = fmt.Sprintf("%%PDF-%s\n%%%s\n", r.Header.Version, `\0xB5\0xB5\0xB5\0xB5`)
	}  else {
		headerString = fmt.Sprintf("%%PDF-%s\n", r.Header.Version)
	}
	r.Buffer.WriteString(headerString)
}

func (r *Raw) writeBody(){
	bodyBuff := bytes.NewBuffer(make([]byte, 0))
	for _, obj := range *r.Body.Objects {
		obj.Offset = uint(len(r.Buffer.Bytes()))
		bodyBuff.WriteString(fmt.Sprintf("%d %d obj\n", obj.ObjectNumber, obj.GenerationNumber))
		bodyBuff.WriteString("<<\n")
		switch obj.ObjectType {
		case TypeCatalog:
			bodyBuff.WriteString("/Type /Catalog\n")
			bodyBuff.WriteString(obj.Data.(string))
		case TypePageTree:
			bodyBuff.WriteString("/Type /Pages\n")
			bodyBuff.WriteString(obj.Data.(string))
		case TypePage:
			bodyBuff.WriteString("/Type /Page\n")
			bodyBuff.WriteString(obj.Data.(string))
		case TypeFont:
			bodyBuff.WriteString("/Type /Font\n")
			bodyBuff.WriteString(obj.Data.(string))
		case TypeStreamObject:
			bodyBuff.WriteString(obj.Data.(string))
			bodyBuff.WriteString("endobj\n")
			r.Buffer.Write(bodyBuff.Bytes())
			bodyBuff.Reset()
			continue
		}
		bodyBuff.WriteString(">>\n")
		bodyBuff.WriteString("endobj\n")
		r.Buffer.Write(bodyBuff.Bytes())
		bodyBuff.Reset()
	}
}

func (r *Raw) writeXRef(){
	xRefBuff := bytes.NewBuffer(make([]byte, 0))
	xRefBuff.WriteString("xref\n")
	xRefBuff.WriteString(fmt.Sprintf("1 %d\n", len(*r.Body.Objects)))
	for _, refItem := range *r.Xref.Objects {
		refItem.Offset = uint(len(r.Buffer.Bytes()))
		inUseFlag := "f"
		if refItem.InUse {
			inUseFlag = "n"
		}
		xRefBuff.WriteString(fmt.Sprintf("%010d %05d %s\n", refItem.RefToBodyObject.Offset, refItem.RefToBodyObject.GenerationNumber, inUseFlag))
		r.Buffer.Write(xRefBuff.Bytes())
		xRefBuff.Reset()
	}
}

func (r *Raw) writeTrailer(){
	trailerBuff := bytes.NewBuffer(make([]byte, 0))
	trailerBuff.WriteString("trailer\n")
	trailerBuff.WriteString(fmt.Sprintf("<< /Size %d\n", r.Trailer.Size))
	trailerBuff.WriteString("/Root 1 0 R\n")
	trailerBuff.WriteString(">>\n")
	trailerBuff.WriteString("startxref\n")
	trailerBuff.WriteString(fmt.Sprintf("%d\n", r.Trailer.LastXref.Offset))
	trailerBuff.WriteString("%%EOF")
	r.Buffer.Write(trailerBuff.Bytes())
}
