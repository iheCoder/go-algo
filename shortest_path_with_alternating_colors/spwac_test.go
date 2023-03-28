package shortest_path_with_alternating_colors

import (
	"reflect"
	"testing"
)

func TestShortestAlternatingPaths(t *testing.T) {
	type td struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
		expected  []int
	}
	tds := []td{

		//{
		//	n:         5,
		//	redEdges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
		//	blueEdges: [][]int{{1, 2}, {2, 3}, {3, 1}},
		//	expected:  []int{0, 1, 2, 3, 7},
		//},
		{
			n:         3,
			blueEdges: [][]int{{0, 1}, {0, 2}},
			redEdges:  [][]int{{1, 2}},
			expected:  []int{0, 1, 1},
		},
	}

	for _, ti := range tds {
		r := shortestAlternatingPaths(ti.n, ti.redEdges, ti.blueEdges)
		if !reflect.DeepEqual(r, ti.expected) {
			t.Fatalf("expected %v got %v", ti.expected, r)
		}
	}
}
