package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":4000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from go server!</h1>")
	})

	fmt.Printf("Starting server at port 4000\n")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
