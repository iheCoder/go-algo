package string_compress

import "strconv"

func compress(chars []byte) int {
	if len(chars) <= 1 {
		return len(chars)
	}

	s := make([]byte, 0)
	rc := chars[0]
	rcount := 1
	// 我本来还想搞虚拟尾节点，但显然并不对，因为append很有可能会改变chars指向的底层数组
	//chars = append(chars, chars[len(chars)-1]+1)
	for i := 1; i < len(chars)+1; i++ {
		if i != len(chars) && rc == chars[i] {
			rcount++
			continue
		}
		s = append(s, rc)
		if rcount != 1 {
			rcb := []byte(strconv.Itoa(rcount))
			s = append(s, rcb...)
		}
		if i != len(chars) {
			rcount = 1
			rc = chars[i]
		}
	}

	copy(chars, s)
	return len(s)
}
