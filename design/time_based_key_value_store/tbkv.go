package time_based_key_value_store

import "sort"

type pair struct {
	timestamp int
	val       string
}

type TimeMap struct {
	m map[string][]*pair
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]*pair)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.m[key] = append(t.m[key], &pair{
		timestamp: timestamp,
		val:       value,
	})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if x, ok := t.m[key]; ok {
		j := sort.Search(len(x), func(i int) bool {
			return x[i].timestamp >= timestamp
		})

		if j == 0 && x[j].timestamp > timestamp {
			return ""
		}

		if j == len(x) || x[j].timestamp > timestamp {
			return x[j-1].val
		}

		return x[j].val
	}

	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
