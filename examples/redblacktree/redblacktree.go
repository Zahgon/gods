package main

import (
	"fmt"

	rbt "github.com/emirpasic/gods/v2/trees/redblacktree"
)

func main() {
	tree := rbt.New[int, string]()

	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")

	fmt.Println(tree)

	_ = tree.Values()
	_ = tree.Keys()

	tree.Remove(2)
	fmt.Println(tree)

	tree.Clear()
	tree.Empty()
	tree.Size()
}
