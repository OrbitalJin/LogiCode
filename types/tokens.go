package types

const (
	TK_ASSIGN TokenType = iota + 1
	TK_IDENT
	TK_SIGNAL
	TK_WRITE
	TK_READ
	TK_AND
	TK_NAND
	TK_OR
	TK_NOR
	TK_XOR
	TK_XNOR
	TK_NOT
	TK_LET
	TK_DECLARESTART
	TK_DECLAREEND
	TK_PROGRAMSTART
	TK_PROGRAMEEND
	TK_BEGIN
	TK_END
	TK_EOF
)

type TokenType int

type Token struct {
	Type TokenType
	Value string
	Symbol string
}

