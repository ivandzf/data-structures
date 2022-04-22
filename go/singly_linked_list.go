package main

type Node[T any] struct {
	Element T
	Next    *Node[T]
}

type LinkedList[T any] struct {
	Head  *Node[T]
	Count uintptr
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{
		Head:  nil,
		Count: 0,
	}
}

func (ll *LinkedList[T]) InsertLast(element T) {
	newNode := Node[T]{element, nil}
	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newNode
	}
	ll.Count++
}
