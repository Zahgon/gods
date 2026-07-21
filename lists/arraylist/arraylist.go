package arraylist

import (
	"github.com/emirpasic/gods/v2/lists"
	"github.com/emirpasic/gods/v2/utils"
)

var _ lists.List[int] = (*List[int])(nil)

type List[T comparable] struct {
	elements []T
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New[T comparable](values ...T) *List[T] { _ = "STUB: not implemented"; return nil }

func (list *List[T]) Add(values ...T) { _ = "STUB: not implemented"; return }

func (list *List[T]) Get(index int) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (list *List[T]) Remove(index int) { _ = "STUB: not implemented"; return }

func (list *List[T]) Contains(values ...T) bool { _ = "STUB: not implemented"; return false }

func (list *List[T]) Values() []T { _ = "STUB: not implemented"; return nil }

func (list *List[T]) IndexOf(value T) int { _ = "STUB: not implemented"; return 0 }

func (list *List[T]) Empty() bool { _ = "STUB: not implemented"; return false }

func (list *List[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func (list *List[T]) Clear() { _ = "STUB: not implemented"; return }

func (list *List[T]) Sort(comparator utils.Comparator[T]) { _ = "STUB: not implemented"; return }

func (list *List[T]) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (list *List[T]) Insert(index int, values ...T) { _ = "STUB: not implemented"; return }

func (list *List[T]) Set(index int, value T) { _ = "STUB: not implemented"; return }

func (list *List[T]) String() string { _ = "STUB: not implemented"; return "" }

func (list *List[T]) withinRange(index int) bool { _ = "STUB: not implemented"; return false }

func (list *List[T]) resize(len, cap int) { _ = "STUB: not implemented"; return }

func (list *List[T]) growBy(n int) { _ = "STUB: not implemented"; return }

func (list *List[T]) shrink() { _ = "STUB: not implemented"; return }
