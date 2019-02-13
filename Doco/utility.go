package Doco

func StdUnitToPoint(val float32, unit Unit) float32 {
	switch unit {
	case UnitMillimeters:
		return float32(val * 2.835)
	case UnitPixels:
		return float32(val * 0.75)
	default:
		return val
	}
}