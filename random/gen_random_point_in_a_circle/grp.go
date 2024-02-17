package gen_random_point_in_a_circle

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius, xc, yc float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius: radius,
		xc:     x_center,
		yc:     y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	// 求的随机的半径
	r := math.Sqrt(rand.Float64())
	// 求的随机的sin、cos
	sin, cos := math.Sincos(2 * math.Pi * rand.Float64())
	x := this.xc + r*sin*this.radius
	y := this.yc + r*cos*this.radius
	return []float64{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
