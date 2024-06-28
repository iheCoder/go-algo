package fraction_add_and_sub

import (
	"fmt"
	"math"
	"strconv"
)

func fractionAddition(expression string) string {
	if len(expression) < 7 {
		return expression
	}
	numerators, denominators := splitFracAndSignal(expression)
	denLcd := getLeastCommonDivisor(denominators)
	var sum int
	for i := 0; i < len(numerators); i++ {
		numerators[i] *= denLcd / denominators[i]
		sum += numerators[i]
	}

	if sum == 0 {
		return "0/1"
	}

	if denLcd < 0 {
		denLcd, sum = -denLcd, -sum
	}

	lcd := gcd(sum, denLcd)
	if lcd < 0 {
		lcd = -lcd
	}
	return fmt.Sprintf("%d/%d", sum/lcd, denLcd/lcd)
}

func splitFracAndSignal(exp string) ([]int, []int) {
	numerators, denominator := make([]int, 0), make([]int, 0)
	var i int
	for j := 0; j < len(exp); j++ {
		if exp[j] == '/' {
			x, _ := strconv.Atoi(exp[i:j])
			numerators = append(numerators, x)
			i = j + 1
		} else if j != 0 && (exp[j] == '+' || exp[j] == '-') {
			x, _ := strconv.Atoi(exp[i:j])
			denominator = append(denominator, x)
			i = j
		}
	}
	x, _ := strconv.Atoi(exp[i:])
	denominator = append(denominator, x)
	return numerators, denominator
}

func getLeastCommonDivisor(arr []int) int {
	if len(arr) < 2 {
		return 0 // There is no concept of LCD for a single number
	}

	result := arr[0]

	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}

	return result
}

// Function to find the greatest common divisor (GCD) of two numbers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the least common multiple (LCM) of two numbers
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}
