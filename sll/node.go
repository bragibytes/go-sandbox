package sll

// Node : The Node type
type Node struct {
	Val  interface{}
	next *Node
}

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
