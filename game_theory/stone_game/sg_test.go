package stone_game

import "testing"

func TestSG(t *testing.T) {
	type testData struct {
		nums     []int
		expected bool
	}
	tds := []testData{
		{
			nums:     []int{5, 3, 4, 5},
			expected: true,
		},
	}

	for i, td := range tds {
		got := stoneGame(td.nums)
		if td.expected != got {
			t.Fatalf("%d nums %d expected %v got %v", i, td.nums, td.expected, got)
		}
	}
}
