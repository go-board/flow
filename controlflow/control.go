package controlflow

type Condition[T any] bool

func New[T any](expr bool) Condition[T] {
	return Condition[T](expr)
}

func (cond Condition[T]) Return(left, right T) T {
	if bool(cond) {
		return left
	}
	return right
}

type ConditionExpr[T any] func() bool

func NewExpr[T any](exprFn func() bool) ConditionExpr[T] {
	return ConditionExpr[T](exprFn)
}

func (cond ConditionExpr[T]) Return(left, right T) T {
	if cond() {
		return left
	}
	return right
}

// Conditional return left if expression is ture, else return right.
//  for example:
// 		a > b ? a : b
//		strings.Contains("a", "b") ? 3 : 1
func Conditional[T any](expression bool, left, right T) T {
	if expression {
		return left
	}
	return right
}

func ConditionalExpr(expression bool, left func (), right func()) {
	if expression {
		left()
	} else {
		right()
	}
}
