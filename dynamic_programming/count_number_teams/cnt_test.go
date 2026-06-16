package count_number_teams

import "testing"

func TestNT(t *testing.T) {
	ratings := []int{2, 5, 3, 4, 1}
	t.Log(numTeams(ratings))
}
