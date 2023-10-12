package types

type TokenType int

type Token struct {
	Type   TokenType
	Literal  string
	Symbol string
	Pos    Pos
}

type Pos struct {
	Row int
	Col int
}

type SymTable []Token

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
	TK_IDENTIFIER
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

var TokenTypeLiterals = map[TokenType]string{
	TK_SEMICOL:      ";",
	TK_SIGNAL:       "SIGNAL",
	TK_ASSIGN:       "<-",
	TK_AND:          "AND",
	TK_NAND:         "NAND",
	TK_OR:           "OR",
	TK_NOR:          "NOR",
	TK_XOR:          "XOR",
	TK_XNOR:         "XNOR",
	TK_NOT:          "NOT",
	TK_IDENTIFIER:        "IDENT",
	TK_LET:          "LET",
	TK_WRITE:        "WRITE",
	TK_READ:         "READ",
	TK_DECLARESTART: "!Declare",
	TK_DECLAREEND:   "!EndDeclare",
	TK_PROGRAMSTART: "!Program",
	TK_PROGRAMEEND:  "!EndProgram",
	TK_BEGIN:        "!Begin",
	TK_END:          "!End",
	TK_EOF:          "EOF",
}

// Reverse the map

var Keywords = map[string]TokenType{}
func init() {
	for k, v := range TokenTypeLiterals {
		Keywords[v] = k
	}
}
	