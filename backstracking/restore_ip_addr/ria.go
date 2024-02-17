package restore_ip_addr

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 {
		return nil
	}

	var ans []string
	var f func(i, rem int, ip string)
	f = func(i, rem int, ip string) {
		if i >= len(s) {
			if rem == 0 {
				ans = append(ans, ip[1:])
			}
			return
		}
		if rem <= 0 {
			return
		}
		if s[i] == '0' {
			f(i+1, rem-1, ip+".0")
			return
		}
		if rem == 1 {
			x, _ := strconv.Atoi(s[i:])
			if x > 255 {
				return
			}
			ans = append(ans, ip[1:]+"."+s[i:])
			return
		}
		for j := i + 1; j <= len(s)-rem+1 && j <= i+3; j++ {
			x, _ := strconv.Atoi(s[i:j])
			if x > 255 {
				continue
			}
			f(j, rem-1, ip+"."+s[i:j])
		}
	}
	f(0, 4, "")
	return ans
}
