package main

import (
	"cmp"

	"github.com/emirpasic/gods/v2/trees/binaryheap"
)

func main() {

	heap := binaryheap.New[int]()
	heap.Push(2)
	heap.Push(3)
	heap.Push(1)
	heap.Values()
	_, _ = heap.Peek()
	_, _ = heap.Pop()
	_, _ = heap.Pop()
	_, _ = heap.Pop()
	_, _ = heap.Pop()
	heap.Push(1)
	heap.Clear()
	heap.Empty()
	heap.Size()

	inverseIntComparator := func(a, b int) int {
		return -cmp.Compare(a, b)
	}
	heap = binaryheap.NewWith(inverseIntComparator)
	heap.Push(2)
	heap.Push(3)
	heap.Push(1)
	heap.Values()
}
