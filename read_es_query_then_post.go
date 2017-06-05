package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "query.txt"
	url := "http://localhost:9200/softphone/logs/_search"

	fmt.Println("\n..reading a file named:", path)
	fmt.Println("\n..query read:")

	dat, err := ioutil.ReadFile(path)
	check(err)

	fmt.Println("\ndat is (as a byte stream):", dat)
	fmt.Println("\ndat is (as a string):", string(dat))

	lines := strings.Split(string(dat), "\n")

	for _, element := range lines {
		fmt.Println(element)
	}

	fmt.Println("\n..posting to", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dat))
	check(err)
	defer resp.Body.Close()

	raw_body, _ := ioutil.ReadAll(resp.Body)
	body_as_string := string(raw_body)

	fmt.Println("\n" + body_as_string)

}
