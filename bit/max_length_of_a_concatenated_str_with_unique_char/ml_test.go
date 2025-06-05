package max_length_of_a_concatenated_str_with_unique_char

import "testing"

func TestML(t *testing.T) {
	arr := []string{"cga", "r", "act", "ers"}
	//t.Log(maxLength(arr))
	arr = []string{"un", "uabcd", "abhifjhreniu"}
	t.Log(maxLength(arr))
}
