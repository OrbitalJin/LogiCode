package main

import (
	"fmt"

	"github.com/OrbitalJin/LogiCode/internal"
)

var source string = "x <- 123";

func main(){
  l := internal.NewLexer(source);
  lexed := l.Lex();
  fmt.Println(lexed);
}
