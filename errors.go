package main

import "fmt"

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, 0, fmt.Errorf("sqrt of negative value (%f)", n)
	}
}

func main() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}
}
