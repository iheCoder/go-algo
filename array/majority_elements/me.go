package majority_elements

import "math/rand"

//func majorityElement(nums []int) int {
//	o := 1
//	rem := nums
//	mi, mv := 0, 0
//	arr := make([][]int, 10)
//	for len(rem) != 0 {
//		for _, num := range rem {
//			p := num%(10^o) - num%(10^(o-1))
//			arr[p%10] = append(arr[p%10], num)
//			if len(arr[p%10]) > mv {
//				mi = p % 10
//				mv = len(arr[p%10])
//			}
//		}
//		rem = arr[mi]
//	}
//}

func majorityElement(nums []int) int {
	var count, candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElementRandom(nums []int) int {
	n := len(nums)
	mc := n / 2
	for {
		candidate := nums[rand.Int63n(int64(n))]
		count := 0
		for _, num := range nums {
			if num == candidate {
				count++
			}
			if count > mc {
				return candidate
			}
		}
	}
}

func majorityElementDivide(nums []int) int {
	var f func(lo, hi int) int
	f = func(lo, hi int) int {
		// 如果仅有一个元素，那么该元素就是众数
		if lo == hi {
			return nums[lo]
		}

		// 递归数组左右边，找到左右众数
		mid := lo + (hi-lo)/2
		left := f(lo, mid)
		right := f(mid+1, hi)

		// 若左右众数相等，直接返回
		if left == right {
			return left
		}

		var lc, rc int
		for i := lo; i <= hi; i++ {
			switch nums[i] {
			case left:
				lc++
			case right:
				rc++
			}
		}

		if lc > rc {
			return left
		}
		return right
	}
	return f(0, len(nums)-1)
}
