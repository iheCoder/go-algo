package replace_words

import "testing"

func TestRW(t *testing.T) {
	ds := []string{"ac", "ab"}
	s := "it is abnormal that this solution is accepted"
	t.Log(replaceWords(ds, s))
}
