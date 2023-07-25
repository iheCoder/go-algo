package target_sum

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < target || -sum > target {
		return 0
	}

	dpa := make(map[int]int)
	dpb := make(map[int]int)
	var cs int
	lastdp := dpb
	curDp := dpa
	lastdp[nums[0]]++
	lastdp[-nums[0]]++
	cs += nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		cs += num
		for j := -cs; j <= cs; j++ {
			curDp[j] = lastdp[j-num] + lastdp[j+num]
		}
		curDp, lastdp = lastdp, curDp
	}
	return lastdp[target]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
