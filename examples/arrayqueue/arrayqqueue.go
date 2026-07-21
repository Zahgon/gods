package main

import aq "github.com/emirpasic/gods/v2/queues/arrayqueue"

func main() {
	queue := aq.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	_ = queue.Values()
	_, _ = queue.Peek()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	queue.Enqueue(1)
	queue.Clear()
	queue.Empty()
	_ = queue.Size()
}
