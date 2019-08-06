package main

import (
	"deadgopher/sandbox/sll"
)

func main() {
	list := sll.NewList(42)
	list.AddSlice([]interface{}{1, 2, 3, 4, 5, 100})

	list.PrintValues()
	list.PopTail()
	list.PrintValues()
}
