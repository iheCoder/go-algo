package complex_number_multiplication

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	xs, ys := strings.Split(num1, "+"), strings.Split(num2, "+")
	xr, _ := strconv.Atoi(xs[0])
	xi, _ := strconv.Atoi(xs[1][:len(xs[1])-1])
	yr, _ := strconv.Atoi(ys[0])
	yi, _ := strconv.Atoi(ys[1][:len(ys[1])-1])

	r := xr*yr - xi*yi
	i := xr*yi + xi*yr
	return fmt.Sprintf("%d+%di", r, i)
}
