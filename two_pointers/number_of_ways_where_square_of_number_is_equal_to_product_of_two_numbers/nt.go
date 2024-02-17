package number_of_ways_where_square_of_number_is_equal_to_product_of_two_numbers

import (
	"fmt"
	"sort"
)

// 找到i所在元素在另一个数组index，然后在index两边尝试找到相等的。可问题在于终止条件是到两端，这个复杂度太高了
func numTriplets(nums1 []int, nums2 []int) int {
	var r int
	m1 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]*nums1[i]]++
	}
	m2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]*nums2[i]]++
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	var sc int
	var ok bool
	for i := 0; i < len(nums1)-1; i++ {
		for j := i + 1; j < len(nums1); j++ {
			if sc, ok = m2[nums1[i]*nums1[j]]; ok {
				r += sc
			}
		}
	}

	for i := 0; i < len(nums2)-1; i++ {
		for j := i + 1; j < len(nums2); j++ {
			if sc, ok = m1[nums2[i]*nums2[j]]; ok {
				r += sc
			}
		}
	}

	return r
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

// 	m1 := make(map[string]int)
//	minNum2 := nums2[0]
//	maxNum2 := nums2[len(nums2)-1]
//	var j,k,index int
//	var value int
//	var ok bool
//	for i := 0; i < len(nums1); i++ {
//		if nums1[i] > maxNum2 || nums1[i] < minNum2 {
//			continue
//		}
//		index=sort.SearchInts(nums2,nums1[i])
//		j,k=index-1,index
//
//		value,ok=m1[getKey(index-1,index)]
//
//
//	}
//
//	m2 := make(map[string]int)

// 		// 相同似乎不能直接加lc。还有一种可能是i, i-1值相同并且该值在另一个nums也存在
//		if i > 0 && nums1[i] == nums1[i-1] {
//			r += lc
//			continue
//		}

// 			// 一个数乘以不一样的数当然不可能等于之前的平方，可是却能等于其他平方
//			//if lc > 0 && nums1[j] != nums1[j-1] {
//			//	break
//			//}
