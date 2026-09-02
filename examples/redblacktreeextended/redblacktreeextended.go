package redblacktreeextended

import (
	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
)

type RedBlackTreeExtended[K comparable, V any] struct {
	*rbt.Tree[K, V]
}

func (tree *RedBlackTreeExtended[K, V]) GetMin() (value V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (tree *RedBlackTreeExtended[K, V]) GetMax() (value V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (tree *RedBlackTreeExtended[K, V]) RemoveMin() (value V, deleted bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (tree *RedBlackTreeExtended[K, V]) RemoveMax() (value V, deleted bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (tree *RedBlackTreeExtended[K, V]) getMinFromNode(node *rbt.Node[K, V]) (foundNode *rbt.Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tree *RedBlackTreeExtended[K, V]) getMaxFromNode(node *rbt.Node[K, V]) (foundNode *rbt.Node[K, V], found bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func print(tree *RedBlackTreeExtended[int, string]) { _ = "STUB: not implemented"; return }

func main() {
	tree := RedBlackTreeExtended[int, string]{rbt.New[int, string]()}

	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")

	print(&tree)

	tree.RemoveMin()
	tree.RemoveMax()
	tree.RemoveMin()
	tree.RemoveMax()

	print(&tree)

}
