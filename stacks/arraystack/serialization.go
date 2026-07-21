package arraystack

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Stack[int])(nil)
var _ containers.JSONDeserializer = (*Stack[int])(nil)

func (stack *Stack[T]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (stack *Stack[T]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (stack *Stack[T]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (stack *Stack[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
