package main

import "github.com/emirpasic/gods/v2/sets/hashset"

func main() {
	set := hashset.New[int]()
	set.Add(1)
	set.Add(2, 2, 3, 4, 5)
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
