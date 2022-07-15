package main

import (
	"fmt"
)

func main() {

	// #1: Nil Map: Read-Only
	var dict map[string]string

	// #2: Map retrieval is O(1) — on average.
	key := "good"

	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)

	// #1B
	fmt.Printf("# of Keys: %d\n", len(dict))

}
