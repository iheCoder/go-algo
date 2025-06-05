package group_the_people_given_the_group_size_they_belong_to

import "sort"

func groupThePeople(groupSizes []int) [][]int {
	type member struct {
		id, groupSize int
	}
	ms := make([]member, 0, len(groupSizes))
	for i, size := range groupSizes {
		ms = append(ms, member{
			id:        i,
			groupSize: size,
		})
	}
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].groupSize > ms[j].groupSize
	})

	var gc int
	var choose []int
	var ans [][]int
	for _, m := range ms {
		if gc <= 0 {
			gc = m.groupSize
		}

		choose = append(choose, m.id)
		gc--

		if gc <= 0 {
			ans = append(ans, choose)
			choose = []int{}
		}
	}

	return ans
}
