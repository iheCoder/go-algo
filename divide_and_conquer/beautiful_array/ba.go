package beautiful_array

var memo = make(map[int][]int)

func beautifulArray(n int) []int {
	var f func(n int) []int
	f = func(n int) []int {
		if x, ok := memo[n]; ok {
			return x
		}

		ans := make([]int, n)
		if n == 1 {
			ans[0] = 1
		} else {
			var t int
			for _, x := range f((n + 1) / 2) {
				ans[t] = 2*x - 1
				t++
			}
			for _, x := range f(n / 2) {
				ans[t] = 2 * x
				t++
			}
		}

		memo[n] = ans
		return ans
	}

	return f(n)
}

//func beautifulArray(n int) []int {
//	m := make([]bool, n+1)
//	var ans []int
//	var f func(x int) bool
//	f = func(x int) bool {
//		prev := 10000
//		if len(ans) > 0 {
//			prev = ans[len(ans)-1]
//		}
//
//		m[x] = true
//		ans = append(ans, x)
//
//		if len(ans) == n {
//			return true
//		}
//
//		for i := 1; i <= n; i++ {
//			if !m[i] && prev+i != 2*x {
//				if f(i) {
//					return true
//				}
//			}
//		}
//
//		ans = ans[:len(ans)-1]
//		return false
//	}
//
//	rand.NewSource(time.Now().UnixMilli())
//	x := rand.Intn(n)
//	f(x + 1)
//	return ans
//}
