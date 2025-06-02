package linkedList

import "fmt"

type DoublyNode struct {
	Data int
	Next *DoublyNode
	Prev *DoublyNode
}

type DoublyLinkedList struct {
	Head *DoublyNode
	Tail *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *DoublyLinkedList) Print() {
	if l.Head == nil {
		fmt.Println("The list is empty")
	}

	// iterate over the list nodes and print them
	current := l.Head
	for current != nil {
		fmt.Printf("%d", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func (l *DoublyLinkedList) PushFront(value int) {
	// Create the node to be added to the list
	newNode := &DoublyNode{
		Data: value,
		Next: l.Head,
		Prev: nil,
	}
		
	// If the list is not empty, we point the old head prev to the new node
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	
	// Point the head to the new node
	l.Head = newNode
	
	// If the list is empty, we should assign the new node as the tail too
	if l.Tail == nil {
		l.Tail = newNode
	}
}

func (l *DoublyLinkedList) PopFront() {
	// TODO
}
