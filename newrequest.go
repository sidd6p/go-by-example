package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Perform a GET request with headers
	getURL := "https://jsonplaceholder.typicode.com/posts/1"
	req, err := http.NewRequest("GET", getURL, nil) // Create a new GET request
	if err != nil {
		fmt.Println("Error creating GET request:", err)
		return
	}

	// Adding custom headers to the GET request
	req.Header.Set("User-Agent", "MyGoClient/1.0")
	req.Header.Set("Accept", "application/json")

	// &http.Client{} creates a new HTTP client with a timeout.
	// The Timeout field specifies the maximum time for the request to complete.
	// If the request takes longer than the specified time, it will be cancelled.
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout set to 10 seconds
	}

	// http.DefaultClient is a predefined HTTP client with default settings.
	// It's a globally available client instance in Go and can be used if you don't need customization.
	// It doesn't allow setting custom timeouts or any other settings.
	// If you need more control, like custom timeouts, proxies, or retries, use a custom client as shown.
	//
	// Example of using http.DefaultClient for a request:
	// response, err := http.DefaultClient.Do(req)  // Using default client without custom timeout

	// Send the GET request using the custom client
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error in GET request:", err)
		return
	}
	defer response.Body.Close() // Close response body to free resources

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading GET response:", err)
		return
	}
	fmt.Println("GET Response:")
	fmt.Println(string(body))

	// Prepare data for the POST request
	postURL := "https://jsonplaceholder.typicode.com/posts"
	postData := map[string]interface{}{
		"title":  "Go HTTP POST Example",
		"body":   "This is a sample request body",
		"userId": 1,
	}

	// Convert Go map to JSON
	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new POST request with headers
	req, err = http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return
	}

	// Setting headers for the POST request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN") // Example of an Authorization header

	// Send the POST request using the same custom client
	response, err = client.Do(req)
	if err != nil {
		fmt.Println("Error in POST request:", err)
		return
	}
	defer response.Body.Close() // Close response body to free resources

	// Read the POST response
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading POST response:", err)
		return
	}

	fmt.Println("POST Response:")
	fmt.Println(string(body))
}
