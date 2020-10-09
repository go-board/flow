package assert

import "fmt"

func Assert(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}

func Assertf(expr bool, format string, args ...interface{}) {
	Assert(expr, fmt.Sprintf(format, args...))
}

func Equal[T comparable](left, right T, msg string) {
	if left != right {
		panic(msg)
	}
}

func Equalf[T comparable](left, right T, format string, args ...interface{}) {
	Equal(left, right, fmt.Sprintf(format, args...))
}

func NotEqual[T comparable](left, right T, msg string) {
	if left == right {
		panic(msg)
	}
}

func NotEqualf[T comparable](left, right T, format string, args ...interface{}) {
	NotEqual(left, right, fmt.Sprintf(format, args...))
}
