package priorityqueue

import (
	"cmp"

	"github.com/emirpasic/gods/v2/queues"
	"github.com/emirpasic/gods/v2/trees/binaryheap"
	"github.com/emirpasic/gods/v2/utils"
)

var _ queues.Queue[int] = (*Queue[int])(nil)

type Queue[T comparable] struct {
	heap       *binaryheap.Heap[T]
	Comparator utils.Comparator[T]
}

func New[T cmp.Ordered]() *Queue[T] { _ = "STUB: not implemented"; return nil }

func NewWith[T comparable](comparator utils.Comparator[T]) *Queue[T] {
	_ = "STUB: not implemented"
	return nil
}

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
