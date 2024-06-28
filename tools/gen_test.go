package tools

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGenNums(t *testing.T) {
	t.Log(genNumsStr(10000, 10000))
}

func genNumsStr(n, limit int) string {
	nums := genRandNums(n, limit)
	ns := make([]string, len(nums))
	for i, num := range nums {
		ns[i] = strconv.Itoa(num)
	}
	return strings.Join(ns, ",")
}

func genRandNums(n, limit int) []int {
	rand.Seed(time.Now().Unix())
	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x := rand.Intn(limit)
		for x < 0 {
			x = rand.Intn(limit)
		}
		//if rand.Int()%2 == 1 {
		//	x = -x
		//}
		r = append(r, x)
	}
	return r
}

func TestGenShuffleNumbers(t *testing.T) {
	n := 100000
	nums := genShuffleNumbers(n)
	ns := make([]string, len(nums))
	for i, num := range nums {
		ns[i] = strconv.Itoa(num)
	}
	t.Log(strings.Join(ns, ","))
}

func genShuffleNumbers(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	for i := 0; i < n; i++ {
		x := rand.Intn(n)
		nums[i], nums[x] = nums[x], nums[i]
	}
	return nums
}
