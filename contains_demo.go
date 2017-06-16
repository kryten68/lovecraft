package main

import (
	"fmt"
)

func main() {

	fmt.Println("Does a value in one slice appear in another?\n .contains() by any other name.\n")

	// Define first slice
	slice_one := []string{"52200", "53233", "54999", "52525"}

	// Define second slice
	slice_two := []string{"32552", "52200", "76554", "87765"}

	// Define a third slice which will contain any matches
	slice_three := []string{}

	// Define a new map to hold values from the first slice.
	temp_map := make(map[string]bool)

	// Iterate over first slice. Print values and add them to the map.
	// Just throwing away the index value with "_".
	for _, v := range slice_one {
		fmt.Println("slice one value:", v)
		temp_map[v] = true
	}

	// Iterate over the second slice.
	//
	for _, v := range slice_two {
		fmt.Println("\nSearching for:", v, "in both slices...")

		if temp_map[v] {
			fmt.Println(v, "is in both slices")
			slice_three = append(slice_three, v)
		} else {
			fmt.Println(v, "is not in both slices")
		}
	}
	println("\n")
	// Print out list_three
	i := len(slice_three)

	if i == 0 {
		fmt.Println("There were no matches found")

	} else if i == 1 {
		fmt.Println("There was one match found:")
		for _, v := range slice_three {
			fmt.Println("list three value:", v)
		}

	} else {
		fmt.Println("There were ", i, "matches found:")
		for _, v := range slice_three {
			fmt.Println("list three value:", v)
		}
	}
}
