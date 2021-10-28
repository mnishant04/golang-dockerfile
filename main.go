package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func main(){
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w , "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw , "Hi Nishant")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}