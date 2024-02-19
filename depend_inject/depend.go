package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
)

func Greet(writter io.Writer, name string) {
	fmt.Fprintf(writter, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":8088", http.HandlerFunc(MyGreeterHandler)))
}