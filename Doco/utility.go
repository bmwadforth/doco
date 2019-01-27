package Doco

func StdUnitToPoint(val float32, unit Unit) float32 {
	switch unit {
	case Milimeters:
		return float32(val * 2.835)
	case Pixels:
		return float32(val * 0.75)
	default:
		return val
	}
}