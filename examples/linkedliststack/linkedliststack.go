package main

import lls "github.com/emirpasic/gods/v2/stacks/linkedliststack"

func main() {
	stack := lls.New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Values()
	_, _ = stack.Peek()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	_, _ = stack.Pop()
	stack.Push(1)
	stack.Clear()
	stack.Empty()
	stack.Size()
}
