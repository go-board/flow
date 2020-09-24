package iter

import "github.com/go-board/flow/option"

type AsyncItem[T any] struct {
	ch <-chan T
}

type AsyncIteator[T any] interface {
	AsyncNext() option.Optional[AsyncItem[T]]
}
