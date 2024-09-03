package pairs_of_songs_with_total_durations_divisible_by_60

import "testing"

func TestPSD60(t *testing.T) {
	ta := []int{30, 20, 150, 100, 40}
	t.Log(numPairsDivisibleBy60(ta))
}
