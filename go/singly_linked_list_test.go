package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	assert.Equal(t, nil, ll.Head)
	assert.Equal(t, uintptr(0), ll.Count)
}
