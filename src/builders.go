package src

import (
	"fmt"
	"unsafe"
)

//TODO:
//Build function that outputs PDF structure formatted

//Document Structure
func (d *Doco) addDocumentCatalog(catalog DocumentCatalog) error {
	d.DocumentCatalog = catalog
	return nil
}

func (d *Doco) addPageTree(tree PageTree) error {
	d.PageTrees = append(d.PageTrees, tree)
	return nil
}

func (d *Doco) addPage(p Page) error {
	d.Pages = append(d.Pages, p)
	return nil
}


//File Structure
func (d *Doco) buildHeader() {
	d.header = DocumentHeader(fmt.Sprintf("%%PDF-1.%s\n%%%s\n", d.Version, `\0xB5\0xB5\0xB5\0xB5`))
	d.currentPosition = uint(len([]byte(d.header)))
}

func (d *Doco) buildBody() {
	d.body.Objects = make([]DocumentObject, 4)
	//Generate First Object In body (Catalog)
	d.body.Objects[0] = DocumentObject{ObjectNumber: 1, GenerationNumber: 0, ByteOffset: d.currentPosition, Dictionary: []map[string]string{{"/Type": "/Catalog", "/Pages": GenerateIndirectReference(IndirectReference{ObjectNumber: 2, GenerationNumber: 0},)}}}
	d.currentPosition = d.currentPosition + uint(unsafe.Sizeof(d.body.Objects[0]))
	//Generate Page Tree
	d.body.Objects[1] = DocumentObject{ObjectNumber: 2, GenerationNumber: 0, ByteOffset: d.currentPosition, Dictionary: []map[string]string{{"/Type": "/Pages", "/Kids": "[3 0 R]", "/Count": "1"},}}
	d.currentPosition = d.currentPosition + uint(unsafe.Sizeof(d.body.Objects[1]))
	//Generate Page Node
	d.body.Objects[2] = DocumentObject{ObjectNumber: 3, GenerationNumber: 0, ByteOffset: d.currentPosition, Dictionary: []map[string]string{{"/Type": "/Page", "/Parent": GenerateIndirectReference(IndirectReference{ObjectNumber: 2, GenerationNumber: 0}), "/MediaBox": "[0 0 612 729]", "/Contents": GenerateIndirectReference(IndirectReference{ObjectNumber: 4, GenerationNumber: 0},),}}}
	d.currentPosition = d.currentPosition + uint(unsafe.Sizeof(d.body.Objects[2]))
	//Generate Page Contents
	d.body.Objects[3] = DocumentObject{ObjectNumber: 4, GenerationNumber: 0, ByteOffset: d.currentPosition, Data: "(Hello World)"}
	d.currentPosition = d.currentPosition + uint(unsafe.Sizeof(d.body.Objects[3]))

	d.body.Count = uint(len(d.body.Objects))
	//TODO:
	//This Should Be function calls, generate catalog, foreach page - generate page (and its content)

}

func (d *Doco) buildCrossRef() {
	d.crossReference = DocumentCrossReferenceTable{FirstObject: 0, Count: uint(len(d.body.Objects))}
	d.crossReference.References = make([]DocumentCrossRefItem, 4)
	for i, obj := range d.body.Objects {
		d.crossReference.References[i] = DocumentCrossRefItem{ByteOffset: obj.ByteOffset, GenerationNumber: obj.GenerationNumber, RefFlag:N}
	}
}

func (d *Doco) buildTrailer() {
	//Point To Root Object (Catalog)
	d.trailer.Root = IndirectReference{ObjectNumber: 1, GenerationNumber: 0}
	d.trailer.Size = uint(len(d.body.Objects))
}
