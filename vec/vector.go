package vec

import (
	"github.com/go-board/flow/iter"
	"github.com/go-board/flow/option"
)

type Vector[T any] []T

func From[T any](ts ...T) Vector[T] { return Vector[T](ts) }

func (v Vector[T]) Iter() iter.Iterator[T] {
	return &vecIter[T]{vec: v}
}

type vecIter[T any] struct {
	vec Vector[T]
}

func (vi *vecIter[T]) Next() option.Optional[T] {
	panic("Not implemented!")
}
