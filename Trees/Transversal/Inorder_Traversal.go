package Transversal

import "fmt"

type NodeData struct {
	data  int
	left  *NodeData
	right *NodeData
}

/*

LEFT - ROOT - RIGHT

*/
// Inorder traversal function
func (n *NodeData) Inorder() {
	if n == nil {
		return
	}

	n.left.Inorder()
	fmt.Print(n.data, " \n")
	n.right.Inorder()
}

func Inorder_Transversal() {
	// Create a binary tree
	n1 := &NodeData{data: 1}
	n2 := &NodeData{data: 2}
	n3 := &NodeData{data: 3}
	n4 := &NodeData{data: 4}
	n5 := &NodeData{data: 5}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5

	// Perform inorder traversal
	fmt.Println("Inorder traversal:")
	n1.Inorder()
}
