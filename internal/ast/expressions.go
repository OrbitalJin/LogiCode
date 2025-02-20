package ast

import t "OrbitalJin/LogiCode/types/tokens"

// Expression is the interface for all expressions
type Expression interface {
	Node
	expressionNode()
}

// Signal is the value of the variable
type SignalExpression struct {
	Token t.Token // the token.SIGNAL token
	Value string
}

func (s *SignalExpression) expressionNode()      {}
func (s *SignalExpression) TokenLiteral() string { return s.Token.Literal }

// BinaryExpression is the expression for binary operations
type BinaryExpression struct {
	Left     Expression
	Operator t.Token
	Right    Expression
}
