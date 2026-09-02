package main

import "github.com/emirpasic/gods/v2/maps/treemap"

func main() {
	m := treemap.New[int, string]()
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")
	_, _ = m.Get(2)
	_, _ = m.Get(3)
	_ = m.Values()
	_ = m.Keys()
	m.Remove(1)
	m.Clear()
	m.Empty()
	m.Size()
}
