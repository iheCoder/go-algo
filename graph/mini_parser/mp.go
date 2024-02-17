package mini_parser

import "unicode"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	val   int
	child *NestedInteger
}

func (n NestedInteger) IsInteger() bool {
	return false
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {
}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func deserializeDep(s string) *NestedInteger {
	var i int
	var f func() *NestedInteger
	f = func() *NestedInteger {
		var num int
		var ln, rn int
		var neg bool
		ni := &NestedInteger{}
		for i < len(s) {
			switch s[i] {
			case '[':
				ln++
				i++
				ni.Add(*(f()))
			case ']':
				rn++
				if num != 0 {
					if neg {
						num = -num
					}
					ni.SetInteger(num)
					neg = false
				}
				if rn > ln {
					i--
					return ni
				}
			case ',':
				if neg {
					num = -num
				}
				ni.SetInteger(num)
				num = 0
				neg = false
			case '-':
				neg = true
			default:
				num = num*10 + convertByteToInt(s[i])
			}
			i++
		}
		if num != 0 {
			if neg {
				num = -num
			}
			ni.SetInteger(num)
		}
		return ni
	}
	return f()
}

func convertByteToInt(b byte) int {
	switch b {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return 0
}

func deserializeDFS(s string) *NestedInteger {
	var i int
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[i] == '[' {
			i++
			for s[i] != ']' {
				ni.Add(*dfs())
				if s[i] == ',' {
					i++
				}
			}
			i++
			return ni
		}

		neg := s[i] == '-'
		if neg {
			i++
		}
		var num int
		for ; i < len(s) && unicode.IsDigit(rune(s[i])); i++ {
			num = num*10 + int(s[i]-'0')
		}
		if neg {
			num = -num
		}
		ni.SetInteger(num)
		return ni
	}
	return dfs()
}

func deserialize(s string) *NestedInteger {
	var i int
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}

		// 1. 若为[，则说明出现list，一直递归直到出现]
		if s[i] == '[' {
			i++
			for s[i] != ']' {
				ni.Add(*dfs())
				if i >= len(s) {
					return ni
				}
				if s[i] == ',' {
					i++
				}
			}
			// 恰好将本轮递归消除]
			i++
			return ni
		}

		var neg bool
		var num int
		// 2. 遍历取得整数。如果碰到,或者]说明整数已经获取完毕
		// 那么如果递归返回，便得到了,的情况会进一步进行递归
		// 如果得到了]说明递归结束
		for ; i < len(s); i++ {
			if s[i] == '-' {
				neg = true
			} else if s[i] == ',' || s[i] == ']' {
				break
			} else {
				num = num*10 + int(s[i]-'0')
			}
		}

		if neg {
			num = -num
		}
		ni.SetInteger(num)
		return ni
	}

	return dfs()
}
