package main

import (
	"fmt"
	"flag"
)

func main() {

	testid := flag.String("testid","nothing provided","a string")
	start := flag.String("start","default","a start datetime")
	end := flag.String("end","default","an end datetime")
	flag.Parse()

	if (*testid == "nothing provided") {
		fmt.Println("Nothing to report. Nothing provided")
	} else {
		fmt.Println("Got for testid param: ", *testid)
	}

	if (*start == "default") {
		fmt.Println("Nothing to report. Nothing provided")
	} else {
		fmt.Println("Got for start param: ", *start)
	}

	if (*end == "default") {
		fmt.Println("Nothing to report. Nothing provided")
	} else {
		fmt.Println("Got for end param: ", *end)
	}

}
