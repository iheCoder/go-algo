package tweet_counts_per_frequency

import "sort"

type TweetCounts struct {
	users map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{users: make(map[string][]int)}
}

func (tc *TweetCounts) RecordTweet(tweetName string, time int) {
	tc.users[tweetName] = append(tc.users[tweetName], time)
}

func (tc *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	var l int
	switch freq {
	case "minute":
		l = 60
	case "hour":
		l = 3600
	case "day":
		l = 86400
	}

	bucketCount := (endTime-startTime)/l + 1
	ans := make([]int, bucketCount)
	times, exists := tc.users[tweetName]
	if !exists {
		return ans
	}

	sort.Ints(times)

	left := sort.Search(len(times), func(i int) bool {
		return times[i] >= startTime
	})
	right := sort.Search(len(times), func(i int) bool {
		return times[i] > endTime
	})

	for i := left; i < right; i++ {
		idx := (times[i] - startTime) / l
		ans[idx]++
	}

	return ans
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
