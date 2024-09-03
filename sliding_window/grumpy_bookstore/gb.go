package grumpy_bookstore

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var initSatisfied int
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			initSatisfied += customers[i]
		}
	}

	var winSatisfied int
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			winSatisfied += customers[i]
		}
	}

	mws := winSatisfied
	for i := minutes; i < n; i++ {
		if grumpy[i-minutes] == 1 {
			winSatisfied -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			winSatisfied += customers[i]
		}

		if mws < winSatisfied {
			mws = winSatisfied
		}
	}

	return mws + initSatisfied
}
