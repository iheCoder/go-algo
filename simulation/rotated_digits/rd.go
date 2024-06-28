package rotated_digits

var numsStatus = []int{
	0,
	0,
	1,
	-1,
	-1,
	1,
	1,
	-1,
	0,
	1,
}

// good：存在任意一个goodNums且不存在无效值
func rotatedDigits(n int) int {
	var ans int
	for i := 1; i <= n; i++ {
		x := i
		sum := 0
		for x > 0 {
			y := x % 10
			if numsStatus[y] == -1 {
				break
			}
			sum += numsStatus[y]
			x /= 10
		}
		if x <= 0 && sum > 0 {
			ans++
		}
	}

	return ans
}
