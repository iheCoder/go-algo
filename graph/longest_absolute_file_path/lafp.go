package longest_absolute_file_path

import "strings"

func lengthLongestPath(input string) int {
	var ans int

	// 1. 建立深度和长度的映射
	dlm := map[int]int{-1: 0}

	// 2. 按换行符进行分离，当前深度长度等于上一个深度+当前长度-t字符的个数
	for _, line := range strings.Split(input, "\n") {
		d := strings.Count(line, "\t")
		dlm[d] = dlm[d-1] + len(line) - d
		// 3. 若包含.说明存在文件，那么求的最大文件路径长度
		if strings.Contains(line, ".") {
			// 加上d是补上/
			ans = maxInt(ans, dlm[d]+d)
		}
	}

	return ans
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
