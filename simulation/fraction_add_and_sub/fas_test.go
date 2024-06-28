package fraction_add_and_sub

import "testing"

func TestSplitFrac(t *testing.T) {
	exp := "-1/2+1/2"
	//t.Log(fractionAddition(exp))
	exp = "-1/2+1/2+1/3"
	exp = "1/3-1/2"
	t.Log(fractionAddition(exp))
}
