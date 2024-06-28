package decoded_string_at_index

import "testing"

func TestDS(t *testing.T) {
	s := "cpmxv8ewnfk3xxcilcmm68d2ygc88daomywc3imncfjgtwj8nrxjtwhiem5nzqnicxzo248g52y72v3yujqpvqcssrofd99lkovg"
	k := 480551547
	t.Log(decodeAtIndexForce(s, k))
}
