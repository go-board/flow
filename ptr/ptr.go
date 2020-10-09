package ptr

// Ptr
func Ptr[T any](t T) *T { return &t }
