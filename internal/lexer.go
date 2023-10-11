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

	// Iterating over lines
	for i, line := range src {
		if !l.isSkippable(line) {
			line := strings.Split(line, " ")
			// Iterating over words
			for j, word := range line {
				word = strings.Trim(word, "\n ")
				// Keywords
				switch word {
					case "WRITE": tokens = append(tokens, types.Token{Type: types.TK_WRITE, Symbol: word})
					case "READ" : tokens = append(tokens, types.Token{Type: types.TK_READ, Symbol: word})
					case "LET"  : tokens = append(tokens, types.Token{Type: types.TK_LET, Symbol: word})
					case "<-"   : tokens = append(tokens, types.Token{Type: types.TK_ASSIGN, Symbol: word})
					case "AND"  : tokens = append(tokens, types.Token{Type: types.TK_AND, Symbol: word})
					case "OR"   : tokens = append(tokens, types.Token{Type: types.TK_OR, Symbol: word})
					case "XOR"  : tokens = append(tokens, types.Token{Type: types.TK_XOR, Symbol: word})
					case "NOT"  : tokens = append(tokens, types.Token{Type: types.TK_NOT, Symbol: word})
					case "NAND" : tokens = append(tokens, types.Token{Type: types.TK_NAND, Symbol: word})
					case "NOR"  : tokens = append(tokens, types.Token{Type: types.TK_NOR, Symbol: word})
					case "XNOR" : tokens = append(tokens, types.Token{Type: types.TK_XNOR, Symbol: word})
					default:
						// Signals
						if l.isDigit(word) {
							tokens = append(tokens, types.Token{Type: types.TK_SIGNAL, Value: word})

						} else if l.isAlpha(word) {
							// Block Delimiters
							if word[0] == '!' {
								tokens = append(tokens, l.lexBlockDeclaration(word, i))
							// Identifiers
							} else {
								tokens = append(tokens, types.Token{Type: types.TK_IDENT, Symbol: word})
							}
						// Invalid Token
						} else {
							l.syntaxError(word, i, j)
						}
				}
	
			}
		}
	}
	// Add EOF token
	tokens = append(tokens, types.Token{Type: types.TK_EOF, Value: "EOF"})

	return tokens
}

func (l *Lexer) preProcessSource() []string {
	// Split into lines
	lines := strings.Split(l.src, "\n")

	// Remove leading/trailing whitespace from each line
	for i, line := range lines {
			lines[i] = strings.TrimSpace(line)
	}

	// Remove empty lines
	filtered := make([]string, 0, len(lines))
	for _, line := range lines {
			if line != "" {
					filtered = append(filtered, line)
			}
	}

	return filtered
}

func (l *Lexer) lexBlockDeclaration(s string, line int) types.Token {
	if len(s) == 1 {
		l.syntaxError(s, line, 1);
	}
	switch s[1:] {
		case "Program": return types.Token{Type: types.TK_PROGRAMSTART, Symbol: s}
		case "EndProgram": return types.Token{Type: types.TK_PROGRAMEEND, Symbol: s}
		case "Declare": return types.Token{Type: types.TK_DECLARESTART, Symbol: s}
		case "EndDeclare": return types.Token{Type: types.TK_DECLAREEND, Symbol: s}
		case "Begin": return types.Token{Type: types.TK_BEGIN, Symbol: s}
		case "End": return types.Token{Type: types.TK_END, Symbol: s}
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
