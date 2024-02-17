package num_of_segments_in_a_str

import (
	"math"
	"strings"
	"testing"
)

func TestSplitSpace(t *testing.T) {
	t.Log(math.MaxInt32)
	s := " hello  world "
	rs := strings.Split(s, " ")
	t.Log(rs)
}
