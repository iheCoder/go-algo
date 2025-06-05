package filter_restaurants_by_vegan_friendly_price_and_distance

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	type resMetric struct {
		id     int
		rating int
	}

	rms := make([]*resMetric, 0, len(restaurants))
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] == 0 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}

		rms = append(rms, &resMetric{
			id:     r[0],
			rating: r[1],
		})
	}

	sort.Slice(rms, func(i, j int) bool {
		return rms[i].rating > rms[j].rating || (rms[i].rating == rms[j].rating && rms[i].id > rms[j].id)
	})

	ans := make([]int, 0, len(rms))
	for _, rm := range rms {
		ans = append(ans, rm.id)
	}

	return ans
}
