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
	ping = constant.Ping
	pong = constant.Pong
)

func replyMessage(message string) string {
	switch message {
	case ping:
		return pong
	case pong:
		return ping
	default:
		return ""
	}
}

func actor(channel chan string, id int, init string) {
	if len(replyMessage(init)) > 0 {
		channel <- replyMessage(init)
	}
	for true {
		message := <-channel
		fmt.Printf("%d:%s\n", id, message)
		channel <- replyMessage(message)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan string)
	go actor(channel, 1, ping)
	go actor(channel, 2, "")

	time.Sleep(10 * time.Second)
}
