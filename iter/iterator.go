package iter

import "github.com/go-board/flow/option"

type Iterator[T any] interface {
	Next() option.Optional[T]
}

type SizedIterator[T any] interface {
	Iterator[T]
	Size() int
}

type DoubleEndedIterator[T any] interface {
	Iterator[T]
	NextBack() option.Optional[T]
}
