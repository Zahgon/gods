package arraylist

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

func (list *List[T]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (list *List[T]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (list *List[T]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (list *List[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
