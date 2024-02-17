package find_min_time_to_finish_all_jobs

import (
	"math"
	"math/bits"
)

func minimumTimeRequired(jobs []int, k int) int {
	if len(jobs) == 1 {
		return jobs[0]
	}

	n := len(jobs)
	m := 1 << n
	// sums[i]是分配i工作的总工作量。即等于最新分配工作加上在最新分配工作之前的总工作量
	sums := make([]int, m)
	for i := 1; i < m; i++ {
		x := bits.TrailingZeros(uint(i))
		y := i ^ 1<<x
		sums[i] = sums[y] + jobs[x]
	}

	// dp[i][j] 表示给前i个人分配工作，工作的分配情况为j时，完成所有工作的最短时间
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i, sum := range sums {
		dp[0][i] = sum
	}

	for i := 1; i < k; i++ {
		for j := 0; j < 1<<n; j++ {
			v := math.MaxInt64
			// (x-1)&j求j子集
			// 关键在于利用了 (x - 1) 的结果会将 x 的最低位的 1 变为 0，而其他位保持不变。然后通过与 j 进行按位与运算，可以确保只保留 j 中相应位上为 1 的位置。
			for x := j; x > 0; x = (x - 1) & j {
				v = minInt(v, maxInt(dp[i-1][j-x], sums[x]))
			}
			dp[i][j] = v
		}
	}
	return dp[k-1][(1<<n)-1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func nearThan(target, x, y int) bool {
	if x == y {
		return false
	}
	return absInt(y-target) < absInt(x-target)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getRightmostPos(x int) int {
	var i int
	for x&1 == 1 {
		x >>= 1
		i++
	}
	return i
}

// 麻烦的点在于就算对于x求的了最接近average的组合，可是如果这个最接近有多个组合，那么如何处理这多个组合就是难题
// 同样差距的组合如果选择小的，那么可能最终结果就是大的，如果选择大的，那么这可能是一个并没有那么接近的组合，却将这个大的当作了答案
// 答案确实存在组合最小的集合中，可是这集合中最小的不一定是答案。因为可能存在最小的xmin的组合中存在必定大于average的组合
// 	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
//	var sum int
//	for _, job := range jobs {
//		sum += job
//	}
//	aver := sum / k
//	n := len(jobs)
//	var mask int
//	r := 1<<31 - 1
//	var ec int
//	for i := 0; i < n; i++ {
//		x := jobs[i]
//		if x >= aver || i == n-1 {
//			if x > aver {
//				r = minInt(r, x)
//			} else if x == aver {
//				ec++
//			}
//			continue
//		}
//		if i != 0 && x == jobs[i-1] {
//			continue
//		}
//		z := 1<<31 - 1
//		for mask = 1; mask < 1<<(n-i-1); mask++ {
//			y := x
//			for steps := 0; steps+i+1 < n; steps++ {
//				if mask>>steps&1 == 0 {
//					continue
//				}
//				y += jobs[steps+i+1]
//				if y > aver {
//					z = minInt(z, y)
//				} else if y == aver {
//					ec++
//				}
//			}
//		}
//		r = minInt(r, z)
//	}
//	if ec >= k {
//		return aver
//	}
//	return r

// 想找到最接近average的组合，但却忽略了组合可能有多个，而不仅仅只有两个
// 	var sum int
//	for _, job := range jobs {
//		sum += job
//	}
//	aver := sum / k
//	var x, xp, mask int
//	n := len(jobs)
//	var best, bestIndex int
//	var r int
//	limit := 1<<n - 1
//	for mask < limit {
//		// 取反+1，那么最低位的1便成了0，所以与只有最低位0以及之后的1会剩下
//		xp = getRightmostPos(mask)
//		x = jobs[xp]
//		best, bestIndex = x, xp
//		mask |= 1 << xp
//		for i := xp + 1; i < n; i++ {
//			if mask>>i&1 > 0 {
//				continue
//			}
//			if nearThan(aver, best, x+jobs[i]) {
//				best = x + jobs[i]
//				bestIndex = i
//			}
//		}
//		mask |= 1 << bestIndex
//		r = maxInt(r, best)
//	}
//	return r

// 通过贪心法将剩下的最大job给空闲最大的这不是最佳方法，比如[5, 5, 4, 4, 4]
// 	if k == 1 {
//		var sum int
//		for _, job := range jobs {
//			sum += job
//		}
//		return sum
//	}
//	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
//	if len(jobs) < k {
//		return jobs[0]
//	}
//	workload := make([]int, k)
//	for i := k - 1; i >= 0; i-- {
//		workload[i] = jobs[k-1-i]
//	}
//
//	for i := k; i < len(jobs); i++ {
//		workload[0] += jobs[i]
//		j, l := 0, 1
//		for l < k && workload[j] > workload[l] {
//			workload[j], workload[l] = workload[l], workload[j]
//			j++
//			l = j + 1
//		}
//	}
//	return workload[k-1]
