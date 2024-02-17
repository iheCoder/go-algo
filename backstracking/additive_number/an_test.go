package additive_number

import "testing"

func TestAN(t *testing.T) {
	b := '0'
	t.Logf("%d", b)
	type testData struct {
		num      string
		expected bool
	}

	tds := []testData{
		//{
		//	"112358",
		//	true,
		//},
		//{
		//	"199100199",
		//	true,
		//},
		//{
		//	"199991111",
		//	false,
		//},
		//{
		//	"11111111111111111111111111111111111",
		//	false,
		//},
		{
			"1023",
			false,
		},
	}

	for i, td := range tds {
		if isAdditiveNumber(td.num) != td.expected {
			t.Fatalf("fail expected %v index %d", td.expected, i)
		}
	}
}
