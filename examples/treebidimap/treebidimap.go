package main

import (
	"github.com/emirpasic/gods/v2/maps/treebidimap"
)

func main() {
	m := treebidimap.New[int, string]()
	m.Put(1, "x")
	m.Put(3, "b")
	m.Put(1, "a")
	m.Put(2, "b")
	_, _ = m.GetKey("a")
	_, _ = m.Get(2)
	_, _ = m.Get(3)
	_ = m.Values()
	_ = m.Keys()
	m.Remove(1)
	m.Clear()
	m.Empty()
	m.Size()
}
