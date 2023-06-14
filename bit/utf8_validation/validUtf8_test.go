package utf8_validation

import "testing"

func TestValidYtf8(t *testing.T) {
	t.Logf("%2b", 145)
	t.Log(0b11111000)
	t.Log(0b10000000)
	data := []int{235, 140, 4}
	t.Log(validUtf8(data))
	data = []int{197, 130, 1}
	t.Log(validUtf8(data))
}
