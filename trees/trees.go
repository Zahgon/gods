package trees

import "github.com/emirpasic/gods/v2/containers"

type Tree[V any] interface {
	containers.Container[V]
}
