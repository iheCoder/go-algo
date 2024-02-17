package gas_stations

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		gasSum, costSum, cnt := 0, 0, 0
		for cnt < len(gas) {
			index := (cnt + i) % len(gas)
			gasSum += gas[index]
			costSum += cost[index]
			if gasSum-costSum < 0 {
				break
			}
			cnt++
		}
		if cnt == len(gas) {
			return i
		} else {
			i = i + cnt + 1
		}
	}
	return -1
}

// 最后还是超出时间了
func canCompleteCircuitForce(gas []int, cost []int) int {
	sums := make([]int, len(gas))
	var diff int
	for i := 0; i < len(gas); i++ {
		sums[i] = gas[i] - cost[i]
		diff += sums[i]
	}
	if diff < 0 {
		return -1
	}
	if diff == 0 && len(sums) == 1 {
		return 0
	}

	var g, index int
	var pass bool
	for i := 0; i < len(sums); i++ {
		if sums[i] <= 0 {
			continue
		}
		g = 0
		pass = true
		for j := 0; j < len(sums); j++ {
			index = (i + j) % len(sums)
			g += sums[index]
			if g < 0 {
				pass = false
				break
			}
		}
		if pass {
			return i
		}
	}
	return -1
}
