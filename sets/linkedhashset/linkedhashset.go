package linkedhashset

import (
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
	"github.com/emirpasic/gods/v2/sets"
)

var _ sets.Set[int] = (*Set[int])(nil)

type Set[T comparable] struct {
	table    map[T]struct{}
	ordering *doublylinkedlist.List[T]
}

var itemExists = struct{}{}

func New[T comparable](values ...T) *Set[T] { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) Add(items ...T) { _ = "STUB: not implemented"; return }

func (set *Set[T]) Remove(items ...T) { _ = "STUB: not implemented"; return }

func (set *Set[T]) Contains(items ...T) bool { _ = "STUB: not implemented"; return false }

func (set *Set[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (set *Set[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (set *Set[T]) Clear() { _ = "STUB: not implemented"; return }

func (set *Set[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) String() string { _ = "STUB: not implemented"; return "" }

func (set *Set[T]) Intersection(another *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) Union(another *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) Difference(another *Set[T]) *Set[T] { _ = "STUB: not implemented"; return nil }
