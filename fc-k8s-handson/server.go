package main

import (
	"fmt"
	"net/http"
	"os"
)
func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprint(w, "Hello, I'm %s. i'm %s.", name, age)
}