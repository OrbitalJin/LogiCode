package ast

// Node is the interface for all nodes of the AST
type Node interface {
	TokenLiteral() string
}

// ProgramNode is the root node of the AST
type Program struct {
	Statements []Statement
}

// Constructor
func NewProgramAST() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

// TokenLiteral returns the token literal of the first statement
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 { // if there are statements
		return p.Statements[0].TokenLiteral() // return the first statement's token literal
	} else {
		return ""
	}
}
