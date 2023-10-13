package types

type TokenType int

// Token is a struct that holds the type and literal of a token
type Token struct {
	Type   TokenType
	Literal  string
	Symbol string
	Pos    Pos
}

// Pos is a struct that holds the position of a token
type Pos struct {
	Row int
	Col int
}

// SymTable is a slice of Token
type SymTable []Token

// TokenType enum
const (
	TK_SEMICOL TokenType = iota + 1
	TK_SIGNAL
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
	// Operators
	OP_ASSIGN 
	OP_AND
	OP_NAND
	OP_OR
	OP_NOR
	OP_XOR
	OP_XNOR
	OP_NOT
)

// KeywordLiterals is a map of TokenType to their string literal
var KeywordLiterals = map[TokenType]string{
	TK_SEMICOL:      ";",
	TK_SIGNAL:       "SIGNAL",
	TK_IDENTIFIER:   "IDENT",
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
	for k, v := range KeywordLiterals {
		Keywords[v] = k
	}
}

// OperatorsLiterals is a map of TokenType to their string literal
var OperatorsLiterals = map[TokenType]string{
	OP_ASSIGN: "<-",
	OP_AND:    "AND",
	OP_NAND:   "NAND",
	OP_OR:     "OR",
	OP_NOR:    "NOR",
	OP_XOR:    "XOR",
	OP_XNOR:   "XNOR",
	OP_NOT:    "NOT",
}

// Reverse the map
var Operators = map[string]TokenType{}
func init() {
	for k, v := range OperatorsLiterals {
		Operators[v] = k
	}
}
	