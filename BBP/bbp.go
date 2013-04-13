package main

import (
	"fmt"
	"math"
	"math/big"
)

func loopCalc(n int) *big.Rat {
	pi := big.NewRat(0, 1)
	for i := 0; i < n; i++ {
		j := int64(i)
		x1 := big.NewRat(4, 8 * j + 1)
		x2 := big.NewRat(2, 8 * j + 4)
		x3 := big.NewRat(1, 8 * j + 5)
		x4 := big.NewRat(1, 8 * j + 6)
		a := big.NewRat(1, 1)
		a.Sub(x1, x2)
		b := big.NewRat(1, 1)
		b.Sub(x3, x4)
		c := big.NewRat(1, 1)
		d := c.Sub(a, b)
		e := big.NewInt(1)
		for k := 0; k < i; k++ {
			e.Mul(e, big.NewInt(16))
		}
		num := d.Num()
		denom := big.NewInt(1)
		denom.Mul(d.Denom(), e)
		f := big.NewRat(1, 1)
		f.SetFrac(num, denom)
		pi.Add(pi, f)
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
	fmt.Println(loopCalc(10))
	fmt.Println(parallelCalc(1000))
}

