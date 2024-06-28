package implement_magic_dict

import "testing"

func TestIMD(t *testing.T) {
	d := Constructor()
	ws := []string{"hello", "leetcode"}
	d.BuildDict(ws)
	d.Search("hhllo")
}
