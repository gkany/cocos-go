package util

// byte -- 8 bits
func PutUvarintByByte(buf []byte, x uint64) int {
	i := 0
	for x > 0xff {
		buf[i] = byte(x)
		x >>= 8
		i++
	}
	buf[i] = byte(x)
	return i + 1
}
