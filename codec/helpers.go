package codec

// make byte array with length 1 << 6; used to test 2-byte length encoding
func byteArray(length int) ([]byte) {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = 0xff
	}
	return b
}