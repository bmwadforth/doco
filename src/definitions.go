package src

import "bytes"

type IDoco interface {
	Build() error
	Save(fileName string) error
}

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


type IndirectReference struct {
	ObjectNumber     uint
	GenerationNumber uint
}

type DocumentObject struct {
	ObjectNumber     uint
	GenerationNumber uint
	Dictionary       []map[string]string
	Data             interface{}
	ByteOffset       uint
}

type DocumentHeader string
type DocumentBody struct {
	Count uint
	Objects []DocumentObject
}

type Flag rune
const (
	F Flag = 102
	N Flag = 110
)
type DocumentCrossRefItem struct {
	ByteOffset uint
	GenerationNumber uint
	RefFlag Flag
}
type DocumentCrossReferenceTable struct {
	FirstObject uint
	Count uint
	References []DocumentCrossRefItem
}

type DocumentTrailer struct {
	Size uint
	Root IndirectReference
}


type Doco struct {
	Version int
	Header DocumentHeader
	Body DocumentBody
	CrossReference DocumentCrossReferenceTable
	Trailer DocumentTrailer

	pageCount uint
	pageSize PageSize
	dimensions DocumentDimensions

	buffer *bytes.Buffer

	size uint
	lastCrOffset uint
}

//Regular Characters
type WhiteSpaceChars rune
const (
	NUL WhiteSpaceChars = 0
	HT WhiteSpaceChars = 9
	LF WhiteSpaceChars = 10
	FF WhiteSpaceChars = 12
	CR WhiteSpaceChars = 13
	SP WhiteSpaceChars = 32
)

type DelimiterChars rune
const (
	LP DelimiterChars = 40
	RP DelimiterChars = 41
	LS DelimiterChars = 60
	GT DelimiterChars = 62
	LSQ DelimiterChars = 91
	RSQ DelimiterChars = 93
	LCB DelimiterChars = 123
	RCB DelimiterChars = 125
	SOLIDUS DelimiterChars = 47
	PERCENT DelimiterChars = 37
)

//EOL = LF

//