package arrayqueue

import "github.com/emirpasic/gods/v2/containers"

var _ containers.ReverseIteratorWithIndex[int] = (*Iterator[int])(nil)

type Iterator[T comparable] struct {
	queue *Queue[T]
	index int
}

func (queue *Queue[T]) Iterator() *Iterator[T] { _ = "STUB: not implemented"; return nil }

func (iterator *Iterator[T]) Next() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) Prev() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (iterator *Iterator[T]) Index() int { _ = "STUB: not implemented"; return 0 }

func (iterator *Iterator[T]) Begin() { _ = "STUB: not implemented"; return }

func (iterator *Iterator[T]) End() { _ = "STUB: not implemented"; return }

func (iterator *Iterator[T]) First() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) Last() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (iterator *Iterator[T]) PrevTo(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}
