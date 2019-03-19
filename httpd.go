package main

import "fmt"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers!")
}

func main() {
	http.Handlefunc("/hello", helloHandler)
}
