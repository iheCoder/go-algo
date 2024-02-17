package design_twitter

// 一种方案用户最新推文插入排序，然后忽略或者去掉不在关注的。那发表的时候岂不是还要找到所有关注自己的，然后去更新吗？
// 另一种找到所有自己关注的，然后得到前10个。时间复杂度有些高啊！
// 排序总体的，然后一个个取看是不是自己关注的
type Twitter struct {
	// 关注关系
	fr map[int]map[int]bool
	// 发表推文
	ts map[int][]*tweet
	// 发表时间id
	tid int
}

type tweet struct {
	id  int
	uid int
	tid int
}

func Constructor() Twitter {
	return Twitter{
		fr:  make(map[int]map[int]bool),
		ts:  make(map[int][]*tweet),
		tid: 0,
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	// 更新本用户推文
	nt := &tweet{
		id:  tweetId,
		uid: userId,
		tid: t.tid,
	}
	t.tid++
	t.ts[userId] = append(t.ts[userId], nt)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	// 若用户不存在，直接返回
	if len(t.ts[userId]) == 0 && len(t.fr[userId]) == 0 {
		return make([]int, 0)
	}

	// 找到所有目标推文
	// 构建长度为10的最新推文队列
	th := make(tweetHeap, 0, 10)
	n := len(t.ts[userId])
	fps := map[int]int{
		userId: n - 1,
	}
	if n > 0 {
		th.insert(t.ts[userId][n-1])
	}

	for fid, _ := range t.fr[userId] {
		n = len(t.ts[fid])
		if n <= 0 {
			continue
		}
		fps[fid] = n - 1
		th.insert(t.ts[fid][n-1])
	}

	// 取最晚推文，并将该最晚推文用户的下一个推文插入队列
	var ans []int
	var i int
	var x *tweet
	for len(ans) < 10 && len(th) > 0 {
		x = th.pop()
		ans = append(ans, x.id)
		fps[x.uid]--
		i = fps[x.uid]
		if i >= 0 {
			th.insert(t.ts[x.uid][i])
		}
	}
	return ans
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if t.fr[followerId] == nil {
		t.fr[followerId] = make(map[int]bool)
	}
	t.fr[followerId][followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	delete(t.fr[followerId], followeeId)
}

type tweetHeap []*tweet

func (h *tweetHeap) insert(t *tweet) {
	n := len(*h)
	*h = append(*h, t)
	if n == 0 {
		return
	}

	i, j := n-1, n
	for i >= 0 && (*h)[i].tid < t.tid {
		h.Swap(i, j)
		i--
		j--
	}
	if len(*h) > 10 {
		*h = (*h)[:n]
	}
}

func (h *tweetHeap) pop() *tweet {
	x := (*h)[0]
	(*h) = (*h)[1:]
	return x
}

func (h *tweetHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

//func (h tweetHeap) Len() int {
//	return len(h)
//}
//
//func (h tweetHeap) Less(i, j int) bool {
//	return h[i].tid > h[j].tid
//}
//

//
//func (h *tweetHeap) Push(x interface{}) {
//	*h = append(*h, x.(*tweet))
//}
//
//func (h *tweetHeap) Pop() interface{} {
//	n := len(*h)
//	x := (*h)[n-1]
//	*h = (*h)[:n-1]
//	return x
//}

// type Twitter struct {
//	// 用户最新推文
//	uth map[int]*tweetHeap
//	// 关注关系
//	fr map[int]map[int]bool
//	// 发表推文
//	ts map[int][]*tweet
//	// 发表时间id
//	tid int
//}
//
//type tweet struct {
//	id  int
//	uid int
//	tid int
//}
//
//func Constructor() Twitter {
//	return Twitter{
//		uth: make(map[int]*tweetHeap),
//		fr:  make(map[int]map[int]bool),
//		ts:  make(map[int][]*tweet),
//		tid: 0,
//	}
//}
//
//func (t *Twitter) PostTweet(userId int, tweetId int) {
//	// 更新本用户推文
//	nt := &tweet{
//		id:  tweetId,
//		uid: userId,
//		tid: t.tid,
//	}
//	t.tid++
//	t.ts[userId] = append(t.ts[userId], nt)
//
//	//
//}
