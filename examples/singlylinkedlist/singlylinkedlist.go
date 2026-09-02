package main

import (
	"cmp"

	sll "github.com/emirpasic/gods/v2/lists/singlylinkedlist"
)

func main() {
	list := sll.New[string]()
	list.Add("a")
	list.Append("b")
	list.Prepend("c")
	list.Sort(cmp.Compare[string])
	_, _ = list.Get(0)
	_, _ = list.Get(100)
	_ = list.Contains("a", "b", "c")
	_ = list.Contains("a", "b", "c", "d")
	list.Remove(2)
	list.Remove(1)
	list.Remove(0)
	list.Remove(0)
	_ = list.Empty()
	_ = list.Size()
	list.Add("a")
	list.Clear()
}
