package serialize_and_deserialize_BST

import (
	"bytes"
	"encoding/binary"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

var nullBytes = []byte{39, 17}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var bb bytes.Buffer
	stack := []*TreeNode{root}
	var p *TreeNode
	for len(stack) != 0 {
		p = stack[0]
		stack = stack[1:]
		if p == nil {
			bb.Write(nullBytes)
		} else {
			valBytes := make([]byte, 2)
			binary.BigEndian.PutUint16(valBytes, uint16(p.Val))
			bb.Write(valBytes)

			stack = append(stack, p.Left)
			stack = append(stack, p.Right)
		}
	}

	return bb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	db := []byte(data)
	root := &TreeNode{Val: int(binary.BigEndian.Uint16(db[:2]))}
	var p *TreeNode
	stack := []*TreeNode{root}
	var val uint16
	for i := 2; i < len(db); i += 2 {
		p = stack[0]
		stack = stack[1:]
		val = binary.BigEndian.Uint16(db[i : i+2])
		if val != 10001 {
			p.Left = &TreeNode{Val: int(val)}
			stack = append(stack, p.Left)
		}
		i += 2
		if i+2 >= len(db) {
			break
		}
		val = binary.BigEndian.Uint16(db[i : i+2])
		if val != 10001 {
			p.Right = &TreeNode{Val: int(val)}
			stack = append(stack, p.Right)
		}
	}

	return root
}
