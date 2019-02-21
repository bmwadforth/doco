package Doco

func CalculatePoints(pageType DocoPageType) (uint, uint) {
	switch pageType {
	case A3:
		panic("Not Implemented Yet")
	case A4:
		return 595, 842
	}

	return 0, 0
}