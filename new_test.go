package reader

import (
	"bytes"
	"io"
	"testing"
)

func Test_New(t *testing.T) {

	var reader io.Reader
	var result *Reader

	reader = bytes.NewReader([]byte{})
	result = New(reader)
	if result.reader != reader {
		t.FailNow()
	}
	if result.bufioReader == nil {
		t.FailNow()
	}
}
