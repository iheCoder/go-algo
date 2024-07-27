package powerful_int

import "math"

var powerCache = make(map[int]int)

func powerfulIntegers(x int, y int, bound int) []int {
	var res int
	var i, j int
	m := make(map[int]bool)
	for {
		j = 0
		is := powInt(x, i)
		if is >= bound {
			break
		}
		res = 0
		for res <= bound {
			res = is + powInt(y, j)
			m[res] = true
			j++

			if y == 1 {
				break
			}
		}
		i++

		if x == 1 {
			break
		}
	}

	r := make([]int, 0, len(m))
	for z, _ := range m {
		if z < 2 || z > bound {
			continue
		}
		r = append(r, z)
	}
	return r
}

func powInt(x, i int) int {
	k := x*101 + i
	if p, exists := powerCache[k]; exists {
		return p
	}

	p := int(math.Pow(float64(x), float64(i)))
	powerCache[k] = p
	return p
}
