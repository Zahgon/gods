package binaryheap

import (
	"cmp"

	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

var _ trees.Tree[int] = (*Heap[int])(nil)

type Heap[T comparable] struct {
	list       *arraylist.List[T]
	Comparator utils.Comparator[T]
}

func New[T cmp.Ordered]() *Heap[T] { _ = "STUB: not implemented"; return nil }

func NewWith[T comparable](comparator utils.Comparator[T]) *Heap[T] {
	_ = "STUB: not implemented"
	return nil
}

func (heap *Heap[T]) Push(values ...T) { _ = "STUB: not implemented"; return }

func (heap *Heap[T]) Pop() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (heap *Heap[T]) Peek() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (heap *Heap[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (heap *Heap[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (heap *Heap[T]) Clear() { _ = "STUB: not implemented"; return }

func (heap *Heap[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (heap *Heap[T]) String() string { _ = "STUB: not implemented"; return "" }

func (heap *Heap[T]) bubbleDown() { _ = "STUB: not implemented"; return }

func (heap *Heap[T]) bubbleDownIndex(index int) { _ = "STUB: not implemented"; return }

func (heap *Heap[T]) bubbleUp() { _ = "STUB: not implemented"; return }

func (heap *Heap[T]) withinRange(index int) bool { _ = "STUB: not implemented"; return false }
