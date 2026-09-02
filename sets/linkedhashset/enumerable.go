package linkedhashset

import "github.com/emirpasic/gods/v2/containers"

var _ containers.EnumerableWithIndex[int] = (*Set[int])(nil)

func (set *Set[T]) Each(f func(index int, value T)) { _ = "STUB: not implemented"; return }

func (set *Set[T]) Map(f func(index int, value T) T) *Set[T] { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) Select(f func(index int, value T) bool) *Set[T] {
	_ = "STUB: not implemented"
	return nil
}

func (set *Set[T]) Any(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (set *Set[T]) All(f func(index int, value T) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (set *Set[T]) Find(f func(index int, value T) bool) (int, T) {
	_ = "STUB: not implemented"
	return 0, *new(T)
}
