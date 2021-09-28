// package reouter handles the endpoints and routes them to the correct handler.
package main

import (
	"net/http"
)

// Router saves the handler for each path and method endpoint
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// FindHanlder finds the handler for the path and method.
func (r *Router) FindHanlder(path, method string) (http.HandlerFunc, bool, bool) {
	// Does path exist?
	_, pathExist := r.rules[path]

	// Does method exist?
	handler, methodExist := r.rules[path][method]

	// Return the handler, path exist and method exist
	return handler, pathExist, methodExist
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Find the handler for the path and method
	handler, exist, methodExist := r.FindHanlder(req.URL.Path, req.Method)
	// Path doesnt exist
	if !exist {
		http.NotFound(w, req)
		return
	}

	// Method doesnt exist
	if !methodExist {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Call the handler
	handler(w, req)
}
