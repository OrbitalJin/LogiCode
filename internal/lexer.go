package internal

import (
	"fmt"
	"strconv"

	"github.com/OrbitalJin/LogiCode/meta"
	"github.com/OrbitalJin/LogiCode/types"
)

/*
Lexer Struct
*/
const prefix string = meta.COMPILER_PREFIX + " Lexer -"

type Lexer struct {
	src     string
	pointer int
	pos     types.Pos
	prefix  string
}

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
	for !l.srcEmpty() {
		tk, err := l.nextToken()
		if err != nil {
			return &table, err
		}
		table = append(table, tk)

		l.pointer++
	}
	return &table, nil
}

// Return the Next Token
func (l *Lexer) nextToken() (types.Token, error) {
	var token types.Token
	ch, err := l.read()

	if !l.isWhiteSpace(ch) {
		switch ch {
		case ";":
			token, err = l.readSemiColumn()
			// case "!":
			// token, err = l.readBlockDelimiter()
		default:
			return token, l.syntaxError(ch)
		}
	}
	return token, err
}

// Tokenize Semi Column
func (l *Lexer) readSemiColumn() (types.Token, error) {
	l.pos.Col++
	token := types.Token{Type: types.TK_SEMICOL, Pos: l.pos}
	return token, nil
}

// TODO
// Tokenize Signals (Integers)
func (l *Lexer) readSignal() (types.Token, error) {
	return types.Token{}, nil
}

// TODO
// Tokenize Identifiers
func (l *Lexer) readIdentifier() (types.Token, error) {
	return types.Token{}, nil
}

// TODO
// Tokenize Block delimiters (e.g. !Program, !Begin)
func (l *Lexer) readBlockDelimiter() (types.Token, error) {
	return types.Token{}, nil
}

// TODO
// Tokenize BitWise Operators (e.g &, |, ~)
func (l *Lexer) readOperator() (types.Token, error) {
	return types.Token{}, nil
}

// Read the character at the current pointer index
func (l *Lexer) read() (string, error) {
	if l.pointer < len(l.src) {
		return string(l.src[l.pointer]), nil
	}
	return "", fmt.Errorf("%s Peek: Out of index", l.prefix)
}

// Peek at the next character
func (l *Lexer) peek(i int) (string, error) {
	if l.pointer+i < len(l.src) {
		return string(l.src[i]), nil
	}
	return "", fmt.Errorf("%s Peek: Out of index", l.prefix)
}

// Syntax error handler
func (l *Lexer) syntaxError(s string) error {
	var err = "%s Untokenizable literal: %s at (%d, %d)"
	return fmt.Errorf(err, l.prefix, s, l.pos.Row, l.pos.Col)
}

// Makes sure that L hasn't reached the end of the HLL src
func (l *Lexer) srcEmpty() bool {
	return len(l.src) <= l.pointer
}

// Checks wheter the token at a certain position is Skippable
func (l *Lexer) isWhiteSpace(s string) bool {
	if s == "\n" {
		l.pos.Row++
	}
	return s == " " || s == "\t " || s == "\n"
}

// Checks wehter a char is alpah
func (l *Lexer) isAlpha(s string) bool {
	return !l.isNumber(s)
}

// Checks wether a character is str representation of an int
func (l *Lexer) isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
