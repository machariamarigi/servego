package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting REST API")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err)
	}
}
