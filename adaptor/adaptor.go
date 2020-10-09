package adaptor

import (
	"github.com/go-board/flow/iter"
	"github.com/go-board/flow/option"
)

// Slice turn slice to an iterator of slice element.
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

func DoubleEndedSlice[T any](slice []T) iter.DoubleEndedIterator[T] {
	return &doubleEndedSliceIter[T]{}
}

type doubleEndedSliceIter[T any] struct {

}

func (dsi *doubleEndedSliceIter[T]) Next() option.Optional[T] {

}

func (dsi *doubleEndedSliceIter[T any]) NextBack() option.Optionalp[T] {

}
