package flatten_nested_list_iterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 *
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	val   int
	l     []*NestedInteger
	isInt bool
}

func (n NestedInteger) IsInteger() bool {
	return n.isInt
}

func (n NestedInteger) GetInteger() int {
	return n.val
}

func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

func (n *NestedInteger) Add(elem NestedInteger) {
	n.l = append(n.l, &elem)
}

func (n NestedInteger) GetList() []*NestedInteger {
	return n.l
}

type NestedIterator struct {
	nl []*NestedInteger
	// li具有歧义，当前为int指向下一个，list指向当前
	li, ii int
	items  []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		nl: nestedList,
	}
}

func (ni *NestedIterator) Next() int {
	cur := ni.nl[ni.li]
	if cur.IsInteger() {
		ni.li++
		return cur.GetInteger()
	}

	x := ni.items[ni.ii]
	ni.ii++
	if ni.ii >= len(ni.items) {
		ni.li++
	}
	return x
}

func fullIterate(nis []*NestedInteger) []int {
	nums := make([]int, 0)
	for _, ni := range nis {
		if ni.IsInteger() {
			nums = append(nums, ni.GetInteger())
		} else {
			nums = append(nums, fullIterate(ni.GetList())...)
		}
	}
	return nums
}

func (ni *NestedIterator) HasNext() bool {
	if ni.li >= len(ni.nl) {
		return false
	}

	if ni.nl[ni.li].IsInteger() {
		return true
	}

	if ni.ii < len(ni.items) {
		return true
	}

	var items []int
	i := ni.li
	ni.ii = 0
	//if ni.li == 0 && len(ni.items) == 0 {
	//	i = -1
	//}
	for len(items) == 0 {
		if i >= len(ni.nl) {
			return false
		}
		next := ni.nl[i]
		if next.IsInteger() {
			ni.li = i
			return true
		}
		items = fullIterate(next.GetList())
		i++
	}
	ni.items = items
	ni.li = i - 1
	return true
}

// 	var items []int
//	for len(items) == 0 {
//		ni.nextJump++
//		if ni.nextJump >= len(ni.nl) {
//			return false
//		}
//		next := ni.nl[ni.nextJump]
//		if next.IsInteger() {
//			return true
//		}
//		items = fullIterate(next.GetList())
//	}
//	ni.nextItems = items
//	return true
