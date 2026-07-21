package maps

import "github.com/emirpasic/gods/v2/containers"

type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
}

type BidiMap[K comparable, V comparable] interface {
	GetKey(value V) (key K, found bool)

	Map[K, V]
}
