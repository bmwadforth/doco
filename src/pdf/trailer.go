package pdf

type Trailer struct {
	Size       uint32
	RootObject ObjectReference
	XrefStart  uint32
}

func (t *Trailer) ToBytes() []byte {

	return make([]byte, 1)
}
