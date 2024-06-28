package reordered_power_of_2

import (
	"sort"
	"strconv"
)

func reorderedPowerOf2(n int) bool {
	power2 := make([]int, 0)
	for i := 0; 1<<i <= 1000000000; i++ {
		power2 = append(power2, 1<<i)
	}

	getKey := func(s []byte) string {
		var key []byte
		for i := 0; i < len(s); i++ {
			key = append(key, '_')
			key = append(key, s[i])
		}

		return string(key)
	}

	m := make(map[string]bool)
	for _, p2 := range power2 {
		s := []byte(strconv.Itoa(p2))
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		m[getKey(s)] = true
	}

	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	key := getKey(s)

	return m[key]
}
