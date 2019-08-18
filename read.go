package reader

// Reads a Line ending exactly with the 'CR'+'LF' Symbols Sequence.
// The two Symbols at the End of the Line (CR+LF) are included into the
// returned Result.
func (r *Reader) ReadLineEndingWithCRLF() ([]byte, error) {

	var err error
	var line []byte
	var linePart []byte
	var lineByte byte
	var success bool

	// We must find the exact Combination (Sequence) of CR and LF, where LF is
	// right after the CR.
	for !success {
		linePart, err = r.bufioReader.ReadBytes(CR)
		if err != nil {
			return []byte{}, err
		}
		lineByte, err = r.bufioReader.ReadByte()
		if err != nil {
			return []byte{}, err
		}
		line = append(line, append(linePart, lineByte)...)
		success = (lineByte == LF)
	}

	return line, nil
}
