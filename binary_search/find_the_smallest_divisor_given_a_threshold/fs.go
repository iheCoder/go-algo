package find_the_smallest_divisor_given_a_threshold

import "sort"

func smallestDivisor(nums []int, threshold int) int {
	sort.Ints(nums)
	end := nums[len(nums)-1]
	x := sort.Search(end, func(i int) bool {
		target := i + 1
		var sum int
		for _, num := range nums {
			sum += upperDivide(num, target)
		}

		return sum <= threshold
	})

	return x + 1
}

func upperDivide(x, y int) int {
	z := x / y
	if x%y != 0 {
		z++
	}
	return z
}
