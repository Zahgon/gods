package main

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/v2/sets/treeset"
)

func main() {
	set := treeset.New[string]()
	set.Add("a", "b", "c")
	it := set.Iterator()

	fmt.Print("\nForward iteration\n")
	for it.Next() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]")
	}

	fmt.Print("\nForward iteration (again)\n")
	for it.Begin(); it.Next(); {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]")
	}

	fmt.Print("\nBackward iteration\n")
	for it.Prev() {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]")
	}

	fmt.Print("\nBackward iteration (again)\n")
	for it.End(); it.Prev(); {
		index, value := it.Index(), it.Value()
		fmt.Print("[", index, ":", value, "]")
	}

	if it.First() {
		fmt.Print("\nFirst index: ", it.Index())
		fmt.Print("\nFirst value: ", it.Value())
	}

	if it.Last() {
		fmt.Print("\nLast index: ", it.Index())
		fmt.Print("\nLast value: ", it.Value())
	}

	seek := func(index int, value string) bool {
		return strings.HasSuffix(value, "b")
	}

	it.Begin()
	for found := it.NextTo(seek); found; found = it.Next() {
		fmt.Print("\nNextTo index: ", it.Index())
		fmt.Print("\nNextTo value: ", it.Value())
	}

	it.End()
	for found := it.PrevTo(seek); found; found = it.Prev() {
		fmt.Print("\nNextTo index: ", it.Index())
		fmt.Print("\nNextTo value: ", it.Value())
	}

}
