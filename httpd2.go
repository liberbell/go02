package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	db     = map[string]interface{}{}
	dblock sync.Mutex
)

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
}
