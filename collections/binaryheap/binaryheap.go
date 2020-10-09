package binaryheap

import "github.com/go-board/flow/collections/vec"

type BinaryHeap[T any] struct {
	data vec.Vector[T]
	cmp  func(T, T) int
}

func Make[T any](cmp func(T, T) int, elems ...T) BinaryHeap[T] {
	h := BinaryHeap[T]{data: vec.Of(elems...), cmp: cmp}
	h.init()
	return h
}

func (h BinaryHeap[T]) init() {
}
