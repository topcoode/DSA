package Methods

import "fmt"

type Structure struct {
	Frist  int
	Second int
}

func (n *Structure) Structure_to_Function(value *Structure) *Structure {
	fmt.Println("In function 01--->", value)
	fmt.Println("structure  02--->", n)
	return value
}
func Methods_Main() {
	var structdata Structure
	data := &Structure{
		Frist:  1,
		Second: 2,
	}

	value := structdata.Structure_to_Function(data)
	fmt.Println("value data---->", value)
}
