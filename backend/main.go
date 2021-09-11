package main

import (
	"fmt"
	"net/http"
)

type ResponseString string

func (rs ResponseString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rs)
}

func main() {
	fmt.Println("starting...")
	http.Handle("/api", ResponseString("Hello World!"))
	http.ListenAndServe(":8080", nil)
}
