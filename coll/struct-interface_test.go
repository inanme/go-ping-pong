package coll

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type MySlice []int

type Rectangle struct {
	h, w int
}

type Square struct {
	Rectangle
}

type Area interface {
	area() int
}

func (rec* Rectangle) area() int {
	return rec.w * rec.h
}
func (mySlice MySlice) area() int {
	var result int
	result = 1
	for _, e := range mySlice {
		result *= e
	}
	return result
}

func MakeSquare(x int) *Square {
	return &Square{Rectangle{x,x}}
}

func giveMeArea(area Area) int {
	return area.area()
}

func Test_square(t *testing.T) {
	sq := MakeSquare(3)
	assert.Equal(t, 9, giveMeArea(sq))
}

func Test_mySlice(t *testing.T) {
	myList := MySlice{1, 2, 3, 4}
	assert.Equal(t, 24, giveMeArea(myList))
}
