package src

import "fmt"

//TODO:
//Build function that outputs PDF structure formatted

func (d *Doco) buildHeader() {
	d.Header = DocumentHeader(fmt.Sprintf("%x%x%x%x%x", `%PDF-1.7`, LF, PERCENT, `\0xB5\0xB5\0xB5\0xB5`, LF))
}

func (d *Doco) buildBody() {
	d.Body.Objects = make([]DocumentObject, 5)
	//Generate First Object In Body (Catalog)
	d.Body.Objects[0] = DocumentObject{ObjectNumber: 1, GenerationNumber: 0, Dictionary: []map[string]string{{"/Type": "/Catalog", "/Pages": GenerateIndirectReference(IndirectReference{ObjectNumber: 2, GenerationNumber: 0},)}}}
	//Generate Page Tree
	d.Body.Objects[1] = DocumentObject{ObjectNumber: 2, GenerationNumber: 0, Dictionary: []map[string]string{{"/Type": "/Pages", "/Kids": "[3 0 R]", "/Count": "1"},}}
	//Generate Page Node
	d.Body.Objects[2] = DocumentObject{ObjectNumber: 3, GenerationNumber: 0, Dictionary: []map[string]string{{"/Type": "/Page", "/Parent": GenerateIndirectReference(IndirectReference{ObjectNumber: 2, GenerationNumber: 0}), "/MediaBox": "[0 0 612 729]", "/Contents": GenerateIndirectReference(IndirectReference{ObjectNumber: 4, GenerationNumber: 0},),}}}
	//Generate Page Contents
	d.Body.Objects[3] = DocumentObject{ObjectNumber: 4, GenerationNumber: 0, Data: "(Hello World)"}

	d.Body.Count = uint(len(d.Body.Objects))
	//TODO:
	//This Should Be function calls, generate catalog, foreach page - generate page (and its content)

}

func (d *Doco) buildCrossRef() {
	d.CrossReference = DocumentCrossReferenceTable{FirstObject: 0, Count: uint(len(d.Body.Objects))}
	d.CrossReference.References = []DocumentCrossRefItem{}

	for i, obj := range d.Body.Objects {
		d.CrossReference.References[i] = DocumentCrossRefItem{ByteOffset: obj.ByteOffset, GenerationNumber: obj.GenerationNumber, RefFlag:N}
	}
}

func (d *Doco) buildTrailer() {
	//Point To Root Object (Catalog)
	d.Trailer.Root = IndirectReference{ObjectNumber: 1, GenerationNumber: 0}
	d.Trailer.Size = uint(len(d.Body.Objects))
}
