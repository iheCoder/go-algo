package life_game

import "testing"

func TestCopyBoard(t *testing.T) {
	board := [][]int{{1, 2, 3}, {4, 5, 6}}
	cb := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		cb[i] = make([]int, len(board[i]))
		copy(cb[i], board[i])
	}
	board[0][0] = 10
	t.Log(cb)
}

func TestLG(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	t.Log(board)
}
