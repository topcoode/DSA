package Transversal

import "fmt"

type NodeDatatype struct {
	data  int
	left  *NodeDatatype
	right *NodeDatatype
}

/*

	LEFT - RIGHT - ROOT

*/
// Postorder traversal function
func (n *NodeDatatype) Postorder() {
	if n == nil {
		return
	}

	n.left.Postorder()
	n.right.Postorder()
	fmt.Print(n.data, " \n")
}

func PostOrderTransversal() {
	// Create a binary tree
	n1 := &NodeDatatype{data: 1}
	n2 := &NodeDatatype{data: 2}
	n3 := &NodeDatatype{data: 3}
	n4 := &NodeDatatype{data: 4}
	n5 := &NodeDatatype{data: 5}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5

	// Perform postorder traversal
	fmt.Println("Postorder traversal:")
	n1.Postorder()
}
