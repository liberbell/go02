package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

func main() {
	p := Point{1, 2}
	p.Move(2, 3)
	fmt.Printf("%+v\n", p)
}
