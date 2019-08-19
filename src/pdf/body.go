package pdf

type Body struct {

}

func (b *Body) ToBytes() []byte {

	return make([]byte, 1)
}