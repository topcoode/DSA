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

func (t *Rootdata) insertdata(data []int) *Rootdata {
	for _, value := range data {
		if t.Rootdata == nil {
			t.Rootdata = &Treestruct{Value: value, Left: nil, Right: nil}
		} else {
			t.Rootdata.insertdata(value)
		}

	}
	return t
}
func (n *Treestruct) insertdata(value int) *Treestruct {
	if n == nil {
		return nil
	} else if value <= n.Value {
		if n.Left == nil { // checking left side of the root node after that only inserting the value into the left side node.
			n.Left = &Treestruct{Value: value, Left: nil, Right: nil}
		} else {
			n.Left.insertdata(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Treestruct{Value: value, Left: nil, Right: nil}
		} else {
			// If the left child exists, the insert method is recursively called on the left child.
			n.Right.insertdata(value)
		}
	}
	return n
}
func (n *Rootdata) Searching_element(element int) {
	var tree Treestruct

	fmt.Println("searching data--------->", tree)

}
func Binary_Tree() {
	var Treestructdata Rootdata
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	element := 6
	Treestructdata.Searching_element(element)
	Treestructdata.insertdata(array)
	fmt.Println("Treestructdata--->", Treestructdata.Rootdata)
	fmt.Println("Treestructdata--->", Treestructdata.Rootdata.Right.Right.Value)

}
