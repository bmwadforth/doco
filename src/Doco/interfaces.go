package Doco

type IDocument interface {
	Save(path string) error
	Write(content string)
	WriteHtml(html string)
	Output() []byte
}
