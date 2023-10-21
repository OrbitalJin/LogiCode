package ast

import t "OrbitalJin/LogiCode/types/tokens"

// Expression is the interface for all expressions
type Expression interface {
	Node
	expressionNode()
}

// Signal is the value of the variable
type Signal struct {
	Token t.Token // the token.SIGNAL token
	Value string
}

func (s *Signal) expressionNode()      {}
func (s *Signal) TokenLiteral() string { return s.Token.Literal }

