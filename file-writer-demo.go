package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write_something(path string, to_be_written string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	check(err)
	f.WriteString(to_be_written)
	f.Close()
}

func main() {

	path := "test.txt"

	fmt.Println("Going to write something to the file now...")

	some_stuff := "just some demo text....\n"
	write_something(path, some_stuff)

}
