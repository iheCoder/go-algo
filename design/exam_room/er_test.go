package exam_room

import (
	"strings"
	"testing"
)

func TestPrintTestCases(t *testing.T) {
	n := 10000
	ss := make([]string, n)
	bs := make([]string, n)

	for i := 0; i < n; i++ {
		ss[i] = "\"seat\""
		bs[i] = "[]"
	}
	t.Log(strings.Join(ss, ","))
	t.Log(strings.Join(bs, ","))
}

func TestExamRoom(t *testing.T) {
	n := 10
	er := Constructor(n)
	er.Seat()
	er.Seat()
	er.Seat()
	er.Seat()
	er.Leave(4)
	er.Seat()
}
