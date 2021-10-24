package filter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func even(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func odd(i int) bool {
	if i%2 == 1 {
		return true
	}
	return false
}

func Test_odd(t *testing.T) {
	assert.Equal(t, []int{1, 3, 5}, Filter([]int{1, 2, 3, 4, 5, 6}, odd), "they should be equal")
}

func Test_even(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6}, Filter([]int{1, 2, 3, 4, 5, 6}, even), "they should be equal")
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
