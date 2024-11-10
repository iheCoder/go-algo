package largest_values_from_labels

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	type pair struct {
		l, val int
	}
	n := len(values)
	ps := make([]pair, 0, n)
	for i := 0; i < n; i++ {
		ps = append(ps, pair{
			l:   labels[i],
			val: values[i],
		})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].val > ps[j].val
	})

	var sum int
	lun := make(map[int]int)

	var i int
	for numWanted > 0 && i < n {
		if lun[ps[i].l] < useLimit {
			lun[ps[i].l]++
			numWanted--
			sum += ps[i].val
		}
		i++
	}

	return sum
}
