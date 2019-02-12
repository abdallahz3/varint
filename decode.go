package varint

func decode(encoded []byte) int64 {
	var result int64

	reverseArray(encoded)

	encodedAsArrayOfBooleans := []bool{}

	for _, b := range encoded {
		// loop over bits
		for i := 0; i < 7; i++ {
			temp := b & (1 << uint(6-i))

			if temp != 0 {
				encodedAsArrayOfBooleans = append(encodedAsArrayOfBooleans, true)
			} else {
				encodedAsArrayOfBooleans = append(encodedAsArrayOfBooleans, false)
			}
		}
	}

	l := len(encodedAsArrayOfBooleans)

	for i, b := range encodedAsArrayOfBooleans {
		if b {
			result = setBit(result, reverseIndex(i, l))
		}
	}

	return result
}
