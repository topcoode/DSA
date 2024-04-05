package Leaf

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func CountLeft_binarytree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := CountLeft_binarytree(root.Left) // Recursively count left nodes
	rightCount := CountLeft_binarytree(root.Right)
	fmt.Println("rightcount--->", rightCount)
	return 1 + leftCount
}

func CountLeftNodes_main() {
	// Constructing a binary tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("The total number of left nodes in the binary tree is:", CountLeft_binarytree(root))
}
