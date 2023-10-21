package parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"OrbitalJin/LogiCode/meta"
	e "OrbitalJin/LogiCode/types/errors"
	t "OrbitalJin/LogiCode/types/tokens"
)

/*
Lexer Struct
*/
type Lexer struct {
	src    string
	ptr    int
	pos    t.Pos
	prefix string
}

// Constructor
func NewLexer(src string) *Lexer {
	return &Lexer{
		prefix: meta.LEXER_PREFIX,
		src:    src,
		ptr:    0,
		pos: t.Pos{
			Row: 1,
			Col: 1,
		},
	}
}

// Lex the HLL source
func (l *Lexer) Lex() (*[]t.Token, error) {
	var table []t.Token
	if l.srcEmpty() {
		return &table, e.Err(e.ERRNO_EMPTY_FILE)
	}

	for l.ptr < len(l.src) {
		ch := l.read()
		// Only try to lex if the current char is not a WhiteSpace
		if !l.isWhiteSpace(ch) {
			l.debug(ch) // Debug

			tk, err := l.nextToken()
			if err != nil {
				return nil, err
			}
			l.pos.Col++
			table = append(table, tk)
			if meta.DEBUG {
				fmt.Print(tk, "\n")
			} // Debug
		} else {
			l.computeNewPos(ch)
		}
		l.ptr++
	}
	if meta.DEBUG {
		fmt.Printf("%v Successfully lexed %d tokens\n", l.prefix, len(table)) // Debug
	}
	table = append(table, t.Token{Type: t.TK_EOF, Pos: l.pos})
	return &table, nil
}

// Returns a token of the currently pointed byte
func (l *Lexer) nextToken() (t.Token, error) {
	var token t.Token
	var ch byte = l.read()
	var err error = nil

	if l.isNumber(ch) {
		token = l.readSignal()
	}
	if l.isAlpha(ch) {
		switch ch {
		case ';':
			token, err = l.readSemiColon()

		case '<':
			token, err = l.readAssignment()

		default:
			token, err = l.readTokenType()
		}
	}
	return token, err
}

// Tokenize Semi Column
func (l *Lexer) readSemiColon() (t.Token, error) {
	return t.Token{Type: t.TK_SEMICOL, Literal: ";", Pos: l.pos}, nil
}

// Tokenize Signals (Integers)
func (l *Lexer) readSignal() t.Token {
	token := t.Token{Pos: l.pos}

	// Loop until the next char is not a number
	for l.ptr < len(l.src) && l.isNumber(l.read()) {
		token.Literal += string(l.read())
		l.ptr++
	}
	l.ptr--

	token.Type = t.TK_SIGNAL
	return token
}

// Tokenize Assignment operator
func (l *Lexer) readAssignment() (t.Token, error) {
	token := t.Token{Pos: l.pos}
	literal := t.LiteralsMap[t.OP_ASSIGN]

	// Check if the next char is a hyphen
	ch, err := l.peek()
	if err != nil {
		return token, e.Err(e.ERRNO_EOF)
	}
	if ch != '-' {
		return token, l.syntaxError(string(ch), literal, e.Errors[e.ERRNO_ILLEGAL_OP])
	}
	token.Type = t.OP_ASSIGN
	token.Literal = literal
	l.ptr++
	return token, nil
}

// Tokenize Identifiers And Keywords/Operators
func (l *Lexer) readTokenType() (t.Token, error) {
	token := t.Token{Pos: l.pos}
	var str string

	// Loop until the next char doesn't verify the isIdentifier condition
	for l.ptr < len(l.src) && l.isIdentifier(l.read()) {
		str += string(l.read())
		l.ptr++
	}
	l.ptr--

	token.Literal = str

	// If the string is a keyword, return the corresponding TK_
	if kw, ok := l.isKeyword(str); ok {
		token.Type = kw
		return token, nil
	}

	// If the strins is an operator, return the corresponding OP_
	if op, ok := l.isOperator(str); ok {
		token.Type = op
		return token, nil
	}

	// Check if the string contains any illegal identifier char
	for _, ch := range str {
		if strings.Contains(t.IllegalIdentifierChars, string(ch)) {
			return token, l.syntaxError(str, "", e.Errors[e.ERRNO_ILLEGAL_ID])
		}
	}

	// If the string is not a keyword nor an operator, and contains not illegal char, then it's an identifier
	token.Type = t.TK_IDENTIFIER
	return token, nil
}

/// Helper Functions

// Checks wether a string is a keyword (e.g. !Program, LET)
func (l *Lexer) isKeyword(str string) (t.TokenType, bool) {
	if _, found := t.Literals[str]; found {
		return t.Literals[str], true
	}
	return -1, false
}

// Checks wether a string is a  BitWise Operators (e.g &, |, ~)
func (l *Lexer) isOperator(s string) (t.TokenType, bool) {
	if _, found := t.Literals[s]; found {
		return t.Literals[s], true
	}
	return -1, false
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
	return l.src[l.ptr]
}

// Peek the next byte
func (l *Lexer) peek() (byte, error) {
	if l.ptr+1 >= len(l.src) {
		return 0, e.Err(e.ERRNO_EOF)
	}
	return l.src[l.ptr+1], nil
}

// Makes sure that the Lexer hasn't reached the end of the HLL src
func (l *Lexer) srcEmpty() bool {
	return len(l.src) <= l.ptr
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

// Syntax error handler
func (l *Lexer) syntaxError(s, suggest, info string) error {
	var err = "%s Untokenizable literal: `%s` at (%d, %d)"
	if info != "" {
		err += "\n--> [Err] " + info
	}
	if suggest != "" {
		if info == "" {
			err += "\n"
		} else {
			err += " - "
		}
		err += "Did you mean `" + suggest + "`?"
	}

	return fmt.Errorf(err, l.prefix, s, l.pos.Row, l.pos.Col)
}

// Debug
func (l *Lexer) debug(ch byte) {
	if meta.DEBUG {
		fmt.Printf("%s DEBUG - Char: %v, Pointer: %d, Pos: (%d, %d) -> ", l.prefix, string(ch), l.ptr, l.pos.Row, l.pos.Col)
	}
}
