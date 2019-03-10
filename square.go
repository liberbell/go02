package main

import "fmt"

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be >=0(was %s)", length)
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}

func (s, *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}
