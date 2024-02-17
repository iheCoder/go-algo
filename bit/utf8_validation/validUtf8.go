package utf8_validation

const (
	bitNumMask = 0x80
)

func validUtf8(data []int) bool {
	for i := 0; i < len(data); i++ {
		if data[i]&bitNumMask != 0 {
			c := 0
			x := data[i]
			for (x>>(7-c))&1 > 0 {
				c++
				if c > 4 {
					return false
				}
			}
			c--
			if c == 0 {
				return false
			}
			for c > 0 {
				i++
				c--
				if i >= len(data) {
					return false
				}
				if data[i]&bitNumMask == 0 {
					return false
				}
			}
		}
	}
	return true
}
