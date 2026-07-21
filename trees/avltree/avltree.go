package avltree

import (
	"cmp"

	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

var _ trees.Tree[int] = (*Tree[string, int])(nil)

type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	Comparator utils.Comparator[K]
	size       int
}

type Node[K comparable, V any] struct {
	Key      K
	Value    V
	Parent   *Node[K, V]
	Children [2]*Node[K, V]
	b        int8
}

func New[K cmp.Ordered, V any]() *Tree[K, V] { _ = "STUB: not implemented"; return nil }

func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (tree *Tree[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) Get(key K) (value V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) Empty() bool { _ = "STUB: not implemented"; return false }

func (tree *Tree[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (n *Node[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Values() []V { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Left() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Right() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tree *Tree[K, V]) Ceiling(key K) (floor *Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tree *Tree[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func (n *Node[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func (tree *Tree[K, V]) put(key K, value V, p *Node[K, V], qp **Node[K, V]) bool {
	_ = "STUB: not implemented"
	return false
}

func (tree *Tree[K, V]) remove(key K, qp **Node[K, V]) bool {
	_ = "STUB: not implemented"
	return false
}

func removeMin[K comparable, V any](qp **Node[K, V], minKey *K, minVal *V) bool {
	_ = "STUB: not implemented"
	return false
}

func putFix[K comparable, V any](c int8, t **Node[K, V]) bool {
	_ = "STUB: not implemented"
	return false
}

func removeFix[K comparable, V any](c int8, t **Node[K, V]) bool {
	_ = "STUB: not implemented"
	return false
}

func singlerot[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func doublerot[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func rotate[K comparable, V any](c int8, s *Node[K, V]) *Node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (tree *Tree[K, V]) bottom(d int) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *Node[K, V]) Prev() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *Node[K, V]) Next() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *Node[K, V]) walk1(a int) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func output[K comparable, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	_ = "STUB: not implemented"
	return
}
