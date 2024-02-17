package teemo_attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	if duration == 0 {
		return 0
	}

	last := timeSeries[0]
	ans := duration
	var span int
	for i := 1; i < len(timeSeries); i++ {
		span = timeSeries[i] - last
		if span > duration {
			span = duration
		}
		ans += span
		last = timeSeries[i]
	}
	return ans
}
