package varint

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	res := encode(42000)
	x := decode(res)

	if x != 42000 {
		t.Fatal("Error, `x` should equal 42000")
	}
}
