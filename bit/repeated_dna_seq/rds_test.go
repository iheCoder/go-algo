package repeated_dna_seq

import "testing"

func TestRDS(t *testing.T) {
	s := "AAAAAAAAAAAAA"
	t.Log(findRepeatedDnaSequencesHashForce(s))
}
