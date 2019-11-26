package main

import "fmt"

import "net/http"

import "log"

func main() {
	fmt.Println("GO Docker Tut")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
