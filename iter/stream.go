package iter

import "github.com/go-board/flow/option"

// Stream wrap a low-layer Iterator for chaining call.
// With Stream, we can do below function call.
// 	stream.
//	  TakeWhile(func(i int) bool { return i > 0 }).
//	  StepBy(3).
//	  Filter(func(idx int, value int) bool { return value%3 == 0 })
type Stream[T any] struct { iter Iterator[T] }

// Streaming create new Stream.
func Streaming[T any](i Iterator[T]) *Stream { return &Stream[T]{iter: i} }

// TakeWhile return a new stream which emit element match the predicate function
func (s *Stream[T]) TakeWhile(predicate func(T) bool) *Stream { panic("Not implemented!") }

// DropWhile return a new stream which omit element match the predicate function
func (s *Stream[T]) DropWhile(predicate func(T) bool) *Stream { panic("Not implemented!") }

func (s *Stream[T]) StepBy(n int) *Stream { panic("Not implemented!") }
func (s *Stream[T]) Skip(n int) *Stream { panic("Not implemented!") }
func (s *Stream[T]) Limit(n int) *Stream { panic("Not implemented!") }
func (s *Stream[T]) ForEach(fn func(T)) { panic("Not implemented!") }
func (s *Stream[T]) Size() int { panic("Not implemented!") }
func (s *Stream[T]) Concat(n *Stream[T]) *Stream { panic("Not implemented!") }
func (s *Stream[T]) SortFn(fn func(T, T) int) *Stream { panic("Not implemented!") }
func (s *Stream[T]) Distinct(fn func(T, T) int) *Stream { panic("Not implemented!") }
func (s *Stream[T]) Nth(n int) option.Optional[T] { panic("Not implemented!") }
func (s *Stream[T]) First() option.Optional[T] { panic("Not implemented!") }
func (s *Stream[T]) Last() option.Optional[T] { panic("Not implemented!") }
func (s *Stream[T]) All(predicate func(T) bool) bool { panic("Not implemented!") }
func (s *Stream[T]) Any(predicate func(T) bool) bool { panic("Not implemented!") }
func (s *Stream[T]) Non(predicate func(T) bool) bool { panic("Not implemented!") }
func (s *Stream[T]) Count(predicate func(T) bool) bool { panic("Not implemented!") }
