package roman_to_int

import "testing"

func TestR2i(t *testing.T) {
	s := "III"
	t.Log(romanToInt(s))
	s = "I"
	t.Log(romanToInt(s))

	s = "II"
	t.Log(romanToInt(s))

	s = "IX"
	t.Log(romanToInt(s))

	s = "XIII"
	t.Log(romanToInt(s))

	s = "CDXL"
	t.Log(romanToInt(s))
	s = "MCMXCIV"
	t.Log(romanToInt(s))
}
