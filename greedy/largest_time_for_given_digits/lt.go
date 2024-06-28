package largest_time_for_given_digits

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(arr []int) string {
	sort.Ints(arr)
	n := len(arr)
	picked := make([]bool, n)

	// hour
	h1 := -1
	var h2, m1, m2 int
	for i := n - 1; i >= 0; i-- {
		if arr[i] > 2 {
			continue
		}
		h1 = arr[i]
		picked[i] = true

		isTwoH1 := h1 == 2
		h2 = -1
		for j := n - 1; j >= 0; j-- {
			if picked[j] {
				continue
			}
			if isTwoH1 && arr[j] > 3 {
				continue
			}
			h2 = arr[j]
			picked[j] = true

			m1 = -1
			for k := n - 1; k >= 0; k-- {
				if picked[k] {
					continue
				}

				if arr[k] > 5 {
					continue
				}

				m1 = arr[k]
				picked[k] = true

				m2 = -1
				for l := n - 1; l >= 0; l-- {
					if picked[l] {
						continue
					}
					m2 = arr[l]
					break
				}

				picked[k] = false
				if m2 < 0 {
					continue
				}

				goto success
			}
			picked[j] = false

			if m1 < 0 {
				continue
			}
		}

		picked[i] = false
		if h2 < 0 {
			continue
		}
	}

success:
	if h1 < 0 || h2 < 0 || m1 < 0 || m2 < 0 {
		return ""
	}

	return fmt.Sprintf("%d%d:%d%d", h1, h2, m1, m2)
}
