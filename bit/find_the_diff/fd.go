package find_the_diff

func findTheDifference(s string, t string) byte {
	var r int
	for _, si := range s {
		r ^= int(si - 'a')
	}
	for _, ti := range t {
		r ^= int(ti - 'a')
	}
	return byte(r + 'a')
}
