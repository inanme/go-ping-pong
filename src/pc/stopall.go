package main

import (
	"fmt"
	"time"
)

func stop(stop chan int, id int) {
	for {
		select {
		case <-stop:
			return
		default:
			fmt.Printf("Waiting %d\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	dummyChannel := make(chan int)
	for i := 0; i < 10; i++ {
		go stop(dummyChannel, i)
	}

	time.Sleep(3 * time.Second)
	close(dummyChannel)
}
