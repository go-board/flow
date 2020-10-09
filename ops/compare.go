package ops

type Order int

const (
	OrderLess Order = iota - 1
	OrderEqual 
	OrderGreat
)

type Comparable[T any] interface {
	Compare(T) Order
}

type CompareFunc[T any] func(T, T) Order
