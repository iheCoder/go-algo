package find_min_time_to_finish_all_jobs

import (
	"testing"
)

func TestFM(t *testing.T) {
	arr := []int{12, 13, 14, 17, 25}
	k := 3
	//t.Log(minimumTimeRequired(arr, k))
	//arr = []int{5, 5, 4, 4, 4}
	//k = 2
	//t.Log(minimumTimeRequired(arr, k))
	//arr = []int{5, 5, 5}
	//k = 3
	//t.Log(minimumTimeRequired(arr, k))
	//arr = []int{2, 3, 5}
	//k = 2
	//t.Log(minimumTimeRequired(arr, k))
	arr = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	k = 4
	t.Log(minimumTimeRequired(arr, k))
}

func TestRightestZero(t *testing.T) {
	for mask := 1; mask < 1<<12; mask++ {
		// chatgpt失败率还挺高的以下算法都无法得到最右边的0所在位置
		//got := (^mask+1)&mask
		//got := 2 + int(math.Log2(float64(mask&-mask)))
		//got := bits.TrailingZeros(uint(mask))
		got := ^mask & (mask + 1)
		expected := getRightPos(mask)
		if got != expected+1 {
			t.Fatalf("mask %d got %d expect %d", mask, got, expected)
		}
	}
	mask := 0
	t.Log((^mask + 1) & mask)
}

func getRightPos(x int) int {
	var i int
	for x&1 == 1 {
		x >>= 1
		i++
	}
	return i
}
