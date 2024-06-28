package subdomain_visit_count

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		cpdomain = strings.ReplaceAll(cpdomain, " ", "")
		var i int
		for ; i < len(cpdomain) && cpdomain[i] >= '0' && cpdomain[i] <= '9'; i++ {
		}
		num, _ := strconv.Atoi(cpdomain[:i])
		d := cpdomain[i:]
		m[d] += num
		n := len(d)
		for j := n - 1; j >= 0; j-- {
			if d[j] == '.' {
				m[d[j+1:n]] += num
			}
		}
	}

	var ans []string
	for s, v := range m {
		ans = append(ans, fmt.Sprintf("%d %s", v, s))
	}
	return ans
}
