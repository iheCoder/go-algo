package letter_tile_possibilities

func numTilePossibilities(tiles string) int {
	ts := make(map[string]int)
	for i := 0; i < len(tiles); i++ {
		ts[string(tiles[i])]++
	}

	ans := make(map[string]bool)
	var f func(s string)
	f = func(s string) {
		for c, b := range ts {
			if b > 0 {
				ts[c]--
				f(s + c)
				ts[c]++
			}
		}
		ans[s] = true
	}
	f("")

	return len(ans) - 1
}
