package main

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type binaryNode[T any] struct {
	Element T
	Left    *binaryNode[T]
	Right   *binaryNode[T]
}

type BinaryRoot[T any] struct {
	Root *binaryNode[T]
}

func NewTree[T any]() BinaryRoot[T] {
	return BinaryRoot[T]{}
}

func (t *BinaryRoot[T]) HashInsert(element T) {
	if t.Root == nil {
		t.Root = &binaryNode[T]{
			Element: element,
		}

		return
	}

	t.Root.hashInsert(element)
}

func (t *BinaryRoot[T]) Print() {
	if t == nil {
		return
	}

	t.Root.print(os.Stdout, 0, "root")
}

func (t *binaryNode[T]) hashInsert(element T) {
	if t == nil {
		return
	}

	hashedElement, err := NewHashKey(element)
	if err != nil {
		log.Panicf("failed to hash the key %v, error: %v", element, err)
	}

	if hashedElement.LessThan(t.Element) {
		if t.Left == nil {
			t.Left = &binaryNode[T]{
				Element: element,
			}
		} else {
			t.Left.hashInsert(element)
		}
	} else {
		if t.Right == nil {
			t.Right = &binaryNode[T]{
				Element: element,
			}
		} else {
			t.Right.hashInsert(element)
		}
	}
}

func (t *binaryNode[T]) print(w io.Writer, tab int, ch string) {
	if t == nil {
		return
	}

	for i := 0; i < tab; i++ {
		fmt.Fprint(w, " ")
	}

	fmt.Fprintf(w, "%s: %+v\n", ch, t.Element)
	t.Left.print(w, tab+2, "left")
	t.Right.print(w, tab+2, "right")
}
