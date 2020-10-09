package controlflow

// Conditional return left if expression is ture, else return right.
//  for example:
// 		a > b ? a : b
//		strings.Contains("a", "b") ? 3 : 1
func Conditional[T any](expression bool, left T, right T) T {
	if expression {
		return left
	}
	return right
}
