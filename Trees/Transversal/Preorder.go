package Transversal

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

/*

ROOT - LEFT - RIGHT

*/
// Preorder traversal function
func (n *Node) Preorder() (int, int) {

	if n == nil {
		return 0, 0
	}
	fmt.Print(n.data, " \n")
	preorderleft, _ := n.left.Preorder()
	preorderright, _ := n.right.Preorder()
	return preorderleft, preorderright
}

func Preorder_data() {
	// Create a binary tree
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	n5 := &Node{data: 5}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5

	// Perform preorder traversal
	fmt.Println("Preorder traversal:")
	n1.Preorder()
}
