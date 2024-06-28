package reorder_data_in_log_files

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		s, t := logs[i], logs[j]
		x := strings.SplitN(s, " ", 2)[1]
		y := strings.SplitN(t, " ", 2)[1]

		isSDigit := unicode.IsDigit(rune(x[0]))
		isTDigit := unicode.IsDigit(rune(y[0]))
		if isSDigit && isTDigit {
			return false
		}

		if !isSDigit && !isTDigit {
			return x < y || (x == y && s < t)
		}

		return !isSDigit
	})

	return logs
}
