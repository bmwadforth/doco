package src

import (
	"fmt"
)

func GenerateIndirectReference(reference IndirectReference) string {
	return fmt.Sprintf("%d %d R", reference.ObjectNumber, reference.GenerationNumber)
}

func StdUnitToPoint(val float32, unit StandardUnit) float32 {
	switch unit {
	case Millimeters:
		return float32(val * 2.835)
	case Pixels:
		return float32(val * 0.75)
	}
	return 0
}