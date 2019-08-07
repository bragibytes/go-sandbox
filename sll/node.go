package sll

import (
	"fmt"
	"reflect"
)

// Node : The Node type
type Node struct {
	Val  interface{}
	next *Node
}

func (n *Node) print() {
	fmt.Printf("%v : %v\n", n.Val, reflect.TypeOf(n.Val))
}

// Recursive node functions
func newNode(n *Node, v interface{}) *Node {
	if n == nil {
		return &Node{
			v,
			nil,
		}
	}
	n.next = newNode(n.next, v)
	return n
}

func search(n *Node, val interface{}) *Node {
	if n == nil || n.Val == val {
		return n
	}
	return search(n.next, val)
}

func printVals(n *Node) {
	if n != nil {
		n.print()
		printVals(n.next)
	}
}

func forEach(n *Node, f func(*Node)) {
	if n != nil {
		f(n)
		forEach(n.next, f)
	}
}
