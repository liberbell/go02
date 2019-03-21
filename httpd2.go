package main

import "sync"

var (
	db     = map[string]interface{}{}
	dblock sync.Mutex
)
