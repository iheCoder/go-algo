package make_file_names_unique

import "fmt"

func getFolderNames(names []string) []string {
	//m := make(map[string]int)
	//for _, name := range names {
	//	m[name]++
	//}
	r := make([]string, len(names))
	rm := make(map[string]int)
	for i, name := range names {
		if _, exists := rm[name]; !exists {
			r[i] = name
			rm[name] = 1
		} else {
			k := fmt.Sprintf("%s(%d)", name, rm[name])
			_, ok := rm[k]
			for ok {
				rm[name]++
				k = fmt.Sprintf("%s(%d)", name, rm[name])
				_, ok = rm[k]
			}
			rm[k]++
			r[i] = k
		}
	}

	return r
}
