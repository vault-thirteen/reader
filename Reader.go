package reader

import (
	"bufio"
	"io"
)

// This Package provides is a convenient Reader for specific Purposes.

// For Example, the built-in Go Language Library provides a Method to read
// from a Reader a single Line ending either with ASCII CR Symbol or with a
// Combination of CR with LF Symbols. It does not provide a Method to read a
// Line which ends exactly with a Combination of CR with LF Symbols! This
// Library helps to get rid of such Stupidity of the Developers of the Language.

// ASCII Symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	reader      io.Reader
	bufioReader *bufio.Reader
}
