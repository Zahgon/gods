package treebidimap

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Map[string, int])(nil)
var _ containers.JSONDeserializer = (*Map[string, int])(nil)

func (m *Map[K, V]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Map[K, V]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
