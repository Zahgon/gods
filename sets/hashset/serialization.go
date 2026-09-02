package hashset

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Set[int])(nil)
var _ containers.JSONDeserializer = (*Set[int])(nil)

func (set *Set[T]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (set *Set[T]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (set *Set[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
