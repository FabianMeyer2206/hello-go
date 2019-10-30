package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Print("Received request.")
	fmt.Fprintf(w, "Hello, {{name}}\n")
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
