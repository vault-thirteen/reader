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

// This Package provides is a convenient Reader for specific Purposes.

// For Example, the built-in Go Language Library provides a Method to read
// from a Reader a single Line ending either with ASCII CR Symbol or with a
// Combination of CR with LF Symbols. It does not provide a Method to read a
// Line which ends exactly with a Combination of CR with LF Symbols! This
// Library helps to get rid of such Stupidity of the Developers of the Language.

import (
	"bufio"
	"io"
)

// ASCII Symbols.
const (
	CR = '\r'
	LF = '\n'
)

type Reader struct {
	reader      io.Reader
	bufioReader *bufio.Reader
}
