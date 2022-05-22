package main

import "fmt"

type Vertex[T any] struct {
	Element T
}

func (v Vertex[T]) String() string {
	return fmt.Sprintf("%v", v.Element)
}

type Graph[T any] struct {
	Vertices []*Vertex[T]
	Edges    map[*Vertex[T]][]*Vertex[T]
}

func NewGraph[T any]() Graph[T] {
	return Graph[T]{
		Vertices: make([]*Vertex[T], 0),
		Edges:    make(map[*Vertex[T]][]*Vertex[T]),
	}
}

func (g *Graph[T]) CreateVertex(element T) *Vertex[T] {
	vertex := new(Vertex[T])
	vertex.Element = element

	g.Vertices = append(g.Vertices, vertex)

	return vertex
}

func (g *Graph[T]) AddEdge(vertex1 *Vertex[T], vertex2 *Vertex[T]) {
	if g.Edges[vertex1] == nil {
		g.Edges[vertex1] = make([]*Vertex[T], 0)
	}

	g.Edges[vertex1] = append(g.Edges[vertex1], vertex2)
}

func (g *Graph[T]) RemoveEdge(vertex1 *Vertex[T], vertex2 *Vertex[T]) {
	for i, v := range g.Edges[vertex1] {
		if v == vertex2 {
			g.Edges[vertex1] = append(g.Edges[vertex1][:i], g.Edges[vertex1][i+1:]...)
		}

		println(g.Edges[vertex1])

		if len(g.Edges[vertex1]) == 0 {
			g.Edges[vertex1] = nil
		}
	}
}

func (g *Graph[T]) RemoveVertex(vertex *Vertex[T]) {
	for i, v := range g.Vertices {
		if v == vertex {
			g.Vertices = append(g.Vertices[:i], g.Vertices[i+1:]...)
			break
		}
	}

	delete(g.Edges, vertex)
	for _, v := range g.Vertices {
		g.RemoveEdge(v, vertex)
	}
}

func (g *Graph[T]) Print() {
	var str string
	for i := 0; i < len(g.Vertices); i++ {
		str += g.Vertices[i].String() + " -> "
		near := g.Edges[g.Vertices[i]]

		for j := 0; j < len(near); j++ {
			str += near[j].String() + " "
		}

		str += "\n"
	}

	fmt.Println(str)
}

func (g *Graph[T]) DepthFirstSearch(vertex *Vertex[T], visited map[*Vertex[T]]bool) {
	visited[vertex] = true
	fmt.Println(vertex.Element)
	for _, v := range g.Edges[vertex] {
		if visited[v] == false {
			g.DepthFirstSearch(v, visited)
		}
	}
}
