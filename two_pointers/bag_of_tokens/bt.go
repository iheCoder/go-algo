package bag_of_tokens

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	if len(tokens) == 0 {
		return 0
	}

	sort.Ints(tokens)

	//mem := make([]int, len(tokens))
	//for k := 0; k < len(tokens)-1; k++ {
	//	mem[k] = tokens[k] + tokens[k+1]
	//}
	//mem[len(tokens)-1] = tokens[len(tokens)-1]

	var score int
	i, j := 0, len(tokens)-1
	for i <= j {
		if power >= tokens[i] {
			power -= tokens[i]
			i++
			score++
		} else if score > 0 {
			if j-i <= 1 || power+tokens[j] <= tokens[i] {
				break
			}
			power += tokens[j]
			j--
			score--
		} else {
			break
		}
	}
	return score
}
