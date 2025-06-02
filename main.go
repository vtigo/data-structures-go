package main

import (
	"github.com/vtigo/data-structures-go/linkedList"
)

func main() {
	list := linkedList.NewLinkedList()

	list.PushFront(5)
	list.PushFront(8)
	list.PushFront(2)
	list.PushBack(2)

	list.Print()

	list.Delete(5)
	list.Print()
}

