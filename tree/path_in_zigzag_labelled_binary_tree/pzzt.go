package path_in_zigzag_labelled_binary_tree

import "math"

func pathInZigZagTree(label int) []int {
	if label == 1 {
		return []int{1}
	}
	n := int(math.Log2(float64(label)))
	fn := float64(n)
	lx := int(math.Pow(2, fn+1)-1+math.Pow(2, fn)) - label
	var flag bool
	r := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		if flag {
			r[n-i] = lx
			flag = false
		} else {
			r[n-i] = label
			flag = true
		}
		label >>= 1
		lx >>= 1
	}

	return r
}
