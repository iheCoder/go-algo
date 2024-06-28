package asteroid_collision

func asteroidCollision(asteroids []int) []int {
	var i int
	var tempSum int
	for {
		for ; i < len(asteroids) && asteroids[i] > 0; i++ {
		}
		if i >= len(asteroids) {
			break
		}

		for j := i - 1; j >= 0; j-- {
			if asteroids[j] <= 0 {
				continue
			}
			tempSum = asteroids[i] + asteroids[j]
			if tempSum < 0 {
				asteroids[j] = 0
				continue
			}
			asteroids[i] = 0
			if tempSum == 0 {
				asteroids[j] = 0
			}
			break
		}
		i++
	}

	ans := make([]int, 0, len(asteroids))
	for _, asteroid := range asteroids {
		if asteroid == 0 {
			continue
		}
		ans = append(ans, asteroid)
	}
	return ans
}
