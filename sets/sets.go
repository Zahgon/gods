package sets

import (
	"github.com/emirpasic/gods/v2/containers"
)

type Set[T comparable] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool

	containers.Container[T]
}
