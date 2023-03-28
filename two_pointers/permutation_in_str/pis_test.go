package permutation_in_str

import (
	"reflect"
	"testing"
)

type testData struct {
	s1       string
	s2       string
	expected bool
}

func TestPISTwoMap(t *testing.T) {
	tds := []testData{
		{
			s1:       "ab",
			s2:       "cgbaddw",
			expected: true,
		},
		{
			s1:       "ab",
			s2:       "agb",
			expected: false,
		},
		{
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			s1:       "ab",
			s2:       "ac",
			expected: false,
		},
		{
			s1:       "adc",
			s2:       "dcda",
			expected: true,
		},
		{
			s1:       "a",
			s2:       "vrfeviua",
			expected: true,
		},
		{
			s1:       "abc",
			s2:       "cccccbabbbaaaa",
			expected: true,
		},
	}

	for i, td := range tds {
		r := checkInclusionTwoMap(td.s1, td.s2)
		if r != td.expected {
			t.Fatalf("index %d expect %v got %v", i, td.expected, r)
		}
	}
}

// BenchmarkReflectCmpMap
// BenchmarkReflectCmpMap-8   	15409162	        76.43 ns/op
func BenchmarkReflectCmpMap(b *testing.B) {
	count := 10000
	m1 := genMap1(count)
	m2 := genMap2(count)
	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(m1, m2)
	}
}

// BenchmarkSelfCmpMap
// BenchmarkSelfCmpMap-8   	438675344	         2.938 ns/op
func BenchmarkSelfCmpMap(b *testing.B) {
	count := 10000
	m1 := genMap1(count)
	m2 := genMap2(count)
	for i := 0; i < b.N; i++ {
		_ = selfCmpMap(m1, m2)
	}
}

func selfCmpMap(m1, m2 map[int]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		r, ok := m2[k]
		if !ok {
			return false
		}
		if r != v {
			return false
		}
	}

	return true
}

func genMap1(count int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < count; i++ {
		m[count] = count
	}
	return m
}

func genMap2(count int) map[int]int {
	m := genMap1(count)
	m[count-1] = 0
	return m
}
