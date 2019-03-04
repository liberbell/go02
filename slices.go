package main

import "fmt"

func main() {
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (Type %T)\n", loons, loons)

	fmt.Println(len(loons)) // 3
	fmt.Println("--------")

	fmt.Println(loons[1])
	fmt.Println("--------")

	fmt.Println(loons[1:])
}
