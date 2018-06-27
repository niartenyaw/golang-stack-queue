package main

import (
	"fmt"

	"github.com/niartenyaw/sequences/queue"
)

func main() {
	// list := linked_list.LinkedList{}

	// list.Insert(1)
	// list.Insert(2)
	// list.Insert(3)
	// list.Insert(4)
	// list.Insert(5)

	// fmt.Println(list.Remove())
	// fmt.Println(list.Remove())
	// fmt.Println(list.Remove())
	// fmt.Println(list.Remove())

	st := queue.Queue{}

	fmt.Printf("%#v\n", st)
	st.Insert(4)
	fmt.Printf("%#v\n", st)
	st.Insert(5)
	fmt.Printf("%#v\n", st)
	st.Insert(6)
	fmt.Printf("%#v\n", st)
	fmt.Println(st.Remove())
	fmt.Printf("%#v\n", st)
	fmt.Println(st.Remove())
	fmt.Printf("%#v\n", st)
	fmt.Println(st.Remove())
	fmt.Printf("%#v\n", st)
}
