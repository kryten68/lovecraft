package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Someobject struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	// Set up the URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Execute GET then close the body
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Echo the current type of the body object
	i := reflect.TypeOf(response.Body)
	fmt.Println("\nReponse is: ", i)

	// Convert the body to array of uint8
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	i = reflect.TypeOf(responseData)

	// Echo the current type of the body object
	fmt.Println("\nresponseData is: ", i)

	// Convert the body to a string
	responseString := string(responseData)

	// Echo the current type of the body object
	i = reflect.TypeOf(responseString)
	fmt.Println("\nresponseString is: ", i)

	// Print the string
	fmt.Println("\nPrinting the json response as a plain string: ")
	fmt.Println(responseString)

	var foobar Someobject
	err = json.Unmarshal(responseData, &foobar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UserID:", foobar.UserID)
	fmt.Println("ID:", foobar.ID)
	fmt.Println("Title:", foobar.Title)
	fmt.Println("Body:", foobar.Body)
}
