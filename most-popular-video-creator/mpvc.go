package most_popular_video_creator

import "sort"

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	max_creators := make([]string, 0)
	creators_views := make(map[string]int)
	creator_max_views := make(map[string]map[string]int)

	max_creators = append(max_creators, creators[0])
	max_creator_views := views[0]
	creators_views[creators[0]] = views[0]
	m0 := make(map[string]int)
	m0[ids[0]] = views[0]
	creator_max_views[creators[0]] = m0
	n := len(creators)

	for i := 1; i < n; i++ {
		creators_views[creators[i]] += views[i]
		if creators_views[creators[i]] > max_creator_views {
			max_creators = []string{creators[i]}
			max_creator_views = creators_views[creators[i]]
		} else if creators_views[creators[i]] == max_creator_views {
			max_creators = append(max_creators, creators[i])
		}

		_, ok := creator_max_views[creators[i]]
		if !ok {
			creator_max_views[creators[i]] = map[string]int{ids[i]: views[i]}
		} else {
			for _, v := range creator_max_views[creators[i]] {
				if views[i] > v {
					creator_max_views[creators[i]] = map[string]int{ids[i]: views[i]}
				} else if views[i] == v {
					creator_max_views[creators[i]][ids[i]] = views[i]
				}
				break
			}
		}
	}

	max_creators_set := make(map[string]bool)
	for _, creator := range max_creators {
		max_creators_set[creator] = true
	}
	r := make([][]string, 0, len(max_creators_set))
	for c, _ := range max_creators_set {
		s := make([]string, 0, len(creator_max_views[c]))
		for id, _ := range creator_max_views[c] {
			s = append(s, id)
		}
		sort.Strings(s)
		r = append(r, []string{c, s[0]})
	}

	return r
}
