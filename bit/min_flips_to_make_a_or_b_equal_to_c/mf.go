package min_flips_to_make_a_or_b_equal_to_c

func minFlips(a int, b int, c int) int {
	var r int
	for i := 0; i < 32; i++ {
		if c>>i&1 > 0 {
			if ((a>>i)|(b>>i))&1 == 0 {
				r++
			}
		} else {
			if ((a>>i)|(b>>i))&1 > 0 {
				if ((a>>i)^(b>>i))&1 > 0 {
					r++
				} else {
					r += 2
				}
			}
		}
	}
	return r
}
