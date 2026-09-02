package queues

import "github.com/emirpasic/gods/v2/containers"

type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
