package hashbidimap

import (
	"github.com/emirpasic/gods/v2/maps"
	"github.com/emirpasic/gods/v2/maps/hashmap"
)

var _ maps.BidiMap[string, int] = (*Map[string, int])(nil)

type Map[K, V comparable] struct {
	forwardMap hashmap.Map[K, V]
	inverseMap hashmap.Map[V, K]
}

func New[K, V comparable]() *Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Get(key K) (value V, found bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

func (m *Map[K, V]) GetKey(value V) (key K, found bool) {
	_ = "STUB: not implemented"
	return *new(K), false
}

func (m *Map[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) Empty() bool { _ = "STUB: not implemented"; return false }

func (m *Map[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (m *Map[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) Values() []V { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) Clear() { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) String() string { _ = "STUB: not implemented"; return "" }
