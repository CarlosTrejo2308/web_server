package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking auth")
			if r.Header.Get("Authorization") != "Bearer 12345" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next(w, r)
		}

	}
}

func Logging() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			next(w, r)
		}
	}
}
