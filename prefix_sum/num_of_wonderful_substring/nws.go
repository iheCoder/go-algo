package num_of_wonderful_substring

func wonderfulSubstrings(word string) int64 {
	masks := 0
	seen := make(map[int]int64)
	seen[0] = 1
	var ans int64

	for i := 0; i < len(word); i++ {
		masks ^= 1 << (word[i] - 'a')
		if c, ok := seen[masks]; ok {
			ans += c
		}
		for j := 0; j < 10; j++ {
			if c, ok := seen[masks^(1<<j)]; ok {
				ans += c
			}
		}
		seen[masks]++
	}

	return ans
}
