package Leaf

import "fmt"

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func Using_tree_Node_struct() {
	Treedata := &Tree{
		Data: 25,
		Left: &Tree{
			Data: 36,
			Left: &Tree{
				Data:  97,
				Left:  nil,
				Right: nil,
			},
			Right: &Tree{
				Data:  99,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Tree{
			Data: 32,
			Left: &Tree{
				Data:  72,
				Left:  nil,
				Right: nil,
			},
			Right: &Tree{
				Data:  72,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(Treedata)
	data := Using_tree_Node_struct_Main(Treedata)
	fmt.Println("no of the tree nodes:", data)

}
func Using_tree_Node_struct_Main(root *Tree) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return Using_tree_Node_struct_Main(root.Left) + Using_tree_Node_struct_Main(root.Right)
}

/*
getLeafCount(node)
1) If node is NULL then return 0.
2) Else If left and right child nodes are NULL return 1.
3) Else recursively calculate leaf count of the tree using below formula.
    Leaf count of a tree = Leaf count of left subtree + Leaf count of right subtree
*/
