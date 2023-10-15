package errors

import (
	"fmt"
	"os"

	"github.com/OrbitalJin/LogiCode/meta"
	t "github.com/OrbitalJin/LogiCode/types/tokens"
)

const (
	ERRNO_FILE_NOT_FOUND int = iota + 1
	ERRNO_NO_INPUT_FILE
	ERRNO_EMPTY_FILE
	ERRNO_ILLEGAL_TK
	ERRNO_ILLEGAL_OP
	ERRNO_ILLEGAL_ID
	ERRNO_EOF
	ERRNO_VOID_TOKENS
	ERRNO_INVALID_PROGRAM
	ERRNO_INVALID_PROGRAM_START
	ERRNO_INVALID_PROGRAM_END
)

var Errors = map[int]string{
	ERRNO_FILE_NOT_FOUND:        meta.COMPILER_PREFIX + " File not found",
	ERRNO_NO_INPUT_FILE:         meta.COMPILER_PREFIX + " No Input file provied",
	ERRNO_EMPTY_FILE:            meta.COMPILER_PREFIX + " The source file is empty",
	ERRNO_EOF:                   meta.COMPILER_NAME + " Unexpected EOF",
	ERRNO_ILLEGAL_OP:            meta.LEXER_PREFIX + " Illegal operator",
	ERRNO_ILLEGAL_ID:            meta.LEXER_PREFIX + " Illegal identifier",
	ERRNO_ILLEGAL_TK:            meta.LEXER_PREFIX + " Illegal token",
	ERRNO_VOID_TOKENS:           meta.PARSER_PREFIX + " No tokens to parse",
	ERRNO_INVALID_PROGRAM:       meta.PARSER_PREFIX + " Invalid program",
	ERRNO_INVALID_PROGRAM_START: meta.PARSER_PREFIX + " Expected " + t.KeywordLiterals[t.TK_PROGRAMSTART] + " keyword",
	ERRNO_INVALID_PROGRAM_END:   meta.PARSER_PREFIX + " Expected " + t.KeywordLiterals[t.TK_PROGRAMEEND] + " keyword",
}

func Err(code int) error {
	return fmt.Errorf(Errors[code])
}

func NewErr(msg string) error {
	return fmt.Errorf(meta.COMPILER_PREFIX + " " + msg)
}

func PutsNewErr(msg string) {
	fmt.Println(meta.COMPILER_PREFIX + " " + msg)
}

func PutsErr(code int) {
	fmt.Println(Errors[code])
}

func Fatal(code int) {
	fmt.Println(Errors[code])
	os.Exit(code)
}
