package iter

import "github.com/go-board/flow/option"

func Map[T, U any](i Iterator[T], fn func(T) U) Iterator[U] {
	return &mapIterator[T, U]{iter: i, mapFn: fn}
}

type mapIterator[T, U] struct {
	iter Iterator[T]
	mapFn func(T) U
}

func (m *mapIterator[T, U]) Next() option.Optional[U] {
	item := m.iter.Next()
	if item.IsSome() {
		return option.Some[U](fn(item))
	}
	return option.None[U]()
}

func Flatten[T any, U Iterator[T]](i Iterator[U]) Iterator[T] {
	return nil
}
