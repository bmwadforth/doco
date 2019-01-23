package src

type PageSize uint
const (
	A0 PageSize = 0
	A1 PageSize = 1
	A2 PageSize = 2
	A3 PageSize = 3
	A4 PageSize = 4
)

type StandardUnit uint
const (
	Pixels      StandardUnit = 0
	Millimeters StandardUnit = 1
)

type DocumentWidth uint
type DocumentHeight uint

type DocumentDimensions struct {
	DocumentHeight
	DocumentWidth
}

type Doco struct {
	pageCount uint
	pageSize PageSize
	dimensions DocumentDimensions
}