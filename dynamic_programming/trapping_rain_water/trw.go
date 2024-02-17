package trapping_rain_water

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	sums := make([]int, len(height))
	var sum int
	for i := 0; i < len(height); i++ {
		sum += height[i]
		sums[i] = sum
	}

	lastIndex, lastMax := 0, height[0]
	var total, cur int
	for i := 1; i < len(height); i++ {
		if height[i] >= lastMax {
			cur = lastMax*(i-lastIndex) - sums[i] + sums[lastIndex]
			if cur > 0 {
				total += cur
			}
			lastIndex = i
			lastMax = height[i]
		}
	}
	return total
}
