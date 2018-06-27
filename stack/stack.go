package stack

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top    *Node
	length int
}

func (list *Stack) Push(val int) {
	node := Node{value: val, next: list.top}
	list.top = &node
	list.length = list.length + 1
}

func (list *Stack) Pop() int {
	if list.length > 0 {
		value := list.top.value
		list.top = list.top.next
		list.length = list.length - 1
		return value
	}
	panic("bad")
}

func (list *Stack) Length() int {
	return list.length
}
