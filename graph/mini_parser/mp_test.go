package mini_parser

import "testing"

func TestMP(t *testing.T) {
	s := "123"
	//t.Log(deserialize(s))
	s = "[123,[456,[789]]]"
	t.Log(deserialize(s))
}
