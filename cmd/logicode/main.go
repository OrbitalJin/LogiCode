package main

import (
	"fmt"
	"os"

	"github.com/OrbitalJin/LogiCode/internal/parser"
	e "github.com/OrbitalJin/LogiCode/types/errors"
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

// args := os.Args[1:] This is for the repl
func main() {
	source := getSource()

	l := parser.NewLexer(source)

	_, err := l.Lex()
	if err != nil {
		fmt.Println(err)
	}
}
