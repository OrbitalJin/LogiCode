package internal

import (
	"fmt"
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
func (l *Lexer) Lex() [] types.Token {
  slice := []types.Token{}
	for _, w := range strings.Split(l.src, " ") {
		if IsDigit(w){
			slice = append(slice, types.Token{
				Type: "INT",
				Value: w,
				Symbol: w,
			})
			
		} else if IsAlpha(w) {
			if IsAppendix(w) {
				slice = append(slice, types.Token{
					Type: "APPENDIX",
				})
			}else {
				slice = append(slice, types.Token{
					Type: "LITERAL",
					Value: w,
					Symbol: w,
				})
			}
		} else {
			fmt.Println("UnTokenizable string: ", w);
		}
	}
  return slice
}

func IsDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsAlpha(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil
}

func IsAppendix(s string) bool {
	return s == "<-"
}
