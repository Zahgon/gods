package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/sets/treeset"
)

func printSet(txt string, set *treeset.Set[int]) { _ = "STUB: not implemented"; return }

func main() {
	set := treeset.New[int]()
	set.Add(2, 3, 4, 2, 5, 6, 7, 8)
	printSet("Initial", set)

	even := set.Select(func(index int, value int) bool {
		return value%2 == 0
	})
	printSet("Even numbers", even)

	foundIndex, foundValue := set.Find(func(index int, value int) bool {
		return value%2 == 0 && value%3 == 0
	})
	if foundIndex != -1 {
		fmt.Println("Number divisible by 2 and 3 found is", foundValue, "at index", foundIndex)
	}

	square := set.Map(func(index int, value int) int {
		return value * value
	})
	printSet("Numbers squared", square)

	bigger := set.Any(func(index int, value int) bool {
		return value > 5
	})
	fmt.Println("Set contains a number bigger than 5 is ", bigger)

	positive := set.All(func(index int, value int) bool {
		return value > 0
	})
	fmt.Println("All numbers are positive is", positive)

	evenNumbersSquared := set.Select(func(index int, value int) bool {
		return value%2 == 0
	}).Map(func(index int, value int) int {
		return value * value
	})
	printSet("Chaining", evenNumbersSquared)
}
