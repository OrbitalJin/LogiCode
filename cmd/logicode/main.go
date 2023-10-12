package main

import (
	"fmt"
	"os"

	"github.com/OrbitalJin/LogiCode/internal"
	"github.com/OrbitalJin/LogiCode/types"
)


func getSource() string {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println(types.ERR_NO_INPUT_FILE)
    os.Exit(types.ERRNO_NO_INPUT_FILE)
  }
  source := readSource("foo.lc")
  return source
}

func readSource(path string) string {
  content, err := os.ReadFile(path)
  if err != nil {
    fmt.Println(types.ERR_FILE_NOT_FOUND)
    os.Exit(1)
  }
  return string(content)
}

// args := os.Args[1:] This is for the repl
func main(){
  source := getSource();
  l := internal.NewLexer(source);
  lexed, err := l.Lex();
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(*lexed)
  }
}
