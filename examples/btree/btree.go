package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/trees/btree"
)

func main() {
	tree := btree.New[int, string](3)

	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")

	fmt.Println(tree)

	_ = tree.Values()
	_ = tree.Keys()

	tree.Remove(2)
	fmt.Println(tree)

	tree.Clear()
	tree.Empty()
	tree.Size()

	tree.Height()
	tree.Left()
	tree.LeftKey()
	tree.LeftValue()
	tree.Right()
	tree.RightKey()
	tree.RightValue()
}
