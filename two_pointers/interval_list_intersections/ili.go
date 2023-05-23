package interval_list_intersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}

	r := make([][]int, 0)
	var is, ie, js, je int
	var i, j int
	var start int
	for i < len(firstList)<<1 && j < len(secondList)<<1 {
		is, js = firstList[i>>1][i&1], secondList[j>>1][j&1]
		ie, je = firstList[(i+1)>>1][(i+1)&1], secondList[(j+1)>>1][(j+1)&1]

		if is <= js {
			if ie < js {
				i += 2
				continue
			}
			start = is
			if js > start {
				start = js
			}
			if ie >= je {
				r = append(r, []int{start, je})
				j += 2
			} else {
				r = append(r, []int{start, ie})
				i += 2
			}
		} else {
			if is > je {
				j += 2
				continue
			}
			start = js
			if is > start {
				start = is
			}
			if ie >= je {
				r = append(r, []int{start, je})
				j += 2
			} else {
				r = append(r, []int{start, ie})
				i += 2
			}
		}
	}

	return r
}
