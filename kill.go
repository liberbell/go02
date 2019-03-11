package main

import (
	"io/ioutil"
	"log"
	"os"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return error.Wrap(err, "can`t open pid file(is server runnig?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warnig: Can`t remove pid file - %s", err)
	}
}
