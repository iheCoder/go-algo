package del_and_earn

func deleteAndEarn(nums []int) int {
	var mv int
	for _, num := range nums {
		mv = maxInt(mv, num)
	}
	sums := make([]int, mv+1)
	for _, num := range nums {
		sums[num] += num
	}
	return rob(sums)
}

func rob(nums []int) int {
	first, second := nums[0], maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		first, second = second, maxInt(first+nums[i], second)
	}
	return second
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 果然超出时间限制了！
// func deleteAndEarn(nums []int) int {
//	m := make(map[int]int)
//	for _, num := range nums {
//		m[num]++
//	}
//
//	var ans, pick int
//	for len(m) > 0 {
//		mx := math.MinInt
//		for k, _ := range m {
//			temp := k - (k-1)*m[k-1] - (k+1)*m[k+1]
//			if temp > mx {
//				mx = temp
//				pick = k
//			}
//		}
//		ans += pick
//		m[pick]--
//		delete(m, pick-1)
//		delete(m, pick+1)
//		if m[pick] == 0 {
//			delete(m, pick)
//		}
//	}
//	return ans
//}
