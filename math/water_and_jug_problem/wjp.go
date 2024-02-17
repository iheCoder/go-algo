package water_and_jug_problem

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}

	if jug1Capacity == 0 || jug2Capacity == 0 {
		return targetCapacity == 0 || jug1Capacity+jug2Capacity == targetCapacity
	}

	return targetCapacity%gcd(jug1Capacity, jug2Capacity) == 0
}

func gcd(x, y int) int {
	rem := x % y
	for rem != 0 {
		x = y
		y = rem
		rem = x % y
	}
	return y
}
