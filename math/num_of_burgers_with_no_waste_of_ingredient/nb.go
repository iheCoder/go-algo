package num_of_burgers_with_no_waste_of_ingredient

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 == 1 {
		return nil
	}

	x := tomatoSlices/2 - cheeseSlices
	y := cheeseSlices - x
	if x < 0 || y < 0 {
		return nil
	}
	return []int{x, y}
}
