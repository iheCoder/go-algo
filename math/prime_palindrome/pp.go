package prime_palindrome

func primePalindrome(n int) int {
	i := n
	for {
		if isPalindrome(i) {
			if isPrime(i) {
				return i
			}
		}

		i++

		if i > 10000000 && i < 100000000 {
			i = 100000000
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed int
	y := x
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return y == reversed
}
