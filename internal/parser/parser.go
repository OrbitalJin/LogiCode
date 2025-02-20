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
	case t.TK_LET:
		return p.parseLetStatement()
	case t.TK_WRITE:
		return p.parseWriteStatement()
	case t.TK_READ:
		return p.parseReadStatement()
	default:
		return nil
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

	stmnt.Value = p.parseExpression()
	return stmnt
}

// Parsing Expressions
func (p *Parser) parseExpression() ast.Expression {
	// Case 1: Signal. Next two tokens are a signal and a semicolon
	if p.peek(1).Type == t.TK_SIGNAL && p.peek(2).Type == t.TK_SEMICOL {
		return &ast.SignalExpression{
			Token: p.currentToken(),
			Value: p.currentToken().Literal,
		}
	}
	// Case 2: Binary Expression. Next three tokens are an expression, an operator, and an expression
	if p.peek(1).Type == t.TK_SIGNAL &&  p.peek(2).Type.IsOperator() && p.peek(3).Type == t.TK_SIGNAL {
		return &ast.BinaryExpression{
			Left: &ast.SignalExpression{
				Token: p.currentToken(),
				Value: p.currentToken().Literal,
			},
			Operator: p.peek(2),
			Right: &ast.SignalExpression{
				Token: p.currentToken(),
				Value: p.currentToken().Literal,
			},
		}
	// Case 3: Unary Expression. Next two tokens are an operator and an expression
	p.peekError(t.TK_SIGNAL)
	return nil
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
func (p *Parser) peek(i int) t.Token {
	if p.ptr+i >= len(p.Tokens) {
		return t.Token{}
	}
	return p.Tokens[p.ptr+i]
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
	msg := fmt.Sprintf(
		"Expected next token to be %s, got %s instead at (%d, %d)",
		t.LiteralsMap[tkt],
		t.LiteralsMap[p.peek(1).Type],
		p.peek(1).Pos.Row,
		p.peek(1).Pos.Col,
	)
	p.errors = append(p.errors, msg)
}

// Return Errors
func (p *Parser) Errors() []string {
	return p.errors
}
