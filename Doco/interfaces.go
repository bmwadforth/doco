package Doco

type DocoSpecimen interface {
	Save(path string) error
	Write(content string)
	Output() []byte
}
