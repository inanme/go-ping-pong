package coll

import (
	"fmt"
	"errors"
	"io"
	"testing"
)

type point struct {
	X, Y float64
}

func (p *point) Close() error {
	fmt.Println("doing nothing")
	p.X = p.X * 2.
	return errors.New("error !")
}

func close2(p io.Closer) {
	p.Close()
}

func close1(p interface{}) {
	switch p.(type) {
	case io.Closer:
		fmt.Println("closing")
		p.(io.Closer).Close()
	default:
		println("none")

	}
}

func Test_should_close(test *testing.T) {
	point := &point{X: 1.1, Y: 2.2}
	close1(point)
}

func Test_should_close2(test *testing.T) {
	point := &point{X: 1.1, Y: 2.2}
	close2(point)
}
