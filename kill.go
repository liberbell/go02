package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return error.Wrap(err, "can`t open pid file(is server runnig?)")
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("warnig: Can`t remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	fmt.Printf("killing server with PID=%d\n", pid)
}
