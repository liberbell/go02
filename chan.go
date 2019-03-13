package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// <-ch
	// fmt.Println("Here")
	go func() {
		ch <- 353
	}()
	val := <-ch
	fmt.Printf("got %d\n", val)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("recieved %d\n", val)
	}
}
