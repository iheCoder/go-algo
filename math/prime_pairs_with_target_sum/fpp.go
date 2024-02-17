package prime_pairs_with_target_sum

var pm = make(map[int]bool)

const maxValue = 1000000

func init() {
	flags := make(map[int]bool)
	for i := 2; i <= maxValue; i++ {
		if flags[i] {
			continue
		}

		pm[i] = true
		for j := i * i; j <= maxValue; j += i {
			flags[j] = true
			if i%4 == 1 {
				flags[i+j] = false
			} else {
				flags[i*i+j] = false
			}
		}
	}
}

func findPrimePairs(n int) [][]int {
	i, j := 2, n-2
	var ans [][]int
	for i <= j {
		if isPrime(i) && isPrime(j) {
			ans = append(ans, []int{i, j})
		}
		i++
		j--
	}
	return ans
}

func isPrime(x int) bool {
	return pm[x]
}
