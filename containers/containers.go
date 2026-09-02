package containers

import (
	"cmp"

	"github.com/emirpasic/gods/v2/utils"
)

type Container[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}

func GetSortedValues[T cmp.Ordered](container Container[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

func GetSortedValuesFunc[T any](container Container[T], comparator utils.Comparator[T]) []T {
	_ = "STUB: not implemented"
	return nil
}
