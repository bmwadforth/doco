package pdf

import (
	"fmt"
)

type Version float32
const (
	Latest Version = 1.7
)

type Header struct {
	Version Version
}

func (h *Header) ToBytes() []byte {
	v := fmt.Sprintf("%sPDF-%.1f", string(PercentSign), h.Version)
	return []byte(v)
}