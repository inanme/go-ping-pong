package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	timer := time.NewTimer(time.Second * 10)

	go func(chan1 chan<- string) {
		for {
			chan1 <- "from chan1"
			time.Sleep(time.Second * 2)
		}
	}(chan1)

	go func(chan2 chan<- string) {
		for {
			chan2 <- "from chan2"
			time.Sleep(time.Second * 3)
		}
	}(chan2)

	for {
		select {
		case v := <-chan1:
			fmt.Println(v)
		case v := <-chan2:
			fmt.Println(v)
		case <-timer.C:
			fmt.Printf("exiting")
			return
		}
	}
}
