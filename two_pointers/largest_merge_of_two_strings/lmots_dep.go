package largest_merge_of_two_strings

func largestMerge4P(word1 string, word2 string) string {
	var il, ir, jl, jr int
	sb := make([]byte, 0)
	for {
		if ir >= len(word1) && jr >= len(word2) {
			sb = append(sb, word1[il:]...)
			sb = append(sb, word2[jl:]...)
			break
		} else if ir >= len(word1) {
			if il != ir {
				ir = il
				jr++
				continue
			}
			sb = append(sb, word2[jl:]...)
			break
		} else if jr >= len(word2) {
			if jl != jr {
				jr = jl
				ir++
				continue
			}
			sb = append(sb, word1[il:]...)
			break
		}

		// w1: ur w2: uru false: ururu true: uurru
		// 纯粹变成一步步打补丁，这是不对的
		if word1[ir] > word2[jr] {
			ir++
			sb = append(sb, word1[il:ir]...)
			il = ir
			jr = jl
		} else if word1[ir] < word2[jr] {
			jr++
			sb = append(sb, word2[jl:jr]...)
			jl = jr
			ir = il
		} else {
			if word1[il] > word2[jr] {
				sb = append(sb, word1[il])
				il++
			}
			if word2[jl] > word1[ir] {
				sb = append(sb, word2[jl])
				jl++
			}
			ir++
			jr++
		}
	}

	return string(sb)
}

// 			ic, jc := i+1, j+1
//			ec := word1[i]
//			for {
//				for ic < len(word1) && word1[ic] == ec {
//					ic++
//				}
//				if ic == len(word1) {
//					ic = len(word1) - 1
//				}
//				for jc < len(word2) && word2[jc] == ec {
//					jc++
//				}
//				if jc == len(word2) {
//					jc = len(word2) - 1
//				}
//				if word1[ic] > word2[jc] {
//					sb = append(sb, word1[i:ic+1]...)
//					i = ic + 1
//					break
//				} else if word1[ic] < word2[jc] {
//					sb = append(sb, word2[j:jc+1]...)
//					j = jc + 1
//					break
//				} else if ic == len(word1)-1 && jc == len(word2)-1 {
//					sb = append(sb, word1[i:]...)
//					sb = append(sb, word2[j:]...)
//					i = len(word1)
//					break
//				} else if ic == len(word1)-1 && jc != len(word2)-1 {
//					if word1[i] > word2[jc] {
//						sb = append(sb, word1[i:]...)
//						sb = append(sb, word2[j:]...)
//						// 如此多的条件判断明显是有问题的，弃之
//					}
//					i = len(word1)
//					break
//				}
//				ec = word1[ic]
//				ic++
//				jc++
//			}

func largestMergeGreedy(word1, word2 string) string {
	n, m := len(word1), len(word2)
	merge := make([]byte, 0, n+m)
	var i, j int
	for i < n || j < m {
		if i < n && word1[i:] > word2[j:] {
			merge = append(merge, word1[i])
			i++
		} else {
			merge = append(merge, word2[j])
			j++
		}
	}

	return string(merge)
}

func largestMergeSuffixArray(word1, word2 string) string {
	//m, n := len(word1), len(word2)

	return ""
}

func buildSuffixArray(text string) []int {
	order := sortChars(text)
	classfiy := computeCharClasses(text, order)
	n := len(text)
	for i := 1; i < n; i <<= 1 {
		order = sortDoubled(text, i, order, classfiy)
		classfiy = updateClasses(order, classfiy, i)
	}
	return order
}

func updateClasses(order []int, classfiy []int, l int) []int {
	n := len(order)
	newClassfiy := make([]int, n)
	newClassfiy[order[0]] = 0
	for i := 1; i < n; i++ {
		cur := order[i]
		prev := order[i-1]
		mid := cur + l
		midPrev := (prev + l) % n
		if classfiy[cur] != classfiy[prev] || classfiy[mid] != classfiy[midPrev] {
			newClassfiy[cur] = newClassfiy[prev] + 1
		} else {
			newClassfiy[cur] = newClassfiy[prev]
		}
	}

	return newClassfiy
}

func sortDoubled(text string, l int, order []int, classfiy []int) []int {
	n := len(text)
	counts := make([]int, n)
	newOrder := make([]int, n)
	for i := 0; i < n; i++ {
		counts[classfiy[i]]++
	}
	for i := 1; i < n; i++ {
		counts[i] += counts[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		start := (order[i] - l + n) % n
		cl := classfiy[start]
		counts[cl]--
		newOrder[counts[cl]] = start
	}
	return newOrder
}

// 重排，将相同字符的排名相同
func computeCharClasses(text string, order []int) []int {
	n := len(text)
	r := make([]int, n)
	r[order[0]] = 0
	for i := 1; i < n; i++ {
		if text[order[i]] != text[order[i-1]] {
			r[order[i]] = r[order[i-1]] + 1
		} else {
			r[order[i]] = r[order[i-1]]
		}
	}
	return r
}

// 返回text字符排序的索引数组
func sortChars(text string) []int {
	n := len(text)
	counts := make([]int, 128)
	orders := make([]int, n)
	for i := 0; i < n; i++ {
		counts[text[i]]++
	}
	for i := 1; i < 128; i++ {
		counts[i] += counts[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		counts[text[i]]--
		orders[counts[text[i]]] = i
	}
	return orders
}
