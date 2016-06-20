package main

import (
	"fmt"
	"strings"
)

// Because Go does not support generic data structures, it is necessary to
// implement our own methods on various collection types (think: map, filter,
// indexOf, etc.)

func IndexOf(collection []string, value string) int {
	for i, val := range collection {
		if val == value {
			return i
		}
	}

	return -1
}

func Contains(collection []string, value string) bool {
	return IndexOf(collection, value) >= 0
}

func Any(collection []string, predicate func(string) bool) bool {
	for _, val := range collection {
		if predicate(val) {
			return true
		}
	}

	return false
}

func All(collection []string, predicate func(string) bool) bool {
	for _, val := range collection {
		if !predicate(val) {
			return false
		}
	}

	return true
}

func Filter(collection []string, predicate func(string) bool) []string {
	filtered := make([]string, 0)

	for _, val := range collection {
		if predicate(val) {
			filtered = append(filtered, val)
		}
	}

	return filtered
}

func Map(collection []string, predicate func(string) string) []string {
	// Rather than creating a 0-sized slice and appending, we use the length of
	// the initial collection to ensure that the size is maintained and assign
	// based on the index within the loop (I suppose this is to handle cases
	// where the value at a specific index is nil or empty?)

	mapped := make([]string, len(collection))

	for i, val := range collection {
		mapped[i] = predicate(val)
	}

	return mapped
}

func main() {
	var animals = []string{"dog", "cat", "mouse", "horse", "elephant", "deer", "elk"}

	fmt.Println(IndexOf(animals, "mouse"))
	fmt.Println(Contains(animals, "elk"))
	fmt.Println(Contains(animals, "puma"))

	// Do any animals start with “e”?
	fmt.Println(Any(animals, func(s string) bool {
		return strings.HasPrefix(s, "e")
	}))

	// Do any animals start with “j”?
	fmt.Println(Any(animals, func(s string) bool {
		return strings.HasPrefix(s, "j")
	}))

	// Grab all animals that start or end with the letter “e”
	fmt.Println(Filter(animals, func(s string) bool {
		return strings.HasPrefix(s, "e") || strings.HasSuffix(s, "e")
	}))

	// Add some definite articles
	fmt.Println(Map(animals, func(s string) string {
		return "The " + s
	}))
}
