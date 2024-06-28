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
	r := genRandNums(30000, 30000)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func TestGenFixedNum(t *testing.T) {
	r := genFixedNum(1, 20)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func genFixedNum(num, count int) []int {
	arr := make([]int, count)
	for i := 0; i < count; i++ {
		arr[i] = num
	}
	return arr
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

func TestGenRandStr(t *testing.T) {
	n, limit := 1, 100000
	r := genRandStr(n, limit)
	for i := 0; i < len(r); i++ {
		r[i] = "\"" + r[i] + "\""
	}
	t.Log(strings.Join(r, ","))
}

func genRandStr(n, limit int) []string {
	r := make([]string, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		l := rand.Intn(limit)
		var s []byte
		for j := 0; j < l+1; j++ {
			x := byte(rand.Intn(26))
			s = append(s, 'a'+x)
		}
		r[i] = string(s)
	}
	return r
}

func TestGenUniqueRandNums(t *testing.T) {
	r := genUniqueRandNums(300, 5001)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func genUniqueRandNums(n, limit int) []int {
	rand.Seed(time.Now().Unix())
	r := make([]int, limit)
	for i := 0; i < limit; i++ {
		r[i] = i + 1
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(limit)
		r[i], r[j] = r[j], r[i]
	}
	return r[:n]
}

func TestGenSortedArray(t *testing.T) {
	r := genSortedArray(50000, 5, true)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func genSortedArray(n, p int, reverse bool) []int {
	//rand.Seed(time.Now().Unix())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		//x := rand.Intn(p)
		//if x == 0 {
		//	nums[i] = i - 1
		//} else {
		//	nums[i] = i
		//}
		if reverse {
			nums[i] = n - i
		} else {
			nums[i] = i
		}
	}
	return nums
}
