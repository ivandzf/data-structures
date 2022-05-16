package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMap(t *testing.T) {
	m := NewHashMap[int]()
	assert.Equal(t, 0, m.Size())
}

func TestInsert(t *testing.T) {
	m := NewHashMap[int]()
	m.Insert("key1", 1)
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, 1, m.Get("key1"))
}

func TestLookup(t *testing.T) {
	m := NewHashMap[int]()
	m.Insert("key1", 1)
	assert.Equal(t, 1, m.Get("key1"))
}

func TestDelete(t *testing.T) {
	m := NewHashMap[int]()
	m.Insert("key1", 1)
	m.Delete("key1")
	assert.Equal(t, 0, m.Size())
}

func TestSize(t *testing.T) {
	m := NewHashMap[int]()
	m.Insert("key1", 1)
	assert.Equal(t, 1, m.Size())

	m.Delete("key1")
	assert.Equal(t, 0, m.Size())
}

func TestIterate(t *testing.T) {
	m := NewHashMap[int]()
	m.Insert("key1", 1)
	m.Insert("key2", 2)
	m.Insert("key3", 3)

	var keys []string
	m.Iterate(func(key string, value int) {
		keys = append(keys, key)
	})

	sort.Strings(keys)

	assert.Equal(t, []string{"key1", "key2", "key3"}, keys)
}
