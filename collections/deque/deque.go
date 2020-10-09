package deque

import (
	"github.com/go-board/flow/option"
)

type Deque[T any] struct {

}

type dequeIter[T any] struct {}

func (d dequeIter[T]) Next() option.Optional[T] {
	return option.None[T]()
}
