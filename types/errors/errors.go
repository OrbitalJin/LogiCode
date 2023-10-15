package errors

import (
	"fmt"
	"os"

	"github.com/OrbitalJin/LogiCode/meta"
)

const (
	ERRNO_FILE_NOT_FOUND int = iota + 1
	ERRNO_NO_INPUT_FILE
	ERRNO_EMPTY_FILE
	ERRNO_ILLEGAL_TK
	ERRNO_ILLEGAL_OP
	ERRNO_ILLEGAL_ID
	ERRNO_UNTOKENIZABLE
	ERRNO_EOF
)

var Errors = map[int]string{
	ERRNO_FILE_NOT_FOUND: meta.COMPILER_PREFIX + " File not found",
	ERRNO_NO_INPUT_FILE:  meta.COMPILER_PREFIX + " No Input file provied",
	ERRNO_EMPTY_FILE:     meta.COMPILER_PREFIX + " The source file is empty",
	ERRNO_ILLEGAL_TK:     meta.COMPILER_PREFIX + " Illegal token",
	ERRNO_ILLEGAL_OP:     meta.COMPILER_PREFIX + " Illegal operator",
	ERRNO_ILLEGAL_ID:     meta.COMPILER_PREFIX + " Illegal identifier",
	ERRNO_UNTOKENIZABLE:  meta.COMPILER_PREFIX + " Syntax Error",
	ERRNO_EOF:            meta.COMPILER_PREFIX + " Unexpected EOF",
}

func Err(code int) error {
	return fmt.Errorf(Errors[code])
}

func PutsErr(code int) {
	fmt.Println(Errors[code])
}

func Fatal(code int) {
	fmt.Println(Errors[code])
	os.Exit(code)
}
