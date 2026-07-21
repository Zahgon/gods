package main

import (
	pq "github.com/emirpasic/gods/v2/queues/priorityqueue"
)

type Element struct {
	name     string
	priority int
}

func byPriority(a, b Element) int { _ = "STUB: not implemented"; return 0 }

func main() {
	a := Element{name: "a", priority: 1}
	b := Element{name: "b", priority: 2}
	c := Element{name: "c", priority: 3}

	queue := pq.NewWith(byPriority)
	queue.Enqueue(a)
	queue.Enqueue(c)
	queue.Enqueue(b)
	_ = queue.Values()
	_, _ = queue.Peek()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	queue.Clear()
	_ = queue.Empty()
	_ = queue.Size()
}
