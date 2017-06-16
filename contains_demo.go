package main

import (
	"fmt"
)

func main() {

	fmt.Println("Does a value in one slice appear in another?\n .contains() by any other name.\n")

	// Define first slice
	list_one := []string{"52200", "53233", "54999", "52525"}

	// Define second slice
	list_two := []string{"32552", "52200", "43237", "32667"}

	// Define a new map to hold values from the first slice.
	list_one_map := make(map[string]bool)

	// Iterate over first slice. Print values and add them to the map.
	// Just throwing away the index value with "_".
	for _, v := range list_one {
		fmt.Println("list one value:", v)
		list_one_map[v] = true
	}

	// Iterate over the second slice.
	//
	for _, v := range list_two {
		fmt.Println("\nSearching for:", v, "in both slices...")

		if list_one_map[v] {
			fmt.Println("HIT:", v, " is in both slices")
		} else {
			fmt.Println("MISS", v, "is not in both slices")
		}
	}
}
