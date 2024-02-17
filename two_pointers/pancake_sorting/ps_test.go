package pancake_sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type testData struct {
	arr []int
}

func TestPS(t *testing.T) {
	tds := []testData{
		{
			arr: []int{3, 2, 4, 1},
		},
		{
			arr: []int{1, 2, 3, 4},
		},
	}

	limit := 100
	rand.Seed(time.Now().Unix())
	var count int
	var td testData
	for i := 0; i < limit; i++ {
		count = rand.Intn(limit)
		td = testData{arr: genRandArr(count)}
		tds = append(tds, td)
	}

	for i, td := range tds {
		r := pancakeSort(td.arr)
		if !sort.IntsAreSorted(td.arr) && len(r) <= 10*len(td.arr) {
			t.Fatalf("index %d got %v", i, td.arr)
		}
	}
}

func genRandArr(count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = i + 1
	}
	rand.Seed(time.Now().Unix())
	var j int
	for i := 0; i < count; i++ {
		j = rand.Intn(count)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
