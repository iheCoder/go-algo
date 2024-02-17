package triangle

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestGenTriangles(t *testing.T) {
	n := 200
	m := make([][]string, n)
	r := make([]string, 0, n)
	for i := 0; i < n; i++ {
		m[i] = make([]string, 0)
		for j := 0; j <= i; j++ {
			x := rand.Intn(1000)
			b := rand.Int()%2 == 1
			if b {
				x = -x
			}
			m[i] = append(m[i], strconv.Itoa(x))
		}

		r = append(r, "["+strings.Join(m[i], ",")+"]")
	}

	t.Log(strings.Join(r, ","))
}
