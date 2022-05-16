package main

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() {
	l := len(*s)
	if l == 0 {
		panic("empty stack")
	}

	*s = (*s)[:l-1]
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Peek() T {
	l := len(*s)
	if l == 0 {
		panic("empty stack")
	}

	return (*s)[l-1]
}

func (s *Stack[T]) Iterate(f func(T)) {
	for i := len(*s) - 1; i >= 0; i-- {
		f((*s)[i])
	}
}
