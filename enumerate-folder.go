package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	files, _ := filepath.Glob("*.*")
	for _, f := range files {
		if f == "es_query.go" {
			fmt.Println("\nfound it!: ", f, "\n")
		} else {
			fmt.Println(f)
		}
	}

}
