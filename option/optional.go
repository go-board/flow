package option

type Optional[T any] struct {
	data *T
}

func Some[T any](t T) Optional[T] {
	return Optional[T]{data: &t}
}
func None[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) IsSome() bool { return o.data != nil }
func (o Optional[T]) IsNone() bool { return !o.IsSome() }

func (o Optional[T]) Value() T { return *o.data }
