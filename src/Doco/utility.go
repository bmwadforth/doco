package Doco

import (
	. "doco/src/Pages"
)

func CalculatePoints(pageType PageType) (uint, uint) {
	switch pageType {
	case A3:
		panic("Not Implemented Yet")
	case A4:
		return 595, 842
	}

	return 0, 0
}