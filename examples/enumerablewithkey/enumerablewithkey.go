package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/maps/treemap"
)

func printMap(txt string, m *treemap.Map[string, int]) { _ = "STUB: not implemented"; return }

func main() {
	m := treemap.New[string, int]()
	m.Put("g", 7)
	m.Put("f", 6)
	m.Put("e", 5)
	m.Put("d", 4)
	m.Put("c", 3)
	m.Put("b", 2)
	m.Put("a", 1)
	printMap("Initial", m)

	even := m.Select(func(key string, value int) bool {
		return value%2 == 0
	})
	printMap("Elements with even values", even)

	foundKey, foundValue := m.Find(func(key string, value int) bool {
		return value%2 == 0 && value%3 == 0
	})
	if foundKey != "" {
		fmt.Println("Element with value divisible by 2 and 3 found is", foundValue, "with key", foundKey)
	}

	square := m.Map(func(key string, value int) (string, int) {
		return key + key, value * value
	})
	printMap("Elements' values squared and letters duplicated", square)

	bigger := m.Any(func(key string, value int) bool {
		return value > 5
	})
	fmt.Println("Map contains element whose value is bigger than 5 is", bigger)

	positive := m.All(func(key string, value int) bool {
		return value > 0
	})
	fmt.Println("All map's elements have positive values is", positive)

	evenNumbersSquared := m.Select(func(key string, value int) bool {
		return value%2 == 0
	}).Map(func(key string, value int) (string, int) {
		return key, value * value
	})
	printMap("Chaining", evenNumbersSquared)
}
