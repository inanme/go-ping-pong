package main

import "fmt"

func split(c chan string, slice []string) {
	for _, e := range slice {
		c <- e
	}
	close(c)
}

func main() {
	c := make(chan string)
	go split(c, []string{"mert", "inan", "hello", "fd"})

	for e := range c {
		fmt.Println(e)
	}
}
