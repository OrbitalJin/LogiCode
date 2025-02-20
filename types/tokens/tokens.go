package tokens

type TokenType int

// Token is a struct that holds the type and literal of a token
type Token struct {
	Type    TokenType
	Literal string
	Pos     Pos
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
	TK_UNKNOWN
	// Operators
	OP_ASSIGN
	OP_AND
	OP_OR
	OP_NOT
	OP_XOR
	OP_NAND
	OP_NOR
	OP_XNOR
)

// KeywordLiterals is a map of TokenType to their string literal
var LiteralsMap = map[TokenType]string{
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
	OP_ASSIGN:       "<-",
	OP_AND:          "AND",
	OP_NAND:         "NAND",
	OP_OR:           "OR",
	OP_NOR:          "NOR",
	OP_XOR:          "XOR",
	OP_XNOR:         "XNOR",
	OP_NOT:          "NOT",
}

// Reverse the map
var Literals = map[string]TokenType{}

func init() {
	for k, v := range LiteralsMap {
		Literals[v] = k
	}
}

// Illegal identifier chars
var IllegalIdentifierChars string = "-!~&|^()"

// Is operator
func (t TokenType) IsOperator() bool {
	return t >= OP_ASSIGN && t <= OP_XNOR
}
