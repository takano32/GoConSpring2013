package main

import (
	"fmt"
	"math"
)

func loopCalc() float64 {
	pi := float64(0)
	for i := 0; i < 10; i++ {
		j := float64(i)
		var x1 float64 = 4 / (8 * j + 1)
		var x2 float64 = 2 / (8 * j + 4)
		var x3 float64 = 1 / (8 * j + 5)
		var x4 float64 = 1 / (8 * j + 6)
		pi += (x1 - x2 - x3 -x4) / math.Pow(16, j)
	}
	return pi
}

func parallelCalc() float64 {
	pi := float64(0)
	k := func(i int, c chan float64) {
		j := float64(i)
		var x1 float64 = 4 / (8 * j + 1)
		var x2 float64 = 2 / (8 * j + 4)
		var x3 float64 = 1 / (8 * j + 5)
		var x4 float64 = 1 / (8 * j + 6)
		c <- (x1 - x2 - x3 -x4) / math.Pow(16, j)
	}
	c := make(chan float64)
	n := int(10)
	for i := 0; i < n; i++ {
		go k(i, c)
	}
	for i := 0; i < n; i++ {
		pi += <- c
	}
	return pi
}

func main() {
	fmt.Println(loopCalc())
	fmt.Println(parallelCalc())
}


