package max_consecutive_ones_iii

func longestOnes(nums []int, k int) int {
	if k == len(nums) {
		return len(nums)
	}
	var r, i, curSum int
	if k == 0 {
		for {
			for i < len(nums) && nums[i] == 0 {
				i++
			}
			for i < len(nums) && nums[i] == 1 {
				curSum++
				i++
			}
			r = maxInt(r, curSum)
			if i >= len(nums) {
				return r
			}
			curSum = 0
		}
	}
	used := 0
	// 1. 只要碰到0就用变1额度，直到无额度且恰好从1出来，然后统计1的数量最大值
	for j := 0; j < len(nums); j++ {
		switch nums[j] {
		// 若为0，有额度则变1；无额度，就统计。并收回到下一个1之间的所有额度
		case 0:
			if used < k {
				used++
				curSum++
			} else {
				// 左边界是0直接收回，若是1直到0的时候再回收
				// 回收之后要重新从j计算
				for nums[i] == 1 {
					i++
					curSum--
				}
				// 要么等于1，不进行收回了，要么used==0，所有的都回收了
				// 是否存在used==0但nums[i]也等于0？如果从i到j一直都是0，那么存在；其他情况不可能
				// 全部回收恐怕是有问题的吧，如果恰好部分回收和相邻的1组合起来最优，那么就显然不对了
				if nums[i] == 0 && used > 0 {
					used--
					i++
					curSum--
				}
				if used < k {
					used++
					curSum++
				}
			}
		// 若为1，统计
		case 1:
			curSum++
		}
		r = maxInt(r, curSum)
	}
	return r
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
