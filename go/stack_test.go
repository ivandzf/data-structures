package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.Equal(t, 0, s.Size())
}

func TestInsertStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())
}

func TestPopStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())

	s.Pop()
	assert.Equal(t, 2, s.Size())
}

func TestIterateStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())

	var values []int
	s.Iterate(func(v int) {
		values = append(values, v)
	})
	assert.Equal(t, []int{3, 2, 1}, values)
}

func TestPeekStack(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Size())
	assert.Equal(t, 3, s.Peek())
}
