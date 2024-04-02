package Trees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
type BinaryTree_struct struct {
	Root *Node
}

func (n *BinaryTree_struct) Binary_Data_function(value int) {
	if n.Root == nil {
		n.Root = &Node{
			Value: value,
			Left:  nil,
			Right: nil,
		}
	} else {
		n.Root.Main_Binartfunc()
	}

}
func (n *Node) Main_Binartfunc() {

}
