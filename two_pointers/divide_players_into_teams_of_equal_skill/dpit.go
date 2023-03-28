package divide_players_into_teams_of_equal_skill

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	sum := skill[0] + skill[len(skill)-1]
	r := skill[0] * skill[len(skill)-1]
	i, j := 1, len(skill)-2
	for i < j {
		if skill[i]+skill[j] != sum {
			return -1
		}
		r += skill[i] * skill[j]
		i++
		j--
	}
	return int64(r)
}
