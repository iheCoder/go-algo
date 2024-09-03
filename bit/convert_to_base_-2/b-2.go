package convert_to_base__2

import "strconv"

func baseNeg2(n int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}

	res := make([]byte, 0)
	for n != 0 {
		rm := n & 1
		res = append(res, '0'+byte(rm))
		n -= rm
		n /= -2
	}

	for i, m := 0, len(res); i < m/2; i++ {
		res[i], res[m-1-i] = res[m-1-i], res[i]
	}
	return string(res)
}
