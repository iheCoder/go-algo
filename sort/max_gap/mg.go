package max_gap

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	n := len(nums)

	var mx int
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	var rank int
	for mx > 0 {
		rank++
		mx /= 10
	}

	bs := make([][]int, 10)
	cbs := make([][]int, 10)
	bs[0] = nums

	base := 1
	count := n
	for i := 0; i < rank; i++ {
		for k, b := range bs {
			for _, bi := range b {
				j := bi / base % 10
				cbs[j] = append(cbs[j], bi)
				count--
			}
			bs[k] = make([]int, 0)
			if count <= 0 {
				break
			}
		}
		bs, cbs = cbs, bs
		count = n
		base *= 10
	}

	var j int
	for _, b := range bs {
		for _, bi := range b {
			nums[j] = bi
			j++
		}
	}

	var ans int
	for i := 1; i < n; i++ {
		d := nums[i] - nums[i-1]
		if d > ans {
			ans = d
		}
	}
	return ans
}
