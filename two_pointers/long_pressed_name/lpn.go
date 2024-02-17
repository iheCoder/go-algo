package long_pressed_name

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}

	var i, j int
	var temp byte
	for {
		if i >= len(name) {
			temp = name[len(name)-1]
			for ; j < len(typed); j++ {
				if typed[j] != temp {
					return false
				}
			}
			return true
		}
		if j >= len(typed) {
			return false
		}
		if name[i] == typed[j] {
			j++
			i++
			continue
		}
		if i == 0 {
			return false
		}
		temp = name[i-1]
		if typed[j] == temp {
			for j < len(typed) && typed[j] == temp {
				j++
			}
			continue
		}
		return false
	}
}
