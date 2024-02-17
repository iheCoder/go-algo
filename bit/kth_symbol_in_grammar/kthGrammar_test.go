package kth_symbol_in_grammar

import "testing"

func TestKthGrammar(t *testing.T) {
	n := 5
	k := 5
	//t.Log(kthGrammar(n, k))
	//n = 3
	//k = 8
	//t.Log(kthGrammar(n, k))
	//n = 17
	//k = 92
	//t.Log(kthGrammar(n, k))
	n = 3
	k = 4
	t.Log(kthGrammar(n, k))
}
