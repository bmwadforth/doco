package src

func New(size PageSize) *Doco {
	return newDoco(size)
}

func newDoco(size PageSize) *Doco {
	doco := Doco{pageSize:size, pageCount:1}

	//Standard Unit Defaults To Millimeters
	switch size {
	case A0:
		doco.dimensions = DocumentDimensions{DocumentWidth: 841, DocumentHeight: 1189}
	case A1:
		doco.dimensions = DocumentDimensions{DocumentWidth: 594, DocumentHeight: 841}
	case A2:
		doco.dimensions = DocumentDimensions{DocumentWidth: 420, DocumentHeight: 594}
	case A3:
		doco.dimensions = DocumentDimensions{DocumentWidth: 297, DocumentHeight: 420}
	case A4:
		doco.dimensions = DocumentDimensions{DocumentWidth: 210, DocumentHeight: 297}
	}

	return &doco
}


//Receivers
func (d *Doco) GetDimensions() (width DocumentWidth, height DocumentHeight) {
	return d.dimensions.DocumentWidth, d.dimensions.DocumentHeight
}