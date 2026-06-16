package longest_palindromic_subsequence

func longestPalindromeManacher(s string) string {
	if len(s) < 2 {
		return s
	}

	t := make([]byte, 0, len(s)*2+1)
	for i := 0; i < len(s); i++ {
		t = append(t, '#')
		t = append(t, s[i])
	}
	t = append(t, '#')

	// p[i] 表示：在转换后的字符串 t 中，以 i 为中心的最大回文半径。
	//
	// 注意：
	//   这里的半径不包含中心点自身。
	//
	// 举例：
	//   t = "#a#b#a#"
	//
	//   以 'b' 为中心时：
	//
	//       # a # b # a #
	//             ↑
	//             i
	//
	//   最长回文是整个 "#a#b#a#"。
	//   从中心 'b' 向左/右可以扩 3 格。
	//
	//   所以 p[i] = 3。
	p := make([]int, len(t))
	// center 表示当前已知“最靠右回文区间”的中心。
	// right 表示当前已知“最靠右回文区间”的右边界。
	//
	// 换句话说，在遍历过程中，我们维护这样一个区间：
	//
	//   [center - p[center], center + p[center]]
	//
	// 其中：
	//   right = center + p[center]
	//
	// 这个区间是目前所有已发现回文中，右边界最远的那个。
	center, right := 0, 0
	// maxLen 表示目前找到的最长回文半径。
	// maxCenter 表示这个最长回文在转换字符串 t 中的中心位置。
	maxL, maxC := 0, 0

	// 遍历转换后的字符串，把每个位置都当成回文中心。
	for i := 0; i < len(t); i++ {

		// 如果 i 落在当前最右回文区间内部：
		//
		//        left          center           right
		//         |--------------|---------------|
		//                        |
		//                      mirror     i
		//
		// 因为 [left, right] 本身是回文，
		// 所以 i 关于 center 的镜像点 mirror 已经被计算过。
		//
		// mirror 的公式：
		//   mirror = center - (i - center)
		//          = 2*center - i
		//
		// 也就是：
		//   i 在 center 右边多少距离，
		//   mirror 就在 center 左边多少距离。
		if i < right {
			mirror := 2*center - i

			// 这里是 Manacher 最关键的一步。
			//
			// 为什么 p[i] 可以先等于 p[mirror]？
			//
			// 因为 i 和 mirror 关于 center 对称，
			// 而 center 对应的大区间本身是回文，
			// 所以 mirror 周围已经验证过的回文部分，
			// 在 i 周围也应该对称成立。
			//
			// 但为什么不能直接写：
			//   p[i] = p[mirror]
			//
			// 因为 mirror 的回文半径可能会越过当前大回文区间的左边界。
			// 一旦越过左边界，右侧对应区域也会越过 right，
			// 那部分还没有被验证过，不能直接信任。
			//
			// 所以 i 能安全复用的最大范围，不能超过 right - i。
			//
			p[i] = minInt(p[mirror], right-i)
		}

		// 第三步：在已有半径基础上继续向外扩展。
		//
		// 如果 i 不在 right 内部，那么 p[i] 初始就是 0。
		// 如果 i 在 right 内部，那么 p[i] 已经复用了一部分镜像结果。
		//
		// 接下来只需要检查更外层字符是否还能匹配。
		//
		// 左边界：
		//   i - p[i] - 1
		//
		// 右边界：
		//   i + p[i] + 1
		//
		// 如果左右字符相等，说明回文半径还能继续扩大。
		for i-p[i]-1 >= 0 && i+p[i]+1 < len(t) && t[i-p[i]-1] == t[i+p[i]+1] {
			p[i]++
		}

		// 第四步：如果以 i 为中心的回文区间超过了当前 right，
		// 就更新当前最靠右回文区间。
		//
		// 新区间为：
		//   [i - p[i], i + p[i]]
		//
		// 如果 i + p[i] 更靠右，
		// 说明后续位置可以利用这个新区间做镜像复用。
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}

		// 第五步：记录全局最长回文。
		//
		// p[i] 是转换字符串中的半径。
		// 很巧的是，它也等于原字符串中的回文长度。
		//
		// 举例：
		//   原字符串 "abba" 长度是 4。
		//   转换后 "#a#b#b#a#"。
		//   以中间 '#' 为中心，p[i] = 4。
		//
		//   原字符串 "aba" 长度是 3。
		//   转换后 "#a#b#a#"。
		//   以 'b' 为中心，p[i] = 3。
		//
		// 所以 maxLen 可以直接作为原字符串里的最长回文长度。
		if p[i] > maxL {
			maxL = p[i]
			maxC = i
		}
	}

	start := (maxC - maxL) / 2
	return s[start : start+maxL]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
