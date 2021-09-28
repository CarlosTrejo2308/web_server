package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHanlder(path, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]

	handler, methodExist := r.rules[path][method]

	return handler, pathExist, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exist, methodExist := r.FindHanlder(req.URL.Path, req.Method)
	if !exist {
		http.NotFound(w, req)
		return
	}

	if !methodExist {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	handler(w, req)

	//http.NotFound(w, req)
}
