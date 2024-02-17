package num_of_steps_to_reduce_a_num_to_zero

func numberOfSteps(num int) int {
	var steps int
	for num > 0 {
		if num&1 > 0 {
			num--
		} else {
			num >>= 1
		}
		steps++
	}
	return steps
}
