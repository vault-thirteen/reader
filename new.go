package reader

import (
	"bufio"
	"io"
)

func New(
	reader io.Reader,
) *Reader {

	var result *Reader

	result = new(Reader)

	result.reader = reader
	result.bufioReader = bufio.NewReader(result.reader)

	return result
}
