package interfaces

import (
	"fmt"
	"strconv"
	"testing"
)

type Name1 interface {
	name() string
}

type Name2 interface {
	name() string
}

type Name3 struct {
	Name string
}

func (n *Name3) name() string {
	return n.Name
}

func name1(p Name1) {
	fmt.Println(p.name())
}

func name2(p Name2) {
	fmt.Println(p.name())
}

func name3(p Name3) {
	fmt.Println(p.name())
}

type Type1 string

func (t Type1) name() string {
	return string(t)
}

type Type2 int

func (t Type2) name() string {
	return strconv.Itoa(int(t))
}

func Test_how_bizarre(t *testing.T) {
	name1(Type2(1))
	name2(Type1("mert"))
}
