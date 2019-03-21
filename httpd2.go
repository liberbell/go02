package main

import "sync"

var (
	db     = map[string]interface{}{}
	dblock sync.Mutex
)

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
