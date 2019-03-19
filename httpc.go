package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can`t call httpbin.org")
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
	fmt.Println("-----------")

	job := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf Bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can`t encode job - %s", err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can`t call httpbin.org")
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}