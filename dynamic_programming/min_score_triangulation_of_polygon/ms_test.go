package min_score_triangulation_of_polygon

import "testing"

func TestMS(t *testing.T) {
	vs := []int{1, 2, 3}
	t.Log(minScoreTriangulation(vs))
}
