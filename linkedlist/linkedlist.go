package main

import (
	"fmt"
)

type Node struct {
	Value string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Insert(value string) {
	new := &Node{Value: value}

	if ll.Head == nil {
		ll.Head = new
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = new
}

func (ll *LinkedList) DeleteLast() {
	if ll.Head == nil {
		return
	}

	if ll.Head.Next == nil {
		ll.Head = nil
		return
	}

	current := ll.Head
	var prevNode *Node
	for current.Next != nil {
		prevNode = current
		current = current.Next
	}
	prevNode.Next = nil
}

// Head -> node1 -> node2
func (ll *LinkedList) Print() {
	if ll.Head == nil {
		fmt.Print("(List is empty!)")
		return
	}

	if ll.Head.Next == nil {
		fmt.Print(ll.Head.Value)
		return
	}

	current := ll.Head
	for current != nil {
		fmt.Print(current.Value)
		current = current.Next
		if current != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
