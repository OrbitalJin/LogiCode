package parser

import (
	"OrbitalJin/LogiCode/internal/ast"
	t "OrbitalJin/LogiCode/types/tokens"
	"fmt"
)

type Parser struct {
	Tokens []t.Token
	errors []string
	ptr    int
}

// Constructor
func NewParser(tks []t.Token) *Parser {
	return &Parser{
		Tokens: tks,
		errors: []string{},
		ptr:    0,
	}
}

// Parse the tokens
func (p *Parser) Parse() (*ast.Program, error) {
	program := ast.NewProgramAST()
	program.Statements = []ast.Statement{}
	for p.ptr < len(p.Tokens) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.advance()
	}
	return program, nil
}

// Parse Statements
func (p *Parser) parseStatement() ast.Statement {
	tk := p.currentToken()
	switch tk.Type {
	case t.TK_LET  : return p.parseLetStatement()
	case t.TK_WRITE: return p.parseWriteStatement()
	case t.TK_READ : return p.parseReadStatement()
	default: return nil
	}
}

// Parses Let Statements
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmnt := &ast.LetStatement{
		Token: p.currentToken(),
	}
	// Assert that the next token is an identifier
	if !p.expectPeek(t.TK_IDENTIFIER) {
		p.peekError(t.TK_IDENTIFIER)
		return nil
	}
	stmnt.Name = &ast.Identifier{
		Token: p.currentToken(),
		Value: p.currentToken().Literal,
	}
	// Assert that the next token is an assignment operator
	if !p.expectPeek(t.OP_ASSIGN) {
		p.peekError(t.OP_ASSIGN)
		return nil
	}
	// Assert that the next token is a signal
	if !p.expectPeek(t.TK_SIGNAL) {
		p.peekError(t.TK_SIGNAL)
		return nil
	}
	stmnt.Value = &ast.Signal{
		Token: p.currentToken(),
		Value: p.currentToken().Literal,
	}
	return stmnt
}

// Parses Write Statements
func (p *Parser) parseWriteStatement() *ast.WriteStatement {
	stmt := &ast.WriteStatement{
		Token: p.currentToken(),
	}
	// Assert that the next token is an IDENTIFIER
	if !p.expectPeek(t.TK_IDENTIFIER) {
		p.peekError(t.TK_IDENTIFIER)
		return nil
	}
	stmt.Name = &ast.Identifier{
		Value: p.currentToken().Literal,
	}
	return stmt
}

// Parses Write Statements
func (p *Parser) parseReadStatement() *ast.ReadStatement {
	stmt := &ast.ReadStatement{
		Token: p.currentToken(),
	}
	// Assert that the next token is an IDENTIFIER
	if !p.expectPeek(t.TK_IDENTIFIER) {
		p.peekError(t.TK_IDENTIFIER)
		return nil
	}
	stmt.Name = &ast.Identifier{
		Value: p.currentToken().Literal,
	}
	return stmt
}

// Advance to the next token
func (p *Parser) advance() bool {
	p.ptr++
	return true
}

// Return Current Token
func (p *Parser) currentToken() t.Token {
	return p.Tokens[p.ptr]
}

// Return Next Token
func (p *Parser) peek() t.Token {
	if p.ptr >= len(p.Tokens) {
		return t.Token{}
	}
	return p.Tokens[p.ptr+1]
}

// Assert that the next token is of type t
func (p *Parser) expectPeek(tkt t.TokenType) bool {
	if p.ptr >= len(p.Tokens) {
		return false
	}
	if p.Tokens[p.ptr+1].Type == tkt {
		return p.advance()
	}
	return false
}

func (p *Parser) peekError(tkt t.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got %s instead at line %d", t.LiteralsMap[tkt], t.LiteralsMap[p.peek().Type], p.peek().Pos.Row)
	p.errors = append(p.errors, msg)
}

// Return Errors
func (p *Parser) Errors() []string {
	return p.errors
}
