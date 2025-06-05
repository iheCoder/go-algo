package rm_sub_folders_from_the_fs

import "strings"

type tries struct {
	ref      int
	children map[string]*tries
}

func newTries() *tries {
	return &tries{
		ref:      -1,
		children: make(map[string]*tries),
	}
}

func removeSubfolders(folder []string) []string {
	root := newTries()

	for i, path := range folder {
		ps := strings.Split(path, "/")
		cur := root
		for _, p := range ps {
			if cur.children[p] == nil {
				cur.children[p] = newTries()
			}
			cur = cur.children[p]
		}

		cur.ref = i
	}

	ans := make([]string, 0)
	var dfs func(ps []string, cur *tries)
	dfs = func(ps []string, cur *tries) {
		if cur.ref != -1 {
			ans = append(ans, ps[cur.ref])
			return
		}

		for _, t := range cur.children {
			dfs(ps, t)
		}
	}
	dfs(folder, root)
	return ans
}
