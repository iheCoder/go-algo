package reorganize_string

import "testing"

func TestRS(t *testing.T) {
	s := "aab"
	t.Log(reorganizeString(s))
	s = "aabb"
	t.Log(reorganizeString(s))
	s = "aaab"
	t.Log(reorganizeString(s))
	s = "babaa"
	t.Log(reorganizeString(s))
	s = "gdhpuhcnavxxmmaopkqkuwhbbfkoypombvmchilxvmhiihjrlzddeekznjeiahjfncebjgmoyptibtjzwnbbhfmwanhxtmmfcindicspqtwexljzzlzlqmrwxjkmoyngttltjbuhvvxxbxtorpchfbahbceqbzkpvsgbjgjfukikvtvdeojxlwqfhfkvgnzuunwuxrbqpnmimtkkehzrzk"
	t.Log(reorganizeString(s))
}
