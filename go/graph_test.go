package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph[string]()
	assert.Equal(t, 0, len(g.Vertices))
	assert.Equal(t, 0, len(g.Edges))
}

func TestCreateVertex(t *testing.T) {
	g := NewGraph[string]()
	a := g.CreateVertex("a")

	assert.Equal(t, 1, len(g.Vertices))
	assert.Equal(t, "a", a.Element)
}

func TestCreateEdge(t *testing.T) {
	g := NewGraph[string]()
	a := g.CreateVertex("a")
	b := g.CreateVertex("b")
	c := g.CreateVertex("c")
	d := g.CreateVertex("d")

	assert.Equal(t, 4, len(g.Vertices))

	g.AddEdge(a, b)
	g.AddEdge(a, c)
	g.AddEdge(b, d)
	g.AddEdge(c, d)
	g.AddEdge(d, a)

	assert.Equal(t, 4, len(g.Edges))

	g.Print()

	g.DepthFirstSearch(d, make(map[*Vertex[string]]bool))
}
