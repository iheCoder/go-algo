package simplify_path

import "testing"

func TestSP(t *testing.T) {
	path := "/home/"
	//t.Log(simplifyPath(path))
	path = "/a/../../b/../c//./////././////"
	t.Log(simplifyPath(path))
}
