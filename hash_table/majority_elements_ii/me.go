package majority_elements_ii

func majorityElement(nums []int) []int {
	m := make(map[int]int)
	n := len(nums)
	t := n / 3

	var r []int
	for _, num := range nums {
		m[num]++
		if m[num] == t+1 {
			r = append(r, num)
			if len(r) == 2 {
				return r
			}
		}
	}

	return r
}

func majorityElementMol(nums []int) []int {
	var e1, e2 int
	var v1, v2 int

	for _, num := range nums {
		if e1 == num {
			v1++
		} else if e2 == num {
			v2++
		} else if v1 == 0 {
			e1 = num
			v1++
		} else if v2 == 0 {
			e2 = num
			v2++
		} else {
			v1--
			v2--
		}
	}

	var cn1, cn2 int
	for _, num := range nums {
		if num == e1 {
			cn1++
		} else if num == e2 {
			cn2++
		}
	}

	t := len(nums) / 3
	var r []int
	if cn1 > t {
		r = append(r, e1)
	}
	if cn2 > t {
		r = append(r, e2)
	}
	return r
}
