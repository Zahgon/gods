package arraylist

import "github.com/emirpasic/gods/v2/containers"

var _ containers.EnumerableWithIndex[int] = (*List[int])(nil)

func (list *List[T]) Each(f func(index int, value T)) { _ = "STUB: not implemented"; return }

func (list *List[T]) Map(f func(index int, value T) T) *List[T] {
	_ = "STUB: not implemented"
	return nil
}

func (list *List[T]) Select(f func(index int, value T) bool) *List[T] {
	_ = "STUB: not implemented"
	return nil
}

func (list *List[T]) Any(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (list *List[T]) All(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (list *List[T]) Find(f func(index int, value T) bool) (int, T) {
	_ = "STUB: not implemented"
	return 0, *new(T)
}
