package iterator_for_combination

type CombinationIterator struct {
	pos      []int
	s        string
	finished bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	pos := make([]int, combinationLength)
	for i := 0; i < combinationLength; i++ {
		pos[i] = i
	}

	return CombinationIterator{
		pos:      pos,
		s:        characters,
		finished: false,
	}
}

func (ci *CombinationIterator) Next() string {
	var ans []byte
	for _, p := range ci.pos {
		ans = append(ans, ci.s[p])
	}

	i := -1
	for j := len(ci.pos) - 1; j >= 0; j-- {
		if ci.pos[j] != len(ci.s)-len(ci.pos)+j {
			i = j
			break
		}
	}

	if i == -1 {
		ci.finished = true
	} else {
		ci.pos[i]++
		for j := i + 1; j < len(ci.pos); j++ {
			ci.pos[j] = ci.pos[j-1] + 1
		}
	}

	return string(ans)
}

func (ci *CombinationIterator) HasNext() bool {
	return !ci.finished
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
