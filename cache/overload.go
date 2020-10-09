package cache


type Vec[T any] struct{}

func (v Vec[T]) + (b Vec[T]) Vec[T] {}

func (v *Vec[T]) += (v Vec[T]) {}

func (v *Vec[T]) Append (item T) {

}

func (v *time.Time) += (d time.Duration) {}

func (v time.Time) + (d time.Duration) time.Time {}

func (v time.Time) !() time.Time {

}

operator func (v *time.Time) !() {

}

func call() {
	time.Now()
}
