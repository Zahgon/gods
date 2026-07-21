package linkedhashmap

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/lists/doublylinkedlist"
)

var _ containers.ReverseIteratorWithKey[string, int] = (*Iterator[string, int])(nil)

type Iterator[K comparable, V any] struct {
	iterator doublylinkedlist.Iterator[K]
	table    map[K]V
}

func (m *Map[K, V]) Iterator() *Iterator[K, V] { _ = "STUB: not implemented"; return nil }

func (iterator *Iterator[K, V]) Next() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) Prev() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) Value() V { _ = "STUB: not implemented"; return *new(V) }

func (iterator *Iterator[K, V]) Key() K { _ = "STUB: not implemented"; return *new(K) }

func (iterator *Iterator[K, V]) Begin() { _ = "STUB: not implemented"; return }

func (iterator *Iterator[K, V]) End() { _ = "STUB: not implemented"; return }

func (iterator *Iterator[K, V]) First() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) Last() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) NextTo(f func(key K, value V) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (iterator *Iterator[K, V]) PrevTo(f func(key K, value V) bool) bool {
	_ = "STUB: not implemented"
	return false
}
