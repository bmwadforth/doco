package Writers

func (c *CharItem) Write() []byte{
	bytes := make([]byte, 1)
	for i := 0; i < int(c.iterations); i++ {
		bytes = append(bytes, byte(c.codePoint))
	}
	return bytes
}