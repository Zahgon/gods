package linkedhashmap

import "github.com/emirpasic/gods/v2/containers"

var _ containers.EnumerableWithKey[string, int] = (*Map[string, int])(nil)

func (m *Map[K, V]) Each(f func(key K, value V)) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Map(f func(key1 K, value1 V) (K, V)) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map[K, V]) Select(f func(key K, value V) bool) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *Map[K, V]) Any(f func(key K, value V) bool) bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V]) All(f func(key K, value V) bool) bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V]) Find(f func(key K, value V) bool) (k K, v V) {
	_ = "STUB: not implemented"
	return *new(K), *new(V)
}
