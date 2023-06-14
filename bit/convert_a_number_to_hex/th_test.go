package convert_a_number_to_hex

import "testing"

func TestHigh3In32(t *testing.T) {
	x := uint32(15)
	t.Log(x << 28)
	t.Log(^(-1))
	t.Log(^0)
}

func TestToHex(t *testing.T) {
	x := 3
	t.Log(toHex(x))
	x = 15
	t.Log(toHex(x))
	x = 16
	t.Log(toHex(x))
	x = -1
	t.Log(toHex(x))
}
