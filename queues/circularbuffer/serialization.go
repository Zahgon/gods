package circularbuffer

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Queue[int])(nil)
var _ containers.JSONDeserializer = (*Queue[int])(nil)

func (queue *Queue[T]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (queue *Queue[T]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (queue *Queue[T]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (queue *Queue[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
