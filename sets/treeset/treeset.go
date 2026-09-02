package treeset

import (
	"cmp"

	"github.com/emirpasic/gods/v2/sets"
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
	"github.com/emirpasic/gods/v2/utils"
)

var _ sets.Set[int] = (*Set[int])(nil)

type Set[T comparable] struct {
	tree *rbt.Tree[T, struct{}]
}

var itemExists = struct{}{}

func New[T cmp.Ordered](values ...T) *Set[T] { _ = "STUB: not implemented"; return nil }

func NewWith[T comparable](comparator utils.Comparator[T], values ...T) *Set[T] {
	_ = "STUB: not implemented"
	return nil
}

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
