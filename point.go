package main

type Point struct {
	x int
	y int
}

func (p Point) move(dx int, dy int) {
	p.x += dx
	p.y += dy
}
