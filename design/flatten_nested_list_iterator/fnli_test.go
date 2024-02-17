package flatten_nested_list_iterator

import "testing"

func TestFNLI(t *testing.T) {
	ni := []*NestedInteger{
		{
			val:   1,
			isInt: true,
		},
		{
			l: []*NestedInteger{
				{
					val:   4,
					isInt: true,
				},
				{
					l: []*NestedInteger{
						{
							val:   6,
							isInt: true,
						},
					},
				},
			},
		},
	}
	solution := Constructor(ni)
	for solution.HasNext() {
		t.Log(solution.Next())
	}
}
