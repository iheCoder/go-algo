package alert_using_same_key_card_three_or_more_times_in_a_one_hour_period

import "sort"

func alertNames(keyName []string, keyTime []string) []string {
	timesMap := make(map[string][]int)
	n := len(keyName)
	for i := 0; i < n; i++ {
		kt := keyTime[i]
		hour := int(kt[0]-'0')*10 + int(kt[1]-'0')
		mins := int(kt[3]-'0')*10 + int(kt[4]-'0')
		tm := hour*60 + mins
		if timesMap[keyName[i]] == nil {
			timesMap[keyName[i]] = []int{tm}
		} else {
			timesMap[keyName[i]] = append(timesMap[keyName[i]], tm)
		}
	}

	for name, times := range timesMap {
		sort.Ints(times)
		timesMap[name] = times
	}

	r := make([]string, 0)
	for name, times := range timesMap {
		if len(times) < 3 {
			continue
		}
		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				r = append(r, name)
				break
			}
		}
	}

	sort.Strings(r)
	return r
}
