package stacks

import "github.com/emirpasic/gods/v2/containers"

type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
}
