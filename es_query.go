package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {
	// Set up the URL
	url := "http://localhost:9200/softphone/logs/_search"
	// Set up the query string
	var es_query_as_string string = `{"_source":["uuid"],"query":{"match_all":{}}}`
	// Encode the query string as a valid buffer for use in 'body'
	var es_query_as_buffer = bytes.NewBufferString(es_query_as_string)
	// Set up the request with method, url and query
	req, _ := http.NewRequest("POST", url, es_query_as_buffer)
	// Set up header properties
	req.Header.Set("X-Custom-Header", "NFT-Client")
	req.Header.Set("Content-Type", "application/json")
	// Create a client
	client := &http.Client{}
	// load the request data into the client and collect the result into 'resp'
	resp, _ := client.Do(req)
	// Close the response stream/body
	defer resp.Body.Close()
	// Parse the response.Body through io/util's ReadAll method.
	raw_body, _ := ioutil.ReadAll(resp.Body)
	// Parse the raw body into a string
	body_as_string := string(raw_body)
	// Print some output based on the string (response)
	fmt.Println("\nresponse Status:", resp.Status)
	fmt.Println("\nresponse Headers:", resp.Header)
	fmt.Println("\n" + body_as_string)
	fmt.Println("\nBody is of type: ", reflect.TypeOf(body_as_string))
}
