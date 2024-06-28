package rm_comments

func removeComments(source []string) []string {
	var ans []string
	var pc bool
	var bs []byte
	for _, s := range source {
		var j int
		for j = 0; j < len(s); j++ {
			if pc {
				if j+1 < len(s) && s[j] == '*' && s[j+1] == '/' {
					j++
					pc = false
				}
			} else {
				if j+1 < len(s) && s[j] == '/' && s[j+1] == '/' {
					break
				} else if j+1 < len(s) && s[j] == '/' && s[j+1] == '*' {
					j++

					pc = true
				} else {
					bs = append(bs, s[j])
				}
			}
		}
		if !pc && len(bs) > 0 {
			ans = append(ans, string(bs))
			bs = []byte{}
		}
	}
	return ans
}
