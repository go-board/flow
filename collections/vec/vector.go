package vec

type Vector[T any] []T

func From[T any](elems ...T) Vector[T] {
	return Vector[T](elems)
}

func FromSlice[T any](slice []T) Vector[T] {
	return From(slice...)
}

func (v Vector[T]) Swap(i, j int) bool {
	if i == j {
		return false
	}
	if i < 0 || i >= len(v) {
		return false
	}
	if i < 0 || i >= len(v) {
		return false
	}
	v[i], v[j] = v[j], v[i]
}

func (v Vector[T]) Prepend(elems ...T) {}
func (v Vector[T]) Append(elems ...T) {}
func (v Vector[T]) Insert(idx int, before bool, elems ...T) {}
func (v Vector[T]) Remove(idx int) {}
func (v Vector[T]) Splice(idx int, n int) {}

func (v Vector[T]) Iter() iter.Iterator[T] {
	return &vecIterator[t]{}
}

type vecIterator[T any] struct {}

func (vi *vecIterator[T]) Next() option.Optional[T] {

}

type ImmutableVector[T any] []T {

}

func (v ImmutableVector[T]) Mutator() Vector[T] {
	panic("Not implemented!")
}

func ImmutableOf[T any](elems ...T) ImmutableVector[T] {
	panic("Not implemented!")
}

type immVecIterator[T any] struct {

}

func (iv *immVecIterator[T]) Next() option.Optional[T] {
	
}
