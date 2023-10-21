package ast

import (
	t "OrbitalJin/LogiCode/types/tokens"
)

// Statement is the interface for all statements
type Statement interface {
	Node
	statementNode()
}

// Identifier is the name of the variable
type Identifier struct {
	Token t.Token // the token.IDENT token
	Value string
}
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// LetStatement is the statement for declaring a variable
type LetStatement struct {
	Token t.Token // the token.LET token
	Name  *Identifier
	Value Expression
}
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// WriteStatement is the statement for writing a variable
type WriteStatement struct {
	Token t.Token // the token.WRITE token
	Name  *Identifier
}
func (ws *WriteStatement) statementNode()       {}
func (ws *WriteStatement) TokenLiteral() string { return ws.Token.Literal }

// ReadStatement is the statement for reading a variable
type ReadStatement struct {
	Token t.Token // the token.READ token
	Name  *Identifier
	Value Expression
}
func (rs *ReadStatement) statementNode()       {}
func (rs *ReadStatement) TokenLiteral() string { return rs.Token.Literal }
