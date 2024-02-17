package fib_num

var m = make([]int, 31)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if m[n] != 0 {
		return m[n]
	}

	r := fib(n-1) + fib(n-2)
	m[n] = r
	return r
}
