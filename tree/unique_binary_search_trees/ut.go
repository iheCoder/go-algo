package unique_binary_search_trees

var box = make(map[int]int)

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	box[0] = 1
	box[1] = 1
	box[2] = 2

	return nt(n)
}

func nt(n int) int {
	var v int
	var ok bool
	if v, ok = box[n]; ok {
		return v
	}

	var sum int
	for i := 1; i <= n; i++ {
		sum += nt(n-i) * nt(i-1)
	}

	box[n] = sum
	return sum
}

// 		left, ok = box[i-1]
//		if !ok {
//			v = 0
//			for j := 1; j <= i-1; j++ {
//				v += nt(i-1-j) * nt(j-1)
//			}
//			box[i-1] = v
//		}
//		left=nt()
//
//		right, ok = box[n-i]
//		if !ok {
//			v = 0
//			v = 0
//			for j := 1; j <= n-i; j++ {
//				v += nt(n-i-j) * nt(j-1)
//			}
//			box[i-1] = v
//		}
