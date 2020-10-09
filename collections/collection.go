package collections

import (
	"github.com/go-board/flow/iter"
)

type Sequence[T any] interface {
	iter.Iterable[T]
	ImmutableSequence[T]

	Add(elems ...T)
	Remove(elems ...T)
}

type ImmutableSequence[T any] interface {
	Size() int
}

type Map[K, V any] interface {
	Put(k K, v V)
	Remove(k K)
}

type ImmutableMap[K, V any] interface {
	HasKey(k K) bool
	KeySet() ImmutableSet[K]
	ValueSet() ImmutableSet[V]
}

type Set[T any] interface {
	Sequence[T]
	ImmutableSet[T]
}

type ImmutableSet[T any] interface {
	ImmutableSequence[T]

	Contains(cmp func(T, T) int, elem T) bool
	ContainsAny(elems ...T) bool
	ContainsAll(elems ...T) bool
	Union(o ImmutableSet[T]) ImmutableSet[T]
	InterSection(o ImmutableSet[T]) ImmutableSet[T]
	Difference(o ImmutableSet[T]) ImmutableSet[T]
	Sorted() vec.Vector[T]
}
