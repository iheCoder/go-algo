package additive_number

import "strconv"

func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	n := len(num)
	for i := 0; i < n-2; i++ {
		if num[0] == '0' && i > 0 {
			return false
		}
		for j := i + 1; j < n-1; j++ {
			if num[i+1] == '0' && j > i+1 {
				break
			}
			for k := j + 1; k < n; k++ {
				if num[j+1] == '0' && k > j+1 {
					break
				}
				if addEqual(num, i, j, k) {
					if k+1 >= n || isAdditiveNumber(num[i+1:]) {
						return true
					}
				}
			}
		}
	}
	return false
}

func addEqual(num string, i, j, k int) bool {
	x, _ := strconv.Atoi(num[:i+1])
	y, _ := strconv.Atoi(num[i+1 : j+1])
	z, _ := strconv.Atoi(num[j+1 : k+1])
	return x+y == z
}

func ianDFS(num string) bool {
	l := make([][]int, 0)
	n := len(num)
	var f func(i int) bool
	var check func(a, b, c []int) bool
	check = func(a, b, c []int) bool {
		ans := make([]int, 0)
		var t int
		for i := 0; i < len(a) || i < len(b); i++ {
			if i < len(a) {
				t += a[i]
			}
			if i < len(b) {
				t += b[i]
			}
			ans = append(ans, t%10)
			t /= 10
		}

		if t > 0 {
			ans = append(ans, t)
		}

		ok := len(c) == len(ans)
		for i := 0; i < len(c); i++ {
			if c[i] != ans[i] {
				ok = false
			}
		}
		return ok
	}

	f = func(i int) bool {
		m := len(l)
		if i >= n {
			return m >= 3
		}

		edge := n
		if num[i] == '0' {
			edge = i + 1
		}
		cur := make([]int, 0)
		for j := i; j < edge; j++ {
			cur = append(cur, byte(num[j])-byte('0'))
			if m < 2 || check(l[m-2], l[m-1], cur) {
				l = append(l, cur)
				if dfs(j + 1) {
					return true
				}
				l = l[:len(l)-1]
			}
		}
		return false
	}

	return f(0)
}
