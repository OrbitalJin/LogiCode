package types

const (
	TK_SEMICOL TokenType = iota + 1
	TK_SIGNAL
	TK_ASSIGN
	TK_AND
	TK_NAND
	TK_OR
	TK_NOR
	TK_XOR
	TK_XNOR
	TK_NOT
	TK_IDENT
	TK_LET
	TK_WRITE
	TK_READ
	TK_DECLARESTART
	TK_DECLAREEND
	TK_PROGRAMSTART
	TK_PROGRAMEEND
	TK_BEGIN
	TK_END
	TK_EOF
)

type TokenType int

type Pos struct {
	Row int
	Col int
}

type Token struct {
	Type   TokenType
	Value  string
	Symbol string
	Pos    Pos
}

type SymTable []Token
