package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
	fmt.Print("Hello world")
}
