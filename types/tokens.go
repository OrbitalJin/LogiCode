package types

const (
	TK_ASSIGN TokenType = iota
	TK_IDENT
	TK_INT
	TK_SIGN
	TK_LPAREN
	TK_RPAREN
	TK_AND
	TK_NAND
	TK_OR
	TK_NOR
	TK_XOR
	TK_XNOR
	TK_NOT
	TK_LET
	TK_EOF
	TK_DECLARESTART
	TK_DECLAREEND
	TK_PROGRAMSTART
	TK_PROGRAMEEND
	TK_BEGIN
	TK_END
)

type TokenType int

type Token struct {
	Type TokenType
	Value string
	Symbol string
}

