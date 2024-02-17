package simplify_path

import "strings"

func simplifyPath(path string) string {
	// 1. 去除多余/
	isSlash := true
	sp := make([]byte, 0, len(path))
	for i := 0; i < len(path); i++ {
		if path[i] != '/' || (!isSlash && i < len(path)-1) {
			sp = append(sp, path[i])
		}
		isSlash = path[i] == '/'
	}
	if len(sp) > 0 && sp[len(sp)-1] == '/' {
		sp = sp[:len(sp)-1]
	}
	path = string(sp)

	// 2. 模拟简化
	stk := make([]string, 0)
	sps := strings.Split(path, "/")
	for _, s := range sps {
		switch s {
		case "..":
			if len(stk) == 0 {
				continue
			}
			stk = stk[:len(stk)-1]
		case ".":
		default:
			stk = append(stk, s)
		}
	}

	// 3. 组合
	r := "/"
	for _, s := range stk {
		r += s + "/"
	}
	if len(r) > 1 {
		r = r[:len(r)-1]
	}
	return r
}
