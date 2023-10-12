package types

import (
	"github.com/OrbitalJin/LogiCode/meta"
)


const (
	ERRNO_FILE_NOT_FOUND int = iota + 1
	ERRNO_NO_INPUT_FILE
	ERRNO_UNTOKENIZABLE
)

const (
	ERR_FILE_NOT_FOUND string = meta.COMPILER_PREFIX + " File not found"
	ERR_NO_INPUT_FILE string = meta.COMPILER_PREFIX + " No Input file provied"
)
