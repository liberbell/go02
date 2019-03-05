package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	val2 := divmod(6, 4)
	fmt.Println(val2)
}
