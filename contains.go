package main

import (
	"fmt"
)

func contains(candidate string) {
	if temp[candidate] {
		fmt.Println("true ", candidate)
	} else {
		fmt.Println("false ", candidate)
	}
}

var temp = make(map[string]bool)

func main() {
	sliceOne := []string{"52200", "53233", "54999", "52525"}
	for _, val := range sliceOne {
		temp[val] = true
	}

	sliceTwo := []string{"42323", "265562", "52200", "62562", "32323", "287878"}
	for _, val := range sliceTwo {
		contains(val)
	}
}
