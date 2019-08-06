package bst

import "fmt"

// Node : BST Node
type Node struct {
	Val   int
	left  *Node
	right *Node
}

func newNode(val int) *Node {
	return &Node{
		val,
		nil,
		nil,
	}
}

// Search : Searches the tree for a value
func Search(n *Node, v int) *Node {
	if n == nil || n.Val == v {
		return n
	}

	if v < n.Val {
		return Search(n.left, v)
	}
	return Search(n.right, v)
}

// Insert : Inserts a value into the tree
func Insert(n *Node, v int) *Node {

	if n == nil {
		return newNode(v)
	}

	if v < n.Val {
		n.left = Insert(n.left, v)
	} else if v > n.Val {
		n.right = Insert(n.right, v)
	}

	return n
}

// InOrder : Prints out the values of the tree in order
func InOrder(n *Node) {
	if n != nil {
		InOrder(n.left)
		fmt.Println(n.Val)
		InOrder(n.right)
	}
}
