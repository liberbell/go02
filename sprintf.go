package main

import "fmt"

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)
	fmt.Println("s = %v (Type %T)\n", n, n)
	fmt.Printf("s = %v (Type %T)\n", s, s)
}
