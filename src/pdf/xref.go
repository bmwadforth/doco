package pdf

type Xref struct {

}

func (x *Xref) ToBytes() []byte {

	return make([]byte, 1)
}