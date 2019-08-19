package pdf

type CodePoint rune
const (
	Null               CodePoint = '\x00'
	HorizontalTab      CodePoint = '\x09'
	LineFeed           CodePoint = '\x0A'
	FormFeed           CodePoint = '\x0C'
	CarriageReturn     CodePoint = '\x0D'
	Space              CodePoint = '\x20'
	LeftParenthesis    CodePoint = '\x28'
	RightParenthesis   CodePoint = '\x29'
	LessThan           CodePoint = '\x3C'
	GreaterThan        CodePoint = '\x3E'
	LeftSquareBracket  CodePoint = '\x5B'
	RightSquareBracket CodePoint = '\x5D'
	LeftCurlyBracket   CodePoint = '\x7B'
	RightCurlyBracket  CodePoint = '\x7D'
	Solidus            CodePoint = '\x2F'
	PercentSign        CodePoint = '\x25'
)

