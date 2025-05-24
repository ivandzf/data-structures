package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStackLimited(t *testing.T) {
	s := NewStackLimited[int](2)

	assert.Equal(t, 2, s.Limit, "limit must be the same")
}

func TestPopStackLimited(t *testing.T) {
	s := NewStackLimited[int](2)
	s.Push(10)
	s.Pop()

	assert.Equal(t, 0, len(s.Stack), "the length must be 0 after pop")

	s.Push(100)
	s.Push(50)
	s.Pop()

	assert.Equal(t, 1, len(s.Stack), "the length must be 1 after pop")
	assert.Equal(t, 100, s.Peek(), "the value must be 100")
}

func TestPushStackLimited(t *testing.T) {
	s := NewStackLimited[int](2)
	s.Push(10)

	assert.Equal(t, 1, len(s.Stack), "the length must be 1")
	assert.Equal(t, 10, s.Peek(), "the value of peek stack must be 10")

	s.Push(50)

	assert.Equal(t, 2, len(s.Stack), "the length must be 2")
	assert.Equal(t, 50, s.Peek(), "the value of peek stack must be 50")
	assert.Equal(t, 10, s.Bottom(), "the value of lower must be 10")

	s.Push(99)

	assert.Equal(t, 2, len(s.Stack), "the length must be 2")
	assert.Equal(t, 99, s.Peek(), "the value of peek stack must be 99")
	assert.Equal(t, 10, s.Bottom(), "the value of lower must be 10")
}
