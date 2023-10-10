package internal

import (
	"github.com/OrbitalJin/LogiCode/types"
)

type Parser struct {
	tokens []types.Token
}

// Constructor
func NewParser(tks []types.Token) *Parser {
	return &Parser {
		tokens: tks,
	}
}

// Parse
func (p *Parser) Parse() {}