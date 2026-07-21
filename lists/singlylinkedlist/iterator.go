package singlylinkedlist

import "github.com/emirpasic/gods/v2/containers"

var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)

type Iterator[T comparable] struct {
	list    *List[T]
	index   int
	element *element[T]
}

func (list *List[T]) Iterator() *Iterator[T] { _ = "STUB: not implemented"; return nil }

func (iterator *Iterator[T]) Next() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) Value() T { _ = "STUB: not implemented"; return *new(T) }

func (iterator *Iterator[T]) Index() int { _ = "STUB: not implemented"; return 0 }

func (iterator *Iterator[T]) Begin() { _ = "STUB: not implemented"; return }

func (iterator *Iterator[T]) First() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[T]) NextTo(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}
