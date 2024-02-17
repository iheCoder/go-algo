package shortest_subarray_with_sum_at_least_k

// 右边界只要窗口不是负收益，那么右边界j都可以继续移动
// 只要左边界i不在负数，那么都可以移动
// 只要i到j区间仍然大于等于k，那么i,j就是可以移动的
func shortestSubarray(nums []int, k int) int {
	r := -1

	var sum, tmpSum int
	sums := make([]int, len(nums))
	for i, num := range nums {
		sum += num
		sums[i] = sum
	}
	var i, j int
	for i < len(nums) && nums[i] <= 0 {
		i++
	}
	// sum 是i到j之间的和
	sum = 0
	for j = i; j < len(nums); j++ {
		sum += nums[j]
		if sum >= k {
			// 会不会存在一种可能窗口i的时候比i+1的时候更小呢？
			// 找到以j为右边界的最大i索引
			tmpSum = nums[j]
			ni := j
			for ni > i && tmpSum < k {
				ni--
				tmpSum += nums[ni]
			}
			sum = tmpSum
			i = ni
			r = minInt(r, j-i+1)
			if r == 1 {
				return 1
			}
			// 若存在大正数、中负数、小正数虽然他们之和也是正数，但是一旦大正数没了，就应该将中负数也删除
			ni = j
			// 如果ni等于i会怎么办呢？说明从i+1到j任何前缀和都是正数
			// 找到去掉nums[i]之后索引ni使得i+1到ni之间的和为负数的数，并且使i更新到ni+1也就是恰好和为正数的索引
			for ni > i && sums[ni]-sums[i] > 0 {
				ni--
			}
			// sum是从i到j区间和，sums[ni]-sums[i]+nums[i]则是i到ni之间的和，因此sum会等于ni+1到j区间的和
			// sum -= sums[ni] - sums[i] - nums[i]为什么得到正确答案20?
			// 最简单的思路其实是不更新i，但还是超时
			sum -= sums[ni] - sums[i] + nums[i]
			i = ni + 1
		} else if sum <= 0 {
			i = j + 1
			sum = 0
		}
	}

	return r
}

func minInt(x, y int) int {
	if x == -1 {
		return y
	}
	if x < y {
		return x
	}
	return y
}
