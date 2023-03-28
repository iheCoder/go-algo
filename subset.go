package algo

import "math/rand"

func DeterministicSubset(backends []string, clientID, subsetSize int) []string {
	subsetCount := len(backends) / subsetSize
	round := clientID / subsetCount
	r := rand.New(rand.NewSource(int64(round)))
	r.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})

	subsetID := clientID % subsetSize
	start := subsetID * subsetSize
	return backends[start : start+subsetSize]
}
