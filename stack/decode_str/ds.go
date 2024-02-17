package decode_str

func decodeString(s string) string {
	n := len(s)
	var i int
	var f func() []byte
	f = func() []byte {
		cur := make([]byte, 0)
		var num int
		for i < n && s[i] != ']' {
			if s[i] >= 'a' && s[i] <= 'z' {
				cur = append(cur, s[i])
			} else if s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
			} else if s[i] == '[' {
				i++
				x := f()
				for j := 0; j < num; j++ {
					cur = append(cur, x...)
				}
				num = 0
			}
			i++
		}
		return cur
	}
	return string(f())
}
