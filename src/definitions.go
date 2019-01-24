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

type Doco struct {
	pageCount uint
	pageSize PageSize
	dimensions DocumentDimensions
	buffer *bytes.Buffer
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