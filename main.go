package main

import (
	"deadgopher/sandbox/sll"
	"fmt"
)

type gslice []interface{}

func main() {
	list := sll.NewList(gslice{1, true, "foo", []string{"hello", "world", "foobar"}, 5.5, sll.NewList(gslice{1, 2, 3, 4})})
	list.Add("Another Test")
	list.PrintValues()

	fmt.Println("The length of the list is : ", list.Length())
	fmt.Println(list.Contains(5.5))

	// Loop over list nodes and mutate values
	// list.ForEach(func(node *sll.Node) {
	// 	switch val := node.Val.(type) {
	// 	case string:
	// 		node.Val = val + " add on to strings"
	// 	case int:
	// 		node.Val = val * 2
	// 	default:
	// 		fmt.Printf("No case for type: %v", reflect.TypeOf(val))
	// 	}
	// })
}
