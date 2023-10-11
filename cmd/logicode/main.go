package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OrbitalJin/LogiCode/internal"
)

func readSource(path string) string {
  content, err := os.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  return string(content)
}

func main(){
  source := readSource("foo.lc");
  l := internal.NewLexer(source);
  lexed := l.Lex();
  fmt.Println(lexed);
}
