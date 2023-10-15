package parser

import (
	"github.com/OrbitalJin/LogiCode/internal/ast"
	"github.com/OrbitalJin/LogiCode/types/errors"
	t "github.com/OrbitalJin/LogiCode/types/tokens"
)

type Parser struct {
	Tokens  []t.Token
	pointer int
}

// Constructor
func NewParser(tks []t.Token) *Parser {
	return &Parser{
		Tokens:  tks,
		pointer: 0,
	}
}

// Parse
func (p *Parser) Parse() (*ast.ProgramNode, error) {
	err := p.isValidStructure()
	return &ast.ProgramNode{}, err
}

// Check whether the program structure is valid
// This loosely enforces the structure of the program allowing other tokens to be present in between
func (p *Parser) isValidStructure() error {
	expectedOrder := []t.TokenType{
		t.TK_PROGRAMSTART,
		t.TK_DECLARESTART,
		t.TK_DECLAREEND,
		t.TK_BEGIN,
		t.TK_END,
		t.TK_PROGRAMEEND,
	}

	currentIndex := 0

	for _, token := range p.Tokens {
		if currentIndex >= len(expectedOrder) {
			break // We've already found all expected tokens; further tokens are allowed.
		}

		if token.Type == expectedOrder[currentIndex] {
			currentIndex++
		}
	}

	if currentIndex != len(expectedOrder) {
		return errors.NewErr("The structure is incomplete. Some expected tokens are missing.")
	}

	return nil
}
