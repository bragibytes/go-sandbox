package sll

import (
	"fmt"
)

// List : A generic singly linked list
type List struct {
	head *Node
}

// NewList creates and returns a new List. You can provide a starting list of values, or nil for an empty list
func NewList(vals []interface{}) *List {
	list := &List{
		nil,
	}

	if vals != nil {
		list.AddSlice(vals)
	}

	return list
}

// Add a node to the list
func (l *List) Add(val interface{}) {
	l.head = newNode(l.head, val)
}

// AddSlice takes a slice of values and adds them to the list
func (l *List) AddSlice(s []interface{}) {
	for _, v := range s {
		l.Add(v)
	}
}

// PopTail removes the node at the end of the list, and returns its Value ( last in first out )
func (l *List) PopTail() (interface{}, bool) {
	if l.head != nil {
		if l.head.next == nil {
			res := l.head.Val
			l.head = nil
			return res, true
		}
		x := l.head
		for x.next.next != nil {
			x = x.next
		}
		res := x.next.Val
		x.next = nil
		return res, true
	}

	return nil, false
}

// PopHead removes the node at the front of the list, and returns its Value ( first in first out )
func (l *List) PopHead() (interface{}, bool) {
	if l.head != nil {
		if l.head.next == nil {
			res := l.head.Val
			l.head = nil
			return res, true
		}
		res := l.head.Val
		l.head = l.head.next
		return res, true
	}

	return nil, false
}

// Slice returns the list as a slice
func (l *List) Slice() []interface{} {
	slice := make([]interface{}, 0)

	if l.head != nil {
		x := l.head
		for x.next != nil {
			slice = append(slice, x.Val)
			x = x.next
		}
		slice = append(slice, x.Val)
	}

	return slice
}

// Contains checks if the list contains the given value
func (l *List) Contains(Val interface{}) bool {
	if l.head != nil {
		x := l.head
		for x.next != nil {
			if x.Val == Val {
				return true
			}
			x = x.next
		}
		if x.Val == Val {
			return true
		}
	}

	return false
}

// PrintValues prints each value in the list to the console
func (l *List) PrintValues() {
	if l.head != nil {
		x := l.head
		for x.next != nil {
			fmt.Println(x.Val)
			x = x.next
		}
		fmt.Println(x.Val)
	} else {
		fmt.Println("List is empty...")
	}
}

// Clear the list of all nodes
func (l *List) Clear() {
	l.head = nil
}

// ForEach loops through the list and performs a custom action on each node
func (l *List) ForEach(f func(*Node)) {
	if l.head != nil {
		x := l.head
		for x.next != nil {
			f(x)
			x = x.next
		}
		f(x)
	}
}

// Reverse the list
func (l *List) Reverse() {
	var prev, current, next *Node = nil, l.head, nil
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}
