package rec_area

import "testing"

func TestRA(t *testing.T) {
	type testData struct {
		ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int
		expected                               int
	}
	tds := []testData{
		{
			-3, 0, 3, 4, 0, -1, 9, 2, 45,
		},
		{
			-2, -2, 2, 2, 3, 3, 4, 4, 17,
		},
	}

	for i, td := range tds {
		got := computeArea(td.ax1, td.ay1, td.ax2, td.ay2, td.bx1, td.by1, td.bx2, td.by2)
		if got != td.expected {
			t.Fatalf("index %d got %d exp %d", i, got, td.expected)
		}
	}
}
