package battleships_in_a_board

import "testing"

func TestBB(t *testing.T) {
	b := [][]byte{
		{'X', 'X', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}
	t.Log(countBattleships(b))
}
