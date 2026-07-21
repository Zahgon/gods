package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/sets/treeset"
)

type User struct {
	id   int
	name string
}

func byID(a, b User) int { _ = "STUB: not implemented"; return 0 }

func main() {
	set := treeset.NewWith(byID)

	set.Add(User{2, "Second"})
	set.Add(User{3, "Third"})
	set.Add(User{1, "First"})
	set.Add(User{4, "Fourth"})

	fmt.Println(set)
}
