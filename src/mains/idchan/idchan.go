package main

import (
	"fmt"
	"time"
)

func makeID(c chan int) {
	i := 0
	for {
		fmt.Printf("Sending %d\n", i)
		c <- i
		i++
	}
}

func main() {
	t := time.NewTimer(3 * time.Second)
	c := make(chan int, 4)
	go makeID(c)
	fmt.Printf("hello %d\n", <-c)
	fmt.Printf("hello %d\n", <-c)
	<-t.C
}
