package unique_paths

import "testing"

// 从规律上来看似乎属于dp，但能不能用数学呢？
// 设x为更小的，y为更大的。那么ans=
func TestUp(t *testing.T) {
	m := 2
	n := 3
	for i := 0; i < 5; i++ {
		t.Log(uniquePathsRecur(m, n+i))
		t.Log(uniquePaths(m, n+i))
	}
	m = 3
	n = 3
	for i := 0; i < 5; i++ {
		t.Log(uniquePathsRecur(m+i, n))
		t.Log(uniquePaths(m, n+i))
	}
	m = 4
	n = 4
	for i := 0; i < 5; i++ {
		t.Log(uniquePathsRecur(m+i, n))
		t.Log(uniquePaths(m, n+i))
	}
}
