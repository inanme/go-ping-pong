package main

import "fmt"

type f1 interface {
	m()
}

type f2 interface {
	m()
}

type m struct{}

func (_ m) m() {
	fmt.Println("m")
}

type n struct{}

func (_ n) m() {
	fmt.Println("n")
}

func goman(f f1) {
	f.m()
}
func main() {
	goman(m{})
	goman(n{})
}
