package internal

import (
	"fmt"
	"strconv"

	"github.com/OrbitalJin/LogiCode/types"
)

/*
Lexer Struct
*/
type Lexer struct {
	src string
	pointer int
	pos types.Pos
	prefix string
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		src: src,
		pointer: -1,
		prefix: "[LGCC] Lexer -",
		pos: types.Pos{
			Row: 1,
			Col: 1,
		},
	}
}

// Lex the HLL source
func (l *Lexer) Lex() (*[]types.Token, error) {
	var table []types.Token
	for l.notEmpty() {
		tk, err := l.nextToken()
		if err != nil {
			return &table, err
		}
		table = append(table, tk)
	}
	return &table, nil
}

// Return the Next Token
func (l *Lexer) nextToken() (types.Token, error) {
	var token types.Token
	var err error = nil
	ch, _ := l.peek(0)

	if !l.isWhiteSpace(ch){
		switch(ch){
			case ";": token = types.Token{Type: types.TK_SEMICOL, Pos: l.pos}
			// case "!": token, err = l.readBlockDelimiter()
			default: return token, l.syntaxError(ch, l.pos)
		}
	}
	return token, err
}

// Tokenize Signals (Integers)
func (l *Lexer) readSignal() (types.Token, error) {
	return types.Token{}, nil
}

// Tokenize Identifiers
func (l *Lexer) readIdentifier() (types.Token, error) {
	return types.Token{}, nil
}

// Tokenize Block delimiters (e.g. !Program, !Begin)
func (l *Lexer) readBlockDelimiter() (types.Token, error) {
	return types.Token{}, nil
}

// Tokenize BitWise Operators (e.g &, |, ~)
func (l *Lexer) readOperator() (types.Token, error) {
	return types.Token{}, nil
}

func (l *Lexer) peek(i int) (string, error) {
	if l.pointer + i < len(l.src){
		return string(l.src[i]), nil
	}
	return "", fmt.Errorf("%s Peek: Out of index", l.prefix)
}

// Syntax error handler
func (l *Lexer) syntaxError(s string, p types.Pos) error {
	var err = "%s Untokenizable literal: %s at (%d, %d)"
	return fmt.Errorf(err, l.prefix, s, p.Row, p.Col)
}

// Makes sure that L hasn't reached the end of the HLL src
func (l *Lexer) notEmpty() bool {
	return l.pointer < len(l.src)
}

// Checks wheter the token at a certain position is Skippable
func (l *Lexer) isWhiteSpace(s string) bool {
	return s == " " || s == "\n" || s == "\t "
}

// Checks wehter a char is alpah 
func (l *Lexer) isAlpha(s string) bool {
	return !l.isNumber(s)
}

// Checks wether a character is str representation of an int
func (l *Lexer) isNumber(s string) bool {
	_, err := strconv.Atoi(s);
	return err == nil
}

