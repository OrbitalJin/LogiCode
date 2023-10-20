package main

import (
	"fmt"
	"os"

	"OrbitalJin/LogiCode/internal/parser"
	e "OrbitalJin/LogiCode/types/errors"
	t "OrbitalJin/LogiCode/types/tokens"
)

func getSource() string {
	args := os.Args[1:]
	if len(args) == 0 {
		e.Fatal(e.ERRNO_NO_INPUT_FILE)
	}
	source := readSource(args[0])
	return source
}

func readSource(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		e.Fatal(e.ERRNO_FILE_NOT_FOUND)
	}
	return string(content)
}

func lex(s string) *[]t.Token {
	l := parser.NewLexer(s)

	tokens, err := l.Lex()
	if err != nil {
		fmt.Println(err)
	}
	return tokens
}

func parse(tks []t.Token) {
	p := parser.NewParser(tks)
	_, err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
}

// args := os.Args[1:] This is for the repl
func main() {
	source := getSource()
	tokens := lex(source)
	parse(*tokens)
}
