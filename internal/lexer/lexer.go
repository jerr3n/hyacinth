package lexer

type Kind int
type Flags int
type Token struct {
	Kind  Kind
	Start int
	End   int
}

const (
	// Invalid tokens
	Invalid Kind = iota
	EOF          // eof
	Semicolon
	// `Kind`s between literalStart -> literalEnd are literals
	literalStart
	Ident
	Lifetime
	Int
	Float
	Duration //5ms, 500s, 50h
	Char
	ByteChar
	Str
	ByteStr
	StrStart // "..{
	StrMid   // { thing }
	StrEnd   // }.."
	FmtSpec
	literalEnd
	// ditto for btwn keywordStart -> keywordEnd
	keywordStart
	Pkg
	Import
	Const
	Let
	Var
	Fn
	Pub
	Struct
	Enum
	Interface
	Type
	Extend
	Use
	Extern
	If
	Else
	Match
	For
	In
	While
	Loop
	Return
	Break
	Continue
	Defer
	Throw
	Try
	Catch
	InOut
	Take
	Shared
	Dyn
	Spawn
	Scope
	Select
	Lock
	As
	With
	Comptime
	Region
	Unsafe
	True
	False
	Null
	keywordEnd
	LParen
	RParen
	LBracket
	RBracket
	LBrace
	RBrace
	Comma
	Colon
	Dot
	At
	Hash
	// FatArrow it's really called that lol https://en.wikipedia.org/wiki/Fat_comma
	FatArrow
	Add
	Subtract
	Multiply
	Divide
	Modulo
	AddWrap
	SubtractWrap
	MultiplyWrap
	DivideWrap
	AddSat
	SubtractSat
	AddEqual
	SubtractEqual
	MultiplyEqual
	DivideEqual
	ModuloEqual
	ShiftLeft
	ShiftRight
	ShiftLeftEqual
	ShiftRightEqual
	EqualEqual
	NotEqual
	LtEqual
	Caret
	QuestionQuestion
	QuestionDot
	DotDot
	DotDotEqual
	Arrow
	AndAnd
	OrOr
	Tilde
	Equal
	Lt
	Gt
	Bang
	Question
	And
	Pipe
	AndEqual
	PipeEqual
	CaretEqual
)
