package src

import (
	"errors"
	"fmt"
)

func ConvertUnit(val int, fromUnit StandardUnit, toUnit StandardUnit) (float32, error){
	if fromUnit == Pixels && toUnit == Millimeters {
		return float32(val) * 0.2645833333, nil
	}

	if fromUnit == Millimeters && toUnit == Pixels {
		return float32(val) * 3.7795275591, nil
	}

	return 0.0, errors.New(fmt.Sprintf("unable to convert unit from %d to unit %d", fromUnit, toUnit))
}

func GenerateIndirectReference(reference IndirectReference) string {
	return fmt.Sprintf("%d %d obj", reference.ObjectNumber, reference.GenerationNumber)
}