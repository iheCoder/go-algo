package power_of_four

var m = map[int]bool{
	1: true, 4: true, 16: true, 64: true, 256: true, 1024: true, 4096: true, 16384: true, 65536: true, 262144: true, 1048576: true, 4194304: true, 16777216: true, 67108864: true, 268435456: true, 1073741824: true,
}

func isPowerOfFour(n int) bool {
	return m[n]
}
