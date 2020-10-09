package ops

type Add[T any] interface {
	Add(T) T
}

type Sub[T any] interface {
	Sub(T) T
}

type Mul[T any] interface {
	Mul(T) T
}

type Div[T any] interface {
	Div(T) T
}

type Rem[T any] interface {
	Rem(T) T
}
