package max_of_absolute_val_expr

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	var ans int
	a, b, c, d := make([]int, 0), make([]int, 0), make([]int, 0), make([]int, 0)
	n := len(arr1)
	for i := 0; i < n; i++ {
		x, y := arr1[i], arr2[i]
		a = append(a, x+y+i)
		b = append(b, x+y-i)
		c = append(c, x-y+i)
		d = append(d, x-y-i)
	}

	ans = maxInts(
		maxInts(a...)-minInts(a...),
		maxInts(b...)-minInts(b...),
		maxInts(c...)-minInts(c...),
		maxInts(d...)-minInts(d...),
	)

	return ans
}

func maxInts(arr ...int) int {
	a := arr[0]
	for i := 1; i < len(arr); i++ {
		if a < arr[i] {
			a = arr[i]
		}
	}

	return a
}

func minInts(arr ...int) int {
	a := arr[0]
	for i := 1; i < len(arr); i++ {
		if a > arr[i] {
			a = arr[i]
		}
	}

	return a
}
