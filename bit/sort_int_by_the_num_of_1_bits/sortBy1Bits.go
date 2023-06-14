package sort_int_by_the_num_of_1_bits

import "sort"

const (
	bit1Mask  = 0x55555555
	bit2Mask  = 0x33333333
	bit4Mask  = 0x0f0f0f0f
	bit8Mask  = 0x00ff00ff
	bit16Mask = 0x0000ffff
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] == arr[j] {
			return true
		}
		il, jl := oneCount(uint(arr[i])), oneCount(uint(arr[j]))
		if il != jl {
			return il < jl
		}
		return arr[i] < arr[j]
	})
	return arr
}

func oneCount(num uint) uint {
	num = num&bit1Mask + num>>1&bit1Mask
	num = num&bit2Mask + num>>2&bit2Mask
	num = num&bit4Mask + num>>4&bit4Mask
	num = num&bit8Mask + num>>8&bit8Mask
	num = num&bit16Mask + num>>16&bit16Mask
	return num
}

// type as []int
//
//func (a as) Swap(i, j int) {
//	a[i], a[j] = a[j], a[i]
//}
//
//func (a as) Less(i, j int) bool {
//	if a[i] == a[j] {
//		return true
//	}
//	return oneCount(uint(a[i])) <= oneCount(uint(a[j]))
//}
//
//func (a as) Len() int {
//	return len(a)
//}
