package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	assert.Equal(t, 0, q.Size())
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 0, q.Size())
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 1, q.Peek())
}

func TestIterateQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var values []int
	q.Iterate(func(i int) {
		values = append(values, i)
	})
	assert.Equal(t, []int{1, 2, 3}, values)
}
