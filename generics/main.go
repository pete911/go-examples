package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("sum of %v: %d\n", ints, Sum(ints))

	floats := []float64{1.0, 2.5, 3.1, 4, 5, 6.9, 7, 8, 9}
	fmt.Printf("sum of %v: %.2f\n", floats, Sum(floats))
}

func Sum[T Number](in []T) T {
	var sum T
	for i := range in {
		sum = sum + in[i]
	}
	return sum
}
