package main

import cb "github.com/emirpasic/gods/v2/queues/circularbuffer"

func main() {
	queue := cb.New[int](3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	_ = queue.Values()
	queue.Enqueue(3)
	_, _ = queue.Peek()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	queue.Enqueue(1)
	queue.Clear()
	queue.Empty()
	_ = queue.Size()
}
