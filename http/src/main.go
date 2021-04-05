package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

	for h, value := range r.Header {
		fmt.Printf("%v:\t\t%v\n", h, value)
	}
	w.Header().Add("test-header", "test-value")
	w.Write([]byte("hello"))
}
