package ways_to_split_array_into_three_subarrays

import "sort"

func waysToSplit(nums []int) (r int) {
	if len(nums) < 3 {
		return 0
	}

	defer func() {
		if r > 1000000007 {
			r = r % 1000000007
		}
	}()

	sums := make([]int, len(nums))
	var sum int
	for i, num := range nums {
		sum += num
		sums[i] = sum
	}
	if sum == 0 {
		n := len(nums) - 2
		r = (n * (n + 1)) / 2
		return
	}

	var si int
	var j, k int
	for i := 0; i < len(nums)-2; i++ {
		si = sums[i]
		j = sort.SearchInts(sums, 2*si)
		for j <= i {
			j++
		}
		if j > len(nums)-2 {
			return r
		}
		k = sort.SearchInts(sums, (sum+si)/2+1)
		// 如果将k视为闭区间，那么当k为len(num)的时候并不能确定sum[k]就是要找的，还是根本不存在该值
		if k >= len(nums) {
			if 2*nums[len(nums)-1] > sum+si {
				continue
			}
			k--
		}
		if k <= j {
			continue
		}
		r += k - j
	}

	return r
}

// 	for i := 0; i < len(nums)-2; i++ {
//		si = sums[i]
//		for j := i + 1; j < len(nums)-1; j++ {
//			sj = sums[j] - si
//			if sj < si {
//				continue
//			}
//			sk = sum - sums[j]
//			if sk < sj {
//				break
//			}
//			r++
//		}
//	}

// 	i, j := 0, len(nums)-2
//	for i < j {
//		si = sums[i]
//		sj = sums[j] - si
//		// 在i++的情况下，j是能够也++的
//		if sj < si {
//			return r
//		}
//		sk = sum - sums[j]
//		if sk < sj {
//			j--
//			continue
//		}
//
//		t = sort.SearchInts(sums, 2*si)
//		if t > i && t <= j {
//			r += j - t
//		}
//		//for j > i && sj >= si {
//		//	r++
//		//	j--
//		//	sj = sums[j] - si
//		//}
//		i++
//	}
