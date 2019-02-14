package Doco

type DocoSpecimen interface {
	Save(path string) error
	Output() []byte
}
