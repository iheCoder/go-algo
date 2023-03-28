package make_file_names_unique

import (
	"reflect"
	"testing"
)

type testData struct {
	names    []string
	expected []string
}

func TestNormal(t *testing.T) {
	tds := []testData{
		//{
		//	names:    []string{"pes", "fifa", "gta", "pes(2019)"},
		//	expected: []string{"pes", "fifa", "gta", "pes(2019)"},
		//},
		//{
		//	names:    []string{"gta", "gta(1)", "gta", "avalon"},
		//	expected: []string{"gta", "gta(1)", "gta(2)", "avalon"},
		//},
		//{
		//	names:    []string{"wano", "wano", "wano", "wano"},
		//	expected: []string{"wano", "wano(1)", "wano(2)", "wano(3)"},
		//},
		//{
		//	names:    []string{"kaido", "kaido(1)", "kaido", "kaido(1)"},
		//	expected: []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"},
		//},
		{
			names:    []string{"kaido", "kaido(1)", "kaido", "kaido(1)", "kaido(2)"},
			expected: []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)", "kaido(2)(1)"},
		},
	}

	for _, td := range tds {
		r := getFolderNames(td.names)
		if !reflect.DeepEqual(r, td.expected) {
			t.Fatalf("expect %v got %v", td.expected, r)
		}
	}
}
