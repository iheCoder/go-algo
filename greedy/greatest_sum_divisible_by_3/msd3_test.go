package greatest_sum_divisible_by_3

import "testing"

func TestMSD3(t *testing.T) {
	nums := []int{3, 6, 5, 1, 8}
	t.Log(maxSumDivThree(nums))
	nums = []int{1, 1}
	t.Log(maxSumDivThree(nums))
	nums = []int{1, 4, 7, 10, 13}
	t.Log(maxSumDivThree(nums))
}

func TestFindMinTwo(t *testing.T) {
	arr := []int{1, 3, 9, 2, 2}
	t.Log(findMinTwo(arr))
	arr = []int{1, 3, 0, 9, 2, 2}
	t.Log(findMinTwo(arr))
	arr = []int{1, 3, 9, 2, 0}
	t.Log(findMinTwo(arr))
	arr = []int{3, 9, 2, 0, 1}
	t.Log(findMinTwo(arr))
}

func findMinTwo(arr []int) (int, int) {
	x, y := 1<<32-1, 1<<32-1
	for _, num := range arr {
		if x > num {
			y = x
			x = num
		} else if y > num {
			y = num
		}
	}
	return x, y
}
