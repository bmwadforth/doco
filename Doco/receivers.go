package Doco

import (
	"fmt"
)

//Doco Receivers
func (d *Doco) addCatalog(catalog DocoCatalog) {
	d.Catalog = catalog
}

func (d *Doco) addPageTree(tree DocoPageTree) {
	d.PageTrees = append(d.PageTrees, tree)
}

func (d *Doco) addPage(page DocoPage) {
	//page.setFont()
	d.Pages = append(d.Pages, page)
}





func (d *Doco) writeHeader() {
	d.buffer.Write([]byte(fmt.Sprintf("%%PDF-%s\n%%%s\n", d.Meta.Version, `\0xB5\0xB5\0xB5\0xB5`)))
}

func (d *Doco) writeBody(){

}

func(d *Doco) writeXref(){

}

func(d *Doco) writeTrailer(){

}