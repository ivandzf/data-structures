package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	tree := NewTree[int]()
	assert.Nil(t, tree.Root)
}

func TestInitRoot(t *testing.T) {
	tree := NewTree[int]()
	tree.HashInsert(10)
	assert.Equal(t, 10, tree.Root.Element)
}

func TestHashInsert(t *testing.T) {
	tree := NewTree[int]()
	tree.HashInsert(10)
	tree.HashInsert(20)
	tree.HashInsert(5)
	tree.HashInsert(90)
	tree.HashInsert(80)
	tree.HashInsert(40)
	tree.HashInsert(32)
	tree.HashInsert(98)
	tree.Print()
	assert.Equal(t, 10, tree.Root.Element)
	assert.Equal(t, 20, tree.Root.Left.Element)
	assert.Equal(t, 5, tree.Root.Left.Left.Element)
}
