package binary_watch

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	var ans []string
	for i := 0; i < 1024; i++ {
		h, m := i>>6, i&63
		if bits.OnesCount(uint(i)) == turnedOn && h < 12 && m < 60 {
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return ans
}
