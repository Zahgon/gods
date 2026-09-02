package btree

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Tree[string, int])(nil)
var _ containers.JSONDeserializer = (*Tree[string, int])(nil)

func (tree *Tree[K, V]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (tree *Tree[K, V]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (tree *Tree[K, V]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
