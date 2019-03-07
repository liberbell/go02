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

}
