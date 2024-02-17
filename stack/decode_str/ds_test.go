package decode_str

import "testing"

func TestDS(t *testing.T) {
	s := "abc3[cd]xyz"
	t.Log(decodeString(s))
}
