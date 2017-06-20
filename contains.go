package main

import (
	"fmt"
)

func contains(candidate string) {
	if Temp[candidate] {
		fmt.Println("true ", candidate)
	} else {
		fmt.Println("false ", candidate)
	}
}

var Temp = make(map[string]bool)

func main() {
	slice_one := []string{"52200", "53233", "54999", "52525"}
	for _, val := range slice_one {
		Temp[val] = true
	}

	slice_two := []string{"42323", "265562", "52200", "62562", "32323", "287878"}
	for _, val := range slice_two {
		contains(val)
	}
}
