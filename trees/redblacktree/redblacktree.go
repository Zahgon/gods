package redblacktree

import (
	"cmp"

	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

var _ trees.Tree[int] = (*Tree[string, int])(nil)

type color bool

const (
	black, red color = true, false
)

type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	size       int
	Comparator utils.Comparator[K]
}

type Node[K comparable, V any] struct {
	Key    K
	Value  V
	color  color
	Left   *Node[K, V]
	Right  *Node[K, V]
	Parent *Node[K, V]
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

func (node *Node[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Values() []V { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Left() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Right() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Floor(key K) (floor *Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tree *Tree[K, V]) Ceiling(key K) (ceiling *Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tree *Tree[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func (node *Node[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func output[K comparable, V any](node *Node[K, V], prefix string, isTail bool, str *string) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) lookup(key K) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (node *Node[K, V]) grandparent() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (node *Node[K, V]) uncle() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (node *Node[K, V]) sibling() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) rotateLeft(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) rotateRight(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) replaceNode(old *Node[K, V], new *Node[K, V]) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) insertCase1(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) insertCase2(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) insertCase3(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) insertCase4(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) insertCase5(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (node *Node[K, V]) maximumNode() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) deleteCase1(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteCase2(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteCase3(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteCase4(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteCase5(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteCase6(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func nodeColor[K comparable, V any](node *Node[K, V]) color {
	_ = "STUB: not implemented"
	return *new(color)
}
