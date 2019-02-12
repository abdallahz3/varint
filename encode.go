package varint

const sevenOnes int64 = 0x7F

func encode(x int64) []byte {
	res := []byte{}

	for i := 0; i < 9; i++ {
		temp := sevenOnes << uint(i*7)
		temp = x & temp
		temp = temp >> uint(i*7)

		if temp == 0 {
			break
		}

		temp = setBit(temp, 7)

		res = append(res, byte(temp))
	}

	res[len(res)-1] = byte(clearBit(int64(res[len(res)-1]), 7))

	return res
}
