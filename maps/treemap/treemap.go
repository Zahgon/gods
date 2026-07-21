package treemap

import (
	"cmp"

	"github.com/emirpasic/gods/v2/maps"
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
	"github.com/emirpasic/gods/v2/utils"
)

var _ maps.Map[string, int] = (*Map[string, int])(nil)

type Map[K comparable, V any] struct {
	tree *rbt.Tree[K, V]
}

func New[K cmp.Ordered, V any]() *Map[K, V] { _ = "STUB: not implemented"; return nil }

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Get(key K) (value V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Empty() bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (m *Map[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) Values() []V { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Min() (key K, value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (m *Map[K, V]) Max() (key K, value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (m *Map[K, V]) Floor(key K) (foundKey K, foundValue V, ok bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (m *Map[K, V]) Ceiling(key K) (foundKey K, foundValue V, ok bool) {
	_ = "STUB: not implemented"
	return *new(K), *new(V), false
}

func (m *Map[K, V]) String() string { _ = "STUB: not implemented"; return "" }
