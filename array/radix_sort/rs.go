package radix_sort

import "math"

func digit(num, exp int) int {
	return (num / exp) % 10
}

func countingSortDigit(nums []int, exp int) {
	counter := make([]int, 10)
	n := len(nums)

	// 统计0～9数字出现次数
	for i := 0; i < n; i++ {
		d := digit(nums[i], exp)
		counter[d]++
	}

	// 计算出现次数的前缀和，以方便计算后面元素所放置位置
	for i := 1; i < 10; i++ {
		counter[i] += counter[i-1]
	}

	// 根据数字找到所在的桶
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		d := digit(nums[i], exp)
		j := counter[d] - 1
		res[j] = nums[i]
		counter[d]--
	}
	for i := 0; i < n; i++ {
		nums[i] = res[i]
	}
}

func radixSort(nums []int) {
	// 找到数组中最大的数
	x := math.MinInt
	for _, num := range nums {
		if num > x {
			x = num
		}
	}

	for exp := 1; exp <= x; exp *= 10 {
		countingSortDigit(nums, exp)
	}
}
