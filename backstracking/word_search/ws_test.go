package word_search

import "testing"

func TestWS(t *testing.T) {
	board := [][]byte{
		{
			'A', 'B', 'C',
		},
		{
			'E', 'A', 'H',
		},
		{
			'G', 'K', 'L',
		},
	}
	word := "AHKL"
	word = "AL"
	word = "AAB"
	word = "ABCHLK"
	board = [][]byte{{'a'}}
	word = "a"
	t.Log(exist(board, word))
}
