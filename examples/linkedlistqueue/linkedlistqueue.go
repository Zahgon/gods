package main

import llq "github.com/emirpasic/gods/v2/queues/linkedlistqueue"

func main() {
	queue := llq.New[int]()
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
