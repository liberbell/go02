package main

import "fmt"

func contentType(url string) (string, error) {
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("Can`t find Content-type header")
	}
	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
