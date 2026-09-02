package binaryheap

import (
	"github.com/emirpasic/gods/v2/containers"
)

var _ containers.JSONSerializer = (*Heap[int])(nil)
var _ containers.JSONDeserializer = (*Heap[int])(nil)

func (heap *Heap[T]) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (heap *Heap[T]) FromJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (heap *Heap[T]) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (heap *Heap[T]) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
