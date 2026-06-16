package stickers_to_spell_word

func minStickers(stickers []string, target string) int {
	n := len(target)
	fullMask := (1 << n) - 1

	queue := []int{0}

	visited := make([]bool, 1<<n)
	visited[0] = true

	var step int
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			mask := queue[0]
			queue = queue[1:]

			if mask == fullMask {
				return step
			}

			for _, sticker := range stickers {
				nextMask := mask
				cnts := [26]int{}
				for j := 0; j < len(sticker); j++ {
					cnts[int(sticker[j]-'a')]++
				}

				for j := 0; j < len(target); j++ {
					if nextMask&(1<<j) != 0 {
						continue
					}

					if cnts[target[j]-'a'] > 0 {
						cnts[target[j]-'a']--
						nextMask |= 1 << j
					}
				}

				if !visited[nextMask] {
					visited[nextMask] = true
					queue = append(queue, nextMask)
				}
			}
		}

		step++
	}

	return -1
}
