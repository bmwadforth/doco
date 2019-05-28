package Doco

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func (doco *Document) Save(path string) error {
	err := ioutil.WriteFile(path, doco.build(), 0777)
	if err != nil {
		return err
	}
	return nil
}

func (doco *Document) Write(content string) {
	//Check If Writing will cause page to leak into new page
	//If yes, then create new page, update pointer to 'current page', and begin writing


	bytesWritten := doco.CurrentPage.Write(content)
	doco.BufferPosition = bytesWritten
}

func (doco *Document) WriteHtml(html string) {
	panic("implement me")
}

func (doco *Document) Output() []byte {
	return doco.build()
}


func (doco *Document) build() []byte {
	buff := bytes.NewBuffer(make([]byte, 0))
	//HEADER
	buff.WriteString(fmt.Sprintf("%%PDF-%.01f\n", doco.Meta.PdfVersion))

	//BODY
	objBytes, objMeta := doco.writeObjects()
	buff.WriteString(string(objBytes))

	//XREF
	buff.WriteString("xref\n")
	startXref := len(buff.Bytes())
	buff.WriteString(fmt.Sprintf("1 %d\n", objMeta.Size))
	for i := 0; i < int(objMeta.Size); i++ {
		buff.WriteString(fmt.Sprintf("%010d %05d %s\n", 0, 0, "f"))
	}

	//TRAILER
	buff.WriteString(fmt.Sprintf("trailer\n<<\n/Size %d\n/Root 1 0 R\n>>\nstartxref\n%d\n%%%%EOF\n", 5, startXref))
	return buff.Bytes()
}


type ObjectMeta struct {
	Size uint
	Offsets []uint
}

func (doco *Document) writeObjects() ([]byte, ObjectMeta) {
	buff := bytes.NewBuffer(make([]byte, 0))

	objNum := 2
	objMeta := ObjectMeta{Size:uint(objNum)}

	if len(doco.Pages) == 1 {
		buff.WriteString("1 0 obj\n<<\n/Type /Catalog\n/Pages 3 0 R\n>>\nendobj\n")
	} else {
		//Foreach Page, Write It into catalog, and increment obj num
	}

	fontRefNum := objNum
	buff.WriteString(fmt.Sprintf("%d 0 obj\n<<\n/Type /Font\n/Subtype /Type1\n/Name /F1\n/BaseFont /Helvetica\n>>\nendobj\n", fontRefNum))

	objNum++

	//Page content ref, needs to be refactored
	contentRefNum := 1000

	data := ""

	for _, pageTree := range doco.Pages {
		childRefs := make([]uint, 0)
		pageBeginNum := objNum + 1
		tempPageBuff := bytes.NewBuffer(make([]byte, 0))

		for _, page := range *pageTree.Children {
			childRefs = append(childRefs, uint(pageBeginNum))
			width, height := CalculatePoints(page.PageType)

			data = page.Body.String()
			tempPageBuff.WriteString(fmt.Sprintf("%d 0 obj\n<<\n/Type /Page\n/Parent %d 0 R\n/MediaBox [%d %d %d %d]\n/Contents %d 0 R\n/Resources << /Font << /F1 %d 0 R >>\n>>\n>>\nendobj\n", pageBeginNum, objNum, 0, 0, width, height, contentRefNum, fontRefNum))
			pageBeginNum++
		}

		var pageTreeChildRef string
		for _, pgeRef := range childRefs {
			pageTreeChildRef = pageTreeChildRef + fmt.Sprintf(" %d 0 R ", pgeRef)
		}
		buff.WriteString(fmt.Sprintf("%d 0 obj\n<</Type /Pages\n/Kids [%s]\n/Count %d\n>>\nendobj\n", objNum, pageTreeChildRef, len(*pageTree.Children)))
		buff.WriteString(tempPageBuff.String())
		objNum = pageBeginNum
	}

	buff.WriteString(fmt.Sprintf("1000 0 obj\n<< /Length 73 >>\nstream\nBT\n/F1 12 Tf\n100 100 Td\n(%s) Tj\nET\nendstream\nendobj\n", data))


	//objMeta.Size = to how ever many objects were 'written'
	objMeta.Size = 5

	return buff.Bytes(), objMeta
}



