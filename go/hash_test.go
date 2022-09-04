package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHashKey(t *testing.T) {
	h, err := NewHashKey(10)
	assert.Nil(t, err)
	assert.NotEmpty(t, h)
}

func TestHashKeyGreaterThan(t *testing.T) {
	h, err := NewHashKey(10)
	assert.Nil(t, err)
	assert.NotEmpty(t, h)

	h2 := h.GreaterThan(1)
	assert.True(t, h2)
}

func TestHashKeyLessThan(t *testing.T) {
	h, err := NewHashKey("abcd")
	assert.Nil(t, err)
	assert.NotEmpty(t, h)

	h2 := h.LessThan("qwer")
	assert.True(t, h2)
}
