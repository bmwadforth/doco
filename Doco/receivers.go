package Doco

import (
	"bytes"
	"fmt"
)

//Core Receivers
func (d *Core) addCatalog(catalog Catalog) {
	d.Catalog = catalog
}

func (d *Core) addPage(page Page) {
	//page.setFont()
	d.Pages = append(d.Pages, page)
}

func (d *Page) generateObj() string {
	pgeBuff := bytes.NewBuffer(make([]byte, 0))
	pgeBuff.WriteString("/Parent 2 0 R\n")
	pgeBuff.WriteString(fmt.Sprintf("/MediaBox [0 0 %f %f]\n", StdUnitToPoint(d.Core.Meta.Dimensions.Width, d.Core.Meta.Unit), StdUnitToPoint(d.Core.Meta.Dimensions.Height, d.Core.Meta.Unit)))
	pgeBuff.WriteString("/Resources << \n")
	pgeBuff.WriteString(fmt.Sprintf("%s", d.Resources.generateObj()))
	pgeBuff.WriteString(">>\n")
	pgeBuff.WriteString("/Contents 10 0 R\n")
	return pgeBuff.String()
}

func (r *PageResources) generateObj() string {
	if r.Font != nil {
		return fmt.Sprintf("/Font\n<<\n /F1 1000 0 R\n>>\n")
	}
	panic("No Page Resources")
}

func (r *PageContents) generateObj() string {
	return ""
}