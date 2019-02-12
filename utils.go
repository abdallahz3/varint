package varint

func reverseArray(buff []byte) {
	for i, j := 0, len(buff)-1; i < j; i, j = i+1, j-1 {
		buff[i], buff[j] = buff[j], buff[i]
	}
}

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	return n &^ (1 << pos)
}

func reverseIndex(i int, relativeTo int) uint {
	return uint(relativeTo - i - 1)
}
