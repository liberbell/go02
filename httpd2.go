package main

import (
	"encoding/json"
	"log"
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
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encode %+v - %s", entry, err)
	}
}

func kvPostHandler(w http.ResponseWriter, r *http.Request) {

}
