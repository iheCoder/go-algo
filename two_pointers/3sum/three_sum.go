package _sum

import (
	"sort"
	"strconv"
)

func threeSumIJK(nums []int) [][]int {
	m := make(map[string]bool)

	sort.Ints(nums)
	r := make([][]int, 0)
	i, j, k := 0, 1, 2
	for ; i < len(nums)-2; i++ {
		for j = i + 1; j < len(nums)-1; j++ {
			for k = j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ri := []int{nums[i], nums[j], nums[k]}
					key := getKey(ri)
					if !m[key] {
						r = append(r, ri)
						m[key] = true
					}
				}
			}
		}
	}

	return r
}

func getKey(a []int) string {
	var s string
	for _, i := range a {
		s = s + strconv.Itoa(i) + "_"
	}
	return s
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	R := make([][]int, 0)

	var l, r int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return R
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r = i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				R = append(R, []int{nums[i], nums[l], nums[r]})
				lv := nums[l]
				for l < r && nums[l] == lv {
					l++
				}
				rv := nums[r]
				for l < r && nums[r] == rv {
					r--
				}

			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return R
}

// 			target := 0 - (nums[i] + nums[j])
//			if _, exists := m[target]; exists {
//				ri := []int{nums[i], nums[j], target}
//				r = append(r, ri)
//			}
