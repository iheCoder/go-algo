package subdomain_visit_count

import (
	"testing"
)

func TestSVC(t *testing.T) {
	cpds := []string{" 9001       discuss.leetcode.com"}
	t.Log(subdomainVisits(cpds))
}
