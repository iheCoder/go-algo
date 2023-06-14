package predict_the_winner

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPW(t *testing.T) {
	type testData struct {
		nums     []int
		expected bool
	}
	tds := []testData{
		{
			nums:     []int{1, 5, 2},
			expected: false,
		},
		{
			nums:     []int{1, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			nums:     []int{1, 2, 2, 1},
			expected: true,
		},
		{
			nums:     []int{1, 5, 2, 4, 6},
			expected: true,
		},
	}
	for i, td := range tds {
		got := PredictTheWinner(td.nums)
		if td.expected != got {
			t.Fatalf("%d nums %d expected %v got %v", i, td.nums, td.expected, got)
		}
	}
}

func TestGenRandNums(t *testing.T) {
	r := genRandNums(100000, 300)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func genRandNums(n, limit int) []int {
	rand.Seed(time.Now().Unix())
	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(limit)
		//x:=1
		sig := rand.Intn(2)
		if sig != 1 {
			x = -x
		}
		//for x == 0 {
		//	x = rand.Intn(limit)
		//}
		r = append(r, x)
	}
	return r
}
