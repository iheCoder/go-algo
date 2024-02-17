package reconstruct_original_digits_from_english

import (
	"math/rand"
	"strings"
	"testing"
)

var words = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func TestGenRandomWordDigit(t *testing.T) {
	n := rand.Intn(1000)
	l := len(words)
	bs := &strings.Builder{}
	for i := 0; i < n; i++ {
		x := rand.Intn(l)
		bs.WriteString(words[x])
	}

	bbs := []byte(bs.String())
	for i := 0; i < len(bbs); i++ {
		x := rand.Intn(len(bbs))
		bbs[i], bbs[x] = bbs[x], bbs[i]
	}
	t.Log(string(bbs))
}
