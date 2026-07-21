package arrayqueue

import (
	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/queues"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

type Queue[T comparable] struct {
	list *arraylist.List[T]
}

func New[T comparable]() *Queue[T] { _ = "STUB: not implemented"; return nil }

func (queue *Queue[T]) Enqueue(value T) { _ = "STUB: not implemented"; return }

func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (queue *Queue[T]) Peek() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (queue *Queue[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (queue *Queue[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (queue *Queue[T]) Clear() { _ = "STUB: not implemented"; return }

func (queue *Queue[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (queue *Queue[T]) String() string { _ = "STUB: not implemented"; return "" }

func (queue *Queue[T]) withinRange(index int) bool { _ = "STUB: not implemented"; return false }
