package ast

type StatementNode struct {
	ActionType string // e.g. Read, Write, Let, Assign
	Target     string // e.g. A, B, C, D
	Express    *ExpressionNode
}
