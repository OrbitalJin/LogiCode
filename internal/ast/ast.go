package ast

import t "OrbitalJin/LogiCode/types/tokens"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// ProgramNode is the root node of the AST
type Program struct {
	Statements []Statement
}

func NewProgramAST() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 { // if there are statements
		return p.Statements[0].TokenLiteral() // return the first statement's token literal
	} else {
		return ""
	}
}

type LetStatement struct {
	Token t.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is the name of the variable
type Identifier struct {
	Token t.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Signal is the value of the variable
type Signal struct {
	Token t.Token // the token.SIGNAL token
	Value string
}

func (s *Signal) expressionNode()      {}
func (s *Signal) TokenLiteral() string { return s.Token.Literal }
