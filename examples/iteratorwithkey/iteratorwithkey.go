package main

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/maps/treemap"
)

func main() {
	m := treemap.New[int, string]()
	m.Put(0, "a")
	m.Put(1, "b")
	m.Put(2, "c")
	it := m.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]")
	}

	fmt.Print("\nForward iteration (again)\n")
	for it.Begin(); it.Next(); {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]")
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]")
	}

	fmt.Print("\nBackward iteration (again)\n")
	for it.End(); it.Prev(); {
		key, value := it.Key(), it.Value()
		fmt.Print("[", key, ":", value, "]")
	}

	if it.First() {
		fmt.Print("\nFirst key: ", it.Key())
		fmt.Print("\nFirst value: ", it.Value())
	}

	if it.Last() {
		fmt.Print("\nLast key: ", it.Key())
		fmt.Print("\nLast value: ", it.Value())
	}

	seek := func(key int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	it.Begin()
	for found := it.NextTo(seek); found; found = it.Next() {
		fmt.Print("\nNextTo key: ", it.Key())
		fmt.Print("\nNextTo value: ", it.Value())
	}

	it.End()
	for found := it.PrevTo(seek); found; found = it.Prev() {
		fmt.Print("\nNextTo key: ", it.Key())
		fmt.Print("\nNextTo value: ", it.Value())
	}

}
