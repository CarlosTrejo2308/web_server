package main

import (
	"net/http"
)

// Server is a http server on a port and handles a router
type Server struct {
	port   string
	router *Router
}

// NewServer returns a new server on a given port
func NewServer(port string) *Server {
	return &Server{port, NewRouter()}
}

// Handle registers a handler on a given path and method
func (s *Server) Hanlde(method, path string, handler http.HandlerFunc) {
	// If the path doesn´t exist, we create one
	if _, ok := s.router.rules[path]; !ok {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	// Add the handler
	s.router.rules[path][method] = handler
}

// AddMiddleware adds a middleware to the router
func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// Call each middleware, and pass the next middleware to it
	for _, m := range middlewares {
		f = m(f)
	}

	// Return the final handler
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	return err
}
