package bitwise_ors_of_subarrays

func subarrayBitwiseORs(arr []int) int {
	set := make(map[int]bool)
	cur := map[int]bool{
		0: true,
	}

	for _, a := range arr {
		tmp := make(map[int]bool)
		for b, _ := range cur {
			tmp[b|a] = true
		}
		tmp[a] = true

		cur = tmp
		for c, _ := range cur {
			set[c] = true
		}
	}

	return len(set)
}
