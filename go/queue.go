package main

type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

func (q *Queue[T]) Dequeue() T {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

func (q *Queue[T]) Iterate(f func(T)) {
	for _, v := range *q {
		f(v)
	}
}
