package Trees

import (
	"fmt"
	"io"
	"os"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	fmt.Println("n.data----->", n.data)
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil { // checking left side of the root node after that only inserting the value into the left side node.
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			// If the left child exists, the insert method is recursively called on the left child.
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func Binarytree_Node() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.root, 0, 'M')
	data := Count_for_Binary(tree.root)
	fmt.Println("count:", data)
}

/*
tree can be traversed in three ways
In-order traversal - left|root|right
Preorder traversal - root|left|right
Post-order ”      - left|right|root
*/
func Count_for_Binary(data *BinaryNode) int {
	if data == nil {
		return 0
	}
	if data.left == nil && data.right == nil {
		return 1
	}
	return Count_for_Binary(data.left) + Count_for_Binary(data.right)
}
