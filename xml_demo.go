package main

import (
	"strings"
	"fmt"
    "io/ioutil"
	"log"
	"net/http"
	"github.com/beevik/etree"
)

func main() {

	// Set up the URL
	url := "http://www.floatrates.com/daily/gbp.xml"

	// Execute GET then close the body
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	//Convert the body to array of uint8
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the body to a string
	responseString:= string(responseData)
	responseString = strings.TrimSpace(responseString)
	responseString = strings.Replace(responseString,`<?xml version="1.0" encoding="ISO-8859-1"?>`,"", 1)

	// Print the string
	fmt.Println("\nPrinting the xml response as a plain string: ")
	fmt.Println(responseString)

	doc := etree.NewDocument()
	err = doc.ReadFromString(responseString)
	if err != nil {
		fmt.Println("Got error: ", err)
	}

	//tt := doc.FindElements(".//item/description")

	//fmt.Println("Title:", tt.Tag, tt.Text() )
	
	for _, t := range doc.FindElements(".//item/targetName") {
    fmt.Println("Desc:", t.Text())
}

}