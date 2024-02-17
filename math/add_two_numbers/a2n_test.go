package add_two_numbers

import "testing"

func TestXChangeWillAffectXPointer(t *testing.T) {
	type simple struct {
		x int
	}

	type containSimple struct {
		s *simple
	}

	s := &simple{}
	cs := &containSimple{s: s}
	t.Log(cs)
	//*s=nil
}
