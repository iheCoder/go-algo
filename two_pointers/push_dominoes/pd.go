package push_dominoes

// 模拟
func pushDominoesAns(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		// 找到一段连续没有被推动的
		for j < n && s[j] == '.' {
			j++
		}

		// 若连续没有被推动两边的都是同一个方向的，那么中间的将会倒下同一个方向
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right {
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' { // 方向相对的时候就从两侧向中间倒
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[j] = 'L'
				i++
				k--
			}
		}

		left = right
		i = j + 1
	}

	return string(s)
}
