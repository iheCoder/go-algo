package divide_array_in_sets_of_k_consecutive_numbers

import "sort"

type dividor struct {
	count, num int
}

func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}

	sort.Ints(nums)

	ds := make([]dividor, 0, n)
	for _, num := range nums {
		if len(ds) > 0 && ds[len(ds)-1].num == num {
			ds[len(ds)-1].count++
		} else {
			ds = append(ds, dividor{
				count: 1,
				num:   num,
			})
		}
	}

	for len(ds) > 0 {
		start := ds[0].num
		ds[0].count--
		if len(ds) < k {
			return false
		}

		for i := 1; i < k && i < len(ds); i++ {
			if ds[i].num != start+1 || ds[i].count <= 0 {
				return false
			}
			ds[i].count--
			start = start + 1
		}

		for len(ds) > 0 && ds[0].count <= 0 {
			ds = ds[1:]
		}
	}

	return true
}
