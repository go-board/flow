package pair

type Pair[A, B any] struct {
	a A
	b B
}

func Make[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a: a, b: b}
}

type Either[Left, Right any] struct {

}

type Tuple2[T1, T2 any] struct {
	t1 T1
	t2 T2
}

type Tuple3[T1, T2, T3 any] struct {

}

type Tuple4[T1, T2, T3, T4 any] struct {}

type Tuple5[T1, T2, T3, T4, T5 any] struct {}

type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {

}


