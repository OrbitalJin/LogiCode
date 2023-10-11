package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/OrbitalJin/LogiCode/types"
)



type Lexer struct {
	src string
}
// Constructor
func NewLexer(src string) *Lexer {
	return &Lexer{
		src: src,
	}
}
// Lex
func (l *Lexer) Lex() []types.Token {
	var tokens []types.Token
	var src []string = l.preProcessSource()

	for line, word := range src {
		word = strings.Trim(word, "\n ")
		if !l.isSkippable(word){
			switch word[0] {
				case '!': tokens = append(tokens, l.lexBlockDeclaration(word, line))
			}
		}
	}
	// add EOF token
	tokens = append(tokens, types.Token{Type: types.TK_EOF, Value: "EOF"})

	return tokens
}

func (l *Lexer) preProcessSource() []string {
	src := strings.Split(l.src, ";")
	if len(src) > 1 {src = src[:len(src) - 1]}
	return src
}

func (l *Lexer) lexBlockDeclaration(s string, line int) types.Token {
	if len(s) == 1 {
		l.syntaxError(s, line, 1);
	}
	switch s[1:] {
		case "Program": return types.Token{Type: types.TK_PROGRAMSTART,}
		case "EndProgram": return types.Token{Type: types.TK_PROGRAMEEND}
		case "Declare": return types.Token{Type: types.TK_DECLARESTART,}
		case "EndDeclare": return types.Token{Type: types.TK_DECLAREEND}
		case "Begin": return types.Token{Type: types.TK_BEGIN,}
		case "End": return types.Token{Type: types.TK_END}
	default:
		l.syntaxError(s, line, 1)
		return types.Token{}
	}
}

func (l *Lexer) syntaxError(w string, line int, row int) {
	fmt.Printf("[Lexer] Syntax Error. \nUntokenizable symbol: %s at %d, %d \n", w, line + 1, row + 1)
	os.Exit(-1)
}

func (l *Lexer) isSkippable(s string) bool {
	return s == "\n" || s == "\t" || s == " "
}

func (l *Lexer) isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func (l *Lexer) isAlpha(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func (l *Lexer) isAppendix(s string) bool {
	return s == "<-"
}
