package btree

import (
	"bytes"
	"cmp"

	"github.com/emirpasic/gods/v2/trees"
	"github.com/emirpasic/gods/v2/utils"
)

var _ trees.Tree[int] = (*Tree[string, int])(nil)

type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	Comparator utils.Comparator[K]
	size       int
	m          int
}

type Node[K comparable, V any] struct {
	Parent   *Node[K, V]
	Entries  []*Entry[K, V]
	Children []*Node[K, V]
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func New[K cmp.Ordered, V any](order int) *Tree[K, V] { _ = "STUB: not implemented"; return nil }

func NewWith[K comparable, V any](order int, comparator utils.Comparator[K]) *Tree[K, V] {
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

func (tree *Tree[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) Height() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) Left() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) LeftKey() interface{} { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) LeftValue() interface{} { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) Right() *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) RightKey() interface{} { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) RightValue() interface{} { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func (entry *Entry[K, V]) String() string { _ = "STUB: not implemented"; return "" }

func (tree *Tree[K, V]) output(buffer *bytes.Buffer, node *Node[K, V], level int) {
	_ = "STUB: not implemented"
	return
}

func (node *Node[K, V]) height() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) isLeaf(node *Node[K, V]) bool { _ = "STUB: not implemented"; return false }

func (tree *Tree[K, V]) isFull(node *Node[K, V]) bool { _ = "STUB: not implemented"; return false }

func (tree *Tree[K, V]) shouldSplit(node *Node[K, V]) bool { _ = "STUB: not implemented"; return false }

func (tree *Tree[K, V]) maxChildren() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) minChildren() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) maxEntries() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) minEntries() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) middle() int { _ = "STUB: not implemented"; return 0 }

func (tree *Tree[K, V]) search(node *Node[K, V], key K) (index int, found bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (tree *Tree[K, V]) searchRecursively(startNode *Node[K, V], key K) (node *Node[K, V], index int, found bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

func (tree *Tree[K, V]) insert(node *Node[K, V], entry *Entry[K, V]) (inserted bool) {
	_ = "STUB: not implemented"
	return false
}

func (tree *Tree[K, V]) insertIntoLeaf(node *Node[K, V], entry *Entry[K, V]) (inserted bool) {
	_ = "STUB: not implemented"
	return false
}

func (tree *Tree[K, V]) insertIntoInternal(node *Node[K, V], entry *Entry[K, V]) (inserted bool) {
	_ = "STUB: not implemented"
	return false
}

func (tree *Tree[K, V]) split(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) splitNonRoot(node *Node[K, V]) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) splitRoot() { _ = "STUB: not implemented"; return }

func setParent[K comparable, V any](nodes []*Node[K, V], parent *Node[K, V]) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) left(node *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) right(node *Node[K, V]) *Node[K, V] { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) leftSibling(node *Node[K, V], key K) (*Node[K, V], int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (tree *Tree[K, V]) rightSibling(node *Node[K, V], key K) (*Node[K, V], int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (tree *Tree[K, V]) delete(node *Node[K, V], index int) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) rebalance(node *Node[K, V], deletedKey K) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) prependChildren(fromNode *Node[K, V], toNode *Node[K, V]) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) appendChildren(fromNode *Node[K, V], toNode *Node[K, V]) {
	_ = "STUB: not implemented"
	return
}

func (tree *Tree[K, V]) deleteEntry(node *Node[K, V], index int) { _ = "STUB: not implemented"; return }

func (tree *Tree[K, V]) deleteChild(node *Node[K, V], index int) { _ = "STUB: not implemented"; return }
