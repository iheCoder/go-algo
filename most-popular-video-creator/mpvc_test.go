package most_popular_video_creator

import (
	"reflect"
	"testing"
)

func TestMPVC(t *testing.T) {
	type td struct {
		creators []string
		ids      []string
		views    []int
		ans      [][]string
	}
	tds := []td{
		//{
		//			[]string{"alice", "bob", "alice", "chris"},
		//	[]string{"one","two","three","four"},
		//	[]int{5,10,5,4},
		//	[][]string{{"alice","one"},{"bob","two"}},
		//},
		{
			[]string{"a", "a"},
			[]string{"aa", "a"},
			[]int{5, 5},
			[][]string{{"a", "a"}},
		},
	}
	for _, ti := range tds {
		r := mostPopularCreator(ti.creators, ti.ids, ti.views)
		if !reflect.DeepEqual(r, ti.ans) {
			t.Fatalf("expect %v got %v", ti.ans, r)
		}
	}
}
