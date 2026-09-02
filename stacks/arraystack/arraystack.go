package arraystack

import (
	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

type Stack[T comparable] struct {
	list *arraylist.List[T]
}

func New[T comparable]() *Stack[T] { _ = "STUB: not implemented"; return nil }

func (stack *Stack[T]) Push(value T) { _ = "STUB: not implemented"; return }

func (stack *Stack[T]) Pop() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (stack *Stack[T]) Peek() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (stack *Stack[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (stack *Stack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (stack *Stack[T]) Clear() { _ = "STUB: not implemented"; return }

func (stack *Stack[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (stack *Stack[T]) String() string { _ = "STUB: not implemented"; return "" }

func (stack *Stack[T]) withinRange(index int) bool { _ = "STUB: not implemented"; return false }
