package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	assert.Equal(t, uintptr(0), ll.Count)
	assert.Equal(t, nil, ll.Head)
}

func TestInsertLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	// insert last
	ll.InsertLast(1)
	assert.Equal(t, uintptr(1), ll.Count)
	assert.Equal(t, 1, ll.GetNthNode(0).Element)

	// insert first
	ll.InsertFirst(10)
	assert.Equal(t, uintptr(2), ll.Count)
	assert.Equal(t, 10, ll.GetNthNode(0).Element)
}

func TestRemoveLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.InsertLast(1)
	ll.InsertLast(10)
	ll.InsertLast(100)
	assert.Equal(t, uintptr(3), ll.Count)

	// remove last
	ll.RemoveLast()
	assert.Equal(t, uintptr(2), ll.Count)
	assert.Equal(t, nil, ll.GetNthNode(2))

	// remove first
	ll.RemoveFirst()
	assert.Equal(t, uintptr(1), ll.Count)
	assert.Equal(t, 10, ll.GetNthNode(0).Element)
}

func TestPrint(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.InsertLast(1)
	ll.InsertLast(10)
	ll.InsertLast(100)
	ll.InsertLast(1000)
	ll.InsertLast(10000)

	ll.Print()
	t.Log("1\n10\n100\n1000\n10000")
}
