package main

// filename does not matter, folder package name is the most important thing
// in the project root level
// GOPATH should not include "src"
// export GOPATH=$PWD && go build hello2
// or go install hello hello2

import (
	"github.com/go-ping-pong/constant"
	"fmt"
	"time"
)

const (
	ping = constant.Ping2
	pong = constant.Pong2
)

func replyMessage(message string) string {
	if message == ping {
		return pong
	} else if message == pong {
		return ping
	} else {
		return ""
	}
}

func actor(channel chan string, init string) {
	if len(replyMessage(init)) > 0 {
		channel <- replyMessage(init)
	}
	for true {
		message := <-channel
		fmt.Println(message)
		channel <- replyMessage(message)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan string)
	go actor(channel, ping)
	go actor(channel, "")

	time.Sleep(10 * time.Second)
}
