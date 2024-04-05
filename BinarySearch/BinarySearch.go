package BinarySearch

import "fmt"

type Treestruct struct {
	Value int
	Left  *Treestruct
	Right *Treestruct
}
type Rootdata struct {
	Rootdata *Treestruct
}

func (n *Rootdata) Binary_Root_Creation(data []int, value int) {
	var datatree Treestruct

	for _, data := range data {
		if n.Rootdata == nil {
			n.Rootdata = &Treestruct{
				Value: data,
				Left:  nil,
				Right: nil,
			}
		} else {
			datatree.Nodes_creation(data)
		}
	}

}
func (n *Treestruct) Nodes_creation(value int) {
	if n == nil {
		return
	} else if value <= n.Value {
		fmt.Println()
	}

}
func Binary_Tree() {
	var Treestructdata Rootdata
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	element := 6

	Treestructdata.Binary_Root_Creation(array, element)
}
