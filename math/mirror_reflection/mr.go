package mirror_reflection

// g是最大公因数，最小公倍数s等于p*q/g
// 则在y方向最终走的圈数为s/p=(p*q)/(p*g)=q/g，因为s是恰好使得kq=p的最小kq
// 在x方向最终走的圈数为
func mirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p /= g
	p %= 2
	q /= g
	q %= 2

	if p == 1 && q == 1 {
		return 1
	}
	if p == 1 {
		return 0
	}
	return 2
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}
