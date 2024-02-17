package validate_ip_addr

import (
	"strconv"
	"strings"
)

const (
	neither = "Neither"
	ipv4    = "IPv4"
	ipv6    = "IPv6"
)

func validIPAddress(queryIP string) string {
	n := len(queryIP)
	if n < 7 || n > 39 {
		return neither
	}

	if strings.Contains(queryIP, ".") {
		ips := strings.Split(queryIP, ".")
		if len(ips) != 4 {
			return neither
		}
		for _, ip := range ips {
			if len(ip) == 0 {
				return neither
			}
			if ip[0] == '0' && len(ip) > 1 {
				return neither
			}
			num, err := strconv.Atoi(ip)
			if err != nil {
				return neither
			}
			if num < 0 || num > 255 {
				return neither
			}
		}
		return ipv4
	} else if strings.Contains(queryIP, ":") {
		if strings.Contains(queryIP, ".") {
			return neither
		}
		ips := strings.Split(queryIP, ":")
		if len(ips) != 8 {
			return neither
		}

		for _, ip := range ips {
			if len(ip) == 0 || len(ip) > 4 {
				return neither
			}
			for i := 0; i < len(ip); i++ {
				if (ip[i] > 'f' && ip[i] <= 'z') || (ip[i] > 'F' && ip[i] <= 'Z') {
					return neither
				}
			}
		}
		return ipv6
	}

	return neither
}
