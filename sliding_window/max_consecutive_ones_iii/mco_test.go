package max_consecutive_ones_iii

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMCO(t *testing.T) {
	t.Log(byte('a'))
	t.Log(byte('A'))
	t.Log(byte('a') - byte('A'))
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	//t.Log(longestOnes(nums, k))
	nums = []int{0, 0, 1, 1, 1, 0, 0}
	k = 0
	//t.Log(longestOnes(nums, k))
	nums = []int{1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1}
	k = 8
	//t.Log(longestOnes(nums, k))
	nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 0
	t.Log(longestOnes(nums, k))
}

func TestGen01(t *testing.T) {
	r := gen01(100000)
	var ss []string
	for _, i := range r {
		ss = append(ss, strconv.Itoa(i))
	}
	t.Log(strings.Join(ss, ","))
}

func gen01(count int) []int {
	r := make([]int, count)
	rand.Seed(time.Now().Unix())
	var x int
	for i := 0; i < count; i++ {
		x = rand.Intn(3)
		if x == 2 {
			r[i] = 1
		}
	}
	return r
}
