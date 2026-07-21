package main

import "github.com/emirpasic/gods/v2/sets/linkedhashset"

func main() {
	set := linkedhashset.New[int]()
	set.Add(5)
	set.Add(4, 4, 3, 2, 1)
	set.Remove(4)
	set.Remove(2, 3)
	set.Contains(1)
	set.Contains(1, 5)
	set.Contains(1, 6)
	_ = set.Values()
	set.Clear()
	set.Empty()
	set.Size()
}
