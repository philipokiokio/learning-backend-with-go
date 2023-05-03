package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my awesome Site! </h1>")
}

func main() {

	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server started at http://localhost:4590!")

	http.ListenAndServe(":4590", nil)

}
