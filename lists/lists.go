package lists

import (
	"github.com/emirpasic/gods/v2/containers"
	"github.com/emirpasic/gods/v2/utils"
)

type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	containers.Container[T]
}
