package main

import "fmt"

type Node struct {
	Data int    // data stored in the node
	Next *Node  // a pointer to the next node in the list
}

type LinkedList struct {
	Head *Node  // a pointer to the head of the list (the first element of the list)
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func (l *LinkedList) Print() {
	if l.Head == nil {
		fmt.Println("The list is empty")
		return
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

func (l *LinkedList) PushFront(value int) {
	// point the pushed Node to the head
	newNode := &Node{
		Data: value,
		Next: l.Head,
	}
	
	// update the list head to point to the new Node
	l.Head = newNode
}

func (l *LinkedList) PopFront() {
	if l.Head != nil {
		// just change the head pointer to point to the next Node
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) PushBack(value int) {
	// initate a Node pointing to nil (it will be the last Node of the list)
	newNode := &Node{
		Data: value,
		Next: nil,
	}
	
	// if the list is empty, just add it as the head
	if l.Head == nil {
		l.Head = newNode
		return
	}
	
	// iterate over the list until we find the tail
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	// add the new node ass the next Node of the current tail
	current.Next = newNode
}

func (l *LinkedList) PopBack() {
	// return if the list is empty
	if l.Head == nil {
		return
	}
	
	// just remove the head if the list has only one Node
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	
	// run the list until whe are in the Node just before the tail
	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	// cut the link before with the tail
	current.Next = nil
}

func (l *LinkedList) Find(key int) *Node {
	current := l.Head
	for current != nil {
		if current.Data == key {
			return current
		}
		current = current.Next
	}

	return nil
}

func (l *LinkedList) Delete(value int) {
	// Return early if the list is empty
	if l.Head == nil {
		return
	}

	// If the head node matches, remove it by advancing the head pointer
	if l.Head.Data == value {
		l.Head = l.Head.Next
		return
	}
	
	// Traverse the list to find the node before the one to delete
	current := l.Head
	for current.Next != nil {
		// If the next node contains the target value
 		if current.Next.Data == value {
			// Skip over the next node to remove it from the list
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

func (l *LinkedList) Length() int {
	count := 0

	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}

	return count
}
