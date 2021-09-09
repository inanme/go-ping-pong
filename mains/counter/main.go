package main

import (
	"fmt"
	"github.com/inanme/counter"
	"github.com/inanme/counter/v2"
	"github.com/inanme/counter/v3"
)

func main() {
	fmt.Println(v2.Counter())
	fmt.Println(banana.Counter())
	fmt.Println(counter.Counter())
}
