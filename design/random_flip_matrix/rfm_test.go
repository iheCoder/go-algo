package random_flip_matrix

import "testing"

func TestRFM(t *testing.T) {
	s := Constructor(10000, 10000)
	for i := 0; i < 10000; i++ {
		s.Flip()
		s.Flip()
		s.Flip()
		s.Flip()
		s.Reset()
	}
}
