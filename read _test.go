////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package reader

import (
	"bytes"
	"io"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ReadLineEndingWithCRLF(t *testing.T) {

	var data []byte
	var err error
	var reader1 io.Reader
	var reader2 *Reader
	var result []byte
	var resultExpected []byte
	var tst *tester.Test

	tst = tester.New(t)

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
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)

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
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.EOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #3. Empty Data.

	// Prepare the Data.
	data = []byte{}
	resultExpected = []byte{}

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = New(reader1)
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeAnError(err)
	tst.MustBeEqual(err.Error(), io.EOF.Error())
	tst.MustBeEqual(result, resultExpected)

	// Test #4. Normal Data. Double Read.

	// Prepare the Data.
	data = []byte("ABC")
	data = append(data, CR, LF)
	data = append(data, []byte("DEF")...)
	data = append(data, CR, LF)
	data = append(data, []byte("XYZ")...)

	// Run the Test.
	reader1 = bytes.NewReader(data)
	reader2 = New(reader1)

	// Part 1.
	resultExpected = []byte("ABC\r\n")
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)

	// Part 2.
	resultExpected = []byte("DEF\r\n")
	result, err = reader2.ReadLineEndingWithCRLF()
	tst.MustBeNoError(err)
	tst.MustBeEqual(result, resultExpected)
}
