package solve_the_equation

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	infiniteSolution = "Infinite solutions"
	noSolution       = "No solution"
)

func solveEquation(equation string) string {
	equation = strings.TrimSpace(equation)
	es := strings.Split(equation, "=")
	l, r := es[0], es[1]
	lx, ln := simple(l)
	rx, rn := simple(r)
	x := lx - rx
	n := rn - ln
	if x == 0 && n == 0 {
		return infiniteSolution
	}
	if x == 0 || n%x != 0 {
		return noSolution
	}

	return fmt.Sprintf("x=%d", n/x)
}

func simple(s string) (int, int) {
	var num, x int
	i := -1
	for j := 0; j < len(s); j++ {
		for j < len(s) && s[j] != '+' && s[j] != '-' {
			if s[j] == 'x' {
				break
			}
			j++
		}

		if j < len(s) && s[j] == 'x' {
			if j-i == 1 {
				if i < 0 || s[i] == '+' {
					x++
				} else {
					x--
				}
			} else {
				if i < 0 {
					i = 0
				}
				y, _ := strconv.Atoi(s[i:j])
				x += y
			}

			j++
		} else {
			if i < 0 {
				i = 0
			}
			y, _ := strconv.Atoi(s[i:j])
			num += y
		}
		i = j
	}
	return x, num
}
