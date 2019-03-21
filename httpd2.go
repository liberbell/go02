package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var (
	db     = map[string]interface{}{}
	dbLock sync.Mutex
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
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value

	sendResponse(entry, w)
}

func kvGetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[4]

	dbLock.Lock()
	defer dbLock.Unlock()
}
