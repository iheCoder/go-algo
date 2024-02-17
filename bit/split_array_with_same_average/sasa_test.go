package split_array_with_same_average

import "testing"

func TestSASA(t *testing.T) {
	arr := []int{14, 32, 134, 11, 356, 77, 34, 23, 666, 224, 67, 34, 57, 62, 66, 22, 99, 11, 34, 1, 0, 33, 112, 44, 13, 55, 11, 55, 33, 73}
	//t.Log(splitArraySameAverage(arr))
	//arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//t.Log(splitArraySameAverage(arr))
	//arr = []int{18, 10, 5, 3}
	//t.Log(splitArraySameAverage(arr))
	arr = []int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}
	t.Log(splitArraySameAverage(arr))
}
