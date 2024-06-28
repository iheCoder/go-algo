package max_distance_to_closest_person

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestMD(t *testing.T) {
	seats := []int{1, 0, 0, 0}
	//t.Log(maxDistToClosest(seats))
	seats = []int{1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0}
	t.Log(maxDistToClosest(seats))
}

func TestGenRanOneZero(t *testing.T) {
	n := int(100000)
	x := rand.Intn(n)
	nums := make([]int, n)
	for i := 0; i < x; i++ {
		y := rand.Intn(n)
		nums[y] = 1
	}

	ns := make([]string, len(nums))
	for i, num := range nums {
		ns[i] = strconv.Itoa(num)
	}
	t.Log(strings.Join(ns, ""))
	//t.Log(strings.Join(ns, ","))
}
