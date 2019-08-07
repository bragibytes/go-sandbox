package sll

import (
	"fmt"
)

// Node : The Node type
type Node struct {
	Val  interface{}
	next *Node
	list *List
}

func (n *Node) print() {
	fmt.Printf("%v : %T\n", n.Val, n.Val)
}

func newNode(n *Node, v interface{}, l *List) *Node {
	if n == nil {

		l.Len++
		return &Node{
			v,
			nil,
			l,
		}
	}
	n.next = newNode(n.next, v, l)
	return n
}

func search(n *Node, val interface{}) *Node {
	if n == nil || n.Val == val {
		return n
	}
	return search(n.next, val)
}

func printr(n *Node) {
	if n != nil {
		n.print()
		printr(n.next)
	}
}
