package src

const (
	pageTree ObjectType = 0
	page ObjectType = 1
	catalog ObjectType = 2
)

const (
	Single DocumentLayout = 0
	OneColumn DocumentLayout = 1
	TwoColumnLeft DocumentLayout = 2
	TwoColumnRight DocumentLayout = 3
	TwoPageLeft DocumentLayout = 4
	TwoPageRight DocumentLayout = 5
)

const (
	A0 PaperSize = 0
	A1 PaperSize = 1
	A2 PaperSize = 2
	A3 PaperSize = 3
	A4 PaperSize = 4
)

const (
	Pixels      StandardUnit = 0
	Millimeters StandardUnit = 1
)

const (
	F Flag = 102
	N Flag = 110
)

const (
	NUL WhiteSpaceChars = 0
	HT WhiteSpaceChars = 9
	LF WhiteSpaceChars = 10
	FF WhiteSpaceChars = 12
	CR WhiteSpaceChars = 13
	SP WhiteSpaceChars = 32
)

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