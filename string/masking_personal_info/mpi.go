package masking_personal_info

import (
	"fmt"
	"strings"
	"unicode"
)

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		ss := strings.Split(s, "@")
		domain := strings.ToLower(ss[1])
		name := strings.ToLower(ss[0])
		return fmt.Sprintf("%s*****%s@%s", string(name[0]), string(name[len(name)-1]), domain)
	}

	digits := extractDigits(s)

	switch len(digits) {
	case 10:
		return fmt.Sprintf("***-***-%s", digits[6:])
	case 11:
		return fmt.Sprintf("+*-***-***-%s", digits[7:])
	case 12:
		return fmt.Sprintf("+**-***-***-%s", digits[8:])
	case 13:
		return fmt.Sprintf("+***-***-***-%s", digits[9:])
	}
	return s
}

func extractDigits(str string) string {
	var result string
	for _, char := range str {
		if unicode.IsDigit(char) {
			result += string(char)
		}
	}
	return result
}
