package parser

import (
	"OrbitalJin/LogiCode/internal/ast"
	t "OrbitalJin/LogiCode/types/tokens"
	"fmt"
)

type Parser struct {
	Tokens []t.Token
	ptr    int
}

// Constructor
func NewParser(tks []t.Token) *Parser {
	return &Parser{
		Tokens: tks,
		ptr:    0,
	}
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := ast.NewProgramAST()
	for p.ptr < len(p.Tokens) {
		tk := p.currentToken()
		switch tk.Type {
		case t.TK_LET:
			letStatement := p.parseLetStatement()
			if letStatement != nil {
				program.Statements = append(program.Statements, letStatement)
			}
		}
		p.ptr++
	}
	return program, nil
}

// Parses Let Statements
func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{
		Token: p.currentToken(),
	}
	// Assert that the next token is an identifier
	if !p.expectPeek(t.TK_IDENTIFIER) {
		fmt.Println(fmt.Sprintf("Expected %s, got %s", t.KeywordLiterals[t.TK_IDENTIFIER], t.KeywordLiterals[p.peek().Type]))
		return nil
	}
	letStatement.Name = &ast.Identifier{
		Token: p.currentToken(),
		Value: p.currentToken().Literal,
	}
	// Assert that the next token is an assignment operator
	if !p.expectPeek(t.OP_ASSIGN) {
		fmt.Println(fmt.Sprintf("Expected %s, got %s", t.OperatorsLiterals[t.OP_ASSIGN], t.KeywordLiterals[p.peek().Type]))
		return nil
	}
	// Assert that the next token is a signal
	if !p.expectPeek(t.TK_SIGNAL) {
		fmt.Println(fmt.Sprintf("Expected %s, got %s", t.KeywordLiterals[t.TK_SIGNAL], t.KeywordLiterals[p.peek().Type]))
		return nil
	}
	letStatement.Value = &ast.Signal{
		Token: p.currentToken(),
		Value: p.currentToken().Literal,
	}
	return letStatement
}

func (p *Parser) currentToken() t.Token {
	if p.ptr >= len(p.Tokens) {

	}
	return p.Tokens[p.ptr]
}

func (p *Parser) nextToken() t.Token {
	if p.ptr >= len(p.Tokens) {
		return t.Token{}
	}
	p.ptr++
	return p.Tokens[p.ptr]
}

func (p *Parser) peek() t.Token {
	if p.ptr >= len(p.Tokens) {
		return t.Token{}
	}
	return p.Tokens[p.ptr+1]
}

func (p *Parser) expectPeek(t t.TokenType) bool {
	if p.ptr >= len(p.Tokens) {
		return false
	}
	if p.Tokens[p.ptr+1].Type == t {
		p.ptr++
		return true
	}
	return false
}
