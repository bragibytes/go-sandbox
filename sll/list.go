package sll

// List : A generic singly linked list
type List struct {
	head *Node
	Len  int
}

// NewList creates and returns a new List. You can provide a starting list of values, or nil for an empty list
func NewList(vals []interface{}) *List {
	list := &List{
		nil,
		0,
	}

	if vals != nil {
		list.AddSlice(vals)
	}

	return list
}

// Add a node to the list
func (l *List) Add(val interface{}) {
	l.head = newNode(l.head, val, l)
}

// AddSlice takes a slice of values and adds them to the list
func (l *List) AddSlice(s []interface{}) {
	for _, v := range s {
		l.Add(v)
	}
}

// PopTail removes the node at the end of the list, and returns its Value ( last in first out )
func (l *List) PopTail() interface{} {
	if l.head != nil {
		if l.head.next == nil {
			res := l.head.Val
			l.head = nil
			l.Len--
			return res
		}
		x := l.head
		for x.next.next != nil {
			x = x.next
		}
		res := x.next.Val
		x.next = nil
		return res
	}

	return nil
}

// PopHead removes the node at the front of the list, and returns its Value ( first in first out )
func (l *List) PopHead() interface{} {
	if l.head != nil {
		res := l.head.Val
		l.head = l.head.next
		l.Len--
		return res
	}
	return nil
}

// Slice returns the list as a slice
func (l *List) Slice() []interface{} {
	var slice []interface{}

	for x := l.head; x != nil; x = x.next {
		slice = append(slice, x.Val)
	}

	return slice
}

// Contains checks if the list contains the given value
func (l *List) Contains(val interface{}) bool {
	node := search(l.head, val)
	return node != nil
}

// PrintValues prints each value in the list to the console
func (l *List) PrintValues() {
	printr(l.head)
}

// Clear the list of all nodes
func (l *List) Clear() {
	l.head = nil
}

// ForEach loops through the list and performs a custom action on each node
func (l *List) ForEach(f func(*Node)) {
	for x := l.head; x != nil; x = x.next {
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
