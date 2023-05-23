package find_the_index_of_the_first_occurrence_in_a_string

import "strings"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	} else if len(needle) == len(haystack) {
		if needle == haystack {
			return 0
		}
		return -1
	}

	return strings.Index(haystack, needle)
}
