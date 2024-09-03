package next_greater_node_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}

	type pair struct {
		val, index int
	}

	ans := make([]int, 1)
	ps := []pair{
		{
			val:   head.Val,
			index: 0,
		},
	}
	p := head.Next
	var i int
	for p != nil {
		i++
		for len(ps) > 0 && p.Val > ps[len(ps)-1].val {
			ans[ps[len(ps)-1].index] = p.Val
			ps = ps[:len(ps)-1]
		}

		ps = append(ps, pair{
			val:   p.Val,
			index: i,
		})

		ans = append(ans, 0)
		p = p.Next
	}

	return ans
}
