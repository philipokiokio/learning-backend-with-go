package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/"{
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<h1>Welcome to my awesome Site! </h1>")

	} else if r.URL.Path == "/contact/"{
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "To get in touch, please send an email to <a href=\"mailto:go@gotunnel.com\">go@gotunnel.com</a>")

	}else{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"<h1>404 Page Not Found</h1>")
	}



}

func main() {

	http.HandleFunc("/", handlerFunc)

	fmt.Println("Server started at http://localhost:4590!")

	http.ListenAndServe(":4590", nil)

}
