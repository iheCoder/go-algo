package rank_teams_by_votes

import "sort"

func rankTeams(votes []string) string {
	ranking := make(map[byte][]int)
	n := len(votes[0])
	for i := 0; i < n; i++ {
		ranking[votes[0][i]] = make([]int, n)
	}

	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			ranking[vote[i]][i]++
		}
	}

	type rankItem struct {
		vid  byte
		rank []int
	}
	results := make([]rankItem, 0, n)
	for k, v := range ranking {
		results = append(results, rankItem{
			vid:  k,
			rank: v,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		for k := 0; k < n; k++ {
			if results[i].rank[k] != results[j].rank[k] {
				return results[i].rank[k] > results[j].rank[k]
			}
		}

		return results[i].vid < results[j].vid
	})

	ans := make([]byte, 0, len(results))
	for _, r := range results {
		ans = append(ans, r.vid)
	}
	return string(ans)
}

//func rankTeams(votes []string) string {
//	m, n := len(votes), len(votes[0])
//	ans := make([]byte, 0, n)
//	selected := make(map[byte]bool)
//
//	type item struct {
//		a     byte
//		rank byte
//	}
//
//	for i := 0; i < n; i++ {
//		items := make([]item, 26)
//		for j := 0; j < m; j++ {
//			a := votes[j][i]
//			if selected[a] {
//				continue
//			}
//
//			itm := &(items[a-'A'])
//			itm.rank++
//			itm.a = a
//		}
//
//		sort.Slice(items, func(i, j int) bool {
//			return items[i].rank > items[j].rank || (items[i].rank == items[j].rank && items[i].a > items[j].a)
//		})
//
//		for _, itm := range items {
//			if itm.rank <= 0 {
//				break
//			}
//			ans = append(ans, itm.a)
//			selected[itm.a] = true
//		}
//
//		if len(ans) >= n {
//			return string(ans)
//		}
//	}
//
//	return string(ans)
//}
