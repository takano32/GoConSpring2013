package main

import (
	"fmt"
	"math"
)

func loopCalc(n int) float64 {
	pi := float64(0)
	for i := 0; i < n; i++ {
		j := float64(i)
		var x1 float64 = 4 / (8 * j + 1)
		var x2 float64 = 2 / (8 * j + 4)
		var x3 float64 = 1 / (8 * j + 5)
		var x4 float64 = 1 / (8 * j + 6)
		pi += (x1 - x2 - x3 -x4) / math.Pow(16, j)
	}
	return pi
}

func parallelCalc(n int) float64 {
	pi := float64(0)
	k := func(i int,procs int, c chan float64) {
		var sum float64
		for ii := procs * i; ii < i + procs; ii++ {
			j := float64(ii)
			var x1 float64 = 4 / (8 * j + 1)
			var x2 float64 = 2 / (8 * j + 4)
			var x3 float64 = 1 / (8 * j + 5)
			var x4 float64 = 1 / (8 * j + 6)
			sum += (x1 - x2 - x3 -x4) / math.Pow(16, j)
		}
		c <- sum
	}
	procs := 100
	c := make(chan float64, 16)
	m := n / procs
	for i := 0; i < m; i++ {
		go k(i, procs, c)
	}
	for i := 0; i < m; i++ {
		pi += <- c
	}
	return pi
}

func main() {
	fmt.Println(loopCalc(1000))
	fmt.Println(parallelCalc(1000))
}

