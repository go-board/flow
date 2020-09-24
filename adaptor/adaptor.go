package adaptor

import (
	"github.com/go-board/flow/iter"
	"github.com/go-board/flow/option"
)

func Slice[T any](slice []T) iter.Iterator[T] {
    return &sliceIter[T]{elems: slice, size: len(slice)}
}

type sliceIter[T any] struct {
    elems 	[]T
	i 		int
	size 	int
}

func (si *sliceIter[T]) Next() option.Optional[T] {
	// todo: implement this.
    return option.None()    
}
