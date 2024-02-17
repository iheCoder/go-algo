package boats_to_save_people

import "sort"

func numRescueBoats(people []int, limit int) int {
	if len(people) == 1 {
		return 1
	}

	sort.Ints(people)

	var count int
	i, j := 0, len(people)-1
	for i < j {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}

		count++
	}
	if i == j {
		count++
	}

	return count
}
