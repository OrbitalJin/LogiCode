package internal

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/OrbitalJin/LogiCode/meta"
	"github.com/OrbitalJin/LogiCode/types"
)

const prefix string = meta.COMPILER_PREFIX + " Lexer -"

/*
Lexer Struct
*/
type Lexer struct {
	src     string
	pointer int
	pos     types.Pos
	prefix  string
}

// Constructor
func NewLexer(src string) *Lexer {
	return &Lexer{
		prefix:  prefix,
		src:     src,
		pointer: 0,
		pos: types.Pos{
			Row: 1,
			Col: 1,
		},
	}
}

// Lex the HLL source
func (l *Lexer) Lex() (*[]types.Token, error) {
	var table []types.Token
	if l.srcEmpty() {
		return &table, fmt.Errorf(types.ERR_EMPTY_FILE)
	}

	for l.pointer < len(l.src) {
		ch := l.read()
		// Only try to lex if the current char is not a WhiteSpace
		if !l.isWhiteSpace(ch) {
			fmt.Println(string(ch))
			tk, err := l.nextToken()
			if err != nil {
				return &table, err
			}
			l.pos.Col++
			table = append(table, tk)
		} else {
			l.computeNewPos(ch)
		}
		l.pointer++
	}
	table = append(table, types.Token{Type: types.TK_EOF, Pos: l.pos})
	return &table, nil
}

// Return a token of the currently pointed byte
func (l *Lexer) nextToken() (types.Token, error) {
	var token types.Token
	var ch byte = l.read()
	var err error = nil

	if l.isNumber(ch) {
		token = l.readSignal()
	}
	if l.isAlpha(ch) {
		switch ch {
		case ';':
			token, err = l.readSemiColumn()

		case '<':
			token, err = l.readAssignment()

		default:
			token, err = l.readIdentifier()
		}
	}
	return token, err
}

// Tokenize Semi Column
func (l *Lexer) readSemiColumn() (types.Token, error) {
	return types.Token{Type: types.TK_SEMICOL, Literal: ";", Pos: l.pos}, nil
}

// Tokenize Signals (Integers)
func (l *Lexer) readSignal() types.Token {
	token := types.Token{Pos: l.pos}

	// Loop until the next char is not a number
	for l.pointer < len(l.src) && l.isNumber(l.read()) {
		token.Literal += string(l.read())
		l.pointer++
	}
	l.pointer--

	token.Type = types.TK_SIGNAL
	return token
}

// Tokenize Assignment operator
func (l *Lexer) readAssignment() (types.Token, error) {
	token := types.Token{Pos: l.pos}
	literal := types.TokenTypeStrings[types.TK_ASSIGN]

	// Check if the next char is a hyphen
	ch, err := l.peek()
	if err != nil {
		return token, fmt.Errorf(types.ERR_EOF)
	}
	if ch != '-' {
		return token, l.syntaxError(ch, literal)
	}
	token.Type = types.TK_ASSIGN
	token.Literal = literal
	l.pointer++
	return token, nil
}

// Tokenize Identifiers
func (l *Lexer) readIdentifier() (types.Token, error) {
	token := types.Token{Pos: l.pos}
	var str string

	// Loop until the next char doesn't verify the isIdentifier condition
	for l.pointer < len(l.src) && l.isIdentifier(l.read()){
		str += string(l.read())
		l.pointer++
	}
	l.pointer--

	token.Type = types.TK_IDENT
	token.Literal = str

	return token, nil
}

// TODO
// Tokenize BitWise Operators (e.g &, |, ~)
func (l *Lexer) readOperator() (types.Token, error) {
	return types.Token{}, nil
}

// TODO
// Tokenize Block delimiters (e.g. !Program, !Begin)
func (l *Lexer) readBlockDelimiter() (types.Token, error) {
	token := types.Token{}

	return token, nil
}

// Skip White Spaces e.g \n, \t, ` `
func (l *Lexer) skipWhiteSpace() {
	for l.pointer < len(l.src) && unicode.IsSpace(rune(l.src[l.pointer])) {
		l.pointer++
	}
}

// Compute the new position
func (l *Lexer) computeNewPos(ch byte) {
	switch ch {
	case '\n':
		l.pos.Row++
		l.pos.Col = 1
	case ' ':
		l.pos.Col++
	}
}

// Reads and returs the current byte
func (l *Lexer) read() byte {
	return l.src[l.pointer]
}

// Peek the next byte
func (l *Lexer) peek() (byte, error) {
	if l.pointer+1 >= len(l.src) {
		return 0, fmt.Errorf(types.ERR_EOF)
	}
	return l.src[l.pointer+1], nil
}

// Syntax error handler
func (l *Lexer) syntaxError(ch byte, suggest string) error {
	var err = "%s Untokenizable literal: %s at (%d, %d)"
	if suggest != "" {
		err += "\nDid you mean: " + suggest + "?"
	}
	return fmt.Errorf(err, l.prefix, string(ch), l.pos.Row, l.pos.Col)
}

// Makes sure that the Lexer hasn't reached the end of the HLL src
func (l *Lexer) srcEmpty() bool {
	return len(l.src) <= l.pointer
}

// Checks wether the token at a certain position is Skippable
func (l *Lexer) isWhiteSpace(ch byte) bool {
	return unicode.IsSpace(rune(ch))
}

// Checks wether a char is alpah
func (l *Lexer) isAlpha(ch byte) bool {
	return !l.isNumber(ch)
}

// Check wether a char is assignable to an identifier
func (l *Lexer) isIdentifier(ch byte) bool {
	return ch != ';' && ch != '<' && (l.isAlpha(ch) || l.isNumber(ch)) && !l.isWhiteSpace(ch)
}

// Checks wether a character is str representation of an int
func (l *Lexer) isNumber(ch byte) bool {
	_, err := strconv.Atoi(string(ch))
	return err == nil
}
