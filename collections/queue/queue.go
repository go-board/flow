package queue

import (
	"github.com/go-board/flow/option"
)

type Queue[T any] interface {
	Push(t T)
	Pop() option.Optional[T]
	Peek() option.Optional[T]
}
