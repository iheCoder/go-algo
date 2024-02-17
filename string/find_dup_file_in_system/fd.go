package find_dup_file_in_system

import (
	"path/filepath"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	m := make(map[string][]string)

	for _, path := range paths {
		fs := strings.Split(path, " ")
		dir := fs[0]
		for i := 1; i < len(fs); i++ {
			var j int
			for j = 0; j < len(fs[i]); j++ {
				if fs[i][j] == '(' {
					break
				}
			}
			f, c := fs[i][:j], fs[i][j+1:len(fs[i])-1]
			m[c] = append(m[c], filepath.Join(dir, f))
		}
	}

	var ans [][]string
	for _, ps := range m {
		if len(ps) > 1 {
			ans = append(ans, ps)
		}
	}
	return ans
}
