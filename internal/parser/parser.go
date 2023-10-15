package parser

import (
	t "github.com/OrbitalJin/LogiCode/types/tokens"
)

type Parser struct {
	Tokens []t.Token
}

// Constructor
func NewParser(tks []t.Token) *Parser {
	return &Parser{
		Tokens: tks,
	}
}

// Parse
// func (p *Parser) Parse() *ast.ProgramNode {
// }
