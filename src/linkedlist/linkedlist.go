package linkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	mHead *Node
}

func NewList() *List {
	return &List{}
}

// Insert value to the list
func (l *List) Insert(value interface{}) {
	l.mHead = &Node{value, l.mHead}
}

// Delete element from the list
func (l *List) Delete(key interface{}) interface{} {
	var (
		tmp  *Node = l.mHead
		prev *Node
	)

	if tmp != nil && tmp.Value == key {
		l.mHead = tmp.Next
		return tmp
	} else {
		for {
			if tmp != nil && tmp.Value != key {
				break
			}
			prev = tmp
			tmp = tmp.Next
		}
		if tmp == nil {
			return nil
		}
		if prev != nil {
			prev.Next = tmp.Next
		} else {
			l.mHead = tmp.Next
		}
	}
	return tmp
}

// Print all of nodes
func (l *List) Print() {
	tmp := l.mHead
	for {
		if tmp == nil {
			break
		}
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Next
	}
	fmt.Print("\n")
}

// Search value in list and returns value if found
func (l *List) Search(value interface{}) interface{} {
	tmp := l.mHead
	for {
		if tmp == nil {
			break
		}
		if tmp.Value == value {
			return tmp.Value
		}
		tmp = tmp.Next
	}
	return nil
}

func (l *List) IsInList(value interface{}) bool {
	tmp := l.mHead
	for {
		if tmp == nil {
			break
		}
		if tmp.Value == value {
			return true
		}
		tmp = tmp.Next
	}
	return false
}

func (l *List) GetHead() *Node {
	return l.mHead
}

func (l *List) Sort() {

}

func (l *List) Reverse() {}
