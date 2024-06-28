package knight_probability_in_chessboard

import "testing"

func TestKpic(t *testing.T) {
	n, k, row, col := 3, 2, 0, 0
	t.Log(knightProbability(n, k, row, col))
}
