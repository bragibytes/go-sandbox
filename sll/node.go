package sll

// Node : The Node type
type Node struct {
	Val  interface{}
	next *Node
}

func newNode(val interface{}) *Node {
	return &Node{
		val,
		nil,
	}
}
