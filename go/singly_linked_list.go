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

func (ll *LinkedList[T]) InsertFirst(element T) {
	newNode := Node[T]{element, nil}
	if ll.Head == nil {
		ll.Head = &newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = &newNode
	}

	ll.Count++
}

func (ll *LinkedList[T]) GetNthNode(index uintptr) *Node[T] {
	if index >= ll.Count {
		return nil
	}

	current := ll.Head
	for i := uintptr(0); i < index; i++ {
		current = current.Next
	}

	return current
}

func (ll *LinkedList[T]) RemoveNthNode(index uintptr) {
	if index >= ll.Count {
		return
	}
	if index == 0 {
		ll.Head = ll.Head.Next
	} else {
		current := ll.Head
		for i := uintptr(0); i < index-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}

	ll.Count--
}

func (ll *LinkedList[T]) RemoveFirst() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	ll.Count--
}

func (ll *LinkedList[T]) RemoveLast() {
	if ll.Head == nil {
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = nil
	ll.Count--
}

func (ll *LinkedList[T]) Print() {
	current := ll.Head

	for current != nil {
		println(current.Element)
		current = current.Next
	}
}
