package main

import (
	"fmt"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers!")
}

func main() {
	http.Handlefunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
