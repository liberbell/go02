package main

import "fmt"

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func newTrade(Symbol string, Volume int, Price float64, Buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can`t be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0(was %d)", volume)
	}

	if price <= 0 {
		return nil, fmt.Errorf("Price must be >= 9(was %d)", price)
	}
	Trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
	}
}
