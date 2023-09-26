package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// When user visits / respond with hello world
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World!")
	// })
	// // when user visit /john respond with hello john
	// http.HandleFunc("/john", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello John!")
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 2 && parts[1] != "" {
			name := parts[1]
			//if the user is jack ----return response error 404 not found
			if name == "jack" {
				// http.NotFound(w, r)
				w.WriteHeader(404)
			} else {
				//say hello to the user based on the incoming request
				fmt.Fprintf(w, "Hello, %s!", name)
			}
		} else {
			fmt.Fprint(w, "Hello, World!")
		}
	})
	http.ListenAndServe(":8090", nil)
}
