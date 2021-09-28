// package Handler is a collection of functions that handle requests (endpoints)
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleRoot is the root handler
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
}

// HanldeHome is the home handler
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

// PostRequest is the handler for POST requests
func PostRequest(w http.ResponseWriter, r *http.Request) {
	// Get a decoder from the request body
	decoder := json.NewDecoder(r.Body)

	// Decode the request body into a struct
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	// Write the response
	fmt.Fprintf(w, "Received: %v\n", metaData)
}

// UserPostRequest is the handler for User POST requests
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	resp, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
