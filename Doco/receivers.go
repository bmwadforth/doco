package Doco

//Core Receivers
func (d *Core) addCatalog(catalog Catalog) {
	d.Catalog = catalog
}

func (d *Core) addPage(page Page) {
	//page.setFont()
	d.Pages = append(d.Pages, page)
}