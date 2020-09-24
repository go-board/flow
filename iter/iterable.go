package iter

type Iterable[T any] interface {
	Iter() Iterator[T]
}
