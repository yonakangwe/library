package util

// EraseByteBuffer wipes content of given byte buffer.
func EraseByteBuffer(b []byte) {
	if b != nil {
		for i := range b {
			b[i] = 0
		}
	}
}

// EraseString wipes content of given string.
func EraseString(s *string) {
	EraseByteBuffer([]byte(*s))
}
