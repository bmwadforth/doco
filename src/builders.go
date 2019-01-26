package src

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
