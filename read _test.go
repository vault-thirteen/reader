package reader

import (
	"bytes"
	"io"
	"testing"
)

func Test_ReadLineEndingWithCRLF(t *testing.T) {

	var data []byte
	var err error
	var reader1 io.Reader
	var reader2 *Reader
	var result []byte
	var resultExpected []byte

	// Test #1. Normal Data.

	// Prepare the Data.
	data = []byte("12")
	data = append(data, CR)
	data = append(data, []byte("34")...)
	data = append(data, LF)
	data = append(data, []byte("5")...)
	data = append(data, CR, LF)
	data = append(data, []byte("67")...)
	resultExpected = data[0:9]

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = New(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	if err != nil {
		t.Error("Test #1")
		t.FailNow()
	}
	if !bytes.Equal(result, resultExpected) {
		t.Error("Test #1")
		t.FailNow()
	}

	// Test #2. Data is not enough.

	// Prepare the Data.
	data = []byte("12")
	data = append(data, CR)
	data = append(data, []byte("34")...)
	data = append(data, LF)
	data = append(data, []byte("5")...)
	data = append(data, LF, CR)
	data = append(data, []byte("67")...)
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = New(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	if err != io.EOF {
		t.Error("Test #2")
		t.FailNow()
	}
	if !bytes.Equal(result, resultExpected) {
		t.Error("Test #2")
		t.FailNow()
	}

	// Test #3. Empty Data.

	// Prepare the Data.
	data = []byte{}
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = New(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	if err != io.EOF {
		t.Error("Test #3")
		t.FailNow()
	}
	if !bytes.Equal(result, resultExpected) {
		t.Error("Test #3")
		t.FailNow()
	}
}
