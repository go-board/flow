package binaryheap

import (
	"github.com/go-board/flow/option"
)

type priorityQueue[T any] struct {

}

func (q priorityQueue[T]) Push(t T) {

}

func (q priorityQueue[T]) Pop() option.Optional[T] {}

func (q priorityQueue[T]) Peek() option.Optional[T] {}
