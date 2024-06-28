package len_of_longest_fib

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestGenIncreasingNumbers(t *testing.T) {
	n := 1000
	arr := genIncreasingNumbers(n)
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(arr[i]))
		sb.WriteString(",")
	}

	s := sb.String()
	t.Log(s[:len(s)-1])
}

func genIncreasingNumbers(n int) []int {
	arr := make([]int, n)
	arr[0] = 1
	last := 1
	for i := 1; i < n; i++ {
		x := rand.Intn(1000000)
		arr[i] = last + x
		last = arr[i]
		if last > 1000000000 {
			panic("too large")
		}
	}

	return arr
}
