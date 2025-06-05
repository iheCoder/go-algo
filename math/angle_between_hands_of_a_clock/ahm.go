package angle_between_hands_of_a_clock

import "math"

func angleClock(hour int, minutes int) float64 {
	minAng, hourAng := float64(6), float64(30)

	hd, md := (float64(hour)+float64(minutes)/60)*hourAng, float64(minutes)*minAng
	d := math.Abs(hd - md)

	return math.Min(d, 360-d)
}
