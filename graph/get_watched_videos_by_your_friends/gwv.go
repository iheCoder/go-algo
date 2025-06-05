package get_watched_videos_by_your_friends

import "sort"

type PSI struct {
	first  string
	second int
}

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	n := len(friends)
	used := make([]bool, n)
	q := []int{id}
	used[id] = true

	for i := 1; i <= level; i++ {
		span := len(q)
		for j := 0; j < span; j++ {
			u := q[0]
			q = q[1:]
			for _, v := range friends[u] {
				if !used[v] {
					q = append(q, v)
					used[v] = true
				}
			}
		}
	}

	freq := make(map[string]int)
	for _, u := range q {
		for _, v := range watchedVideos[u] {
			freq[v]++
		}
	}

	videos := make([]PSI, 0, len(freq))
	for v, c := range freq {
		videos = append(videos, PSI{
			first:  v,
			second: c,
		})
	}
	sort.Slice(videos, func(i, j int) bool {
		if videos[i].second == videos[j].second {
			return videos[i].first < videos[j].first
		}

		return videos[i].second < videos[j].second
	})

	ans := make([]string, 0, len(videos))
	for _, video := range videos {
		ans = append(ans, video.first)
	}

	return ans
}

//func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
//	stk := map[int]struct{}{
//		id: {},
//	}
//	var l int
//	var bfs func()
//	bfs = func() {
//		if l == level {
//			return
//		}
//
//		tmp := make(map[int]struct{})
//		for i, _ := range stk {
//			for _, fr := range friends[i] {
//				tmp[fr] = struct{}{}
//			}
//		}
//		stk = tmp
//		l++
//		bfs()
//	}
//	bfs()
//
//	ans := make([]string, 0)
//	exists := make(map[string]struct{})
//	for i := range stk {
//		for _, v := range watchedVideos[i] {
//			if _, e := exists[v]; e {
//				continue
//			}
//
//			ans = append(ans, v)
//			exists[v] = struct{}{}
//		}
//	}
//
//	type pair struct {
//		v     string
//		count int
//	}
//	m := make(map[string]*pair)
//	for _, video := range watchedVideos {
//		for _, v := range video {
//			if m[v] == nil {
//				m[v] = &pair{
//					v:     v,
//					count: 0,
//				}
//			}
//			m[v].count++
//		}
//	}
//
//	sort.Slice(ans, func(i, j int) bool {
//		return m[ans[j]].count < m[ans[j]].count || (m[ans[j]].count == m[ans[j]].count && ans[i] < ans[j])
//	})
//
//	return ans
//}
