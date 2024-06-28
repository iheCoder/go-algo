package find_eventual_safe_states

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestGenRanGraph(t *testing.T) {
	n := 10000
	gs := generateRandGraphs(n)
	var sb strings.Builder
	sb.WriteString("[")
	for _, g := range gs {
		si := strings.Builder{}
		si.WriteString("[")
		for j := 0; j < len(g); j++ {
			si.WriteString(strconv.Itoa(g[j]))
			if j < len(g)-1 {
				si.WriteString(",")
			}
		}
		si.WriteString("]")
		sb.WriteString(si.String())
		sb.WriteString(",")
	}
	ans := sb.String()
	ans = strings.TrimSuffix(ans, ",")
	t.Log(ans + "]")
}

func generateRandGraphs(n int) [][]int {
	var ans [][]int
	for i := 0; i < n; i++ {
		limit := rand.Intn(n)
		item := make([]int, 0, limit)
		for j := 0; j < limit; j++ {
			x := rand.Intn(n)
			item = append(item, x)
		}
		ans = append(ans, item)
	}
	return ans
}
