package Writers

import "fmt"

func WriteHeader(pdfVersion float32) string {
	return fmt.Sprintf("%%PDF-%.01f\n", pdfVersion)
}

func WriteBody() []byte {
	return []byte{}
}

func WriteCrossReference() []byte {
	return []byte{}
}

func WriteTrailer() []byte {
	return []byte{}
}

