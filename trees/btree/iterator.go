package btree

import "github.com/emirpasic/gods/v2/containers"

var _ containers.ReverseIteratorWithKey[string, int] = (*Iterator[string, int])(nil)

type Iterator[K comparable, V any] struct {
	tree     *Tree[K, V]
	node     *Node[K, V]
	entry    *Entry[K, V]
	position position
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

func (tree *Tree[K, V]) Iterator() *Iterator[K, V] { _ = "STUB: not implemented"; return nil }

func (iterator *Iterator[K, V]) Next() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) Prev() bool { _ = "STUB: not implemented"; return false }

func (iterator *Iterator[K, V]) Value() V { _ = "STUB: not implemented"; return *new(V) }

func (iterator *Iterator[K, V]) Key() K { _ = "STUB: not implemented"; return *new(K) }

func (iterator *Iterator[K, V]) Node() *Node[K, V] { _ = "STUB: not implemented"; return nil }

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
