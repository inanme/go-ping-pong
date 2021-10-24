package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		select {
		case <-time.Tick(d):
			ch <- i
			i++
		}
	}
}

func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	data := make(chan int)
	out := make(chan int)
	go producer(data, 100*time.Millisecond)
	go producer(data, 250*time.Millisecond)
	go reader(out)
	for i := range data {
		out <- i
	}
}
