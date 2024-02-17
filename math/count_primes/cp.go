package count_primes

import "math"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	ip := make([]bool, n)
	for i := 2; i < n; i++ {
		ip[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if ip[i] {
			for j := i * i; j < n; j += i {
				ip[j] = false
			}
		}
	}

	var count int
	for i := 2; i < n; i++ {
		if ip[i] {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	max := int(math.Sqrt(float64(n)))
	for i := 5; i <= max; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
