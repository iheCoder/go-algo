package verify_preorder_serialization_of_a_binary_tree

import "testing"

func TestVP(t *testing.T) {
	po := "1,#,#"
	t.Log(isValidSerialization(po))
}
