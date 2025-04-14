package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	getURL := "https://jsonplaceholder.typicode.com/posts/1"
	response, err := http.Get(getURL)
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer response.Body.Close() // Ensures that the response body is closed to free up resources.

	// io.ReadAll reads from an io.Reader (response.Body in this case) and returns a byte slice ([]byte).
	// It essentially converts the response body (stream data) into a slice of bytes for easy manipulation.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading GET response:", err)
		return
	}
	fmt.Println(string(body)) // Convert byte slice to string for readable output.

	// Prepare data for the POST request
	postURL := "https://jsonplaceholder.typicode.com/posts"
	postData := map[string]interface{}{
		"title":  "Go HTTP POST Example",
		"body":   "This is a sample request body",
		"userId": 1,
	}

	// json.Marshal converts a Go data structure (map in this case) into a JSON byte slice.
	// This is necessary because HTTP POST requests require JSON-formatted payloads.
	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// bytes.NewBuffer creates a new buffer (like a wrapper around the byte slice)
	// that implements the io.Reader interface, which is required for the HTTP request body.
	response, err = http.Post(postURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return
	}
	defer response.Body.Close() // Ensures that the response body is properly closed.

	// Again, io.ReadAll reads from an io.Reader (response.Body) and converts it into a byte slice ([]byte).
	// This allows us to work with the full response data in memory rather than reading it in chunks.
	body, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading POST response:", err)
		return
	}

	fmt.Println("POST Response:")
	fmt.Println(string(body)) // Convert byte slice to string for readable output.
}
