package main

import (
	"deadgopher/sandbox/sll"
)

func main() {
	list := sll.NewList([]interface{}{1, true, "foo", []string{"hello", "world", "foobar"}, 5.5, -65})
	list.Add("Another Test")
	list.PrintValues()
	// list.ForEach(func(node *sll.Node) {
	// 	switch val := node.Val.(type) {
	// 	case string:
	// 		//string stuff
	// 	case int:
	// 		//int stuff
	// 	default:
	// 		fmt.Printf("No case for tyep: %v", reflect.TypeOf(val))
	// 	}
	// })
}
