package short_encoding_of_words

import "testing"

func TestSEW(t *testing.T) {
	ws := []string{"time", "me", "bell"}
	t.Log(minimumLengthEncoding(ws))
}
