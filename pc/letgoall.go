package main

import (
	"fmt"
	"time"
)

func start(start chan int, id int) {
	<-start
	fmt.Printf("mine is %d\n", id)
}

func main() {
	dummyChannel := make(chan int)
	for i := 0; i < 10; i++ {
		go start(dummyChannel, i)
	}

	time.Sleep(1 * time.Second)
	close(dummyChannel)
	time.Sleep(2 * time.Second)
}
