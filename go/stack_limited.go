package main

type StackLimited[T any] struct {
	Limit int
	Stack []T
}

func NewStackLimited[T any](limit int) StackLimited[T] {
	return StackLimited[T]{
		Limit: limit,
		Stack: []T{},
	}
}

func (s *StackLimited[T]) Push(v T) {
	// if the length reached the limit, then pop the latest
	if len(s.Stack) >= s.Limit {
		s.Pop()
	}

	s.Stack = append(s.Stack, v)
}

func (s *StackLimited[T]) Pop() {
	l := len(s.Stack)
	if l == 0 {
		panic("empty stack")
	}

	s.Stack = s.Stack[:l-1]
}

func (s *StackLimited[T]) Peek() T {
	l := len(s.Stack)
	if l == 0 {
		panic("empty stack")
	}

	return s.Stack[l-1]
}

func (s *StackLimited[T]) Bottom() T {
	l := len(s.Stack)
	if l == 0 {
		panic("empty stack")
	}

	return s.Stack[0]
}

func (s *StackLimited[T]) Iterate(f func(T)) {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		f(s.Stack[i])
	}
}
